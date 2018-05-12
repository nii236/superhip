package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/nii236/superhip/models"
)

const getListURL = "/get/list"
const getURL = "/get"
const getManyURL = "/get/many"
const getManyReferenceURL = "/get/many/reference"
const createURL = "/create"
const updateURL = "/update"
const updateManyURL = "/update/many"
const deleteURL = "/delete"
const deleteManyURL = "/delete/many"

const permissionResource = "permissions"
const roleResource = "roles"
const schoolResource = "schools"
const teamResource = "teams"
const userResource = "users"
const studentResource = "students"

// Client is the HTTP client
type Client struct {
	baseURL string
	*http.Client
}

// New returns a new Client
func New(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		Client:  http.DefaultClient,
	}
}

// Do will start a request and unmarshal the result
func (c *Client) Do(resource string, path string, payload io.Reader, target interface{}) error {
	u, err := url.Parse(c.baseURL + "/" + resource + path)
	if err != nil {
		return err
	}
	resp, err := c.Post(u.String(), "application/json", payload)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("non 200 response: " + strconv.Itoa(resp.StatusCode))
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.Unmarshal(b, target)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) create(resource string, model interface{}) ([]byte, error) {
	b, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	req := &models.CreateRequest{
		Data: b,
	}

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	result := &models.Response{}

	err = c.Do(resource, createURL, bytes.NewReader(payload), result)
	if err != nil {
		return nil, err
	}

	return result.Data.MarshalJSON()
}

func (c *Client) list(resource string, model interface{}) ([]byte, error) {
	req := &models.GetListRequest{}

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	result := &models.Response{}

	err = c.Do(resource, getListURL, bytes.NewReader(payload), result)
	if err != nil {
		return nil, err
	}

	return result.Data.MarshalJSON()
}
