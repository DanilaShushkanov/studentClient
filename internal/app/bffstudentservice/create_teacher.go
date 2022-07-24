package bffstudentservice

import (
	"context"
	"github.com/danilashushkanov/studentClient/internal/app/bffstudentservice/adapters"
	"github.com/danilashushkanov/studentClient/internal/dto"
)

func (s *Service) CreateTeacher(ctx context.Context, req *dto.TeacherItem) (*dto.TeacherItem, error) {
	teacher, err := s.studentServiceClient.TeacherClient.CreateTeacher(ctx, adapters.CreateTeacherItemToPb(req))
	if err != nil {
		return nil, err
	}

	return adapters.TeacherItemFromPb(teacher), nil
}
