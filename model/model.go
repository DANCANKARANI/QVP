package model

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
    ID               uuid.UUID      `gorm:"type:varchar(36);primary_key"`
    FullName         string         `json:"full_name" gorm:"size:255"`
    Email            string         `json:"email" gorm:"size:255; unique"`
    PhoneNumber      string         `json:"phone_number" gorm:"size:255;unique"`
    CountryCode      string         `json:"country_code" gorm:"size:10"`
    Password         string         `json:"password" gorm:"size:255"`
    ConfirmPassword  string         `json:"confirm_password" gorm:"size:255"`
    ResetCode        string         `json:"reset_code"`
    CodeExpirationTime time.Time    `json:"code_expiration_time"`
    ProfilePhotoPath string         `json:"profile_photo_path" gorm:"size:1024"`
    CreatedAt        time.Time      `json:"created_at" gorm:"autoCreateTime"`
    DeletedAt        time.Time      `json:"deleted_at" gorm:"autoCreateTime"`
    UpdatedAt        time.Time      `json:"updated_at" gorm:"autoCreateTime"`
    ImageID          uuid.UUID      `json:"image_id" gorm:"type:uuid"`
    Image            Image          `gorm:"foreignKey:ImageID;references:ID;constraint:onUpdate:CASCADE,onDelete:SET NULL;"`
    Dependants       []Dependant    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    Payment          []Payment      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    Notification     []Notification `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
    CreatedAt time.Time         `json:"created_at" gorm:"autoCreateTime"`
    DeletedAt time.Time         `json:"deleted_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time         `json:"updated_at" gorm:"autoCreateTime"`
    ReceivedBy uuid.UUID        `json:"received_by"`
    Dependants     []Dependant  `gorm:"foreignKey:InsuranceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Payment struct{
    ID      uuid.UUID           `json:"id" gorm:"type:varchar(36);primary_key"`
    Amount  float64             `json:"amount" gorm:"type:decimal(13,4)"`
    Narration string            `json:"narration" gorm:"type:varchar(255)"`
    Reference string            `json:"reference" gorm:"type:varchar(255)"`
    ResponseDescription string  `json:"response_description" gorm:"type:text"`
    TransactionTime time.Time   `json:"transaction_time" gorm:"autoCreateTime"`
    CreatedAt time.Time         `json:"created_at" gorm:"autoCreateTime"`
    DeletedAt time.Time         `json:"deleted_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time         `json:"updated_at" gorm:"autoCreateTime"`
    ReceivedBy uuid.UUID        `json:"received_by"`
    UserID  uuid.UUID           `json:"user_id" gorm:"type:varchar(36)"`
    PaymentMethodID uuid.UUID   `json:"payment_method_id" gorm:"type:varchar(36)"`
    User User                   `json:"user"`
    PaymentMethod PaymentMethod `json:"payment_method"`
}
type PaymentMethod struct {
    ID      uuid.UUID           `json:"id" gorm:"type:varchar(36);primary_key"`
    Title   string              `json:"title" gorm:"type:varchar(36); unique"`
    IconUrl string              `json:"icon_url" gorm:"type:varchar(1024)"`
    CreatedAt time.Time         `json:"created_at" gorm:"autoCreateTime"`
    DeletedAt time.Time         `json:"deleted_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time         `json:"updated_at" gorm:"autoCreateTime"`
    ReceivedBy uuid.UUID        `json:"received_by"`
    Payments []Payment          `gorm:"foreignKey:PaymentMethodID;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;reference:ID"`
}
type Notification struct {
    ID        uuid.UUID         `json:"id" gorm:"type:varchar(36);primary_key"`
    UserID    uuid.UUID         `json:"user_id" gorm:"type:varchar(36)"`
    Type      string            `json:"type" gorm:"type:varchar(36)"`
    Message   string            `json:"message" gorm:"type:varchar(256)"`
    IsRead    bool              `json:"is_read"`
    CreatedAt time.Time         `json:"created_at" gorm:"autoCreateTime"`
    DeletedAt time.Time         `json:"deleted_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time         `json:"updated_at" gorm:"autoCreateTime"`
    ReceivedBy uuid.UUID        `json:"received_by"`
}



type Prescription struct {
    ID               uuid.UUID      `json:"id" gorm:"type:varchar(36);primary_key"`
    QuoteNumber      string         `json:"quote_number" gorm:"size:255"`
    SubTotal         float64        `json:"sub_total" gorm:"type:decimal(8,2)"`
    VAT              float64        `json:"vat" gorm:"type:decimal(8,2)"`
    Total            float64        `json:"total" gorm:"type:decimal(8,2)"`
    CreatedAt        time.Time      `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt        time.Time      `json:"updated_at" gorm:"autoCreateTime"`
    DeletedAt        *time.Time     `json:"deleted_at" gorm:"index"`
    DeliveryDetails  string         `json:"delivery_details" gorm:"type:text"`
    UserValidatedAt  *time.Time     `json:"user_validated_at" gorm:""`
    UserValidatedBy  uuid.UUID      `json:"user_validated_by" gorm:""`
    UserApprovedAt   *time.Time     `json:"user_approved_at" gorm:""`
    UserApprovedBy   uuid.UUID      `json:"user_approved_by" gorm:""`
    AdminValidateAt  *time.Time     `json:"admin_validate_at" gorm:""`
    AdminValidateBy  uuid.UUID      `json:"admin_validate_by" gorm:""`
    AdminApprovedAt  *time.Time     `json:"Admin_approved_at" gorm:""`
    AdminApprovedBy  uuid.UUID      `json:"admin_approved_by" gorm:""`
    DeliveredAt      *time.Time     `json:"delivered_at" gorm:""`
    DeliveredBy      uuid.UUID      `json:"delivered_by" gorm:""`
    UserID           uuid.UUID      `gorm:"not null"`
    BranchID         uuid.UUID      `gorm:"not null"`

    // Relationships
    User            User            `gorm:"foreignKey:UserID"`
    Branch          Branch          `gorm:"foreignKey:BranchID"`
}
type Branch struct {
    ID uuid.UUID                    `json:"id" gorm:"type:varchar(36);primary_key"`
    Name string                     `json:"name" gorm:"type:varchar(255)"`
    Location string                 `json:"location" gorm:"type:varchar(255)"`
    Description string              `json:"description" gorm:"type:text"`
    CreatedAt   time.Time           `json:"created_at" gorm:"type:time; autoCreateTime"`
    UpdatedAt   time.Time           `json:"updated_at" gorm:"type:time; autoCreateTime"`
}

type Image struct {
    ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
    OriginalName string    `json:"original_name" gorm:"type:varchar(255)"`
    Path         string    `json:"path" gorm:"type:varchar(255)"`
    Thumbnail    string    `json:"thumbnail" gorm:"type:varchar(255)"`
    Type         string    `json:"type" gorm:"type:varchar(255)"`
    CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type Admin struct {
    ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
    Name     string    `json:"name" gorm:"type:varchar(255)"`
    ImageID  uuid.UUID `json:"image_id" gorm:"type:uuid"`
    Image    Image     `gorm:"foreignKey:ImageID;references:ID;constraint:onUpdate:CASCADE,onDelete:SET NULL;"`
}
