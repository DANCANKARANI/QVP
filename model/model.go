package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type User struct{
	ID      uuid.UUID 				`gorm:"type:char(36);primary_key"`
	FullName 		string 			`json:"full_name" gorm:"size:255"`
	Email 		string 				`json:"email" gorm:"size:255"`
	PhoneNumber string 				`json:"phone_number" gorm:"size:255"`
	CountryCode string 				`json:"country_code" gorm:"size:10"`
	Password 	string 				`json:"password" gorm:"size:255"`
	ConfirmPassword string 			`json:"confirm_password" gorm:"size:255"`
}
type Prescription struct{
	gorm.Model
	ID			uuid.UUID			`json:"id" gorm:"primaryKey;autoIncrement:true"`
	//foreignKey
	UserId 		uuid.UUID 			`json:"user_id"`
	User 		User   				`gorm:"foreignKey:UserId; reference: ID"`
	BranchID	uuid.UUID			`json:"branch_id"`
	QuoteNumber string				`json:"quote_number" sql:"type:varchar(255)"`
	SubTotal	decimal.Decimal		`json:"sub_total" sql:"type:decimal(8,2);"`
	Vat 		decimal.Decimal		`json:"vat" sql:"type:decimal(8,2)"`
	Total	    decimal.Decimal		`json:"total" sql:"type:decimal(8,2)"`
	DeliveryDetails string			`json:"delivery_details" sql:"type:text"`
	UserValidatedAt time.Time		`json:"user_validated_at" sql:"type:timestamp"`
	//foreignKey
	UserValidatedBy uuid.UUID		`json:"user_validated_by"`
	UserApprovedAt time.Time		`json:"user_approved_at" sql:"type:timestamp"`
	//foreignKey
	UserApprovedBy uuid.UUID		`json:"user_approved_by"`
	AdminValidateAt time.Time		`json:"admin_validate_at" sql:"type:timestamp"`
	//foreignKey
	AdminValidateBy uuid.UUID		`json:"admin_validate_by"`
	AdminApprovedAt time.Time		`json:"admin_approved_at" sql:"type:timestamp"`
	//foreignKey
	AdminApprovedBy uuid.UUID		`json:"admin_approved_by"`
	DeliveredAt time.Time			`json:"delivered_at" sql:"type:timestamp"`
	//foreignKey
	DeliveredBy uuid.UUID			`json:"delivered_by"`

}

type Admin struct {
	gorm.Model
	ID uuid.UUID					`json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string						`json:"name" sql:"type:varchar(255)"`
	Email string					`json:"email" sql:"type:varchar(255)"`
	Phone	string  				`json:"phone" sql:"type:varchar(255)"`
	EmailVerifiedAt time.Time 		`json:"email_verified_at"`
	Password	string 				`json:"password" sql:"type:varchar(255)"`
	RememberToken string 			`json:"remember_token" sql:"type:varchar(100)"`
	//foreignKey
	CurrentTeamId	uuid.UUID		`json:"current_team_id"`
	ProfilePhotoPath string 		`json:"profile_photo_path" sql:"type:varchar(2048)"`
}

type Branch struct {
	gorm.Model
	ID uuid.UUID					`json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string						`json:"name" sql:"type:varchar(255); not null"`
	Location string					`json:"location" sql:"type:varchar(255)"`
	Description string				`json:"description" sql:"type:text"`
}

type Team struct {
	gorm.Model
	ID uuid.UUID					`json:"id" gorm:"primaryKey;autoIncrement:true"`
	//foreignKey
	UserID uuid.UUID				`json:"user_id"`
	Name	string					`json:"name" sql:"type:varchar(255); not null"`
	PersonalTeam int64				`json:"personal_team" sql:"type:tinyint(1)"`
}
type Image struct{
	gorm.Model
	ID uuid.UUID			`json:"id" gorm:"primaryKey;autoIncrement:true"`
	File string				`json:"file" sql:"type:varchar(255); not null"`
	OriginalName string		`json:"original_name" sql:"type:varchar(255); not null"`
	Thumb string 			`json:"thumb" sql:"type:varchar(255)"`
	Path  string			`json:"path" sql:"type:varchar(255); not null"`
	Type  string 			`json:"type" sql:"type:varchar(255); not null"`
	FileType string 		`json:"file_type" sql:"type:varchar(255); not null"`
	ImageableType string 	`json:"imageable_type" sql:"type:varchar(255)"`
	//foreignKey
	ImageableID	uint64		`json:"imageable_id"`
}

