package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/lsduenasb/desafio-go-bases/internal/tickets"
)

func createTicketsList(records [][] string)(ticketList []tickets.Ticket) {
	for _, row:= range records{
		var record tickets.Ticket 
		var structFields = []string{}
		for _, col:= range row {
			structFields = append(structFields, col)		
		}
		// fill structure
		idp, _:= strconv.Atoi(structFields[0])
		pricep, _:= strconv.ParseFloat(structFields[4], 64)
		record = tickets.Ticket{
			Id: idp,
			Name: structFields[1],
			Email: structFields[2],
			Country: structFields[3],
			Datetime: structFields[4],
			Price: pricep,
		}
		//fmt.Println(record)
		ticketList = append(ticketList, record)
	}
	return 
}

func ReadCSVFile()[][]string{
	file, err:= os.Open("./tickets.csv")
	if err!= nil{
		panic("File can´t be found")
	}
	defer file.Close()

	// reader
	reader := csv.NewReader(file)
	// read the records
	records, err := reader.ReadAll()

	if err!= nil{
		panic("Can´t read records")
	}
	return records
}

var TicketList = []tickets.Ticket{} 

func main() {

	// read file to get records into a slice of slices
	records:= ReadCSVFile()
	// create ticket list of structures 
	TicketList = createTicketsList(records)

	// R1 
	total, err := tickets.GetTotalTickets("Brazil", TicketList)
	if err != nil {
		fmt.Println("Error total")
	}
	fmt.Println("Total", total)

	// R3
	average, err:= tickets.AverageDestination(float64(total), TicketList)
	if err != nil {
		fmt.Println("Error average ")
	}
	fmt.Println("Average", average)
}
