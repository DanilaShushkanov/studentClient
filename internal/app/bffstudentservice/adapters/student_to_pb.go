package adapters

import (
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	"github.com/danilashushkanov/studentClient/internal/dto"
)

func CreateStudentItemToPb(student *dto.StudentItem) *api.CreateStudentRequest {
	req := &api.CreateStudentRequest{
		FullName: student.FullName,
		Age:      student.Age,
		Salary:   student.Salary,
	}
	for _, teacher := range student.Teachers {
		req.Teachers = append(req.Teachers, CreateTeacherItemToPb(teacher))
	}

	return req
}

func UpdateStudentItemToPb(student *dto.StudentItem) *api.UpdateStudentRequest {
	req := &api.UpdateStudentRequest{
		Id:       student.ID,
		FullName: student.FullName,
		Age:      student.Age,
		Salary:   student.Salary,
	}
	for _, teacher := range student.Teachers {
		req.Teachers = append(req.Teachers, UpdateTeacherItemToPb(teacher))
	}

	return req
}

func UpdateTeacherItemToPb(teacher *dto.TeacherItem) *api.UpdateTeacherRequest {
	return &api.UpdateTeacherRequest{
		Id:           teacher.ID,
		FullName:     teacher.FullName,
		PositionType: PositionTypeToPb(teacher.PositionType),
	}
}

func CreateTeacherItemToPb(teacher *dto.TeacherItem) *api.CreateTeacherRequest {
	return &api.CreateTeacherRequest{
		FullName:     teacher.FullName,
		PositionType: PositionTypeToPb(teacher.PositionType),
		StudentId:    teacher.StudentID,
	}
}

func ListStudentToPb(listDto *dto.ListStudentAndTeachersRequest) *api.ListStudentRequest {
	return &api.ListStudentRequest{
		StudentIds: listDto.StudentIds,
	}
}

func ListTeacherToPb(listDto *dto.ListStudentAndTeachersRequest) *api.ListTeacherRequest {
	return &api.ListTeacherRequest{
		TeacherIds: listDto.TeacherIds,
	}
}

func PositionTypeToPb(positionType dto.PositionType) api.PositionType {
	switch positionType {
	case dto.PositionTypePostgraduate:
		return api.PositionType_POSTGRADUATE
	case dto.PositionTypeAssistant:
		return api.PositionType_ASSISTANT
	case dto.PositionTypeDean:
		return api.PositionType_DEAN
	default:
		return api.PositionType_DEAN
	}
}
