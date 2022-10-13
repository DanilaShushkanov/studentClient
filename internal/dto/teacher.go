package dto

type Teachers struct {
	Teachers []*TeacherItem
	Count    int64
}

type TeacherItem struct {
	ID           string
	PositionType PositionType
	FullName     string
	StudentID    string
}

type PositionType string

const (
	PositionTypePostgraduate PositionType = "POSTGRADUATE"
	PositionTypeAssistant    PositionType = "ASSISTANT"
	PositionTypeDean         PositionType = "DEAN"
)
