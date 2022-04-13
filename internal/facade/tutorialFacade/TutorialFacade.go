package tutorialFacade

import (
	"danceonline/internal/entity"
	"danceonline/internal/facade/purchaseTutorialFacade"
	"danceonline/internal/facade/subscriptionFacade"
	"danceonline/internal/facade/videoFacade"
	"danceonline/internal/repository/purchaseSubscriptionRepository"
	"danceonline/internal/repository/purchaseTutorialRepository"
	"danceonline/internal/repository/tutorialRepository"
	"danceonline/internal/repository/userRepository"
	"danceonline/internal/viewmodel"
	"danceonline/internal/viewmodel/tutorialViewModel"
	"fmt"
	"os"
	"strconv"
)

func CheckAccess(id_of_user int, id_of_tutorial int) viewmodel.JsonAnswerStatus {

	tutorial, errFindTutorial := tutorialRepository.FindById(id_of_tutorial)
	if errFindTutorial != nil {
		fmt.Println(errFindTutorial)
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "when_try_find_tutorial",
		}
	}
	if tutorial == nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "tutorial_not_found",
		}
	}

	purchaseTutorialLiteViewModel, errFindPT := purchaseTutorialFacade.GetLastAvailableLiteViewModelByIdOfUserAndIdOfTutorial(id_of_user, id_of_tutorial)
	if errFindPT != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "when_try_find_purchaseTutorial",
		}
	}
	if purchaseTutorialLiteViewModel == nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "no_access",
		}
	}

	return viewmodel.JsonAnswerStatus{
		Status:                        "success",
		PurchaseTutorialLiteViewModel: *purchaseTutorialLiteViewModel,
	}
}

func GetLesson(id_of_user int, id_of_tutorial int) viewmodel.JsonAnswerStatus {

	tutorial, errFindTutorial := tutorialRepository.FindById(id_of_tutorial)
	if errFindTutorial != nil {
		fmt.Println(errFindTutorial)
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "when_try_find_tutorial",
		}
	}
	if tutorial == nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "tutorial_not_found",
		}
	}

	purchaseTutorialLiteViewModel, errFindPT := purchaseTutorialFacade.GetLastAvailableLiteViewModelByIdOfUserAndIdOfTutorial(id_of_user, id_of_tutorial)
	if errFindPT != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "when_try_find_purchaseTutorial",
		}
	}
	if purchaseTutorialLiteViewModel == nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "no_access",
		}
	}

	videoTutorialLessonViewModel, errVLVM := videoFacade.GetVideoTutorialLessonViewModel(id_of_tutorial)
	if errVLVM != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "when_try_find_video",
		}
	}

	if videoTutorialLessonViewModel == nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "not_find",
		}
	}

	tutorialLessonViewModel := tutorialViewModel.TutorialLessonViewModel{
		Id:                           tutorial.Id,
		Name:                         tutorial.Name,
		PosterSrc:                    GetPosterSrc(tutorial.Id),
		VideoTutorialLessonViewModel: *videoTutorialLessonViewModel,
	}
	return viewmodel.JsonAnswerStatus{
		Status:                  "success",
		TutorialLessonViewModel: tutorialLessonViewModel,
	}
}

func BuyPrepare(id_of_tutorial int) viewmodel.JsonAnswerStatus {

	tutorial, errFind := tutorialRepository.FindById(id_of_tutorial)
	if errFind != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "not_found",
		}
	}

	//check, is access already available?

	subscriptionLiteViewModels, errListSubsriptions := subscriptionFacade.ListAllLiteActive()
	if errListSubsriptions != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "when_try_find_subscriptions",
		}
	}

	return viewmodel.JsonAnswerStatus{
		Status: "success",
		TutorialBuyViewModel: tutorialViewModel.TutorialBuyViewModel{
			Id:                         tutorial.Id,
			Name:                       tutorial.Name,
			Price:                      tutorial.Price,
			SubscriptionLiteViewModels: subscriptionLiteViewModels,
		},
	}
}

func ListAllActiveByFilter(id_of_user int, page int, id_of_style int) []tutorialViewModel.TutorialPreviewViewModel {

	//page, err := strconv.Atoi(pageStr)
	//page = 0

	var purchaseTutorials []entity.PurchaseTutorial
	var isPurchaseSubscription bool = false
	if id_of_user != 0 {
		user, errFindUser := userRepository.FindById(id_of_user)
		if errFindUser != nil {
			return nil
		}
		if user != nil {
			purchaseTutorials = purchaseTutorialRepository.ListAllActiveByIdOfUser(id_of_user)
			purchaseSubscription, errFindPurchaseSubscription := purchaseSubscriptionRepository.FindLastAvailableByIdOfUser(id_of_user)
			if errFindPurchaseSubscription != nil {
				return nil
			}
			if purchaseSubscription != nil {
				isPurchaseSubscription = true
			}
		}
	}

	var tutorialPreviewViewModels []tutorialViewModel.TutorialPreviewViewModel
	tutorials, errSearchTutorials := tutorialRepository.ListAllActiveByFilter(page, id_of_style)
	if errSearchTutorials != nil {
		return nil
	}

	var isActive int = 0
	for _, tutorial := range tutorials {

		isActive = 0
		if isPurchaseSubscription {
			isActive = 1
		} else if len(purchaseTutorials) > 0 {
			for _, purchaseTutorial := range purchaseTutorials {
				//fmt.Println("Check purchaseTutorial.Id_of_tutorial:", purchaseTutorial.Id_of_tutorial)
				if tutorial.Id == purchaseTutorial.Id_of_tutorial {
					isActive = 1
					break
				}
			}
		}

		var tutorialPreviewViewModel tutorialViewModel.TutorialPreviewViewModel = tutorialViewModel.TutorialPreviewViewModel{
			Id:        tutorial.Id,
			Name:      tutorial.Name,
			PosterSrc: GetPosterSrc(tutorial.Id),
			Price:     tutorial.Price,
			IsActive:  isActive,
		}
		tutorialPreviewViewModels = append(tutorialPreviewViewModels, tutorialPreviewViewModel)
	}
	return tutorialPreviewViewModels
}

func GetPosterSrc(id_of_teacher int) string {
	infoOfExistFile, err := os.Stat("./static/uploads/tutorial/" + strconv.Itoa(id_of_teacher) + ".jpg")
	if os.IsNotExist(err) || infoOfExistFile.IsDir() {
		return ""
	}
	return "static/uploads/tutorial/" + strconv.Itoa(id_of_teacher) + ".jpg"
}
