package model

import (
	"time"

	"github.com/google/uuid"
)
type User struct {
    ID               uuid.UUID `gorm:"type:varchar(36);primary_key"`
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
    ID          uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key"`
    FullName    string    `json:"full_name" gorm:"size:255"`
    PhoneNumber string    `json:"phone_number" gorm:"size:255"`
    Relationship string   `json:"relationship" gorm:"size:50"`
    MemberNumber string   `json:"member_number" gorm:"size:255"`
    Status      string    `json:"status"`
    UploadedDate time.Time `json:"uploaded_date"`
    Comments  string      `json:"comments" gorm:"type:text"`
    InsuranceID uuid.UUID `json:"insurance_id" gorm:"type:varchar(36)"`
    UserID      uuid.UUID `json:"user_id" gorm:"type:varchar(36)"`
    User        User      `json:"user"`
}

type Insurance struct {
    ID uuid.UUID                `json:"id" gorm:"type:varchar(36);primary_key"`
    InsuranceName string        `json:"insurance_name" gomr:"type:varchar(50);not null"`
    PhotoPath string            `json:"photo_path" gorm:"size:1024"`
    Dependants     []Dependant  `gorm:"foreignKey:InsuranceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
