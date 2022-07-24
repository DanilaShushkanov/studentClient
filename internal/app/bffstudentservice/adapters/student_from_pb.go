package adapters

import (
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	"github.com/danilashushkanov/studentClient/internal/dto"
)

func StudentItemFromPb(student *api.Student) *dto.StudentItem {
	dtoStudent := &dto.StudentItem{
		ID:       student.GetId(),
		FullName: student.GetFullName(),
		Age:      student.GetAge(),
		Salary:   student.GetSalary(),
		Teachers: make([]*dto.TeacherItem, 0, len(student.GetTeachers())),
	}
	for _, teacher := range student.GetTeachers() {
		dtoStudent.Teachers = append(dtoStudent.Teachers, TeacherItemFromPb(teacher))
	}

	return dtoStudent
}

func TeacherItemFromPb(teacher *api.Teacher) *dto.TeacherItem {
	return &dto.TeacherItem{
		ID:           teacher.GetId(),
		FullName:     teacher.GetFullName(),
		PositionType: PositionTypeFromPb(teacher.GetPositionType()),
		StudentID:    teacher.GetStudentId(),
	}
}

func StudentsFromPb(students []*api.Student) *dto.Students {
	countStudents := len(students)
	studentsDto := &dto.Students{
		Students: make([]*dto.StudentItem, 0, countStudents),
		Count:    int64(countStudents),
	}

	for _, student := range students {
		studentsDto.Students = append(studentsDto.Students, StudentItemFromPb(student))
	}

	return studentsDto
}

func TeachersFromPb(teachers []*api.Teacher) *dto.Teachers {
	countTeachers := len(teachers)
	teachersDto := &dto.Teachers{
		Teachers: make([]*dto.TeacherItem, 0, countTeachers),
		Count:    int64(countTeachers),
	}

	for _, teacher := range teachers {
		teachersDto.Teachers = append(teachersDto.Teachers, TeacherItemFromPb(teacher))
	}

	return teachersDto
}

func PositionTypeFromPb(positionType api.PositionType) dto.PositionType {
	switch positionType {
	case api.PositionType_POSTGRADUATE:
		return dto.PositionTypePostgraduate
	case api.PositionType_ASSISTANT:
		return dto.PositionTypeAssistant
	case api.PositionType_DEAN:
		return dto.PositionTypeDean
	default:
		return "type"
	}
}
