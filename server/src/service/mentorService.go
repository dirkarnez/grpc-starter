package service

import (
	pb "github.com/dirkarnez/mentor/proto"
	"golang.org/x/net/context"
)

type mentorService struct {
}

func NewMentorService() *mentorService {
	return &mentorService{}
}

func (s *mentorService) ListLessons(req *pb.ListLessonsRequest, resp pb.MentorService_ListLessonsServer) error {
	lessons := make(chan *pb.Lesson)

	ids := []int32{2, 3, 4}
	for _, id := range ids {
		go func(id int32) {
			lessons <- &pb.Lesson{
				Id:    id,
				By:    "value.By",
				Score: 100,
				Time:  123,
				Title: "value.Title",
				Url:   "value.Url",
			}
		}(id)
	}

	defer close(lessons)
	for lesson := range lessons {
		resp.Send(&pb.ListLessonsResponse{
			Lesson: lesson,
		})
	}

	return nil
}

func (s *mentorService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}