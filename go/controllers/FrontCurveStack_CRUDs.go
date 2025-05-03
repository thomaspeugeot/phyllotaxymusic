// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/thomaspeugeot/phyllotaxymusic/go/models"
	"github.com/thomaspeugeot/phyllotaxymusic/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __FrontCurveStack__dummysDeclaration__ models.FrontCurveStack
var __FrontCurveStack_time__dummyDeclaration time.Duration

var mutexFrontCurveStack sync.Mutex

// An FrontCurveStackID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFrontCurveStack updateFrontCurveStack deleteFrontCurveStack
type FrontCurveStackID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FrontCurveStackInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFrontCurveStack updateFrontCurveStack
type FrontCurveStackInput struct {
	// The FrontCurveStack to submit or modify
	// in: body
	FrontCurveStack *orm.FrontCurveStackAPI
}

// GetFrontCurveStacks
//
// swagger:route GET /frontcurvestacks frontcurvestacks getFrontCurveStacks
//
// # Get all frontcurvestacks
//
// Responses:
// default: genericError
//
//	200: frontcurvestackDBResponse
func (controller *Controller) GetFrontCurveStacks(c *gin.Context) {

	// source slice
	var frontcurvestackDBs []orm.FrontCurveStackDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFrontCurveStacks", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFrontCurveStack.GetDB()

	_, err := db.Find(&frontcurvestackDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	frontcurvestackAPIs := make([]orm.FrontCurveStackAPI, 0)

	// for each frontcurvestack, update fields from the database nullable fields
	for idx := range frontcurvestackDBs {
		frontcurvestackDB := &frontcurvestackDBs[idx]
		_ = frontcurvestackDB
		var frontcurvestackAPI orm.FrontCurveStackAPI

		// insertion point for updating fields
		frontcurvestackAPI.ID = frontcurvestackDB.ID
		frontcurvestackDB.CopyBasicFieldsToFrontCurveStack_WOP(&frontcurvestackAPI.FrontCurveStack_WOP)
		frontcurvestackAPI.FrontCurveStackPointersEncoding = frontcurvestackDB.FrontCurveStackPointersEncoding
		frontcurvestackAPIs = append(frontcurvestackAPIs, frontcurvestackAPI)
	}

	c.JSON(http.StatusOK, frontcurvestackAPIs)
}

// PostFrontCurveStack
//
// swagger:route POST /frontcurvestacks frontcurvestacks postFrontCurveStack
//
// Creates a frontcurvestack
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFrontCurveStack(c *gin.Context) {

	mutexFrontCurveStack.Lock()
	defer mutexFrontCurveStack.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFrontCurveStacks", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFrontCurveStack.GetDB()

	// Validate input
	var input orm.FrontCurveStackAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create frontcurvestack
	frontcurvestackDB := orm.FrontCurveStackDB{}
	frontcurvestackDB.FrontCurveStackPointersEncoding = input.FrontCurveStackPointersEncoding
	frontcurvestackDB.CopyBasicFieldsFromFrontCurveStack_WOP(&input.FrontCurveStack_WOP)

	_, err = db.Create(&frontcurvestackDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFrontCurveStack.CheckoutPhaseOneInstance(&frontcurvestackDB)
	frontcurvestack := backRepo.BackRepoFrontCurveStack.Map_FrontCurveStackDBID_FrontCurveStackPtr[frontcurvestackDB.ID]

	if frontcurvestack != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), frontcurvestack)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, frontcurvestackDB)
}

// GetFrontCurveStack
//
// swagger:route GET /frontcurvestacks/{ID} frontcurvestacks getFrontCurveStack
//
// Gets the details for a frontcurvestack.
//
// Responses:
// default: genericError
//
//	200: frontcurvestackDBResponse
func (controller *Controller) GetFrontCurveStack(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFrontCurveStack", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFrontCurveStack.GetDB()

	// Get frontcurvestackDB in DB
	var frontcurvestackDB orm.FrontCurveStackDB
	if _, err := db.First(&frontcurvestackDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var frontcurvestackAPI orm.FrontCurveStackAPI
	frontcurvestackAPI.ID = frontcurvestackDB.ID
	frontcurvestackAPI.FrontCurveStackPointersEncoding = frontcurvestackDB.FrontCurveStackPointersEncoding
	frontcurvestackDB.CopyBasicFieldsToFrontCurveStack_WOP(&frontcurvestackAPI.FrontCurveStack_WOP)

	c.JSON(http.StatusOK, frontcurvestackAPI)
}

// UpdateFrontCurveStack
//
// swagger:route PATCH /frontcurvestacks/{ID} frontcurvestacks updateFrontCurveStack
//
// # Update a frontcurvestack
//
// Responses:
// default: genericError
//
//	200: frontcurvestackDBResponse
func (controller *Controller) UpdateFrontCurveStack(c *gin.Context) {

	mutexFrontCurveStack.Lock()
	defer mutexFrontCurveStack.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFrontCurveStack", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFrontCurveStack.GetDB()

	// Validate input
	var input orm.FrontCurveStackAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var frontcurvestackDB orm.FrontCurveStackDB

	// fetch the frontcurvestack
	_, err := db.First(&frontcurvestackDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	frontcurvestackDB.CopyBasicFieldsFromFrontCurveStack_WOP(&input.FrontCurveStack_WOP)
	frontcurvestackDB.FrontCurveStackPointersEncoding = input.FrontCurveStackPointersEncoding

	db, _ = db.Model(&frontcurvestackDB)
	_, err = db.Updates(&frontcurvestackDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	frontcurvestackNew := new(models.FrontCurveStack)
	frontcurvestackDB.CopyBasicFieldsToFrontCurveStack(frontcurvestackNew)

	// redeem pointers
	frontcurvestackDB.DecodePointers(backRepo, frontcurvestackNew)

	// get stage instance from DB instance, and call callback function
	frontcurvestackOld := backRepo.BackRepoFrontCurveStack.Map_FrontCurveStackDBID_FrontCurveStackPtr[frontcurvestackDB.ID]
	if frontcurvestackOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), frontcurvestackOld, frontcurvestackNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the frontcurvestackDB
	c.JSON(http.StatusOK, frontcurvestackDB)
}

// DeleteFrontCurveStack
//
// swagger:route DELETE /frontcurvestacks/{ID} frontcurvestacks deleteFrontCurveStack
//
// # Delete a frontcurvestack
//
// default: genericError
//
//	200: frontcurvestackDBResponse
func (controller *Controller) DeleteFrontCurveStack(c *gin.Context) {

	mutexFrontCurveStack.Lock()
	defer mutexFrontCurveStack.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFrontCurveStack", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFrontCurveStack.GetDB()

	// Get model if exist
	var frontcurvestackDB orm.FrontCurveStackDB
	if _, err := db.First(&frontcurvestackDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&frontcurvestackDB)

	// get an instance (not staged) from DB instance, and call callback function
	frontcurvestackDeleted := new(models.FrontCurveStack)
	frontcurvestackDB.CopyBasicFieldsToFrontCurveStack(frontcurvestackDeleted)

	// get stage instance from DB instance, and call callback function
	frontcurvestackStaged := backRepo.BackRepoFrontCurveStack.Map_FrontCurveStackDBID_FrontCurveStackPtr[frontcurvestackDB.ID]
	if frontcurvestackStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), frontcurvestackStaged, frontcurvestackDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
