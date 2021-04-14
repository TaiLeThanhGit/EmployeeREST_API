package Types

type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	BirthDay  Date   `json:birthday`
}

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type Response4Update struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}
