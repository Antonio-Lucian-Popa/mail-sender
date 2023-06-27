package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/smtp"
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

	sendMailSimple(&newEmail)

	email.IndentedJSON(http.StatusCreated, newEmail)
}

func sendMailSimple(email *mail) {
	fmt.Println(email)
	auth := smtp.PlainAuth(
		"",
		"fast.transanto@gmail.com",
		"vrnugowggzvkkqye",
		"smtp.gmail.com",
	)

	// Compose the email message
	headers := make(map[string]string)
	headers["From"] = email.EMAIL
	headers["To"] = "fast.transanto@gmail.com"
	headers["Subject"] = email.SUBJECT

	msg := ""
	for k, v := range headers {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + email.MESSAGE

	// Send the email
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
	router.Run("localhost:8080")
}
