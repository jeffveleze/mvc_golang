package entities

type User struct {
	Id          int     `json:id`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	CreatedDate []uint8 `json:"created_date"`
	Password    string  `json:"password"`
	Token       string  `json:token`
}
