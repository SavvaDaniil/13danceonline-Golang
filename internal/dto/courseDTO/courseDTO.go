package courseDTO

type CourseIdDTO struct {
	Id_of_course int `json:"id_of_course"`
}

type CourseSearchDTO struct {
	Page int `json:"page"`
}

type CourseLessonDTO struct {
	Id_of_course     int `json:"id_of_course"`
	Number_of_lesson int `json:"number_of_lesson"`
}
