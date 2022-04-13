package purchaseTutorialViewModel

import (
	"danceonline/internal/viewmodel/tutorialViewModel"
)

type PurchaseTutorialLiteViewModel struct {
	Id                      int                                      `json:"id"`
	Id_of_user              int                                      `json:"id_of_user"`
	Id_of_payment           *int                                     `json:"id_of_payment"`
	TutorialMicroViewModel  tutorialViewModel.TutorialMicroViewModel `json:"tutorialMicroViewModel"`
	Days                    int                                      `json:"days"`
	Date_of_add             string                                   `json:"date_of_add"`
	Date_of_activation      string                                   `json:"date_of_activation"`
	Date_of_must_be_used_to string                                   `json:"date_of_must_be_used_to"`
}
