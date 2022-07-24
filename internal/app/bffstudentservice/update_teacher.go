package bffstudentservice

import (
	"context"
	"github.com/danilashushkanov/studentClient/internal/app/bffstudentservice/adapters"
	"github.com/danilashushkanov/studentClient/internal/dto"
)

func (s *Service) UpdateTeacher(ctx context.Context, req *dto.TeacherItem) (*dto.TeacherItem, error) {
	teacher, err := s.studentServiceClient.TeacherClient.PatchTeacher(ctx, adapters.UpdateTeacherItemToPb(req))
	if err != nil {
		return nil, err
	}

	return adapters.TeacherItemFromPb(teacher), nil
}
