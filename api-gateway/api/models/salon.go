package models

import pbu "github.com/barber_shop/api-gateway/genproto/users_service"

type SalonRequest struct {
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email,omitempty"`
	Rating      int32  `json:"rating,omitempty"`
	Address     string `json:"address,omitempty"`
	Latitude    string `json:"latitude,omitempty"`
	Longitude   string `json:"longitude,omitempty"`
	StartTime   string `json:"start_time,omitempty"`
	EndTime     string `json:"end_time,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
}

type Salon struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email,omitempty"`
	Rating      int32  `json:"rating,omitempty"`
	Address     string `json:"address,omitempty"`
	Latitude    string `json:"latitude,omitempty"`
	Longitude   string `json:"longitude,omitempty"`
	StartTime   string `json:"start_time,omitempty"`
	EndTime     string `json:"end_time,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type GetListSalonsResponse struct {
	Salons []*Salon `json:"salons"`
	Count     int64       `json:"count"`
}

func ParsSalonToProtoStruct(salon *SalonRequest) *pbu.Salon {
	return &pbu.Salon{
		Name:        salon.Name,
		PhoneNumber: salon.PhoneNumber,
		Email:       salon.Email,
		Rating:      salon.Rating,
		Address:     salon.Address,
		Latitude:    salon.Latitude,
		Longitude:   salon.Longitude,
		StartTime:   salon.StartTime,
		EndTime:     salon.EndTime,
		ImageUrl:    salon.ImageUrl,
	}
}

func ParsSalonFromProtoStruct(salon *pbu.Salon) *Salon {
	return &Salon{
		Id:          salon.Id,
		Name:        salon.Name,
		PhoneNumber: salon.PhoneNumber,
		Email:       salon.Email,
		Rating:      salon.Rating,
		Address:     salon.Address,
		Latitude:    salon.Latitude,
		Longitude:   salon.Longitude,
		StartTime:   salon.StartTime,
		EndTime:     salon.EndTime,
		ImageUrl:    salon.ImageUrl,
		CreatedAt:   salon.CreatedAt,
		UpdatedAt:   salon.UpdatedAt,
	}
}

func ParsListSalonsFromProtoStruct(salons []*pbu.Salon) (rSalons []*Salon) {
	for _, s := range salons {
		rSalon := Salon{
			Id:          s.Id,
			Name:        s.Name,
			PhoneNumber: s.PhoneNumber,
			Email:       s.Email,
			Rating:      s.Rating,
			Address:     s.Address,
			Latitude:    s.Latitude,
			Longitude:   s.Longitude,
			StartTime:   s.StartTime,
			EndTime:     s.EndTime,
			ImageUrl:    s.ImageUrl,
			CreatedAt:   s.CreatedAt,
			UpdatedAt:   s.UpdatedAt,
		}
		rSalons = append(rSalons, &rSalon)
	}
	return
}
