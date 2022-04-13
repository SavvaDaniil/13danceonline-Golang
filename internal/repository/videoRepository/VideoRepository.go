package videoRepository

import (
	"danceonline/config"
	"danceonline/internal/entity"
	"database/sql"
	"fmt"
)

func FindById(id_of_video int) (*entity.Video, error) {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		//panic("Error database connection")
		return nil, err
	}
	defer database.Close()

	var video entity.Video = entity.Video{}
	errScan := database.QueryRow("SELECT id, hash, status, date_of_add, date_of_update FROM XXXXXXXXXXXXXXXXXXXXX.video WHERE id = ? ORDER BY id DESC", id_of_video).Scan(&video.Id, &video.Hash, &video.Status, &video.Date_of_add, &video.Date_of_update)
	if errScan != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			fmt.Println("videoRepository errScan:", errScan)
			return nil, errScan
		}
	}

	return &video, nil
}

func FindByIdOfTutorial(id_of_tutorial int) (*entity.Video, error) {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		//panic("Error database connection")
		return nil, err
	}
	defer database.Close()

	var video entity.Video = entity.Video{}
	errScan := database.QueryRow("SELECT video.id, video.hash, video.status, video.date_of_add, video.date_of_update FROM XXXXXXXXXXXXXXXXXXXXX.video INNER JOIN connection_tutorial_to_video ON connection_tutorial_to_video.id_of_tutorial = ? AND connection_tutorial_to_video.id_of_video = video.id WHERE video.status = '1' ORDER BY video.id DESC LIMIT 1", id_of_tutorial).Scan(&video.Id, &video.Hash, &video.Status, &video.Date_of_add, &video.Date_of_update)
	if errScan != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			fmt.Println("videoRepository errScan:", errScan)
			return nil, errScan
		}
	}

	return &video, nil
}

func ListAllConnectedToCourse(id_of_course int) ([]entity.Video, error) {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		//panic("Error database connection")
		return nil, err
	}
	defer database.Close()

	rows, err := database.Query("SELECT video.id, video.hash, video.status, video.date_of_add, video.date_of_update FROM XXXXXXXXXXXXXXXXXXXXX.video INNER JOIN connection_course_to_video ON connection_course_to_video.id_of_course = ? AND connection_course_to_video.id_of_video = video.id WHERE video.status = '1' ORDER BY connection_course_to_video.order_in_list ASC", id_of_course)
	if err != nil {
		//panic("Error query")
		return nil, err
	}
	defer rows.Close()

	var videos []entity.Video

	for rows.Next() {
		video := entity.Video{}
		err := rows.Scan(&video.Id, &video.Hash, &video.Status, &video.Date_of_add, &video.Date_of_update)
		if err != nil {
			fmt.Println(err)
			continue
		}
		videos = append(videos, video)
	}

	return videos, nil
}
