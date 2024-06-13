package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Josh2604/literaluser/models"
	"github.com/Josh2604/literaluser/secretm"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DBConnect() error {
	Db, errdb := sql.Open("mysql", ConStr(SecretModel))
	if errdb != nil {
		fmt.Println(fmt.Errorf("error opening db: %v", err))
		return err
	}
	errdb = Db.Ping()

	if errdb != nil {
		fmt.Println(fmt.Errorf("error ping on db: %v", err))
		return err
	}

	fmt.Println("Connected to DB")
	return nil
}

func ConStr(claves models.SecretRDSJson) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", claves.Username, claves.Password, claves.Host, "literal")
	fmt.Println("check ConStr: ", claves.Username)
	return dsn
}
