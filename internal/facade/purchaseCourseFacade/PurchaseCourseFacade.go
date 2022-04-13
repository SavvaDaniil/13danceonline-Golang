package purchaseCourseFacade

import (
	"danceonline/internal/entity"
	//"danceonline/internal/repository/paymentRepository"
	"danceonline/internal/repository/courseRepository"
	"danceonline/internal/repository/purchaseCourseRepository"
	"danceonline/internal/repository/userRepository"
	"danceonline/internal/viewmodel/courseViewModel"
	"danceonline/internal/viewmodel/purchaseCourseViewModel"
	"fmt"
)

func Add(id_of_user int, id_of_payment int, id_of_course int) (*entity.PurchaseCourse, error) {

	user, errFindUser := userRepository.FindById(id_of_user)
	if errFindUser != nil {
		fmt.Println("purchaseCourseFacade add errFindUser:", errFindUser)
		return nil, errFindUser
	}

	course, errFindCourse := courseRepository.FindById(id_of_course)
	if errFindCourse != nil {
		fmt.Println("purchaseCourseFacade add errFindCourse:", errFindCourse)
		return nil, errFindCourse
	}

	purchaseCourse, errorAddPT := purchaseCourseRepository.Add(user.Id, id_of_payment, course.Id, course.Days)
	if errorAddPT != nil {
		fmt.Println("purchaseCourseFacade add errorAddPT:", errorAddPT)
		return nil, errorAddPT
	}
	return purchaseCourse, nil
}

func GetLastAvailableLiteViewModelByIdOfUserAndIdOfCourse(id_of_user int, id_of_course int) (*purchaseCourseViewModel.PurchaseCourseLiteViewModel, error) {

	course, errFindCourse := courseRepository.FindById(id_of_course)
	if errFindCourse != nil {
		return nil, errFindCourse
	}

	purchaseCourse, errFindPT := purchaseCourseRepository.FindLastAvailableByIdOfUserAndIdOfCourse(id_of_user, id_of_course)
	if errFindPT != nil {
		fmt.Println(errFindPT)
		return nil, errFindPT
	}
	if purchaseCourse == nil {
		return nil, nil
	}
	//fmt.Println("Founded purchaseCourse, Date_of_activation:", purchaseCourse.Date_of_activation)

	if purchaseCourse.Date_of_activation == nil {
		//fmt.Println("Try activate purchaseCourse")
		errActivation := purchaseCourseRepository.Activate(purchaseCourse)
		if errActivation != nil {
			fmt.Println(errActivation)
			panic(errActivation)
		}
	}

	purchaseLiteViewModel := ToLiteViewModel(*course, *purchaseCourse)

	return purchaseLiteViewModel, nil
}

func ToLiteViewModel(course entity.Course, purchaseCourse entity.PurchaseCourse) *purchaseCourseViewModel.PurchaseCourseLiteViewModel {

	purchaseLiteViewModel := purchaseCourseViewModel.PurchaseCourseLiteViewModel{}

	courseMicroViewModel := courseViewModel.CourseMicroViewModel{
		Id:   course.Id,
		Name: course.Name,
	}

	purchaseLiteViewModel = purchaseCourseViewModel.PurchaseCourseLiteViewModel{
		Id:                      purchaseCourse.Id,
		Id_of_user:              purchaseCourse.Id_of_user,
		Id_of_payment:           purchaseCourse.Id_of_payment,
		CourseMicroViewModel:    courseMicroViewModel,
		Days:                    purchaseCourse.Days,
		Date_of_add:             *(purchaseCourse.Date_of_add),
		Date_of_activation:      *(purchaseCourse.Date_of_activation),
		Date_of_must_be_used_to: *(purchaseCourse.Date_of_must_be_used_to),
	}
	return &purchaseLiteViewModel
}
