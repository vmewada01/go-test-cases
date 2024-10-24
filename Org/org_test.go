package org

import (
	"context"
	"fmt"
	"testing"
)

func TestOrg(t *testing.T) {
	fmt.Println("Take the organization")

	tests := []struct {
		name      string
		user      *Organisation
		expectErr bool
	}{
		{
			name: "valid user",
			user: &Organisation{
				ID:             1,
				Name:           "Alice",
				Description:    "alice@example.com",
				EmployeesCount: 30,
				Location:       "Somewhere",
				Employees: []Employee{
					{ID: 1, Name: "Alice", Position: "Software Engineer"},
					{ID: 2, Name: "Bob", Position: "Project Manager"},
					{ID: 3, Name: "Charlie", Position: "Senior Software Engineer"},
				},
				Projects: []Project{
					{ID: 1, Name: "Project A", Description: "Description A"},
				},
			},
			expectErr: false,
		},
		{
			name: "valid user",
			user: &Organisation{
				ID:             1,
				Name:           "ASDF",
				Description:    "aliSADFSADFce@example.com",
				EmployeesCount: 30,
				Location:       "ASDFSDF",
				Employees: []Employee{
					{ID: 1, Name: "Alice", Position: "ASDF Engineer"},
					{ID: 2, Name: "Bob", Position: "ProjecASDFSDFt Manager"},
					{ID: 3, Name: "Charlie", Position: "SADF Software Engineer"},
				},
				Projects: []Project{
					{ID: 1, Name: "Project A", Description: "DeASFDSADFscription A"},
				},
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Org(context.Background(), tt.user)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
		})
	}
}
