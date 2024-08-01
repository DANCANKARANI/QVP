package model

import (
	"encoding/json"
	"log"
	"reflect"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Middleware to set userID in GORM context
func SetUserIDMiddleware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("user_id").(*uuid.UUID)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: Invalid user ID")
		}
		log.Println("SetUserIDMiddleware - userID:", userID)
		id := *userID

		// Add userID to GORM session
		tx := db.Session(&gorm.Session{Context: c.Context()}).Set("userID", id)
		
		c.Locals("db", tx)
		return c.Next()
	}
}

// Hook after update to create audit log
func AfterUpdate(tx *gorm.DB) {
	if err := CreateAuditLog(tx, "UPDATE", tx.Statement.ReflectValue); err != nil {
		log.Printf("Error creating audit log: %v", err)
	}
}

// Function to create audit log
func CreateAuditLog(tx *gorm.DB, event string, model reflect.Value) error {
	oldValues, _ := tx.Statement.Get("old_values")
	newValues, err := json.Marshal(model.Interface())
	if err != nil {
		return err
	}

	userID, _ := tx.Get("userID")
	id:=userID.(uuid.UUID)
	auditableType := getTypeName(model.Interface())
	auditableID := getModelID(model)

	audit := Audit{
		UserType:      "user",
		UserID:        id,
		Event:         event,
		AuditableType: auditableType,
		AuditableID:   auditableID,
		OldValues:     oldValues.(string),
		NewValues:     string(newValues),
		Url:           "", // Set to the actual URL from context if needed
		IpAddress:     "", // Set to the actual IP address from context if needed
		UserAgent:     "", // Set to the actual User-Agent from context if needed
		Tags:          "",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return tx.Create(&audit).Error
}


// Helper function to get model ID
func getModelID(model reflect.Value) uuid.UUID {
	if model.Kind() == reflect.Ptr {
		model = model.Elem()
	}
	idField := model.FieldByName("ID")
	if idField.IsValid() && idField.CanInterface() {
		switch v := idField.Interface().(type) {
		case uuid.UUID:
			return v
		case string:
			id, err := uuid.Parse(v)
			if err != nil {
				log.Printf("Error parsing string ID to UUID: %v", err)
				return uuid.Nil
			}
			return id
		default:
			log.Printf("Unsupported ID type: %T", v)
			return uuid.Nil
		}
	}
	return uuid.Nil
}

// Helper function to get type name of model
func getTypeName(model interface{}) string {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}

// Function to register GORM hooks
func RegisterHooks() {
	db.Callback().Update().After("gorm:after_update").Register("after_update_audit", AfterUpdate)
}
