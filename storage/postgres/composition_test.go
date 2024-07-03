package postgres

import (
	"composition_service/config"
	pb "composition_service/genproto"
	"fmt"
	"reflect"
	"testing"
)

func TestCompositionRepository_CreateComposition(t *testing.T) {
	cnf := config.Config{}
	db, err := ConnectionDb(&cnf)
	if err != nil {

		panic(err)
	}

	comp := NewCompositionRepository(db)
	fmt.Println(comp)
	composition := &pb.CreateCompositionRequest{
		Title:       "salom",
		Description: "hozir",
		Status:      "in_progress",
		UserId:      "d3dcbdff-de1c-452d-94da-2bb783f1016a",
	}
	fmt.Println(composition)
	response, err := comp.CreateComposition(composition)
	if err != nil {
		fmt.Println("+++++++++++++++")
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.Void{})
	}

}
func TestCompositionRepository_UpdateComposition(t *testing.T) {
	cnf := config.Config{}
	db, err := ConnectionDb(&cnf)
	if err != nil {

		panic(err)
	}

	comp := NewCompositionRepository(db)
	fmt.Println(comp)
	composition := &pb.UpdateCompositionRequest{
		Id:          "ad6b59c5-8d12-4b9e-ab1e-4d3ba7620c01",
		Title:       "xayr",
		Description: "hozir nima qilyapsan",
		Status:      "in_progress",
		UserId:      "d3dcbdff-de1c-452d-94da-2bb783f1016a",
	}
	fmt.Println(composition)
	response, err := comp.UpdateComposition(composition)
	if err != nil {
		fmt.Println("+++++++++++++++")
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.Void{})
	}
}
func TestCompositionRepository_DeleteComposition(t *testing.T) {
	cnf := config.Config{}
	db, err := ConnectionDb(&cnf)
	if err != nil {

		panic(err)
	}

	comp := NewCompositionRepository(db)
	fmt.Println(comp)
	id := pb.IdRequest{Id: "cb6b231e-290b-430d-823f-0dc2096eee02"}
	response, err := comp.DeleteComposition(&id)
	if err != nil {
		fmt.Println("+++++++++++++++")
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.Void{})
	}

}

//	func TestCompositionRepository_GetCompositionById(t *testing.T) {
//		cnf := config.Config{}
//		db, err := ConnectionDb(&cnf)
//		if err != nil {
//
//			panic(err)
//		}
//
//		comp := NewCompositionRepository(db)
//		fmt.Println(comp)
//		id := pb.IdRequest{Id: "ad6b59c5-8d12-4b9e-ab1e-4d3ba7620c01"}
//		response, err := comp.GetCompositionById(&id)
//		expected := pb.CompositionResponse{
//			Id:          "ad6b59c5-8d12-4b9e-ab1e-4d3ba7620c01",
//			UserId:      "d3dcbdff-de1c-452d-94da-2bb783f1016a",
//			Title:       "xayr",
//			Description: "hozir nima qilyapsan",
//			Status:      "in_progress",
//		}
//		if err != nil {
//			fmt.Println("+++++++++++++++")
//			panic(err)
//		}
//		if !reflect.DeepEqual(response, &expected) {
//			t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.CompositionResponse{})
//		}
//		//fmt.Println(response, "))))))))))))))))))))")
//	}
func TestCompositionRepository_GetCompositionByUserId(t *testing.T) {
	cnf := config.Config{}
	db, err := ConnectionDb(&cnf)
	if err != nil {

		panic(err)
	}

	comp := NewCompositionRepository(db)
	fmt.Println(comp)
	expected := pb.CompositionsResponse{
		CompositionsResponse: nil,
	}
	id := pb.IdRequest{
		Id: "d3dcbdff-de1c-452d-94da-2bb783f1016a",
	}
	response, err := comp.GetCompositionByUserId(&id)
	if err != nil {
		fmt.Println("++++++++++")
		panic(err)
	}

	if reflect.DeepEqual(response, &expected) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.CompositionResponse{})
	}

}

func TestTrackRepository_CreateTrack(t *testing.T) {
	cnf := config.Config{}
	db, err := ConnectionDb(&cnf)
	if err != nil {

		panic(err)
	}

	trac := NewTrackRepository(db)
	fmt.Println(trac)
	track := pb.CreateTrackRequest{
		CompositionId: "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		UserId:        "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		Title:         "salom",
		FileUrl:       "jksdlkja;fsjk",
	}
	response, err := trac.CreateTrack(&track)
	if err != nil {
		fmt.Println("_________")
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.CompositionResponse{})

	}

}
func TestCompositionRepository_DeleteComposition2(t *testing.T) {
	cnf := config.Config{}
	db, err := ConnectionDb(&cnf)
	if err != nil {

		panic(err)
	}

	trac := NewTrackRepository(db)
	fmt.Println(trac)
	tracDelete := pb.DeleteTrackRequest{
		TrackId:       "cfd06c0a-aa56-4755-9e1e-3779647f916e",
		CompositionId: "cfd06c0a-aa56-4755-9e1e-3779647f916e",
	}
	response, err := trac.DeleteTrack(&tracDelete)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.CompositionResponse{})

	}

}
func TestTrackRepository_UpdateTrack(t *testing.T) {
	cnf := config.Config{}
	db, err := ConnectionDb(&cnf)
	if err != nil {

		panic(err)
	}

	trac := NewTrackRepository(db)
	fmt.Println(trac)
	trackUpdate := pb.UpdateTrackRequest{
		Id:            "cfd06c0a-aa56-4755-9e1e-3779647f916e",
		CompositionId: "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		Userid:        "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		Title:         "salom",
		FileUrl:       "+++++++++++++++++++++++",
	}
	response, err := trac.UpdateTrack(&trackUpdate)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.CompositionResponse{})

	}
}

//func TestTrackRepository_GetTrack(t *testing.T) {
//	cnf := config.Config{}
//	db, err := ConnectionDb(&cnf)
//	if err != nil {
//
//		panic(err)
//	}
//
//	trac := NewTrackRepository(db)
//
//	getTrack := pb.GetTrackRequest{
//		CompositionId: "d3dcbdff-de1c-452d-94da-2bb783f1016a",
//		Userid:        "d3dcbdff-de1c-452d-94da-2bb783f1016a",
//		Title:         "salom",
//		FileUrl:       "+++++++++++++++++++++++",
//	}
//
//	fmt.Println(trac)
//	expected := pb.TracksResponse{}
//
//	response, err := trac.GetTrack(&getTrack)
//	if err != nil {
//		fmt.Println("+++++++", err)
//		panic(err)
//	}
//	//if !reflect.DeepEqual(response, &expected) {
//	//	t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.CompositionResponse{})
//	//
//	//}
//}
