package dto

type Students struct {
	Students []*StudentItem
	Count    int64
}

type StudentItem struct {
	ID       int64
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
	StudentIds []int64
	TeacherIds []int64
}
