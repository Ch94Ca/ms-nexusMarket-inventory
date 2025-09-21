// Package domain contains domain entities and repository contracts for the inventory microservice.
package domain

type UserRepository interface {
	GetByID(id string) (*User, error)
}
