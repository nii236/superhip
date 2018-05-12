package client

import "github.com/nii236/superhip/models"

// RoleCreate is the client method for RoleCreate
func (c *Client) RoleCreate(role *models.Role) ([]byte, error) {
	return c.create(roleResource, role)
}
