package db

import (
	"fmt"

	"github.com/Josh2604/literaluser/models"
	"github.com/Josh2604/literaluser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Starting signup")
	err := DBConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentence := fmt.Sprintf("INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES (%s, %s, %s)", sig.UserEmail, sig.UserUUID, tools.DateMysql())

	_, errExec := Db.Exec(sentence)
	if errExec != nil {
		fmt.Println("error insert user into users")
		return err
	}

	fmt.Println("user inserted into users")
	return nil
}
