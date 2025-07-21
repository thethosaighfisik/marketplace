package domain 

type User struct{
	ID string
	Password string 
	Email string 
}

func CreateUser(Email, HashOfPassword string) User{
	return User{Email:Email, Password:HashOfPassword}
}