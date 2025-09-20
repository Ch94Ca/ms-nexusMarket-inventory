// Package domain contains domain entities and repository contracts for the inventory microservice.
package domain

import "time"

type User struct {
	ID        string
	UserName  string
	FullName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
