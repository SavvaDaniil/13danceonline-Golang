package paymentRepository

import (
	"danceonline/config"
	"danceonline/internal/entity"
	"database/sql"
	"fmt"
	"time"
)

func Add(user entity.User, price int) (*entity.Payment, error) {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		//panic("Error database connection")
		return nil, err
	}
	defer database.Close()

	var id_of_payment int64 = 0
	var date_of_add = time.Now()
	resInsert, errInsert := database.Exec("INSERT INTO XXXXXXXXXXXXXXXXXXXXX.payment (id_of_user, price, status, date_of_add, date_of_update) VALUES (?, ?, ?, ?, ?)", user.Id, price, 0, date_of_add, date_of_add)
	if errInsert != nil {
		panic(errInsert.Error())
		//return nil, errInsert
	}

	id_of_payment, errLastInsert := resInsert.LastInsertId()
	if errLastInsert != nil {
		fmt.Println(errLastInsert)
		return nil, errLastInsert
	}

	var payment entity.Payment = entity.Payment{}
	errScan := database.QueryRow("SELECT id, id_of_user, price, status, date_of_add, date_of_update, date_of_done FROM XXXXXXXXXXXXXXXXXXXXX.payment WHERE id = ? ORDER BY id DESC", id_of_payment).Scan(&payment.Id, &payment.Id_of_user, &payment.Price, &payment.Status, &payment.Date_of_add, &payment.Date_of_update, &payment.Date_of_done)
	if errScan != nil {
		panic(errScan.Error())
	}

	return &payment, nil
}

func FindById(id_of_payment int) (*entity.Payment, error) {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		//panic("Error database connection")
		return nil, err
	}
	defer database.Close()

	var payment entity.Payment = entity.Payment{}
	errScan := database.QueryRow("SELECT id, id_of_user, price, status, date_of_add, date_of_update, date_of_done FROM XXXXXXXXXXXXXXXXXXXXX.payment WHERE id = ? ORDER BY id DESC", id_of_payment).Scan(&payment.Id, &payment.Id_of_user, &payment.Price, &payment.Status, &payment.Date_of_add, &payment.Date_of_update, &payment.Date_of_done)
	if errScan != nil {
		if errScan == sql.ErrNoRows {
			return nil, nil
		} else {
			panic(errScan.Error())
		}
	}

	return &payment, nil
}

func FindLastNotActiveByIdOfPaymentAndIdOfUser(id_of_user int) (*entity.Payment, error) {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		//panic("Error database connection")
		return nil, err
	}
	defer database.Close()

	var payment entity.Payment = entity.Payment{}
	errScan := database.QueryRow("SELECT id, id_of_user, price, status, date_of_add, date_of_update, date_of_done FROM XXXXXXXXXXXXXXXXXXXXX.payment WHERE id_of_user = ? AND status = '0' ORDER BY id LIMIT 1", id_of_user).Scan(&payment.Id, &payment.Id_of_user, &payment.Price, &payment.Status, &payment.Date_of_add, &payment.Date_of_update, &payment.Date_of_done)
	if errScan != nil {
		if errScan == sql.ErrNoRows {
			return nil, nil
		} else {
			panic(errScan.Error())
		}
	}

	return &payment, nil
}

func UpdatePriceInDb(payment *entity.Payment, price int) error {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		return err
	}
	defer database.Close()

	date_of_update := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("UpdatePriceInDb price: ", price)
	_, errUpdate := database.Exec("UPDATE XXXXXXXXXXXXXXXXXXXXX.payment SET price = ? , date_of_update = ? WHERE id = ?", price, date_of_update, payment.Id)
	if errUpdate != nil {
		fmt.Println(errUpdate)
		return errUpdate
	}

	return nil
}

func UpdatePaymentConfirmPayed(payment *entity.Payment, dateOfDone time.Time) error {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		return err
	}
	defer database.Close()

	_, errUpdate := database.Exec("UPDATE XXXXXXXXXXXXXXXXXXXXX.payment SET status = ? , date_of_done = ? WHERE id = ?", 1, dateOfDone, payment.Id)
	if errUpdate != nil {
		fmt.Println("paymentRepository UpdatePaymentConfirmPayed:", errUpdate)
		return errUpdate
	}

	date_of_done := dateOfDone.Format("2006-01-02 15:04:05")
	payment.Status = 1
	payment.Date_of_done = &date_of_done

	return nil
}
