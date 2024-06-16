package user_handler

import (

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"main.go/model"
	"main.go/utilities"
)
type ResponseDependant struct{
	ID		uuid.UUID			`json:"id"`
	FullName 	string			`json:"full_name"`
	PhoneNumber string 			`json:"phone_number"`
	Relationship string 		`json:"relationship"`
	MemberNumber string 		`json:"member_number"`
	Status 		string 			`json:"status"`
	InsuaranceID uuid.UUID		`json:"insurance_id"`
	UserID	uuid.UUID			`json:"user_id"`
					
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
		return utilities.ShowError(c,"failed to get dependants",fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"dependants retrieved successsfully",fiber.StatusOK,
	ResponseDependant{
		ID: existingDependants.ID,
		FullName: existingDependants.FullName,
		PhoneNumber: existingDependants.PhoneNumber,
		Status: existingDependants.Status,
		Relationship: existingDependants.Relationship,
		UserID: existingDependants.UserID,
		MemberNumber: existingDependants.MemberNumber,
		InsuaranceID: existingDependants.InsuranceID,
	})
}