package dto

type Teachers struct {
	Teachers []*TeacherItem
	Count    int64
}

type TeacherItem struct {
	ID           int64
	PositionType PositionType
	FullName     string
	StudentID    int64
}

type PositionType string

const (
	PositionTypePostgraduate PositionType = "POSTGRADUATE"
	PositionTypeAssistant    PositionType = "ASSISTANT"
	PositionTypeDean         PositionType = "DEAN"
)
