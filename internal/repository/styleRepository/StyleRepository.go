package styleRepository

import (
	"danceonline/config"
	"danceonline/internal/entity"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ListAll() []entity.Style {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		panic("Error database connection")
	}
	defer database.Close()

	rows, err := database.Query("SELECT * FROM XXXXXXXXXXXXXXXXXXXXX.style")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	styles := []entity.Style{}

	for rows.Next() {
		p := entity.Style{}
		err := rows.Scan(&p.Id, &p.Link, &p.Name, &p.Description, &p.Order_in_list, &p.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		styles = append(styles, p)
	}

	return styles
}

func ListAllActive() []entity.Style {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		panic("Error database connection")
	}
	defer database.Close()

	rows, err := database.Query("SELECT * FROM XXXXXXXXXXXXXXXXXXXXX.style WHERE status = '1' ORDER BY name")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	styles := []entity.Style{}

	for rows.Next() {
		p := entity.Style{}
		err := rows.Scan(&p.Id, &p.Link, &p.Name, &p.Description, &p.Order_in_list, &p.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		styles = append(styles, p)
	}

	return styles
}
