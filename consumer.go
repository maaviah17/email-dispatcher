package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup){
	
	defer wg.Done()
	for recipient := range ch {

		smtpHost := "localhost"
		smtpPort := "1025"

		// formattedMsg := fmt.Sprintf("To : %s\r\nSubject : Hi %s,Testing Email\r\n\r\n%s\r\n",recipient.Email, recipient.Name, "Just testing our email campaign tool")		
		// msg := []byte(formattedMsg)

		msg, err := executeTempelate(recipient)
		if err != nil{
			fmt.Printf("worker : %d, error parsing tempelate for %s", id, recipient.Email)
			// can add better error handling 
			continue
		}
		fmt.Printf("Worker %d : Sending mail to %s⏳ \n", id , recipient.Email)

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "dhriti@coder.com", []string{recipient.Email}, []byte(msg))
		if err != nil{
			log.Fatal(err)
		}

		time.Sleep(50*time.Microsecond)

		fmt.Printf("Worker %d : Sent mail to %s✅ \n", id , recipient.Email)
	}

}