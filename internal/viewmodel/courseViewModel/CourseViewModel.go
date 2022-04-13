package courseViewModel

import (
	//"danceonline/internal/viewmodel/subscriptionViewMode"
	"danceonline/internal/viewmodel/videoViewModel"
)

type CourseLessonViewModel struct {
	Id                         int                                       `json:"id"`
	Name                       string                                    `json:"name"`
	PosterSrc                  string                                    `json:"posterSrc"`
	VideoCourseLessonViewModel videoViewModel.VideoCourseLessonViewModel `json:"videoCourseLessonViewModel"`
}

type CoursePreviewViewModel struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	PosterSrc string `json:"posterSrc"`
	Price     int    `json:"price"`
	IsActive  int    `json:"isActive"`
}

type CourseMicroViewModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CourseBuyViewModel struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	//SubscriptionLiteViewModels []subscriptionViewModel.SubscriptionLiteViewModel `json:"subscriptionLiteViewModels"`
}
