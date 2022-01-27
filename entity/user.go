package entity

type User struct {
	ID       int    `db:"id" json:"id"`
	Fullname string `db:"fullname" json:"fullname,omitempty"`
	Username string `db:"username" json:"username,omitempty"`
	Phone    string `db:"phone" json:"phone,omitempty"`
	Email    string `db:"email" json:"email,omitempty"`
	Password string `db:"password" json:"password,omitempty"`
}
