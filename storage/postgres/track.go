package postgres

import (
	pb "composition_service/genproto"
	"composition_service/help"
	"composition_service/models"
	"database/sql"
	"time"
)

type TrackRepository struct {
	Db *sql.DB
}

func NewTrackRepository(db *sql.DB) *CompositionRepository {
	return &CompositionRepository{Db: db}
}

func (repo TrackRepository) CreateTrack(track *pb.CreateTrackRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("insert into compositions(composition_id,user_id,title,file_url,created_at)", track.CompositionId, track.UserId, track.Title, track.FileUrl, time.Now())
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (repo TrackRepository) UpdateTrack(track *pb.UpdateTrackRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update tracks set composition_id=$1,user_id=$2,title=$3,file_url=$4,updated_at=$5 where id=$6 and deleted_at=0 and composition_id=$7)", track.CompositionId, track.Userid, track.Title, track.FileUrl, time.Now(), track.Id, track.CompositionId)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, err
}
func (repo TrackRepository) DeleteTrack(track *pb.DeleteTrackRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update tracks set deleted_at=$1 where id=$2 and deleted_at is null and  composition_id=$3)", time.Now(), track.TrackId, track.CompositionId)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (repo TrackRepository) GetTrackById(compositionId string, trackId string) (*models.Track, error) {
	rows, err := repo.Db.Query("select composition_id,user_id ,title, file_url from tracks  where id=$1 and deleted_at is null and composition_id=$2)", trackId, compositionId)
	if err != nil {
		return nil, err
	}
	var track models.Track
	for rows.Next() {
		err := rows.Scan(&track.Composition_id, &track.User_Id, &track.Title, &track.File_Url)
		if err != nil {
			return nil, err
		}
		return &track, nil
	}
	return nil, err
}
func (repo TrackRepository) GetTrack(track *pb.GetTrackRequest) (*pb.TracksResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	filter := ""
	if len(track.CompositionId) > 0 {
		params["composition_id"] = track.CompositionId
		filter += " and composition_id = :composition_id "

	}
	if len(track.Userid) > 0 {
		params["user_id"] = track.Userid
		filter += " and user_id = :user_id "

	}
	if len(track.FileUrl) > 0 {
		params["file_url"] = track.FileUrl
		filter += " and file_url = :file_url "

	}
	if len(track.Title) > 0 {
		params["title"] = track.Title
		filter += " and title = :title "

	}

	if track.LimitOffset.Limit > 0 {
		params["limit"] = track.LimitOffset.Limit
		limit = ` LIMIT :limit`

	}
	if track.LimitOffset.Offset > 0 {
		params["offset"] = track.LimitOffset.Offset
		limit = ` OFFSET :offset`

	}
	query := "select composition_id,user_id,title ,file_url from tracks  where  deleted_at is null and composition_id=$1"

	query = query + filter + limit + offset
	query, arr = help.ReplaceQueryParams(query, params)
	arr = append(arr, track.CompositionId)
	rows, err := repo.Db.Query(query, arr...)
	var tracks []*pb.TrackResponse
	for rows.Next() {
		var track pb.TrackResponse
		err := rows.Scan(&track.CompositionId, &track.Userid, &track.Title, &track.FileUrl)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, &track)
	}
	return &pb.TracksResponse{TracksResponse: tracks}, err
}
