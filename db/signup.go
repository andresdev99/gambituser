package db

import (
	"fmt"
	"gambituser/models"
	"gambituser/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Starting Signing Up")

	err := DbConnect()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer Db.Close()
	sentence := fmt.Sprintf(
		"INSERT INTO users (User_UUID, User_Email, User_DateAdd) VALUES ('%s', '%s', '%s')",
		sig.UserUUID, sig.UserEmail, tools.MYSQLDate())
	fmt.Println(sentence)

	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("SignUp Done")
	return nil
}
