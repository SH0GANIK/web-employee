package model

type Employee struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Phone      string `json:"phone"`
	CompanyId  int    `json:"companyId"`
	Passport   `json:"passport"`
	Department `json:"department"`
}
type Passport struct {
	Type   string
	Number string
}
type Department struct {
	Name  string
	Phone string
}
