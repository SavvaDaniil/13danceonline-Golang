package purchaseTutorialFacade

import (
	"danceonline/internal/entity"
	"danceonline/internal/repository/purchaseTutorialRepository"
	"danceonline/internal/repository/tutorialRepository"
	"danceonline/internal/repository/userRepository"
	"danceonline/internal/viewmodel/purchaseTutorialViewModel"
	"danceonline/internal/viewmodel/tutorialViewModel"
	"fmt"
)

func Add(id_of_user int, id_of_payment int, id_of_tutorial int) (*entity.PurchaseTutorial, error) {

	user, errFindUser := userRepository.FindById(id_of_user)
	if errFindUser != nil {
		fmt.Println("purchaseTutorialFacade add errFindUser:", errFindUser)
		return nil, errFindUser
	}

	tutorial, errFindTutorial := tutorialRepository.FindById(id_of_tutorial)
	if errFindTutorial != nil {
		fmt.Println("purchaseTutorialFacade add errFindTutorial:", errFindTutorial)
		return nil, errFindTutorial
	}

	purchaseTutorial, errorAddPT := purchaseTutorialRepository.Add(user.Id, id_of_payment, tutorial.Id, tutorial.Days)
	if errorAddPT != nil {
		fmt.Println("purchaseTutorialFacade add errorAddPT:", errorAddPT)
		return nil, errorAddPT
	}
	return purchaseTutorial, nil
}

func GetLastAvailableLiteViewModelByIdOfUserAndIdOfTutorial(id_of_user int, id_of_tutorial int) (*purchaseTutorialViewModel.PurchaseTutorialLiteViewModel, error) {

	tutorial, errFindTutorial := tutorialRepository.FindById(id_of_tutorial)
	if errFindTutorial != nil {
		return nil, errFindTutorial
	}

	purchaseTutorial, errFindPT := purchaseTutorialRepository.FindLastAvailableByIdOfUserAndIdOfTutorial(id_of_user, id_of_tutorial)
	if errFindPT != nil {
		fmt.Println(errFindPT)
		return nil, errFindPT
	}
	if purchaseTutorial == nil {
		return nil, nil
	}

	if purchaseTutorial.Date_of_activation == nil {
		errActivation := purchaseTutorialRepository.Activate(purchaseTutorial)
		if errActivation != nil {
			fmt.Println(errActivation)
			panic(errActivation)
		}
	}

	purchaseLiteViewModel := ToLiteViewModel(*tutorial, *purchaseTutorial)

	return purchaseLiteViewModel, nil
}

func ToLiteViewModel(tutorial entity.Tutorial, purchaseTutorial entity.PurchaseTutorial) *purchaseTutorialViewModel.PurchaseTutorialLiteViewModel {

	purchaseLiteViewModel := purchaseTutorialViewModel.PurchaseTutorialLiteViewModel{}

	tutorialMicroViewModel := tutorialViewModel.TutorialMicroViewModel{
		Id:   tutorial.Id,
		Name: tutorial.Name,
	}

	purchaseLiteViewModel = purchaseTutorialViewModel.PurchaseTutorialLiteViewModel{
		Id:                      purchaseTutorial.Id,
		Id_of_user:              purchaseTutorial.Id_of_user,
		Id_of_payment:           purchaseTutorial.Id_of_payment,
		TutorialMicroViewModel:  tutorialMicroViewModel,
		Days:                    purchaseTutorial.Days,
		Date_of_add:             *(purchaseTutorial.Date_of_add),
		Date_of_activation:      *(purchaseTutorial.Date_of_activation),
		Date_of_must_be_used_to: *(purchaseTutorial.Date_of_must_be_used_to),
	}
	return &purchaseLiteViewModel
}
