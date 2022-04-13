package purchaseTutorialRepository

import (
	"danceonline/config"
	"danceonline/internal/entity"
	"database/sql"
	"fmt"
	"time"
)

func FindById(id_of_purchase_tutorial int) (*entity.PurchaseTutorial, error) {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	var pt entity.PurchaseTutorial = entity.PurchaseTutorial{}
	errScan := database.QueryRow("SELECT id, id_of_user, id_of_payment, id_of_tutorial, days, active, date_of_add, date_of_activation, date_of_must_be_used_to FROM XXXXXXXXXXXXXXXXXXXXX.purchase_tutorial WHERE id_of_user = ? ORDER BY id DESC", id_of_purchase_tutorial).Scan(&pt.Id, &pt.Id_of_user, &pt.Id_of_payment, &pt.Id_of_tutorial, &pt.Days, &pt.Active, &pt.Date_of_add, &pt.Date_of_activation, &pt.Date_of_must_be_used_to)
	if errScan != nil {
		return nil, errScan
	}

	return &pt, nil

}

func Add(id_of_user int, id_of_payment int, id_of_tutorial int, days int) (*entity.PurchaseTutorial, error) {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	var id_of_purchase_tutorial int64 = 0
	date_of_add := time.Now().Format("2006-01-02 15:04:05")
	resInsert, errInsert := database.Exec("INSERT INTO XXXXXXXXXXXXXXXXXXXXX.purchase_tutorial (id_of_user, id_of_payment, id_of_tutorial, days, active, date_of_add) VALUES (?, ?, ?, ?, ?, ?)", id_of_user, id_of_payment, id_of_tutorial, days, 1, date_of_add)
	if errInsert != nil {
		//panic(errInsert.Error())
		return nil, errInsert
	}

	id_of_purchase_tutorial, errLastInsert := resInsert.LastInsertId()
	if errLastInsert != nil {
		return nil, errLastInsert
	}

	var pt entity.PurchaseTutorial = entity.PurchaseTutorial{}
	errScan := database.QueryRow("SELECT id, id_of_user, id_of_payment, id_of_tutorial, days, active, date_of_add, date_of_activation, date_of_must_be_used_to FROM XXXXXXXXXXXXXXXXXXXXX.purchase_tutorial WHERE id = ? ORDER BY id DESC", id_of_purchase_tutorial).Scan(&pt.Id, &pt.Id_of_user, &pt.Id_of_payment, &pt.Id_of_tutorial, &pt.Days, &pt.Active, &pt.Date_of_add, &pt.Date_of_activation, &pt.Date_of_must_be_used_to)
	if errScan != nil {
		return nil, errScan
	}

	return &pt, nil
}

func ListAllByIdOfUser(id_of_user int) []entity.PurchaseTutorial {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		panic("Error database connection")
	}
	defer database.Close()

	rows, err := database.Query("SELECT id, id_of_user, id_of_payment, id_of_tutorial, days, active, date_of_add, date_of_activation, date_of_must_be_used_to FROM XXXXXXXXXXXXXXXXXXXXX.purchase_tutorial WHERE id_of_user = ? AND active = '1' ORDER BY id DESC", id_of_user)
	if err != nil {
		panic("Error query")
	}
	defer rows.Close()

	var purchaseTutorials []entity.PurchaseTutorial

	for rows.Next() {
		pt := entity.PurchaseTutorial{}
		err := rows.Scan(&pt.Id, &pt.Id_of_user, &pt.Id_of_payment, &pt.Id_of_tutorial, &pt.Days, &pt.Active, &pt.Date_of_add, &pt.Date_of_activation, &pt.Date_of_must_be_used_to)
		if err != nil {
			fmt.Println(err)
		}
		purchaseTutorials = append(purchaseTutorials, pt)
	}

	return purchaseTutorials
}

func ListAllActiveByIdOfUser(id_of_user int) []entity.PurchaseTutorial {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		panic("Error database connection")
	}
	defer database.Close()

	dateNow := time.Now().Format("2006-01-02 15:04:05")
	rows, err := database.Query("SELECT id, id_of_user, id_of_payment, id_of_tutorial, days, active, date_of_add, date_of_activation, date_of_must_be_used_to FROM XXXXXXXXXXXXXXXXXXXXX.purchase_tutorial WHERE id_of_user = ? AND active = '1' AND (date_of_activation IS NULL OR date_of_must_be_used_to > ? OR date_of_must_be_used_to = ?) ORDER BY id DESC", id_of_user, dateNow, dateNow)
	if err != nil {
		panic("Error query")
	}
	defer rows.Close()

	var purchaseTutorials []entity.PurchaseTutorial

	for rows.Next() {
		pt := entity.PurchaseTutorial{}
		err := rows.Scan(&pt.Id, &pt.Id_of_user, &pt.Id_of_payment, &pt.Id_of_tutorial, &pt.Days, &pt.Active, &pt.Date_of_add, &pt.Date_of_activation, &pt.Date_of_must_be_used_to)
		if err != nil {
			fmt.Println(err)
		}
		purchaseTutorials = append(purchaseTutorials, pt)
	}

	return purchaseTutorials
}

func FindLastAvailableByIdOfUserAndIdOfTutorial(id_of_user int, id_of_tutorial int) (*entity.PurchaseTutorial, error) {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		//panic("Error database connection")
		return nil, err
	}
	defer database.Close()

	var pt entity.PurchaseTutorial = entity.PurchaseTutorial{}
	dateNow := time.Now().Format("2006-01-02 15:04:05")
	errScan := database.QueryRow("SELECT id, id_of_user, id_of_payment, id_of_tutorial, days, active, date_of_add, date_of_activation, date_of_must_be_used_to FROM XXXXXXXXXXXXXXXXXXXXX.purchase_tutorial WHERE id_of_user = ? AND active = '1' AND id_of_tutorial = ? AND (date_of_activation IS NULL OR date_of_must_be_used_to > ? OR date_of_must_be_used_to = ?) ORDER BY id DESC", id_of_user, id_of_tutorial, dateNow, dateNow).Scan(&pt.Id, &pt.Id_of_user, &pt.Id_of_payment, &pt.Id_of_tutorial, &pt.Days, &pt.Active, &pt.Date_of_add, &pt.Date_of_activation, &pt.Date_of_must_be_used_to)
	if errScan != nil {
		if errScan == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, errScan
		}
	}

	return &pt, nil
}

func Activate(purchaseTutorial *entity.PurchaseTutorial) error {
	...
}
