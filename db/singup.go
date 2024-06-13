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
	// TODO: delete logs
	fmt.Println("Before sentence: ", sig.UserEmail, sig.UserUUID)
	sentence := fmt.Sprintf("INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('%s','%s','%s')", sig.UserEmail, sig.UserUUID, tools.DateMysql())
	fmt.Println("sentence: ", sentence)

	_, errExec := Db.Exec(sentence)
	if errExec != nil {
		fmt.Println("error insert user into users")
		return err
	}

	fmt.Println("user inserted into users")
	return nil
}
