package postgres

import (
	pb "composition_service/genproto"
	"database/sql"
	"time"
)

type CompositionRepository struct {
	Db *sql.DB
}

func NewCompositionRepository(db *sql.DB) *CompositionRepository {
	return &CompositionRepository{Db: db}
}

func (repo CompositionRepository) CreateComposition(composition *pb.CreateCompositionRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("insert into compositions(user_id,title,decription,status,created_at)", composition.UserId, composition.Title, composition.Description, composition.Status, time.Now())
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (repo CompositionRepository) UpdateComposition(composition *pb.UpdateCompositionRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update compositions set user_id=$1,title=$2,decription=$3,status=$4,updated_at=$5 where id=$6 and deleted_at=0)", composition.UserId, composition.Title, composition.Description, composition.Status, time.Now(), composition.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (repo CompositionRepository) DeleteComposition(id *pb.IdRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update compositions set deleted_at=$1 where id=$2 and deleted_at is null)", time.Now(), id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (repo CompositionRepository) GetCompositionById(id *pb.IdRequest) (*pb.CompositionResponse, error) {
	rows, err := repo.Db.Query("select user_id,title ,decription, status from compositions  where id=$1 and deleted_at is null)", id)
	if err != nil {
		return nil, err
	}
	var composition pb.CompositionResponse
	for rows.Next() {
		err := rows.Scan(&composition.UserId, &composition.Title, &composition.Description, &composition.Status)
		if err != nil {
			return nil, err
		}
		return &composition, nil
	}
	return nil, err
}

//func (repo CompositionRepository) GetComposition(composition *pb.) (*[]models.Composition, error) {
//	var (
//		params = make(map[string]interface{})
//		arr    []interface{}
//		limit  string
//		offset string
//	)
//	filter := ""
//	if len(composition.Title) > 0 {
//		params["title"] = composition.Title
//		filter += " and title = :title "
//
//	}
//	if len(composition.Description) > 0 {
//		params["description"] = composition.Description
//		filter += " and description = :description "
//
//	}
//	if len(composition.Status) > 0 {
//		params["status"] = composition.Description
//		filter += " and status = :status "
//
//	}
//	if len(composition.User_id) > 0 {
//		params["user_id"] = composition.Description
//		filter += " and user_id = :user_id "
//
//	}
//
//	if composition.Limit > 0 {
//		params["limit"] = composition.Limit
//		limit = ` LIMIT :limit`
//
//	}
//	if composition.Offset > 0 {
//		params["offset"] = composition.Offset
//		limit = ` OFFSET :offset`
//
//	}
//	query := "select user_id,title ,decription, status from compositions  where  deleted_at is null"
//
//	query = query + filter + limit + offset
//	query, arr = help.ReplaceQueryParams(query, params)
//	rows, err := repo.Db.Query(query, arr...)
//	var compositions []models.Composition
//	for rows.Next() {
//		var composition models.Composition
//		err := rows.Scan(&composition.User_id, &composition.Title, &composition.Description, &composition.Status)
//		if err != nil {
//			return nil, err
//		}
//		compositions = append(compositions, composition)
//	}
//	return &compositions, err
//}

func (repo CompositionRepository) GetCompositionByUserId(userId *pb.IdRequest) (*pb.CompositionsResponse, error) {
	rows, err := repo.Db.Query("select title ,decription, status from compositions  where user_id=$1 and deleted_at is null)", userId)
	if err != nil {
		return nil, err
	}
	var compositions []*pb.CompositionResponse
	for rows.Next() {
		var composition pb.CompositionResponse
		err := rows.Scan(&composition.UserId, &composition.Title, &composition.Description, &composition.Status)
		if err != nil {
			return nil, err
		}
		compositions = append(compositions, &composition)
	}
	return &pb.CompositionsResponse{CompositionsResponse: compositions}, err
}
