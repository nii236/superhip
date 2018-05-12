package client

import (
	"github.com/nii236/superhip/models"
)

// SchoolGetList is the client method for SchoolGetList
func (c *Client) SchoolGetList(req *models.GetListRequest) ([]byte, error) {
	return c.list(schoolResource, &models.School{})
}

// SchoolCreate is the client method for SchoolCreate
func (c *Client) SchoolCreate(school *models.School) ([]byte, error) {
	return c.create(schoolResource, school)
}
