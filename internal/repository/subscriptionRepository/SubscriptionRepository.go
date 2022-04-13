package subscriptionRepository

import (
	"danceonline/config"
	"danceonline/internal/entity"
	"database/sql"
	"fmt"
)

func FindById(id_of_subscription int) (*entity.Subscription, error) {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		//panic("Error database connection")
		return nil, err
	}
	defer database.Close()

	var subscription entity.Subscription = entity.Subscription{}
	errScan := database.QueryRow("SELECT id, name, price, days, status, order_in_list FROM XXXXXXXXXXXXXXXXXXXXX.subscription WHERE status = '1' ORDER BY order_in_list DESC", id_of_subscription).Scan(&subscription.Id, &subscription.Name, &subscription.Price, &subscription.Days, &subscription.Status, &subscription.OrderInList)
	if errScan != nil {
		if errScan == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, errScan
		}
	}

	return &subscription, nil
}

func ListAllActive() ([]entity.Subscription, error) {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		//panic("Error database connection")
		return nil, err
	}
	defer database.Close()

	rows, err := database.Query("SELECT id, name, price, days, status, order_in_list FROM XXXXXXXXXXXXXXXXXXXXX.subscription WHERE status = '1' ORDER BY order_in_list DESC")
	if err != nil {
		//panic("Error query")
		return nil, err
	}
	defer rows.Close()

	var subscriptions []entity.Subscription

	for rows.Next() {
		s := entity.Subscription{}
		err := rows.Scan(&s.Id, &s.Name, &s.Price, &s.Days, &s.Status, &s.OrderInList)
		if err != nil {
			fmt.Println(err)
		}
		subscriptions = append(subscriptions, s)
	}

	return subscriptions, nil
}
