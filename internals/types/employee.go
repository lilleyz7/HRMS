package types

type Employee struct {
	FirstName   string       `json:"firstname"`
	LastName    string       `json:"lastname"`
	StartDate   string       `json:"startdate"`
	ExitDate    string       `json:"exitdate"`
	Salary      int          `json:"salary"`
	Infractions []Infraction `json:"infractions"`
	SickDays    int          `json:"sickdays"`
	PaidDays    int          `json:"paiddays"`
}

type Infraction struct {
	Date    string `json:"date"`
	Message string `json:"message"`
}
