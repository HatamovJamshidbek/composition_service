package service

import (
	pb "composition_service/genproto"
	"composition_service/storage/postgres"
	"context"
	"google.golang.org/grpc"
)

type CompositionService struct {
	DB *postgres.CompositionRepository
	Db *postgres.TrackRepository
	pb.UnimplementedCompositionServiceServer
}

func NewCompositionService(db *postgres.CompositionRepository, db2 *postgres.TrackRepository) *CompositionService {
	return &CompositionService{DB: db}
}
func (service *CompositionService) CreateComposition(ctx context.Context, in *pb.CreateCompositionRequest, opts ...grpc.CallOption) (*pb.Void, error) {
	response, err := service.DB.CreateComposition(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *CompositionService) UpdateComposition(ctx context.Context, in *pb.UpdateCompositionRequest) (*pb.Void, error) {
	response, err := service.DB.UpdateComposition(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *CompositionService) DeleteComposition(ctx context.Context, in *pb.IdRequest) (*pb.Void, error) {
	response, err := service.DB.DeleteComposition(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *CompositionService) GetCompositionByUserid(ctx context.Context, in *pb.IdRequest) (*pb.CompositionsResponse, error) {
	response, err := service.DB.GetCompositionByUserId(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *CompositionService) GetCompositionById(ctx context.Context, in *pb.IdRequest) (*pb.CompositionResponse, error) {
	response, err := service.DB.GetCompositionById(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *CompositionService) CreateTrack(ctx context.Context, in *pb.CreateTrackRequest) (*pb.Void, error) {
	response, err := service.Db.CreateTrack(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *CompositionService) GetTrack(ctx context.Context, in *pb.GetTrackRequest) (*pb.TracksResponse, error) {
	response, err := service.Db.GetTrack(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *CompositionService) UpdateTrack(ctx context.Context, in *pb.UpdateTrackRequest) (*pb.Void, error) {
	response, err := service.Db.UpdateTrack(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *CompositionService) DeleteTrack(ctx context.Context, in *pb.DeleteTrackRequest) (*pb.Void, error) {
	response, err := service.Db.DeleteTrack(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
