package user

type User struct {
	ID           string `json:"id" bson:"_id"`
	Name         string `json:"name" bson:"name"`
	Surname      string `json:"surname" bson:"surname"`
	PasswordHash string `json:"-" bson:"password"`
	Email        string `json:"email" bson:"email"`
}

type UserDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *User) ToDTO() UserDTO {
	return UserDTO{
		ID:       u.ID,
		Name:     u.Name,
		Surname:  u.Surname,
		Password: u.PasswordHash,
		Email:    u.Email,
	}
}

func (u *UserDTO) ToUser() User {
	return User{
		ID:           u.ID,
		Name:         u.Name,
		Surname:      u.Surname,
		PasswordHash: u.Password,
		Email:        u.Email,
	}
}
