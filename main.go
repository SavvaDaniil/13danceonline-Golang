package main

import (
	"danceonline/internal/facade/courseFacade"
	"danceonline/internal/facade/paymentFacade"
	"danceonline/internal/facade/styleFacade"
	"danceonline/internal/facade/subscriptionFacade"
	"danceonline/internal/facade/tutorialFacade"
	"danceonline/internal/facade/userFacade"
	"danceonline/internal/viewmodel"
	//"danceonline/internal/viewmodel/styleViewModel"
	//"danceonline/internal/viewmodel/tutorialViewModel"
	"danceonline/internal/DTO/courseDTO"
	"danceonline/internal/DTO/paymentDTO"
	"danceonline/internal/DTO/tutorialDTO"
	"danceonline/internal/DTO/userDTO"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func resNotAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json")
	jsonAnswerStatus := viewmodel.JsonAnswerStatus{
		Status: "error",
		Errors: "not_auth",
	}
	jsonResp, err := json.Marshal(jsonAnswerStatus)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func defaultJsonRequestProperties(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		} else {
			http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
			return
		}
	}
}

func main() {

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	var orig = http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	var wrapped = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		//w.WriteHeader(http.StatusOK)
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		orig.ServeHTTP(w, r)
		//w.Header().Set("Access-Control-Allow-Credentials", "true")
		//w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
		//w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		//w.Header().Set("Content-Type", "application/x-mpegURL")
	})
	http.Handle("/static/", wrapped)

	var origVideo = http.StripPrefix("/static/uploads/video", http.FileServer(http.Dir("./static/uploads/video")))
	var wrappedVideo = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/x-mpegURL")
		w.WriteHeader(http.StatusOK)
		origVideo.ServeHTTP(w, r)
	})
	http.Handle("/static/uploads/video", wrappedVideo)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//for k, v := range r.Header {
		//fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
		//}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Status Created"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/api/style/list/micro", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		var jsonAnswerStatus viewmodel.JsonAnswerStatus = styleFacade.JsonListAllActiveMicro()

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/api/tutorial/search", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		var tutorialSearchDTO tutorialDTO.TutorialSearchDTO
		err := json.NewDecoder(r.Body).Decode(&tutorialSearchDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		id_of_user, _ := userFacade.ParseJWT(r)

		var jsonAnswerStatus viewmodel.JsonAnswerStatus = viewmodel.JsonAnswerStatus{
			Status:                    "success",
			Errors:                    "",
			TutorialPreviewViewModels: tutorialFacade.ListAllActiveByFilter(id_of_user, tutorialSearchDTO.Page, tutorialSearchDTO.Id_of_style),
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	//MARK: UserController

	http.HandleFunc("/api/user/registration", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		username := r.Form.Get("username")
		password := r.Form.Get("password")

		var jsonAnswerStatus viewmodel.JsonAnswerStatus = userFacade.Registration(username, password)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/api/user/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		var userLoginDTO userDTO.UserLoginDTO
		err := json.NewDecoder(r.Body).Decode(&userLoginDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		jsonAnswerStatus := userFacade.Login(userLoginDTO.Username, userLoginDTO.Password)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/api/user/forget", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		var userForgetDTO userDTO.UserForgetDTO
		err := json.NewDecoder(r.Body).Decode(&userForgetDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		var jsonAnswerStatus viewmodel.JsonAnswerStatus = userFacade.Forget(userForgetDTO.Step, userForgetDTO.Username, userForgetDTO.Code, userForgetDTO.ForgetId)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/api/user/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(nil)
			return
		}

		id_of_user, errAuth := userFacade.ParseJWT(r)
		if errAuth != nil {
			resNotAuth(w, r)
			return
		}
		r.ParseForm()

		jsonAnswerStatus := userFacade.GetProfile(id_of_user)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/api/user/update", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(nil)
			return
		}

		id_of_user, errAuth := userFacade.ParseJWT(r)
		if errAuth != nil {
			resNotAuth(w, r)
			return
		}
		r.ParseForm()

		var userProfileDTO userDTO.UserProfileDTO
		err := json.NewDecoder(r.Body).Decode(&userProfileDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		jsonAnswerStatus := userFacade.Update(id_of_user, userProfileDTO.Username, userProfileDTO.PasswordNew, userProfileDTO.PasswordCurrent, userProfileDTO.Firstname)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/api/user/secret", func(w http.ResponseWriter, r *http.Request) {

		id_of_user, errAuth := userFacade.ParseJWT(r)
		if errAuth != nil {
			resNotAuth(w, r)
			return
		}
		fmt.Println("id_of_user: ", id_of_user)

		jsonAnswerStatus := viewmodel.JsonAnswerStatus{
			Status: "success",
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	//MARK: SubscriptionController
	http.HandleFunc("/api/subscription/list_all_active", func(w http.ResponseWriter, r *http.Request) {

		var jsonAnswerStatus viewmodel.JsonAnswerStatus
		subscriptionLiteViewModel, err := subscriptionFacade.ListAllLiteActive()
		if err != nil {
			jsonAnswerStatus = viewmodel.JsonAnswerStatus{
				Status: "error",
				Errors: "db_connection",
			}
		} else {
			jsonAnswerStatus = viewmodel.JsonAnswerStatus{
				Status:                     "success",
				SubscriptionLiteViewModels: subscriptionLiteViewModel,
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	//MARK: TutorialController
	http.HandleFunc("/api/tutorial/buy/prepare", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		var tutorialIdDTO tutorialDTO.TutorialIdDTO
		err := json.NewDecoder(r.Body).Decode(&tutorialIdDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		var jsonAnswerStatus viewmodel.JsonAnswerStatus = tutorialFacade.BuyPrepare(tutorialIdDTO.Id_of_tutorial)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/api/tutorial/check_access", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		var tutorialIdDTO tutorialDTO.TutorialIdDTO
		err := json.NewDecoder(r.Body).Decode(&tutorialIdDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		var jsonAnswerStatus viewmodel.JsonAnswerStatus
		id_of_user, errAuth := userFacade.ParseJWT(r)
		if errAuth != nil {
			jsonAnswerStatus = viewmodel.JsonAnswerStatus{
				Status: "error",
				Errors: "not_auth",
			}
		} else {
			jsonAnswerStatus = tutorialFacade.CheckAccess(id_of_user, tutorialIdDTO.Id_of_tutorial)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/api/tutorial/lesson/get", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		var tutorialIdDTO tutorialDTO.TutorialIdDTO
		err := json.NewDecoder(r.Body).Decode(&tutorialIdDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		var jsonAnswerStatus viewmodel.JsonAnswerStatus
		id_of_user, errAuth := userFacade.ParseJWT(r)
		if errAuth != nil {
			jsonAnswerStatus = viewmodel.JsonAnswerStatus{
				Status: "error",
				Errors: "not_auth",
			}
		} else {
			jsonAnswerStatus = tutorialFacade.GetLesson(id_of_user, tutorialIdDTO.Id_of_tutorial)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	//MARK: CourseController

	http.HandleFunc("/api/course/search", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		var courseSearchDTO courseDTO.CourseSearchDTO
		err := json.NewDecoder(r.Body).Decode(&courseSearchDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		id_of_user, _ := userFacade.ParseJWT(r)

		var jsonAnswerStatus viewmodel.JsonAnswerStatus = viewmodel.JsonAnswerStatus{
			Status:                  "success",
			Errors:                  "",
			CoursePreviewViewModels: courseFacade.ListAllActiveByFilter(id_of_user, courseSearchDTO.Page),
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/api/course/check_access", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		var courseIdDTO courseDTO.CourseIdDTO
		err := json.NewDecoder(r.Body).Decode(&courseIdDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		var jsonAnswerStatus viewmodel.JsonAnswerStatus
		id_of_user, errAuth := userFacade.ParseJWT(r)
		if errAuth != nil {
			jsonAnswerStatus = viewmodel.JsonAnswerStatus{
				Status: "error",
				Errors: "not_auth",
			}
		} else {
			jsonAnswerStatus = courseFacade.CheckAccess(id_of_user, courseIdDTO.Id_of_course)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/api/course/buy/prepare", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		var courseIdDTO courseDTO.CourseIdDTO
		err := json.NewDecoder(r.Body).Decode(&courseIdDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		var jsonAnswerStatus viewmodel.JsonAnswerStatus = courseFacade.BuyPrepare(courseIdDTO.Id_of_course)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/api/course/lesson/get", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		var courseLessonDTO courseDTO.CourseLessonDTO
		err := json.NewDecoder(r.Body).Decode(&courseLessonDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		var jsonAnswerStatus viewmodel.JsonAnswerStatus
		id_of_user, errAuth := userFacade.ParseJWT(r)
		if errAuth != nil {
			jsonAnswerStatus = viewmodel.JsonAnswerStatus{
				Status: "error",
				Errors: "not_auth",
			}
		} else {
			jsonAnswerStatus = courseFacade.GetLesson(id_of_user, courseLessonDTO.Id_of_course, courseLessonDTO.Number_of_lesson)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	//MARK: PaymentController
	http.HandleFunc("/api/payment/init", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		var jsonAnswerStatus viewmodel.JsonAnswerStatus
		id_of_user, errAuth := userFacade.ParseJWT(r)
		if errAuth != nil {
			jsonAnswerStatus = viewmodel.JsonAnswerStatus{
				Status: "error",
				Errors: "not_auth",
			}
		} else {
			var paymentInitDTO paymentDTO.PaymentInitDTO
			err := json.NewDecoder(r.Body).Decode(&paymentInitDTO)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				fmt.Println(err)
				return
			}

			jsonAnswerStatus = paymentFacade.GenerateOrGetLastNotPayed(id_of_user, paymentInitDTO.Id_of_tutorial, paymentInitDTO.Id_of_subscription, paymentInitDTO.Id_of_course)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/api/payment/result", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		var paymentResultRobokassaDTO paymentDTO.PaymentResultRobokassaDTO
		err := json.NewDecoder(r.Body).Decode(&paymentResultRobokassaDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		jsonAnswerStatus := paymentFacade.ResultPayment(paymentResultRobokassaDTO.Id_of_payment)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(jsonAnswerStatus)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)

}
