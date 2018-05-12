package client

import (
	"github.com/nii236/superhip/models"
)

// UserCreate creates a user
func (c *Client) UserCreate(user *models.User) ([]byte, error) {
	return c.create(userResource, user)
}
