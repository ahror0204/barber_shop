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

type GetSalonsParams struct {
	Page   int
	Limit  int
	Search string
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
	ID          string     `db:"id"`
	FirstName   string     `db:"first_name"`
	LastName    string     `db:"last_name"`
	PhoneNumber string     `db:"phone_number"`
	Email       string     `db:"email"`
	UserName    string     `db:"user_name"`
	Password    string     `db:"passward"`
	Gender      string     `db:"gender"`
	ImageURL    string     `db:"image_url"`
	CreatedAT   time.Time  `db:"created_at"`
	UpdatedAT   *time.Time `db:"updated_at"`
	DeletedAT   time.Time  `db:"deleted_at"`
}

type PersonalWorkImage struct {
	ID             string
	StaffID        string
	ImageURl       string
	SequenceNumber int8
	CreatedAT      time.Time
}

type GetUsersParams struct {
	Page   int
	Limit  int
	Search string
}

type AllUsers struct {
	Users []User
	Count int64
}
