package models

type User struct {
	User_id    int        `json:"User_id"`
	Name       string     `json:"Name"`
	Surname    string     `json:"Surname"`
	AFM        int        `json:"AFM"`
	AMKA       int        `json:"AMKA"`
	Profession Profession `json:"Profession"`
	Email      string     `json:"Email"`
	Password   string     `json:"Password"`
}

type Profession struct {
	Role_id int    `json:"Role_id"`
	Title   string `json:"Title"`
}