type QuoteDetail struct {
	gorm.Association
	ID uuid.UUID				`json:"id" gorm:"primaryKey;autoIncrement:true"`
	//foreignKey
	PrescriptionId uuid.UUID	`json:"prescription_id"`
	Prescription Prescription 	`gorm:"foreignKey:PrescriptionId; reference: ID"`
	Description string			`json:"description" sql:"type:varchar(255); not null"`
	Unit 	float64				`json:"unit" sql:"type:double; not null"`
	Quantity float64			`json:"quantity" sql:"type:double; not null"`
	Measure  string				`json:"measure" sql:"type:varchar(255); not null"`
	Price	float64				`json:"price" sql:"type:double; not null"`
	Discount float64 			`json:"discount" sql:"type:double; not null"`
	Vat 	float64 			`json:"vat" sql:"type:double; not null"`
	Total  float64 				`json:"total" sql:"type:double; not null"`
}
type TeamInvitation struct {
	gorm.Model
	ID uuid.UUID 			`json:"id" gorm:"primaryKey; autoIncrement:true"`
	//foreignKey
	TeamId uuid.UUID 		`json:"team_id"`
	Team Team				`gorm:"foreignKey:TeamID; reference:ID"`
	Email string 			`json:"email" sql:"type:varchar(255)"`
	Role string 			`json:"role" sql:"type:varchar(255)"`
}

type Payments struct {
	gorm.Model
	ID uuid.UUID 				`json:"id" gorm:"primaryKey;autoIncrement:true"`
	//foreignKey
	UserId uuid.UUID 			`json:"user_id"`
	User User					`gorm:"foreignKey:UserId; reference: ID"`
	//foreignKey
	MethodID uuid.UUID 			`json:"method_id"`
	Method PaymentMethod 		`gorm:"foreignKey:MethodId; reference: ID"`
	PrescriptionId uuid.UUID 	`json:"prescription_id"`
	Prescription Prescription 	`gorm:"foreignKey:PrescriptionId; reference: ID"`
	//foreignKey
	ReceivedBy uuid.UUID 		`json:"received_by"`
	Amount float64 				`json:"amount" sql:"type:decimal(13,4)"`
	Reference string 			`json:"reference" sql:"type:varchar(255)"`
	Narration string 			`json:"narration" sql:"type:varchar(255)"`
	ResponseDescription string 	`json:"response_description" sql:"type:varchar(255)"`
	TransactionTime	time.Time 	`json:"transaction_time" sql:"type:timestamp"`
}


type PaymentMethod struct {
	gorm.Model
	ID uuid.UUID 	`json:"id" gorm:"primaryKey; autoIncrement:true; not null"`
	Title string 	`json:"title" sql:"type:varchar(255)"`
	Icon string 	`json:"icon" sql:"type:text"`
}

type TeamUser struct {
	gorm.Model
	ID uuid.UUID 		`json:"id" gorm:"primaryKey; autoIncrement:true"`
	TeamID 	uuid.UUID 	`json:"team_id"`
	UserID uuid.UUID 	`json:"user_id"`
	Role string 		`json:"role" sql:"type:varchar(255)"`
	Team Team 			`gorm:"foreignKey:TeamID; reference:ID; not null"`
	User User 			`gorm:"foreignKey:UserID; reference:ID; not null"`
}

type Insurance struct {
	gorm.Model
	ID uuid.UUID 		`json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string			`json:"name" sql:"type:varchar(255)"`
	PhotoPath string	`json:"photoPath" sql:"type:varchar(2048)"`
	Description string 	`json:"description" sql:"type:text"`
}

type InsuranceUser struct {
	gorm.Model
	ID uuid.UUID			`json:"id" gorm:"primaryKey;autoIncrement:true"`
	InsuranceID uuid.UUID 	`json:"insurance_id"`
	UserID uuid.UUID		`json:"user_id"`
}

type Insurancer struct {
	gorm.Model
	ID uuid.UUID 				`json:"id" gorm:"primaryKey;autoIncrement:true"`
	InsuranceId uuid.UUID 		`json:"insurance_id"`
	Insurance Insurance			`gorm:"foreignKey:InsuranceId; reference:ID"`
	Name string 				`json:"name" sql:"type:varchar(255); not null"`
	Email string 				`json:"email" sql:"type:varchar(255); not null"`
	Phone string 				`json:"phone" sql:"type:varchar(255); not null"`
	EmailVerifiedAt time.Time 	`json:"email_verified_at" `
	Password string 			`json:"password" sql:"type:varchar(255)"`
	RememberToken string  		`json:"remember_token" sql:"type:varchar(100)"`
	CurrentTeamId	uuid.UUID 	`json:"current_team_id"`
	Team Team 					`gorm:"foreignKey: CurrentTeamId; reference: ID"`
	ProfilePhotoPath	string 	`json:"profile_photo_path" sql:"type:varchar(2048)"`
}

type Dependant struct{

}