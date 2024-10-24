package user

import (
	"context"
	"testing"
)

func TestAddUser(t *testing.T) {
	tests := []struct {
		name      string
		user      *User
		expectErr bool
	}{
		{
			name:      "valid user",
			user:      &User{Name: "Alice", Email: "alice@example.com", Age: 30, Password: "password123"},
			expectErr: true,
		},
		{
			name:      "empty name",
			user:      &User{Name: "", Email: "bob@example.com", Age: 25, Password: "password123"},
			expectErr: false,
		},
		{
			name:      "invalid email",
			user:      &User{Name: "Bob", Email: "invalid-email", Age: 25, Password: "password123"},
			expectErr: true,
		},
		{
			name:      "negative age",
			user:      &User{Name: "Charlie", Email: "charlie@example.com", Age: -5, Password: "password123"},
			expectErr: true,
		},
		{
			name:      "short password",
			user:      &User{Name: "David", Email: "david@example.com", Age: 0, Password: "123"},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := AddUser(context.Background(), tt.user)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
		})
	}
}
