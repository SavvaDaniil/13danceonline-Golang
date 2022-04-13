package tutorialDTO

type TutorialIdDTO struct {
	Id_of_tutorial int `json:"id_of_tutorial"`
}

type TutorialSearchDTO struct {
	Page        int `json:"page"`
	Id_of_style int `json:"id_of_style"`
}
