package utilities

import (
	"encoding/json"
	"time"

	"github.com/DANCANKARANI/QVP/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
var db = database.ConnectDB()
type Audit struct{
    ID              uuid.UUID           `json:"id" gorm:"type:varchar(36);primary_key"`
    UserType        string              `json:"user_type" gorm:"type:varchar(255)"`
    UserID          uuid.UUID           `json:"user_id" gorm:"type:varchar(36); default:NULL"`
    Event           string              `json:"event" gorm:"type:varchar(255)"`
    AuditableType   string              `json:"auditable_type" gorm:"type:varchar(255)"`
    AuditableID     uuid.UUID           `json:"auditable_id" gorm:"type:varchar(255)"`
    OldValues       string              `json:"old_values" gorm:"type:text"`
    NewValues       string              `json:"new_values" gorm:"type:text"`
    Url             string              `json:"url" gorm:"type:text"`
    IpAddress       string              `json:"ip_address" gorm:"type:varchar(45)"`
    UserAgent       string              `json:"user_agent" gorm:"type:varchar(1024)"`
    Tags            string              `json:"tags" gorm:"type:varchar(255)"`
    CreatedAt       time.Time           `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt       time.Time           `json:"updated_at" gorm:"autoUpdateTime"`
    DeletedAt       gorm.DeletedAt      `json:"deleted_at" gorm:"index"`
}
/*
logs audits to the audit table
@params event
@params user_id, user_type, auditable_type, auditable_id, old_values, new_values
*/
func LogAudit(event string, user_id uuid.UUID, user_type string, auditable_type string, auditable_id uuid.UUID, oldValues interface{}, newValues interface{}, ctx *fiber.Ctx) error {
	oldValuesJSON, _ := json.Marshal(oldValues)
	newValuesJSON, _ := json.Marshal(newValues)

	audit := Audit{
		ID:            uuid.New(),
		UserType:      user_type,
		UserID:        user_id,
		Event:         event,
		AuditableType: auditable_type,
		AuditableID:   auditable_id,
		OldValues:     string(oldValuesJSON),
		NewValues:     string(newValuesJSON),
		Url:           ctx.OriginalURL(),
		IpAddress:     ctx.IP(),
		UserAgent:     string(ctx.Request().Header.UserAgent()),
		CreatedAt:     time.Now(),
	}

	return db.Create(&audit).Error
}
