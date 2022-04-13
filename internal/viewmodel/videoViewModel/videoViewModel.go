package videoViewModel

type VideoTutorialLessonViewModel struct {
	Id             int    `json:"id"`
	Id_of_tutorial int    `json:"id_of_tutorial"`
	VideoSrc       string `json:"videoSrc"`
}

type VideoCourseLessonViewModel struct {
	Id             int    `json:"id"`
	Id_of_course   int    `json:"id_of_course"`
	NumberOfLesson int    `json:"number_of_lesson"`
	IdsOfVideos    []int    `json:"ids_of_video"`
	VideoSrc       string `json:"videoSrc"`
}
