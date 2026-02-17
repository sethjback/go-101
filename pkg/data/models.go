package data

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
}
