package main

import (
	"fmt"
	"sync"
)

type Recipient struct {
	Name string
	Email string
}

func main(){

	fmt.Println("welcome to email dispatcher")

	//will use an unbuffered channel[capacity=0] || 
	// sender will be blocked until the reciever isnt ready
	recipientChannel := make(chan Recipient)

	go func(){
		loadRecipient("./emails-2.csv", recipientChannel)
	}()
	
	var wg sync.WaitGroup
	workerCount := 5

	for i := 1; i<=workerCount; i++ {
		wg.Add(1)
		go emailWorker(i,recipientChannel, &wg)
	}

	wg.Wait()

}