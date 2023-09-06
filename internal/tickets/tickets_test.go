package tickets

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var ticketList = []Ticket {
	{
		Id:       1,
		Name:     "Tait Mc Caughan",
		Email:    "tmc0@scribd.com",
		Country:  "Finland",
		Datetime: "17:11",
		Price:    785.0,
	},
	{
		Id:       2,
		Name:     "Padget McKee",
		Email:    "pmckee1@hexun.com",
		Country:  "China",
		Datetime: "20:19",
		Price:    537.0,
	},
	{
		Id:       3,
		Name:     "Yalonda Jermyn",
		Email:    "yjermyn2@omniture.com",
		Country:  "China",
		Datetime: "18:11",
		Price:    579.0,
	},
	{
		Id:       71,
		Name:     "Marcie Goodricke",
		Email:    "mgoodricke1y@cbc.ca",
		Country:  "Finland",
		Datetime: "4:24",
		Price:    1201.0,
	},
}
func TestGetTotalTickets(t *testing.T) {
	t.Run("Get total tickets from specific country", func(t *testing.T) {
		// Arrange
		country := "Finland"
		expectedResult := 2
		
		// Act
		obtainedResult, err := GetTotalTickets(country, ticketList)

		//Assert
		if err == nil {
			assert.Equal(t, expectedResult, obtainedResult)
		}
	})
	t.Run("Get total tickets from specific country", func(t *testing.T) {
		// Arrange
		country := "China"
		expectedResult := 2
		
		// Act
		obtainedResult, err := GetTotalTickets(country, ticketList)

		//Assert
		if err == nil {
			assert.Equal(t, expectedResult, obtainedResult)
		}
	})
}

func TestAverageDestination(t *testing.T) {
	t.Run("Get average destination", func(t *testing.T) {
		// Arrange
		country:= "Finland"
		total,_ := GetTotalTickets(country, ticketList)
		expectedResult := 50.0
		
		// Act
		obtainedResult, err := AverageDestination(float64(total), ticketList)

		//Assert
		if err == nil {
			assert.Equal(t, expectedResult, obtainedResult)
		}
	})
	t.Run("Get average destination", func(t *testing.T) {
		// Arrange
		country:= "China"
		total,_ := GetTotalTickets(country, ticketList)
		expectedResult := 50.0
		
		// Act
		obtainedResult, err := AverageDestination(float64(total), ticketList)

		//Assert
		if err == nil {
			assert.Equal(t, expectedResult, obtainedResult)
		}
	})
}