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
var __SpiralLine__dummysDeclaration__ models.SpiralLine
var __SpiralLine_time__dummyDeclaration time.Duration

var mutexSpiralLine sync.Mutex

// An SpiralLineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSpiralLine updateSpiralLine deleteSpiralLine
type SpiralLineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SpiralLineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSpiralLine updateSpiralLine
type SpiralLineInput struct {
	// The SpiralLine to submit or modify
	// in: body
	SpiralLine *orm.SpiralLineAPI
}

// GetSpiralLines
//
// swagger:route GET /spirallines spirallines getSpiralLines
//
// # Get all spirallines
//
// Responses:
// default: genericError
//
//	200: spirallineDBResponse
func (controller *Controller) GetSpiralLines(c *gin.Context) {

	// source slice
	var spirallineDBs []orm.SpiralLineDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralLines", "Name", stackPath)
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
	db := backRepo.BackRepoSpiralLine.GetDB()

	_, err := db.Find(&spirallineDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spirallineAPIs := make([]orm.SpiralLineAPI, 0)

	// for each spiralline, update fields from the database nullable fields
	for idx := range spirallineDBs {
		spirallineDB := &spirallineDBs[idx]
		_ = spirallineDB
		var spirallineAPI orm.SpiralLineAPI

		// insertion point for updating fields
		spirallineAPI.ID = spirallineDB.ID
		spirallineDB.CopyBasicFieldsToSpiralLine_WOP(&spirallineAPI.SpiralLine_WOP)
		spirallineAPI.SpiralLinePointersEncoding = spirallineDB.SpiralLinePointersEncoding
		spirallineAPIs = append(spirallineAPIs, spirallineAPI)
	}

	c.JSON(http.StatusOK, spirallineAPIs)
}

// PostSpiralLine
//
// swagger:route POST /spirallines spirallines postSpiralLine
//
// Creates a spiralline
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSpiralLine(c *gin.Context) {

	mutexSpiralLine.Lock()
	defer mutexSpiralLine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSpiralLines", "Name", stackPath)
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
	db := backRepo.BackRepoSpiralLine.GetDB()

	// Validate input
	var input orm.SpiralLineAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spiralline
	spirallineDB := orm.SpiralLineDB{}
	spirallineDB.SpiralLinePointersEncoding = input.SpiralLinePointersEncoding
	spirallineDB.CopyBasicFieldsFromSpiralLine_WOP(&input.SpiralLine_WOP)

	_, err = db.Create(&spirallineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSpiralLine.CheckoutPhaseOneInstance(&spirallineDB)
	spiralline := backRepo.BackRepoSpiralLine.Map_SpiralLineDBID_SpiralLinePtr[spirallineDB.ID]

	if spiralline != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spiralline)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spirallineDB)
}

// GetSpiralLine
//
// swagger:route GET /spirallines/{ID} spirallines getSpiralLine
//
// Gets the details for a spiralline.
//
// Responses:
// default: genericError
//
//	200: spirallineDBResponse
func (controller *Controller) GetSpiralLine(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralLine", "Name", stackPath)
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
	db := backRepo.BackRepoSpiralLine.GetDB()

	// Get spirallineDB in DB
	var spirallineDB orm.SpiralLineDB
	if _, err := db.First(&spirallineDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spirallineAPI orm.SpiralLineAPI
	spirallineAPI.ID = spirallineDB.ID
	spirallineAPI.SpiralLinePointersEncoding = spirallineDB.SpiralLinePointersEncoding
	spirallineDB.CopyBasicFieldsToSpiralLine_WOP(&spirallineAPI.SpiralLine_WOP)

	c.JSON(http.StatusOK, spirallineAPI)
}

// UpdateSpiralLine
//
// swagger:route PATCH /spirallines/{ID} spirallines updateSpiralLine
//
// # Update a spiralline
//
// Responses:
// default: genericError
//
//	200: spirallineDBResponse
func (controller *Controller) UpdateSpiralLine(c *gin.Context) {

	mutexSpiralLine.Lock()
	defer mutexSpiralLine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSpiralLine", "Name", stackPath)
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
	db := backRepo.BackRepoSpiralLine.GetDB()

	// Validate input
	var input orm.SpiralLineAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spirallineDB orm.SpiralLineDB

	// fetch the spiralline
	_, err := db.First(&spirallineDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spirallineDB.CopyBasicFieldsFromSpiralLine_WOP(&input.SpiralLine_WOP)
	spirallineDB.SpiralLinePointersEncoding = input.SpiralLinePointersEncoding

	db, _ = db.Model(&spirallineDB)
	_, err = db.Updates(&spirallineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spirallineNew := new(models.SpiralLine)
	spirallineDB.CopyBasicFieldsToSpiralLine(spirallineNew)

	// redeem pointers
	spirallineDB.DecodePointers(backRepo, spirallineNew)

	// get stage instance from DB instance, and call callback function
	spirallineOld := backRepo.BackRepoSpiralLine.Map_SpiralLineDBID_SpiralLinePtr[spirallineDB.ID]
	if spirallineOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), spirallineOld, spirallineNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spirallineDB
	c.JSON(http.StatusOK, spirallineDB)
}

// DeleteSpiralLine
//
// swagger:route DELETE /spirallines/{ID} spirallines deleteSpiralLine
//
// # Delete a spiralline
//
// default: genericError
//
//	200: spirallineDBResponse
func (controller *Controller) DeleteSpiralLine(c *gin.Context) {

	mutexSpiralLine.Lock()
	defer mutexSpiralLine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSpiralLine", "Name", stackPath)
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
	db := backRepo.BackRepoSpiralLine.GetDB()

	// Get model if exist
	var spirallineDB orm.SpiralLineDB
	if _, err := db.First(&spirallineDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&spirallineDB)

	// get an instance (not staged) from DB instance, and call callback function
	spirallineDeleted := new(models.SpiralLine)
	spirallineDB.CopyBasicFieldsToSpiralLine(spirallineDeleted)

	// get stage instance from DB instance, and call callback function
	spirallineStaged := backRepo.BackRepoSpiralLine.Map_SpiralLineDBID_SpiralLinePtr[spirallineDB.ID]
	if spirallineStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spirallineStaged, spirallineDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
