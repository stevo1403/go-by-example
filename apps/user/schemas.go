package user

type UserSchema struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

type UserOutSchema struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type UserLoginSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserOut struct {
	User UserOutSchema `json:"user"`
}

func (u User) to_schema() UserOutSchema {
	return UserOutSchema{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
	}
}

func (_u *UserSchema) from_schema() User {
	return User{
		FirstName: _u.FirstName,
		LastName:  _u.LastName,
		Email:     _u.Email,
		Phone:     _u.Phone,
		Password:  _u.Password,
	}
}

func (_u *UserSchema) to_object() User {
	return _u.from_schema()
}

type UsersOut struct {
	Users []UserOutSchema `json:"users"`
}

type LoginOut struct {
	User  UserOutSchema `json:"user"`
	Token string        `json:"token"`
}
