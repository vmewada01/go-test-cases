package org

import (
	"context"
	"errors"
	"fmt"
)

type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

type Organisation struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	Location       string     `json:"location"`
	Employees      []Employee `json:"employees"`
	Projects       []Project  `json:"projects"`
	EmployeesCount int        `json:"employees_count"`
}

func Org(ctx context.Context, og *Organisation) (org *Organisation, err error) {
	fmt.Println("Org package")

	// Check the values and return them after checking
	if og.ID < 1 || len(og.Name) == 0 {
		return nil, errors.New("invalid org id or name")
	}

	org = &Organisation{
		ID:             og.ID,
		Name:           og.Name,
		Description:    og.Description,
		Location:       og.Location,
		Employees:      og.Employees,
		Projects:       og.Projects,
		EmployeesCount: len(og.Employees), // Now using the length correctly
	}

	if len(og.Employees) > 0 {
		for _, e := range og.Employees {
			if e.ID < 1 || len(e.Name) == 0 {
				return nil, errors.New("invalid employee id or name")
			}
		}
	}

	if len(og.Projects) > 0 {
		for _, p := range og.Projects {
			if p.ID < 1 || len(p.Name) == 0 {
				return nil, errors.New("invalid project id or name")
			}
		}
	}

	return org, nil
}
