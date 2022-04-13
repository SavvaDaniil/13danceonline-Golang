package paymentFacade

import (
	"danceonline/internal/entity"
	"danceonline/internal/facade/purchaseCourseFacade"
	"danceonline/internal/facade/purchaseTutorialFacade"
	"danceonline/internal/model/paymentModel"
	"danceonline/internal/repository/courseRepository"
	"danceonline/internal/repository/paymentRepository"
	"danceonline/internal/repository/subscriptionRepository"
	"danceonline/internal/repository/tutorialRepository"
	"danceonline/internal/repository/userRepository"
	"danceonline/internal/viewmodel"
	"danceonline/internal/viewmodel/paymentViewModel"
	"encoding/gob"
	"fmt"
	"os"
	"strconv"
	"time"
)

func GenerateOrGetLastNotPayed(id_of_user int, id_of_tutorial int, id_of_subscription int, id_of_course int) viewmodel.JsonAnswerStatus {

	...
}

func ResultPayment(id_of_payment int) viewmodel.JsonAnswerStatus {

	...

}

func serializeToFile(payment entity.Payment, id_of_tutorial int, id_of_subscription int, id_of_course int) error {
	paymentPrepare := paymentModel.PaymentPrepare{
		Id:                 payment.Id,
		Id_of_user:         payment.Id_of_user,
		Price:              payment.Price,
		Status:             0,
		Id_of_tutorial:     id_of_tutorial,
		Id_of_subscription: id_of_subscription,
		Id_of_course:       id_of_course,
	}
	var dirPath string = "./XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	var filePath string = dirPath + "/" + strconv.Itoa(payment.Id) + ".gob"
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		errCreateDir := os.Mkdir(dirPath, 0775)
		if errCreateDir != nil {
			fmt.Println("PaymentFacade serializeToFile fail create dir error: ", err)
			return err
		}
	}

	err := writeGob(filePath, paymentPrepare)
	if err != nil {
		fmt.Println("PaymentFacade serializeToFile error: ", err)
		return err
	}
	return nil
}

func ToMicroViewModel(payment entity.Payment) *paymentViewModel.PaymentMicroViewModel {
	paymentMicroViewModel := paymentViewModel.PaymentMicroViewModel{
		Id:         payment.Id,
		Id_of_user: payment.Id_of_user,
		Price:      payment.Price,
	}
	return &paymentMicroViewModel
}

func writeGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

//func readGob(filePath string, object interface{}) error {
func readGob(filePath string) (*paymentModel.PaymentPrepare, error) {
	var paymentPrepare *paymentModel.PaymentPrepare
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("readGob read file", err)
		file.Close()
		return nil, err
	}

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&paymentPrepare)
	if err != nil {
		fmt.Println("readGob", err)
		file.Close()
		return nil, err
	}

	file.Close()
	return paymentPrepare, nil
}
