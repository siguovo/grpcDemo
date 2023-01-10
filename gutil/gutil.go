package gutil

type User struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

func InitRoot() User {
	return User{Name: "root", Level: 5}
}

func (u *User) GetLevel() int {
	return u.Level
}

func (u *User) SetName(n string) {
	u.Name = n
}
