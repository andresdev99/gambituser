package db

import (
	"database/sql"
	"fmt"
	"gambituser/models"
	"gambituser/secretm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = Db.Ping()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Connected Successfully")
	return nil
}

func ConnStr(keys models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = keys.UserName
	authToken = keys.Password
	dbEndpoint = keys.Host
	dbName = "gambit"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}
