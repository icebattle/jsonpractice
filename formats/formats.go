package formats

//Person - this is for learning
type Person struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

//PersonWithResides - learning how to deal with embedded objects
type PersonWithResides struct {
	Fname   string  `json:"fname"`
	Lname   string  `json:"lname"`
	Resides Resides `json:"Resides"`
}

//Resides - this should always be embedded
type Resides struct {
	City    string `json:"city"`
	Country string `json:"country"`
}
