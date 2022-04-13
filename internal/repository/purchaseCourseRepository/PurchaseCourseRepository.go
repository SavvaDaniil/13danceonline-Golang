package purchaseCourseRepository

import (
	"danceonline/config"
	"danceonline/internal/entity"
	"database/sql"
	"fmt"
	"time"
)

func FindById(id_of_purchase_course int) (*entity.PurchaseCourse, error) {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	var pt entity.PurchaseCourse = entity.PurchaseCourse{}
	errScan := database.QueryRow("SELECT id, id_of_user, id_of_payment, id_of_course, days, active, date_of_add, date_of_activation, date_of_must_be_used_to FROM XXXXXXXXXXXXXXXXXXXXX.purchase_course WHERE id_of_user = ? ORDER BY id DESC", id_of_purchase_course).Scan(&pt.Id, &pt.Id_of_user, &pt.Id_of_payment, &pt.Id_of_course, &pt.Days, &pt.Active, &pt.Date_of_add, &pt.Date_of_activation, &pt.Date_of_must_be_used_to)
	if errScan != nil {
		return nil, errScan
	}

	return &pt, nil

}

func Add(id_of_user int, id_of_payment int, id_of_course int, days int) (*entity.PurchaseCourse, error) {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	var id_of_purchase_course int64 = 0
	date_of_add := time.Now().Format("2006-01-02 15:04:05")
	resInsert, errInsert := database.Exec("INSERT INTO XXXXXXXXXXXXXXXXXXXXX.purchase_course (id_of_user, id_of_payment, id_of_course, days, active, date_of_add) VALUES (?, ?, ?, ?, ?, ?)", id_of_user, id_of_payment, id_of_course, days, 1, date_of_add)
	if errInsert != nil {
		//panic(errInsert.Error())
		return nil, errInsert
	}

	id_of_purchase_course, errLastInsert := resInsert.LastInsertId()
	if errLastInsert != nil {
		return nil, errLastInsert
	}

	var pt entity.PurchaseCourse = entity.PurchaseCourse{}
	errScan := database.QueryRow("SELECT id, id_of_user, id_of_payment, id_of_course, days, active, date_of_add, date_of_activation, date_of_must_be_used_to FROM XXXXXXXXXXXXXXXXXXXXX.purchase_course WHERE id = ? ORDER BY id DESC", id_of_purchase_course).Scan(&pt.Id, &pt.Id_of_user, &pt.Id_of_payment, &pt.Id_of_course, &pt.Days, &pt.Active, &pt.Date_of_add, &pt.Date_of_activation, &pt.Date_of_must_be_used_to)
	if errScan != nil {
		return nil, errScan
	}

	return &pt, nil
}

func ListAllByIdOfUser(id_of_user int) []entity.PurchaseCourse {
	...
}

func ListAllActiveByIdOfUser(id_of_user int) []entity.PurchaseCourse {
	...
}

func FindLastAvailableByIdOfUserAndIdOfCourse(id_of_user int, id_of_course int) (*entity.PurchaseCourse, error) {
	...
}

func Activate(purchaseCourse *entity.PurchaseCourse) error {
	...

}
