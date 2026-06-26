package custom_marshal

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	aux := Alias(*u)
	aux.Name = strings.ToUpper(aux.Name)
	aux.Email = strings.ToUpper(aux.Email)
	aux.Password = strings.ToUpper(aux.Password)
	return json.Marshal(&aux)
}

func (u *User) UnmarshalJSON(data []byte) error {
	type Alias User
	err := json.Unmarshal(data, (*Alias)(u))
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	u.Name = strings.ToLower(u.Name)
	u.Email = strings.ToLower(u.Email)
	u.Password = strings.ToLower(u.Password)
	return nil
}
