package main

import (
	"encoding/csv"
	// "fmt"
	"os"
)

func loadRecipient(filePath string, ch chan Recipient) error {
	defer close(ch)
	file,err := os.Open(filePath)
	if err != nil{
		return err
	}

	defer file.Close()

	r := csv.NewReader(file)
	records,err := r.ReadAll()
	if err != nil{
		return err
	} 

	
	for _,record := range records[1:]{
		//send -> consumer (will be done using channels)
		ch <- Recipient{
			Name: record[0],
			Email: record[1],
		}
	}	

	return nil
}