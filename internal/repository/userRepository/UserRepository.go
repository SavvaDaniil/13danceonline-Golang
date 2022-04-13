package userRepository

import (
	"danceonline/config"
	"danceonline/internal/component"
	"danceonline/internal/entity"
	"database/sql"
	//"fmt"
	"time"
)

func FindByUsername(username string) (*entity.User, error) {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		//panic("Error database connection")
		return nil, err
	}
	defer database.Close()

	var u entity.User = entity.User{}
	errScan := database.QueryRow("SELECT id, username, password, authKey, accessToken, active, firstname, date_of_add, forgetCode, forgetTry, date_forget_last_try FROM XXXXXXXXXXXXXXXXXXXXX.user WHERE username = ? ORDER BY id DESC", username).Scan(
		&u.Id,
		&u.Username,
		&u.Password,
		&u.AuthKey,
		&u.AccessToken,
		&u.Active,
		&u.Firstname,
		&u.DateOfAdd,
		&u.ForgetCode,
		&u.ForgetTry,
		&u.DateForgetLastTry,
	)
	if errScan != nil {
		//panic("Error when user FindByUsername: " + errScan.Error())
		return nil, err
	}

	return &u, nil
}

func FindByUsernameExceptId(id_of_user int, username string) (*entity.User, error) {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		//panic("Error database connection")
		return nil, err
	}
	defer database.Close()

	var u entity.User = entity.User{}
	errScan := database.QueryRow("SELECT id, username, password, authKey, accessToken, active, firstname, date_of_add, forgetCode, forgetTry, date_forget_last_try FROM XXXXXXXXXXXXXXXXXXXXX.user WHERE username = ? AND id != ? ORDER BY id DESC", username, id_of_user).Scan(
		&u.Id,
		&u.Username,
		&u.Password,
		&u.AuthKey,
		&u.AccessToken,
		&u.Active,
		&u.Firstname,
		&u.DateOfAdd,
		&u.ForgetCode,
		&u.ForgetTry,
		&u.DateForgetLastTry,
	)
	if errScan != nil {
		if errScan == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &u, nil
}

func FindById(id_of_user int) (*entity.User, error) {
	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	var u entity.User = entity.User{}
	errScan := database.QueryRow("SELECT id, username, password, authKey, accessToken, active, firstname, date_of_add, forgetCode, forgetTry, date_forget_last_try FROM XXXXXXXXXXXXXXXXXXXXX.user WHERE id = ? ORDER BY id DESC", id_of_user).Scan(
		&u.Id,
		&u.Username,
		&u.Password,
		&u.AuthKey,
		&u.AccessToken,
		&u.Active,
		&u.Firstname,
		&u.DateOfAdd,
		&u.ForgetCode,
		&u.ForgetTry,
		&u.DateForgetLastTry,
	)
	//, &u.Password, &u.AuthKey, &u.AccessToken, &u.Active, &u.Fio, &u.DateOfAdd, &u.ForgetCode, &u.ForgetTry, &u.ForgetHash, &u.ForgetCount, &u.ForgetLast
	if errScan != nil {
		if errScan == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &u, nil
}

func Add(username string, passwordHash string) (*entity.User, error) {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	var id_of_user int64 = 0
	var authKey string = randomComponent.GenerateRandomString(32)
	resInsert, errInsert := database.Exec("INSERT INTO XXXXXXXXXXXXXXXXXXXXX.user (username, password, authKey) VALUES (?, ?, ?)", username, passwordHash, authKey)
	if errInsert != nil {
		panic(errInsert.Error())
		//return nil, errInsert
	}

	id_of_user, errLastInsert := resInsert.LastInsertId()
	if errLastInsert != nil {
		return nil, errLastInsert
	}

	var u entity.User = entity.User{}
	errScan := database.QueryRow("SELECT XXXXXXXXXXXXXXXXXXXXX.user.id, XXXXXXXXXXXXXXXXXXXXX.user.username, XXXXXXXXXXXXXXXXXXXXX.user.password FROM XXXXXXXXXXXXXXXXXXXXX.user WHERE id = ? ORDER BY id DESC", id_of_user).Scan(&u.Id, &u.Username, &u.Password)
	if errScan != nil {
		return nil, errScan
	}

	return &u, nil
}

func Update(id_of_user int, username string, password string, firstname string) error {
	...
}

func ForgetStep0Update(id_of_user int, date_forget_last_try time.Time, forgetTry int, codeNew string) error {
	...
}

func ForgetWrongUpdate(id_of_user int, date_forget_last_try time.Time, forgetTry int) error {
	...
}

func ForgetSuccessUpdate(id_of_user int, passwordHash string, authKey string) error {
	...
}
