package client

import (
	"github.com/nii236/superhip/models"
)

// StudentGetList is the client method for StudentGetList
func (c *Client) StudentGetList() ([]byte, error) {
	return c.list(studentResource, &models.Student{})
}

// StudentCreate is the client method for StudentCreate
func (c *Client) StudentCreate(student *models.Student) ([]byte, error) {
	return c.create(studentResource, student)
}
