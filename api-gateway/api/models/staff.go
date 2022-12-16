package models

import pbu "github.com/barber_shop/api-gateway/genproto/users_service"

type Staff struct {
	Id          string `json:"id,omitempty"`
	SalonId     string `json:"salon_id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	Type        string `json:"type"`
	ImageUrl    string `json:"image_url,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	DeletedAt   string `json:"deleted_at,omitempty"`
}

type StaffRequest struct {
	SalonId     string `json:"salon_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	UserName    string `json:"user_name"`
	Password    string `json:"passward"`
	Type        string `json:"type" binding:"required,oneof=superadmin user"`
	ImageURL    string `json:"image_url"`
}

type UpdateStaffRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	UserName    string `json:"user_name"`
	Password    string `json:"passward"`
	ImageURL    string `json:"image_url"`
}

type GetListStaffResponse struct {
	Staff []*Staff `json:"staff"`
	Count     int64       `json:"count"`
}

func ParsStaffToProtoStruct(staff *StaffRequest) *pbu.Staff {
	return &pbu.Staff{
		SalonId:     staff.SalonId,
		FirstName:   staff.FirstName,
		LastName:    staff.LastName,
		PhoneNumber: staff.PhoneNumber,
		Email:       staff.Email,
		UserName:    staff.UserName,
		Password:    staff.Password,
		Type:        staff.Type,
		ImageUrl:    staff.ImageURL,
	}
}

func ParsUpdateStaffToProtoStruct(staff *UpdateStaffRequest) *pbu.Staff {
	return &pbu.Staff{
		FirstName:   staff.FirstName,
		LastName:    staff.LastName,
		PhoneNumber: staff.PhoneNumber,
		Email:       staff.Email,
		UserName:    staff.UserName,
		Password:    staff.Password,
		ImageUrl:    staff.ImageURL,
	}
}

func ParsStaffFromProtoStruct(staff *pbu.Staff) *Staff {
	return &Staff{
		Id:          staff.Id,
		SalonId:     staff.SalonId,
		FirstName:   staff.FirstName,
		LastName:    staff.LastName,
		PhoneNumber: staff.PhoneNumber,
		Email:       staff.Email,
		UserName:    staff.UserName,
		Password:    staff.Password,
		Type:        staff.Type,
		ImageUrl:    staff.ImageUrl,
		CreatedAt:   staff.CreatedAt,
		UpdatedAt:   staff.UpdatedAt,
	}
}


func ParsListStaffFromProtoStruct(staff []*pbu.Staff) (resp []*Staff) {
	for _, s := range staff {
		st := Staff{
			Id:          s.Id,
			SalonId: 	 s.SalonId,
			FirstName:   s.FirstName,
			LastName:    s.LastName,
			PhoneNumber: s.PhoneNumber,
			Email:       s.Email,
			UserName:    s.UserName,
			Password:    s.Password,
			Type:        s.Type,
			ImageUrl:    s.ImageUrl,
			CreatedAt:   s.CreatedAt,
			UpdatedAt:   s.UpdatedAt,
		}
		resp = append(resp, &st)
	}
	return
}