package userFacade

import (
	//"danceonline/internal/entity"
	"danceonline/internal/component"
	"danceonline/internal/repository/userRepository"
	"danceonline/internal/viewmodel"
	"danceonline/internal/viewmodel/userViewModel"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var hmacSampleSecret = []byte{'x' ...}

func GetProfile(id_of_user int) viewmodel.JsonAnswerStatus {

	user, err := userRepository.FindById(id_of_user)
	if err != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "user_not_found",
		}
	}

	return viewmodel.JsonAnswerStatus{
		UserProfileViewModel: userViewModel.UserProfileViewModel{
			Firstname: *(user.Firstname),
			Username:  user.Username,
		},
		Status: "success",
	}
}

func Update(id_of_user int, username string, passwordNew string, passwordCurrent string, firstname string) viewmodel.JsonAnswerStatus {

	if username == "" {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "username_cannot_be_null",
		}
	}

	user, err := userRepository.FindById(id_of_user)
	if err != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "user_not_found",
		}
	}

	if user.Username != username {
		userAlreadyExist, errAlreadyExist := userRepository.FindByUsernameExceptId(id_of_user, username)
		if errAlreadyExist != nil {
			return viewmodel.JsonAnswerStatus{
				Status: "error",
				Errors: "username_already_found_err",
			}
		}
		if userAlreadyExist != nil {
			return viewmodel.JsonAnswerStatus{
				Status: "error",
				Errors: "username_already_exist",
			}
		}
		//fmt.Println("username_already_exist not: ", userAlreadyExist)
	} else {
		username = ""
	}

	if passwordNew != "" {
		errPassword := bcrypt.CompareHashAndPassword([]byte(*(user.Password)), []byte(passwordCurrent))
		if errPassword != nil {
			return viewmodel.JsonAnswerStatus{
				Status: "error",
				Errors: "wrong",
			}
		}
	}

	errUpdate := userRepository.Update(id_of_user, username, passwordNew, firstname)
	if errUpdate != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "error_when_try_update",
		}
	}

	return viewmodel.JsonAnswerStatus{
		Status: "success",
	}
}

func Login(username string, password string) viewmodel.JsonAnswerStatus {

	if username == "" || password == "" {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "no_data",
		}
	}

	//fmt.Println("username: " + username)
	user, errFindByUsername := userRepository.FindByUsername(username)
	if errFindByUsername != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "try_find_by_username",
		}
	}
	if user == nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "wrong",
		}
	}


	errPassword := bcrypt.CompareHashAndPassword([]byte(*(user.Password)), []byte(password))
	if errPassword != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "wrong",
		}
	}

	accessToken, errJWT := generateJWT(user.Id)
	if errJWT != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "jwt_generate_wrong",
		}
	}

	return viewmodel.JsonAnswerStatus{
		Status:      "success",
		Errors:      "",
		AccessToken: accessToken,
	}
}

func Registration(username string, password string) viewmodel.JsonAnswerStatus {

	if username == "" || password == "" {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "no_data",
		}
	}

	userAlready, errFindByUsername := userRepository.FindByUsername(username)
	if errFindByUsername != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "when_try_to_find",
		}
	}
	if userAlready != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "username_already_exist",
		}
	}

	
	...
	...
	...

	accessToken, errJWT := generateJWT(user.Id)
	if errJWT != nil {
		return viewmodel.JsonAnswerStatus{
			Status: "error",
			Errors: "jwt_generate_wrong",
		}
	}

	return viewmodel.JsonAnswerStatus{
		Status:      "success",
		Errors:      "",
		AccessToken: accessToken,
	}
}

func Forget(step int, username string, code string, forget_id int) viewmodel.JsonAnswerStatus {
	...
	...
}

func generateJWT(id_of_user int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  strconv.Itoa(id_of_user),
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	accessToken, err := token.SignedString(hmacSampleSecret)

	//fmt.Println(accessToken, err)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func ParseJWT(r *http.Request) (int, error) {

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		return 0, fmt.Errorf("no bearer")
	}
	tokenString := splitToken[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed parse")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id_of_user, errIdOfUser := strconv.Atoi(claims["id"].(string))
		if errIdOfUser != nil {
			return 0, fmt.Errorf("failed get id_of_user")
		}

		return id_of_user, nil
	}

	return 0, fmt.Errorf("failed parse")
}
