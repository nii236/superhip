package main

import (
	"encoding/json"
	"fmt"

	"github.com/manveru/faker"

	"github.com/nii236/superhip/client"
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

const schoolResource = "schools"
const teamResource = "teams"
const userResource = "users"
const studentResource = "students"

type Seeder struct {
	*client.Client
	faker *faker.Faker
}

func main() {
	c := client.New("http://localhost:8080")
	f, err := faker.New("en")
	if err != nil {
		fmt.Println(err)
		return
	}
	seeder := &Seeder{c, f}
	err = seeder.seedPermissions()
	if err != nil {
		fmt.Println("seedPermissions:", err)
		return
	}
	err = seeder.seedRoles()
	if err != nil {
		fmt.Println("seedRoles:", err)
		return
	}
	err = seeder.seedSchools()
	if err != nil {
		fmt.Println("seedSchools:", err)
		return
	}
	err = seeder.seedUsers()
	if err != nil {
		fmt.Println("seedUsers:", err)
		return
	}
	err = seeder.seedTeams()
	if err != nil {
		fmt.Println("seedTeams:", err)
		return
	}
	err = seeder.seedStudents()
	if err != nil {
		fmt.Println("seedStudents:", err)
		return
	}
}

func (c *Seeder) seedPermissions() error {
	var err error
	_, err = c.Client.PermissionCreate(&models.Permission{
		Name: "manage schools",
	})
	_, err = c.Client.PermissionCreate(&models.Permission{
		Name: "manage teams",
	})
	_, err = c.Client.PermissionCreate(&models.Permission{
		Name: "manage students",
	})
	if err != nil {
		return fmt.Errorf("could not create model: %s", err)
	}
	return nil
}
func (c *Seeder) seedRoles() error {
	var err error
	_, err = c.Client.RoleCreate(&models.Role{
		Name: "teacher",
	})
	_, err = c.Client.RoleCreate(&models.Role{
		Name: "admin",
	})
	if err != nil {
		return fmt.Errorf("could not create model: %s", err)
	}
	return nil
}
func (c *Seeder) seedSchools() error {
	for i := 0; i < 10; i++ {
		school := &models.School{
			Name: c.faker.CompanyName(),
		}
		_, err := c.Client.SchoolCreate(school)
		if err != nil {
			return fmt.Errorf("could not create model: %s", err)
		}
	}
	return nil
}

func (c *Seeder) seedUsers() error {
	for i := 0; i < 10; i++ {
		user := &models.User{
			FirstName: c.faker.FirstName(),
			LastName:  c.faker.LastName(),
			Email:     c.faker.Email(),
			Password:  c.faker.Characters(10),
			SchoolIDs: models.UUIDArray{},
		}
		c.Client.UserCreate(user)
	}
	return nil
}

func (c *Seeder) seedTeams() error {
	b, err := c.Client.SchoolGetList(&models.GetListRequest{})
	if err != nil {
		return fmt.Errorf("could not fetch list: %s", err)
	}
	schools := models.SchoolList{}
	err = json.Unmarshal(b, &schools)
	if err != nil {
		return fmt.Errorf("could not unmarshal list: %s", err)
	}
	for _, school := range schools {
		for i := 0; i < 10; i++ {
			team := &models.Team{
				SchoolID: school.ID,
				Name:     c.faker.CompanyCatchPhrase(),
			}
			_, err := c.Client.TeamCreate(team)
			if err != nil {
				return fmt.Errorf("could not create model: %s", err)
			}
		}
	}
	return nil
}

func (c *Seeder) seedStudents() error {
	b, err := c.Client.SchoolGetList(&models.GetListRequest{})
	if err != nil {
		return fmt.Errorf("could not fetch list: %s", err)
	}
	schools := models.SchoolList{}
	err = json.Unmarshal(b, &schools)
	if err != nil {
		return fmt.Errorf("could not unmarshal list: %s", err)
	}
	for _, school := range schools {
		for i := 0; i < 10; i++ {
			student := &models.Student{
				SchoolID: school.ID,
				Name:     c.faker.FirstName(),
			}
			_, err := c.Client.StudentCreate(student)
			if err != nil {
				return fmt.Errorf("could not create model: %s", err)
			}
		}
	}
	return nil
}
