package tutorialRepository

import (
	"danceonline/config"
	"danceonline/internal/entity"
	"database/sql"
	"fmt"
)

func FindById(id_of_tutorial int) (*entity.Tutorial, error) {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		//panic("Error database connection")
		return nil, err
	}
	defer database.Close()

	var tutorial entity.Tutorial = entity.Tutorial{}
	errScan := database.QueryRow("SELECT id, name, days, price, status, beginner, intermediate FROM XXXXXXXXXXXXXXXXXXXXX.tutorial WHERE id = ? ORDER BY id DESC", id_of_tutorial).Scan(&tutorial.Id, &tutorial.Name, &tutorial.Days, &tutorial.Price, &tutorial.Status, &tutorial.Beginner, &tutorial.Intermediate)
	if errScan != nil {
		return nil, errScan
	}

	return &tutorial, nil
}

func ListAll(page int) []entity.Tutorial {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		panic("Error database connection")
	}
	defer database.Close()

	rows, err := database.Query("SELECT XXXXXXXXXXXXXXXXXXXXX.tutorial.id, XXXXXXXXXXXXXXXXXXXXX.tutorial.name FROM XXXXXXXXXXXXXXXXXXXXX.tutorial ORDER BY order_in_list DESC")
	if err != nil {
		panic("Error query")
	}
	defer rows.Close()

	var tutorials []entity.Tutorial

	for rows.Next() {
		t := entity.Tutorial{}
		err := rows.Scan(&t.Id, &t.Name)
		if err != nil {
			fmt.Println(err)
		}
		tutorials = append(tutorials, t)
	}

	return tutorials
}

func ListAllActiveByFilter(page int, id_of_style int) ([]entity.Tutorial, error) {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		panic("Error database connection")
	}
	defer database.Close()

	//SELECT * FROM `tutorial` AS t WHERE (1 = 0 AND t.id_of_style = '0' OR t.id_of_style = '1');
	rows, err := database.Query("SELECT id, name, days, price, status, beginner, intermediate FROM XXXXXXXXXXXXXXXXXXXXX.tutorial WHERE status = '1' AND ( 0 = ? OR id_of_style = ? ) ORDER BY order_in_list DESC", id_of_style, id_of_style)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			fmt.Println("ListAllActiveByFilter err:", err)
			return nil, err
		}
	}
	defer rows.Close()

	var tutorials []entity.Tutorial

	for rows.Next() {
		t := entity.Tutorial{}
		err := rows.Scan(
			&t.Id,
			&t.Name,
			&t.Days,
			&t.Price,
			&t.Status,
			&t.Beginner,
			&t.Intermediate,
		)
		if err != nil {
			fmt.Println(err)
		}
		tutorials = append(tutorials, t)
	}

	return tutorials, nil
}
