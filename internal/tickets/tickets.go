package tickets

type Ticket struct {
	Id int
	Name string
	Email string
	Country string
	Datetime string
	Price float64
}


// R1
func GetTotalTickets(destination string, ticketList []Ticket) (counter int, e error) {
	counter = 0
	if ticketList == nil {
		panic(e)
	}
	for _, ticket := range ticketList{
		if ticket.Country == destination {
			counter++
		}
	}
	return
}

// R2
//func GetMornings(time string) (int error) {
	
//}

// R3
func AverageDestination(total float64, ticketList []Ticket) (average float64, err error) {
	if ticketList == nil {
		panic(err)
	}
	average = (total / float64(len(ticketList)))*100
	return
}
