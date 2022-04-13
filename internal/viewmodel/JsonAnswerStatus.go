package viewmodel

import (
	"danceonline/internal/viewmodel/courseViewModel"
	"danceonline/internal/viewmodel/paymentViewModel"
	"danceonline/internal/viewmodel/purchaseCourseViewModel"
	"danceonline/internal/viewmodel/purchaseTutorialViewModel"
	"danceonline/internal/viewmodel/styleViewModel"
	"danceonline/internal/viewmodel/subscriptionViewModel"
	"danceonline/internal/viewmodel/tutorialViewModel"
	"danceonline/internal/viewmodel/userViewModel"
)

type JsonAnswerStatus struct {
	Status                        string                                                  `json:"status"`
	Errors                        string                                                  `json:"errors"`
	AccessToken                   string                                                  `json:"accessToken"`
	ForgetId                      int                                                     `json:"forget_id"`
	StyleMicroViewModels          []styleViewModel.StyleMicroViewModel                    `json:"styleMicroViewModels"`
	StyleLiteViewModels           []styleViewModel.StyleLiteViewModel                     `json:"styleLiteViewModels"`
	TutorialPreviewViewModels     []tutorialViewModel.TutorialPreviewViewModel            `json:"tutorialPreviewViewModels"`
	TutorialBuyViewModel          tutorialViewModel.TutorialBuyViewModel                  `json:"tutorialBuyViewModel"`
	TutorialLessonViewModel       tutorialViewModel.TutorialLessonViewModel               `json:"tutorialLessonViewModel"`
	CourseBuyViewModel            courseViewModel.CourseBuyViewModel                      `json:"courseBuyViewModel"`
	CoursePreviewViewModels       []courseViewModel.CoursePreviewViewModel                `json:"coursePreviewViewModels"`
	CourseLessonViewModel         courseViewModel.CourseLessonViewModel                   `json:"courseLessonViewModel"`
	UserProfileViewModel          userViewModel.UserProfileViewModel                      `json:"userProfileViewModel"`
	SubscriptionLiteViewModels    []subscriptionViewModel.SubscriptionLiteViewModel       `json:"subscriptionLiteViewModels"`
	PurchaseTutorialLiteViewModel purchaseTutorialViewModel.PurchaseTutorialLiteViewModel `json:"purchaseTutorialLiteViewModel"`
	PurchaseCourseLiteViewModel   purchaseCourseViewModel.PurchaseCourseLiteViewModel     `json:"purchaseCourseLiteViewModel"`
	PaymentMicroViewModel         paymentViewModel.PaymentMicroViewModel                  `json:"paymentMicroViewModel"`
}
