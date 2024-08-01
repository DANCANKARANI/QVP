package audit

// import (
// 	"github.com/DANCANKARANI/QVP/model"
// 	"github.com/DANCANKARANI/QVP/utilities"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// )

// //add audit handler
// func AddAuditHandler(c *fiber.Ctx) error {
// 	response, err := model.AddAudit(c)
// 	if err != nil{
// 		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
// 	}
// 	return utilities.ShowSuccess(c,"successfully added audits",fiber.StatusOK,response)
// }
// //update audit handler
// func UpdateAuditHandler(c *fiber.Ctx)error{
// 	id, _:=uuid.Parse(c.Params("id"))
// 	code,response, err := model.UpdateAudit(c,id)
// 	if err != nil{
// 		return utilities.ShowError(c,err.Error(),code)
// 	}
// 	return utilities.ShowSuccess(c,"successfuly updated audits",code,response)
// }

// //get audits handler

// func GetUserAuditsHandler(c *fiber.Ctx)error{
// 	id,_ := model.GetAuthUserID(c)
// 	 response,code, err:=model.GetUserAudits(id)
// 	if err != nil{
// 		return utilities.ShowError(c,err.Error(),code)
// 	}
// 	return utilities.ShowSuccess(c,"successfully retrieved user's audits",code,response)
// }

// //delete audits handler
// func DeleteAuditHandler(c *fiber.Ctx)error{
// 	id,_:=uuid.Parse(c.Params("id"))
// 	code, err := model.DeleteAudit(id)
// 	if err != nil{
// 		return utilities.ShowError(c,err.Error(),code)
// 	}
// 	return utilities.ShowMessage(c,"audit deleted successfull",code)
// }