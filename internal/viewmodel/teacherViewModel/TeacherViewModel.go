package teacherViewModel

type TeacherLiteViewModel struct {
	Id        int
	Name      string
	PosterSrc string
}

type TeacherMicroViewModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
