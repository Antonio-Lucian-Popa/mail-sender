package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
)

type mail struct {
	NAME    string `json:"name"`
	EMAIL   string `json:"email"`
	SUBJECT string `json:"subject"`
	MESSAGE string `json:"message"`
}

func createEmail(event mail) error {
	sendMailSimple(event)
	return nil
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
		log.Println(err)
	}
}

func main() {
	if os.Getenv("NETLIFY_LAMBDA") == "true" {
		lambda.Start(createEmail)
	} else {
		router := gin.Default()
		router.POST("/.netlify/functions/sendMail", func(c *gin.Context) {
			var newEmail mail

			if err := c.BindJSON(&newEmail); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
				return
			}

			if err := createEmail(newEmail); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
				return
			}

			c.JSON(http.StatusCreated, newEmail)
		})
		router.Run(":8888")
	}
}
