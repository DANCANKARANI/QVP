package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type contextKey string

const fiberCtxKey contextKey = "fiberCtx"

// Function to inject Fiber context into GORM context
func InjectFiberCtxMiddleware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create a new context with the Fiber context stored in it
		ctx := context.WithValue(c.Context(), fiberCtxKey, c)
		// Use this new context in the GORM session
		tx := db.Session(&gorm.Session{Context: ctx})
		// Store the session with the modified context in Fiber context
		c.Locals("db", tx)
		return c.Next()
	}
}

// Hook after update to create audit log
func AfterUpdate(tx *gorm.DB) {
	// Extract Fiber context from GORM context
	if ctx := tx.Statement.Context; ctx != nil {
		if fiberCtx, ok := ctx.Value(fiberCtxKey).(*fiber.Ctx); ok {
			if err := CreateAuditLog(tx, fiberCtx, "UPDATE", tx.Statement.ReflectValue); err != nil {
				log.Printf("Error creating audit log: %v", err)
			}
		} else {
			log.Println("Fiber context not found in GORM context")
		}
	} else {
		log.Println("Context not found in GORM statement")
	}
}

// Function to create audit log
func CreateAuditLog(tx *gorm.DB, c *fiber.Ctx, event string, model reflect.Value) error {
	oldValues, _ := tx.Statement.Get("old_values")
	newValues, err := json.Marshal(model.Interface())
	if err != nil {
		return err
	}

	// Retrieve userID from Fiber context
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return fmt.Errorf("userID not found in context")
	}

	// Retrieve IP address from Fiber context
	ip, ok := c.Locals("ip_address").(string)
	if !ok {
		ip = "unknown"
	}

	auditableType := getTypeName(model.Interface())
	auditableID := getModelID(model)

	audit := Audit{
		UserType:      "user",
		UserID:        userID,
		Event:         event,
		AuditableType: auditableType,
		AuditableID:   auditableID,
		OldValues:     oldValues.(string),
		NewValues:     string(newValues),
		Url:           c.OriginalURL(),
		IpAddress:     ip,
		UserAgent:     c.Get("User-Agent"),
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
