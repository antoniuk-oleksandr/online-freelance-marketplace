package body

type SignUpBody struct {
	Email     string `json:"email" db:"email"`
	FirstName string `json:"firstName" db:"first_name"`
	Password  string `json:"password" db:"password"`
	Surname   string `json:"surname" db:"surname"`
	Username  string `json:"username" db:"username"`
}
