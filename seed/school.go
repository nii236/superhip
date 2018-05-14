package main

// SchoolCSV is a row in the school CSV
type SchoolCSV struct {
	Code                int     `csv:"Code"`
	SchoolName          string  `csv:"School Name"`
	Street              string  `csv:"Street"`
	Suburb              string  `csv:"Suburb"`
	State               string  `csv:"State"`
	Postcode            string  `csv:"Postcode"`
	PostalStreet        string  `csv:"Postal Street"`
	PostalSuburb        string  `csv:"Postal Suburb"`
	PostalState         string  `csv:"Postal State"`
	PostalPostcode      string  `csv:"Postal Postcode"`
	Latitude            float64 `csv:"Latitude"`
	Longitude           float64 `csv:"Longitude"`
	CourierCode         string  `csv:"Courier Code"`
	Phone               string  `csv:"Phone"`
	EducationRegion     string  `csv:"Education Region"`
	BroadClassification string  `csv:"Broad Classification"`
	ClassificationGroup string  `csv:"Classification Group"`
	KIN                 int     `csv:"KIN"`
	PPR                 int     `csv:"PPR"`
	Y01                 int     `csv:"Y01"`
	Y02                 int     `csv:"Y02"`
	Y03                 int     `csv:"Y03"`
	Y04                 int     `csv:"Y04"`
	Y05                 int     `csv:"Y05"`
	Y06                 int     `csv:"Y06"`
	Y07                 int     `csv:"Y07"`
	Y08                 int     `csv:"Y08"`
	Y09                 int     `csv:"Y09"`
	Y10                 int     `csv:"Y10"`
	Y11                 int     `csv:"Y11"`
	Y12                 int     `csv:"Y12"`
	USE                 int     `csv:"USE"`
	TotalStudents       int     `csv:"Total Students"`
}
