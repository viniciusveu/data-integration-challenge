package models

type Company struct {
	Name string `json:"name"`
	Zip string `json:"zip"`
	Site string `json:"site"`
}

type UpdateCompanyInput struct {
	Name  string `json:"name"`
	Zip string `json:"zip"`  
	Site string `json:"site"`  
}

type CreateCompanyInput struct {
	Name  string `json:"name"`
	Zip string `json:"zip"`  
	Site string `json:"site"`  
}