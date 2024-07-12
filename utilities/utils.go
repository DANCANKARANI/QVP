package utilities

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GenerateImageUrl(c *fiber.Ctx) (string, error) {
	uploadsDir := "./uploads"
	file, err := c.FormFile("image")
	if err != nil {
		log.Fatal(err.Error())
		return "",errors.New("failed to upload the image")
	}
	fileName := fmt.Sprintf("%s_%s",time.Now(),file.Filename)
	savePath := filepath.Join(uploadsDir, SanitizeFileName(fileName))
	if err := c.SaveFile(file,savePath); err != nil {
		log.Fatal(err.Error())
		return "",errors.New("failed to save the to upload the image")
	}
	imageURL := fmt.Sprintf("http://localhost:3000/uploads/%s",fileName)
	fmt.Println(imageURL)
	return imageURL,nil
}
func SanitizeFileName(fileName string) string {
    // Define a regular expression to match invalid characters
    invalidChars := regexp.MustCompile(`[<>:"/\\|?*]+`)
    return invalidChars.ReplaceAllString(fileName, "_")
}

