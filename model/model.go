package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
    ID               uuid.UUID `gorm:"type:varchar(36);primary_key"`
    FullName         string    `json:"full_name" gorm:"size:255"`
    Email            string    `json:"email" gorm:"size:255; unique"`
    PhoneNumber      string    `json:"phone_number" gorm:"size:255;unique"`
    CountryCode      string    `json:"country_code" gorm:"size:10"`
    Password         string    `json:"password" gorm:"size:255"`
    ConfirmPassword  string    `json:"confirm_password" gorm:"size:255"`
    ResetCode        string    `json:"reset_code"`
    CodeExpirationTime time.Time `json:"code_expiration_time"`
    ProfilePhotoPath string    `json:"profile_photo_path" gorm:"size:1024"`
    Dependants       []Dependant `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    Payment          []Payment  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
    Insurance  Insurance `json:"insurance"`
}

type Insurance struct {
    ID uuid.UUID                `json:"id" gorm:"type:varchar(36);primary_key"`
    InsuranceName string        `json:"insurance_name" gomr:"type:varchar(50);not null"`
    PhotoPath string            `json:"photo_path" gorm:"size:1024"`
    Description string          `json:"description" gorm:"type:text"`
    Dependants     []Dependant  `gorm:"foreignKey:InsuranceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Payment struct{
    ID      uuid.UUID           `json:"id" gorm:"type:varchar(36);primary_key"`
    Amount  float64             `json:"amount" gorm:"type:decimal(13,4)"`
    Narration string            `json:"narration" gorm:"type:varchar(255)"`
    Reference string            `json:"reference" gorm:"type:varchar(255)"`
    ResponseDescription string  `json:"response_description" gorm:"type:text"`
    UserID  uuid.UUID           `json:"user_id" gorm:"type:varchar(36)"`
    PaymentMethodID uuid.UUID   `json:"payment_method_id" gorm:"type:varchar(36)"`
    User User                   `json:"user"`
    PaymentMethod PaymentMethod `json:"payment_method"`
}
type PaymentMethod struct {
    ID      uuid.UUID     `json:"id" gorm:"type:varchar(36);primary_key"`
    Title   string        `json:"title" gorm:"type:varchar(36); unique"`
    IconUrl string        `json:"icon_url" gorm:"type:varchar(1024)"`
    Payments []Payment    `gorm:"foreignKey:PaymentMethodID;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;reference:ID"`
}


