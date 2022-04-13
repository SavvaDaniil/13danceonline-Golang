package purchaseSubscriptionRepository

import (
	"danceonline/config"
	"danceonline/internal/entity"
	"database/sql"
	"fmt"
	"time"
)

func FindById(id_of_purchase_subscription int) (*entity.PurchaseSubscription, error) {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	var pt entity.PurchaseSubscription = entity.PurchaseSubscription{}
	errScan := database.QueryRow("SELECT id, id_of_user, id_of_payment, Id_of_subscription, days, active, date_of_add, date_of_activation, date_of_must_be_used_to FROM XXXXXXXXXXXXXXXXXXXXX.purchase_subscription WHERE id_of_user = ? ORDER BY id DESC", id_of_purchase_subscription).Scan(&pt.Id, &pt.Id_of_user, &pt.Id_of_payment, &pt.Id_of_subscription, &pt.Days, &pt.Active, &pt.Date_of_add, &pt.Date_of_activation, &pt.Date_of_must_be_used_to)
	if errScan != nil {
		return nil, errScan
	}

	return &pt, nil

}

func Add(id_of_user int, id_of_payment int, Id_of_subscription int, days int) (*entity.PurchaseSubscription, error) {

	...
}

func FindAllByIdOfUser(id_of_user int) []entity.PurchaseSubscription {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		panic("Error database connection")
	}
	defer database.Close()

	rows, err := database.Query("SELECT id, id_of_user, id_of_payment, Id_of_subscription, days, active, date_of_add, date_of_activation, date_of_must_be_used_to FROM XXXXXXXXXXXXXXXXXXXXX.purchase_subscription WHERE id_of_user = ? ORDER BY id DESC", id_of_user)
	if err != nil {
		panic("Error query")
	}
	defer rows.Close()

	var PurchaseSubscriptions []entity.PurchaseSubscription

	for rows.Next() {
		pt := entity.PurchaseSubscription{}
		err := rows.Scan(&pt.Id, &pt.Id_of_user, &pt.Id_of_payment, &pt.Id_of_subscription, &pt.Days, &pt.Active, &pt.Date_of_add, &pt.Date_of_activation, &pt.Date_of_must_be_used_to)
		if err != nil {
			fmt.Println(err)
		}
		PurchaseSubscriptions = append(PurchaseSubscriptions, pt)
	}

	return PurchaseSubscriptions
}

func FindLastAvailableByIdOfUser(id_of_user int) (*entity.PurchaseSubscription, error) {
	...
}

func Activate(PurchaseSubscription *entity.PurchaseSubscription) error {

	if PurchaseSubscription.Date_of_activation == nil {
		var days int = PurchaseSubscription.Days
		if days <= 0 {
			days = 1
		}

		dateOfNow := time.Now()
		dateOfActivationStr := dateOfNow.Format("2006-01-02 15:04:05")
		dateOfMustBeUsedToStr := dateOfNow.AddDate(0, 0, days).Format("2006-01-02")

		PurchaseSubscription.Date_of_activation = &dateOfActivationStr
		PurchaseSubscription.Date_of_must_be_used_to = &dateOfMustBeUsedToStr

		database, err := sql.Open("mysql", config.DbConnection)
		if err != nil {
			//panic("Error database connection")
			return err
		}
		defer database.Close()

		_, errUpdate := database.Exec("UPDATE XXXXXXXXXXXXXXXXXXXXX.purchase_subscription SET date_of_activation = ?, date_of_must_be_used_to = ? WHERE id = ? ORDER BY id DESC", PurchaseSubscription.Date_of_activation, PurchaseSubscription.Date_of_must_be_used_to, PurchaseSubscription.Id)
		if errUpdate != nil {
			return errUpdate
		}

	}

	return nil
}
