package userService

type Provider string
const (
	Google Provider = "google"
)

type User struct {
	ID 			string 			`json:"user_id" db:"user_id"`
	CreatedAt 	string 			`json:"created_at" db:"created_at"`
	Provider 	Provider 		`json:"provider" db:"provider"`
	Email 		string 			`json:"email" db:"email"`
	Name 		string			`json:"name" db:"name"`
}
