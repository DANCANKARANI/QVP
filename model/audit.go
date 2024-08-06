package model

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/google/uuid"
	"gorm.io/gorm"
)
func LogChanges(db *gorm.DB, event string, model interface{}) {
    modelType := reflect.TypeOf(model)
    modelValue := reflect.ValueOf(model)

    // Handle pointers to models
    if modelValue.Kind() == reflect.Ptr {
        modelValue = modelValue.Elem()
        modelType = modelType.Elem()
    }

    idField := modelValue.FieldByName("ID")
    if !idField.IsValid() {
        log.Fatal("ID field not found in model")
    }
    idValue := idField.Interface()

    // Create a new transaction for reading
    tx := db.Begin()
    if err := tx.Error; err != nil {
        log.Fatalf("Failed to begin transaction: %v", err)
    }
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    oldModel := reflect.New(modelType).Interface()
    err := tx.Where("id = ?", idValue).First(oldModel).Error
    if err != nil {
        log.Printf("Error retrieving old model: %v", err)
        tx.Rollback()
        return
    }

    tx.Commit() // End read transaction

    oldValues, _ := json.Marshal(oldModel)
    newValues, _ := json.Marshal(model)

    audit := Audit{
        UserType:       "system",
        UserID:         uuid.Nil,
        Event:          event,
        AuditableType:  modelType.Name(),
        AuditableID:    idValue.(uuid.UUID),
        OldValues:      string(oldValues),
        NewValues:      string(newValues),
    }

    // Start a new transaction for writing
    tx = db.Begin()
    if err := tx.Error; err != nil {
        log.Fatalf("Failed to begin transaction: %v", err)
    }

    if err := tx.Create(&audit).Error; err != nil {
        log.Fatalf("Failed to create audit record: %v", err)
    }

    tx.Commit() // Commit write transaction
}


// registers hooks
func RegisterHooks() {
	// db.Callback().Create().Before("gorm:before_update").Register("audit:before_update", func(tx *gorm.DB) {
	// 	log.Println("before update")
	// })
	// db.Callback().Create().After("gorm:after_create").Register("audit:after_create", func(tx *gorm.DB) {
	// 	LogChanges(tx, "create")
	// })
	// db.Callback().Update().After("gorm:after_update").Register("audit:after_update", func(tx *gorm.DB) {
	// 	LogChanges(tx, "update")
	// })
	// db.Callback().Delete().After("gorm:after_delete").Register("audit:after_delete", func(tx *gorm.DB) {
	// 	LogChanges(tx, "delete")
	// })
}
