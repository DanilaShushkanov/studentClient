package dto

type Students struct {
	Students []*StudentItem
	Count    int64
}

type StudentItem struct {
	ID       string
	FullName string
	Age      int64
	Salary   int64
	Teachers []*TeacherItem
}

type ListStudentAndTeachersResponse struct {
	StudentList *Students
	TeacherList *Teachers
}

type ListStudentAndTeachersRequest struct {
	StudentIds []string
	TeacherIds []string
}
