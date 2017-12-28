package auth

import "github.com/riesinger/freecloud/models"

// CredentialsProvider is an interface for various credential sources like Databases or alike
type CredentialsProvider interface {
	GetUserByID(uid int) (*models.User, error)
	CreateUser(user *models.User) error
}