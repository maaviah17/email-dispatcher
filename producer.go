package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func loadRecipient(filePath string) error {
	
	file,err := os.Open(filePath)
	if err != nil{
		return err
	}

	r := csv.NewReader(file)
	records,err := r.ReadAll()
	if err != nil{
		return err
	} 
	
	for _,record := range records[1:]{
		fmt.Println(record)
	}

	return nil
}