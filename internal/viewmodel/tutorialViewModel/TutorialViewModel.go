package tutorialViewModel

import (
	"danceonline/internal/viewmodel/subscriptionViewModel"
	"danceonline/internal/viewmodel/videoViewModel"
)

type TutorialLessonViewModel struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	PosterSrc string `json:"posterSrc"`
	//VideoSrc  *string `json:"videoSrc"`
	VideoTutorialLessonViewModel videoViewModel.VideoTutorialLessonViewModel `json:"videoTutorialLessonViewModel"`
}

type TutorialPreviewViewModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	//DescLeft               string                                   `json:"descLeft"`
	//TeacherMicroViewModels []teacherViewModel.TeacherMicroViewModel `json:"teacherMicroViewModels"`
	PosterSrc string `json:"posterSrc"`
	//StyleMicroViewModel    []styleViewModel.StyleMicroViewModel     `json:"styleMicroViewModel"`
	Price    int `json:"price"`
	IsActive int `json:"isActive"`
}

type TutorialMicroViewModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TutorialBuyViewModel struct {
	Id                         int                                               `json:"id"`
	Name                       string                                            `json:"name"`
	Price                      int                                               `json:"price"`
	SubscriptionLiteViewModels []subscriptionViewModel.SubscriptionLiteViewModel `json:"subscriptionLiteViewModels"`
}
