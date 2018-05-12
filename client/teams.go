package client

import (
	"github.com/nii236/superhip/models"
)

// TeamCreate is the client method for TeamCreate
func (c *Client) TeamCreate(team *models.Team) ([]byte, error) {
	return c.create(teamResource, team)
}
