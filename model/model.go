package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)
type User struct {
    ID               uuid.UUID `gorm:"type:char(36);primary_key"`
    FullName         string    `json:"full_name" gorm:"size:255"`
    Email            string    `json:"email" gorm:"size:255"`
    PhoneNumber      string    `json:"phone_number" gorm:"size:255"`
    CountryCode      string    `json:"country_code" gorm:"size:10"`
    Password         string    `json:"password" gorm:"size:255"`
    ConfirmPassword  string    `json:"confirm_password" gorm:"size:255"`
    ResetCode        string    `json:"reset_code"`
    CodeExpirationTime time.Time `json:"code_expiration_time"`
    ProfilePhotoPath string    `json:"profile_photo_path" gorm:"size:1024"`
    Dependants       []Dependant `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Dependant struct {
    ID          uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
    FullName    string    `json:"full_name" gorm:"size:255"`
    PhoneNumber string    `json:"phone_number" gorm:"size:255"`
    Relationship string   `json:"relationship" gorm:"size:50"`
    MemberNumber string   `json:"member_number" gorm:"size:255"`
    Status      string    `json:"status"`
    InsuranceID uuid.UUID `json:"insurance_id" gorm:"type:char(36)"`
    UserID      uuid.UUID `json:"user_id" gorm:"type:char(36)"`
    Insurance   Insurance `gorm:"foreignKey:InsuranceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    User        User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Insurance struct {
	gorm.Model
    ID uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
}
