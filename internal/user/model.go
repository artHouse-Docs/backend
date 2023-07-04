package user

type User struct {
	ID           string `json:"id" bson:"_id"`
	Name         string `json:"name" bson:"name"`
	Surname      string `json:"surname" bson:"surname"`
	PasswordHash string `json:"password" bson:"password"`
	Email        string `json:"email" bson:"email"`
}
