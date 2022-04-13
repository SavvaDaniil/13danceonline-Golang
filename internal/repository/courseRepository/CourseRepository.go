package courseRepository

import (
	"danceonline/config"
	"danceonline/internal/entity"
	"database/sql"
	"fmt"
)

func FindById(id_of_course int) (*entity.Course, error) {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		//panic("Error database connection")
		return nil, err
	}
	defer database.Close()

	var course entity.Course = entity.Course{}
	errScan := database.QueryRow("SELECT id, name, days, price, status, beginner, intermediate, order_in_list FROM XXXXXXXXXXXXXXXXXXXXX.Course WHERE id = ? ORDER BY id DESC", id_of_course).Scan(&course.Id, &course.Name, &course.Days, &course.Price, &course.Status, &course.Beginner, &course.Intermediate, &course.Order_in_list)
	if errScan != nil {
		return nil, errScan
	}

	return &course, nil
}

func ListAll(page int) []entity.Course {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		panic("Error database connection")
	}
	defer database.Close()

	rows, err := database.Query("SELECT XXXXXXXXXXXXXXXXXXXXX.Course.id, XXXXXXXXXXXXXXXXXXXXX.Course.name FROM XXXXXXXXXXXXXXXXXXXXX.Course ORDER BY order_in_list DESC")
	if err != nil {
		panic("Error query")
	}
	defer rows.Close()

	var courses []entity.Course

	for rows.Next() {
		course := entity.Course{}
		err := rows.Scan(
			&course.Id,
			&course.Name,
			&course.Days,
			&course.Price,
			&course.Status,
			&course.Beginner,
			&course.Intermediate,
			&course.Order_in_list,
		)
		if err != nil {
			fmt.Println(err)
		}
		courses = append(courses, course)
	}

	return courses
}

func ListAllActiveByFilter(page int) ([]entity.Course, error) {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		panic("Error database connection")
	}
	defer database.Close()

	rows, err := database.Query("SELECT id, name, days, price, status, beginner, intermediate, order_in_list FROM XXXXXXXXXXXXXXXXXXXXX.Course WHERE status = '1' ORDER BY order_in_list DESC")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			//fmt.Println("ListAllActiveByFilter err:", err)
			return nil, err
		}
	}
	defer rows.Close()

	var courses []entity.Course

	for rows.Next() {
		course := entity.Course{}
		err := rows.Scan(
			&course.Id,
			&course.Name,
			&course.Days,
			&course.Price,
			&course.Status,
			&course.Beginner,
			&course.Intermediate,
			&course.Order_in_list,
		)
		if err != nil {
			fmt.Println(err)
		}
		courses = append(courses, course)
	}

	return courses, nil
}
