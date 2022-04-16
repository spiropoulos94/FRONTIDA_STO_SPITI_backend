package models

type User struct {
	User_id    int    `json:"User_id"`
	Name       string `json:"Name"`
	Surname    string `json:"Surname"`
	AFM        int    `json:"AFM"`
	AMKA       int    `json:"AMKA"`
	Profession string `json:"Profession"`
	Email      string `json:"Email"`
	Password   string `json:"Password"`
}
