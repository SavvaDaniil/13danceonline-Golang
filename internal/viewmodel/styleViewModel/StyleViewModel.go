package styleViewModel

import (
	"danceonline/internal/viewmodel/teacherViewModel"
)

type StyleLiteViewModel struct {
	Id          int    `json:"id"`
	Index       int    `json:"index"`
	IsIndexDiv5 bool   `json:"isIndexDiv5"`
	Link        string `json:"link"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type StylePreviewWithTeachersLiteViewModel struct {
	Id           int
	Name         string
	TeacherLites []teacherViewModel.TeacherLiteViewModel
}

type StyleMicroViewModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
