package main

import "fmt"

type Recipient struct {
	Name string
	Email string
}

func main(){

	fmt.Println("welcome to email dispatcher")
	loadRecipient("./emails-2.csv")
}