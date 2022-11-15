package structures

import "time"

type Salon struct {
	ID          string
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Gender      string
	Password    string
	ImageURL    string
	CreatedAT   time.Time
	UpdatedAT   time.Time
	DeletedAT   time.Time
}

type Staff struct {
	ID                 string
	Name               string
	PhoneNumber        string
	Email              string
	Rating             string
	Address            string
	Latitude           string
	Longitude          string
	StartTime          time.Time
	EndTime            time.Time
	ImageURL           string
	PersonalWorkImages []PersonalWorkImage
	CreatedAT          time.Time
	UpdatedAT          time.Time
	DeletedAT          time.Time
}

type User struct {
	ID          string
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Gender      string
	Password    string
	ImageURL    string
	CreatedAT   time.Time
	UpdatedAT   time.Time
	DeletedAT   time.Time
}

type PersonalWorkImage struct {
	ID             string
	StaffID        string
	ImageURl       string
	SequenceNumber int8
	CreatedAT      time.Time
}
