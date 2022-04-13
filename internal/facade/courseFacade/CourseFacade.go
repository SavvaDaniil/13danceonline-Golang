package courseFacade

import (
	"danceonline/internal/entity"
	"danceonline/internal/facade/purchaseCourseFacade"
	"danceonline/internal/facade/videoFacade"
	"danceonline/internal/repository/courseRepository"
	"danceonline/internal/repository/purchaseCourseRepository"
	"danceonline/internal/repository/userRepository"
	"danceonline/internal/viewmodel"
	"danceonline/internal/viewmodel/courseViewModel"
	"fmt"
	"os"
	"strconv"
)

func CheckAccess(id_of_user int, id_of_course int) viewmodel.JsonAnswerStatus {

	course, errFindCourse := courseRepository.FindById(id_of_course)
	if errFindCourse != nil {
		fmt.Println(errFindCourse)
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "when_try_find_course",
		}
	}
	if course == nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "course_not_found",
		}
	}

	purchaseCourseLiteViewModel, errFindPT := purchaseCourseFacade.GetLastAvailableLiteViewModelByIdOfUserAndIdOfCourse(id_of_user, id_of_course)
	if errFindPT != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "when_try_find_purchase_course",
		}
	}
	if purchaseCourseLiteViewModel == nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "no_access",
		}
	}

	return viewmodel.JsonAnswerStatus{
		Status:                      "success",
		PurchaseCourseLiteViewModel: *purchaseCourseLiteViewModel,
	}
}

func GetLesson(id_of_user int, id_of_course int, number_of_lesson int) viewmodel.JsonAnswerStatus {

	course, errFindCourse := courseRepository.FindById(id_of_course)
	if errFindCourse != nil {
		fmt.Println(errFindCourse)
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "when_try_find_course",
		}
	}
	if course == nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "course_not_found",
		}
	}

	purchaseCourseLiteViewModel, errFindPT := purchaseCourseFacade.GetLastAvailableLiteViewModelByIdOfUserAndIdOfCourse(id_of_user, id_of_course)
	if errFindPT != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "when_try_find_purchaseCourse",
		}
	}
	if purchaseCourseLiteViewModel == nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "no_access",
		}
	}

	videoCourseLessonViewModel, errVLVM := videoFacade.GetVideoCourseLessonViewModel(id_of_course, number_of_lesson)
	if errVLVM != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "when_try_find_video",
		}
	}

	if videoCourseLessonViewModel == nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "not_find",
		}
	}

	courseLessonViewModel := courseViewModel.CourseLessonViewModel{
		Id:                         course.Id,
		Name:                       course.Name,
		PosterSrc:                  GetPosterSrc(course.Id),
		VideoCourseLessonViewModel: *videoCourseLessonViewModel,
	}

	return viewmodel.JsonAnswerStatus{
		Status:                "success",
		CourseLessonViewModel: courseLessonViewModel,
	}
}

func BuyPrepare(id_of_course int) viewmodel.JsonAnswerStatus {

	course, errFind := courseRepository.FindById(id_of_course)
	if errFind != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "not_found",
		}
	}

	return viewmodel.JsonAnswerStatus{
		Status: "success",
		CourseBuyViewModel: courseViewModel.CourseBuyViewModel{
			Id:    course.Id,
			Name:  course.Name,
			Price: course.Price,
			//SubscriptionLiteViewModels: subscriptionLiteViewModels,
		},
	}
}

func ListAllActiveByFilter(id_of_user int, page int) []courseViewModel.CoursePreviewViewModel {

	var purchaseCourses []entity.PurchaseCourse
	var isPurchaseSubscription bool = false
	if id_of_user != 0 {
		user, errFindUser := userRepository.FindById(id_of_user)
		if errFindUser != nil {
			return nil
		}
		if user != nil {
			purchaseCourses = purchaseCourseRepository.ListAllActiveByIdOfUser(id_of_user)

		}
	}

	var coursePreviewViewModels []courseViewModel.CoursePreviewViewModel
	courses, errSearchcourses := courseRepository.ListAllActiveByFilter(page)
	if errSearchcourses != nil {
		return nil
	}

	var isActive int = 0
	for _, course := range courses {

		isActive = 0
		if isPurchaseSubscription {
			isActive = 1
		} else if len(purchaseCourses) > 0 {
			for _, purchaseCourse := range purchaseCourses {
				//fmt.Println("Check purchaseCourse.id_of_course:", purchaseCourse.id_of_course)
				if course.Id == purchaseCourse.Id_of_course {
					isActive = 1
					break
				}
			}
		}

		var coursePreviewViewModel courseViewModel.CoursePreviewViewModel = courseViewModel.CoursePreviewViewModel{
			Id:        course.Id,
			Name:      course.Name,
			PosterSrc: GetPosterSrc(course.Id),
			Price:     course.Price,
			IsActive:  isActive,
		}
		coursePreviewViewModels = append(coursePreviewViewModels, coursePreviewViewModel)
	}
	return coursePreviewViewModels
}

func GetPosterSrc(id_of_course int) string {
	infoOfExistFile, err := os.Stat("./static/uploads/course/" + strconv.Itoa(id_of_course) + ".jpg")
	if os.IsNotExist(err) || infoOfExistFile.IsDir() {
		return ""
	}
	return "static/uploads/course/" + strconv.Itoa(id_of_course) + ".jpg"
}
