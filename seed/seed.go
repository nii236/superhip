package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/gocarina/gocsv"
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
	files packr.Box
}

func main() {
	c := client.New("http://localhost:8080/api")
	f, err := faker.New("en")
	if err != nil {
		fmt.Println(err)
		return
	}

	b := packr.NewBox("./files")

	seeder := &Seeder{c, f, b}
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
		Name: "Can manage schools",
	})
	_, err = c.Client.PermissionCreate(&models.Permission{
		Name: "Can manage teams",
	})
	_, err = c.Client.PermissionCreate(&models.Permission{
		Name: "Can manage students",
	})
	_, err = c.Client.PermissionCreate(&models.Permission{
		Name: "Can manage users",
	})
	_, err = c.Client.PermissionCreate(&models.Permission{
		Name: "Can manage roles",
	})
	_, err = c.Client.PermissionCreate(&models.Permission{
		Name: "Can manage permissions",
	})
	if err != nil {
		return fmt.Errorf("could not create model: %s", err)
	}
	return nil
}
func (c *Seeder) seedRoles() error {
	var err error
	_, err = c.Client.RoleCreate(&models.Role{
		Name: "Teacher",
	})
	_, err = c.Client.RoleCreate(&models.Role{
		Name: "Admin",
	})
	if err != nil {
		return fmt.Errorf("could not create model: %s", err)
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

// seedSchools will add all schools to the DB
func (c *Seeder) seedSchools() error {
	schoolfile := c.files.Bytes("schools.csv")
	rows := []*SchoolCSV{}
	err := gocsv.UnmarshalBytes(schoolfile, &rows)
	if err != nil {
		return err
	}
	for _, row := range rows {
		_, err = c.Client.SchoolCreate(&models.School{Name: strings.Title(strings.ToLower(row.SchoolName))})
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 	req := fmt.Sprintf(`
		// 	mutation {
		// 		addSchool(school: {
		// 			code: %d,
		// 			schoolName: "%v",
		// 			street: "%v",
		// 			suburb: "%v",
		// 			state: "%v",
		// 			postcode: "%v",
		// 			postalStreet: "%v",
		// 			postalSuburb: "%v",
		// 			postalState: "%v",
		// 			postalPostcode: "%v",
		// 			latitude: %v,
		// 			longitude: %v,
		// 			courierCode: "%v",
		// 			phone: "%v",
		// 			educationRegion: "%v",
		// 			broadClassification: "%v",
		// 			classificationGroup: "%v",
		// 			kin: %d,
		// 			ppr: %d,
		// 			y01: %d,
		// 			y02: %d,
		// 			y03: %d,
		// 			y04: %d,
		// 			y05: %d,
		// 			y06: %d,
		// 			y07: %d,
		// 			y08: %d,
		// 			y09: %d,
		// 			y10: %d,
		// 			y11: %d,
		// 			y12: %d,
		// 			use: %d,
		// 			totalStudents: %d,
		// 		})
		// 	}
		// `,
		// 		row.Code,
		// 		strings.Title(strings.ToLower(row.SchoolName)),
		// 		strings.Title(strings.ToLower(row.Street)),
		// 		strings.Title(strings.ToLower(row.Suburb)),
		// 		strings.Title(strings.ToLower(row.State)),
		// 		strings.Title(strings.ToLower(row.Postcode)),
		// 		strings.Title(strings.ToLower(row.PostalStreet)),
		// 		strings.Title(strings.ToLower(row.PostalSuburb)),
		// 		strings.Title(strings.ToLower(row.PostalState)),
		// 		strings.Title(strings.ToLower(row.PostalPostcode)),
		// 		row.Latitude,
		// 		row.Longitude,
		// 		strings.Title(strings.ToLower(row.CourierCode)),
		// 		strings.Title(strings.ToLower(row.Phone)),
		// 		strings.Title(strings.ToLower(row.EducationRegion)),
		// 		strings.Title(strings.ToLower(row.BroadClassification)),
		// 		strings.Title(strings.ToLower(row.ClassificationGroup)),
		// 		row.KIN,
		// 		row.PPR,
		// 		row.Y01,
		// 		row.Y02,
		// 		row.Y03,
		// 		row.Y04,
		// 		row.Y05,
		// 		row.Y06,
		// 		row.Y07,
		// 		row.Y08,
		// 		row.Y09,
		// 		row.Y10,
		// 		row.Y11,
		// 		row.Y12,
		// 		row.USE,
		// 		row.TotalStudents,
		// 	)
		// 	q := handler.Query{
		// 		Query: req,
		// 	}

		// 	b, err := json.Marshal(q)
		// 	if err != nil {
		// 		return err
		// 	}

		// resp, err := http.Post("http://localhost:8000/graphql", "text/plain", bytes.NewReader(b))
		// if err != nil {
		// 	return err
		// }
		// if resp.StatusCode != 200 {
		// 	return fmt.Errorf("non 200 response: %d", resp.StatusCode)
		// }
		// respStruct := &Response{}

		// respBody, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	return err
		// }
		// defer resp.Body.Close()
		// json.Unmarshal(respBody, respStruct)
		// if len(respStruct.Errors) > 0 {
		// 	for _, errMsg := range respStruct.Errors {
		// 		log.Println(errMsg.Locations)
		// 		log.Println(errMsg.Message)
		// 		panic("Could not seed school")
		// 	}
		// }
	}

	return nil

}
