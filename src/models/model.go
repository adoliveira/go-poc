package models

// User representa um modelo de usuário de exemplo
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// IsValid verifica se os dados do usuário são válidos
func (u User) IsValid() bool {
	return u.ID > 0 && u.Name != ""
}
