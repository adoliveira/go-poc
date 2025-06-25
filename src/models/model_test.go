package models

import "testing"

func TestUser_IsValid(t *testing.T) {
	tests := []struct {
		user  User
		valid bool
	}{
		{User{ID: 1, Name: "Exemplo"}, true},
		{User{ID: 0, Name: "Exemplo"}, false},
		{User{ID: 1, Name: ""}, false},
		{User{ID: 0, Name: ""}, false},
	}

	for _, tt := range tests {
		if tt.user.IsValid() != tt.valid {
			t.Errorf("esperado %v para %+v", tt.valid, tt.user)
		}
	}
}
