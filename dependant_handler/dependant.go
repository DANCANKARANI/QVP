package dependant_handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"main.go/database"
	"main.go/model"
	"main.go/utilities"
)
var db = database.ConnectDB()
var country_code = "KE"
type ResponseDependant struct{
	ID		uuid.UUID			`json:"id"`
	FullName 	string			`json:"full_name"`
	PhoneNumber string 			`json:"phone_number"`
	Relationship string 		`json:"relationship"`
	MemberNumber string 		`json:"member_number"`
	Status 		string 			`json:"status"`
	InsuranceID uuid.UUID		`json:"insurance_id"`
	UserID	uuid.UUID			`json:"user_id"`
	User 	model.User			
}

func RegisterDependantAccount(c *fiber.Ctx) error {
	db.AutoMigrate(&model.Dependant{})
	ID:=uuid.New()
	body :=model.Dependant{}
	if err := c.BodyParser(&body); err !=nil {
		return c.JSON(fiber.Map{"error":err.Error()})
	}
	//check if the dependant exists
	dependantExist,_,_ := model.DependantExist(c,body.PhoneNumber)
	if dependantExist{
		return utilities.ShowError(c,"User with this phone number already exists",fiber.StatusConflict)
	}
	//validate phone number
	_,err := utilities.ValidatePhoneNumber(body.PhoneNumber,country_code)
	if err !=nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusAccepted)
	}
	user_id :=c.Locals("user_id")
	id,ok := user_id.(*uuid.UUID)
	if !ok{
		return utilities.ShowError(c,"failed conversion",fiber.StatusInternalServerError)
	}
	user_id=*id
	result :=db.Create(&model.Dependant{
		ID: ID,
		FullName: body.FullName,
		PhoneNumber: body.PhoneNumber,
		Relationship: body.Relationship,
		MemberNumber: body.MemberNumber,
		Status: body.Status,
		UploadedDate: time.Now(),
		UserID:user_id.(uuid.UUID),
	})
	if result.Error != nil {
		return utilities.ShowError(c,"failed to register the dependant",fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"dependant registered successfully",fiber.StatusOK)
}

//get dependants handler
func GetDependantsHandler(c *fiber.Ctx)error{
	//check if the dependant exist
	user_id :=c.Locals("user_id")
	id,ok := user_id.(*uuid.UUID)
	if !ok{
		return utilities.ShowError(c,"failed conversion",fiber.StatusInternalServerError)
	}
	user_id=*id
	existingDependants,err := model.GetAllDependants(c,user_id.(uuid.UUID))
	if err != nil{
		return utilities.ShowError(c,"The user has no dependants",fiber.StatusInternalServerError)
	}
	//response
	var response []ResponseDependant
    for _, dependant := range existingDependants {
        resUser := model.User{
            FullName:    dependant.User.FullName,
            PhoneNumber: dependant.User.PhoneNumber,
            Email:      dependant.User.Email,
        }

        responseDependant := ResponseDependant{
            ID:           dependant.ID,
            FullName:     dependant.FullName,
            PhoneNumber:  dependant.PhoneNumber,
            Status:       dependant.Status,
            Relationship: dependant.Relationship,
            UserID:       dependant.UserID,
            MemberNumber: dependant.MemberNumber,
            InsuranceID:  dependant.InsuranceID,
            User:         resUser,
        }
		
        response = append(response, responseDependant)
    }
	did:=c.Locals("d_id")
	return utilities.ShowSuccess(c, did, fiber.StatusOK, response)
}

//get user id to set in the url
func GetDependantID(c *fiber.Ctx)error{
	dependant_id,err :=model.GetDependantID(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}

	return utilities.ShowMessage(c,dependant_id.String(),fiber.StatusOK)
}