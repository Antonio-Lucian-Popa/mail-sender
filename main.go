package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/smtp"
	"os"
)

type mail struct {
	NAME    string `json:"name"`
	EMAIL   string `json:"email"`
	SUBJECT string `json:"subject"`
	MESSAGE string `json:"message"`
}

func createEmail(email *gin.Context) {
	var newEmail mail

	if err := email.BindJSON(&newEmail); err != nil {
		return
	}

	sendMailSimple(newEmail)

	email.IndentedJSON(http.StatusCreated, newEmail)
}

func sendMailSimple(email mail) {
	fmt.Println(email)
	auth := smtp.PlainAuth(
		"",
		email.EMAIL,
		"vrnugowggzvkkqye",
		"smtp.gmail.com",
	)

	msg := email.MESSAGE

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		email.EMAIL,
		[]string{"fast.transanto@gmail.com"},
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	router := gin.Default()
	router.POST("/sendMail", createEmail)
	router.Run(os.Getenv("APP_URL"))
}
