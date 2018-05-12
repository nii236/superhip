package client

import "github.com/nii236/superhip/models"

// PermissionCreate is the client method for PermissionCreate
func (c *Client) PermissionCreate(permission *models.Permission) ([]byte, error) {
	return c.create(permissionResource, permission)
}
