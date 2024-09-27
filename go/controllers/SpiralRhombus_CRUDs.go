// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/thomaspeugeot/phylotaxymusic/go/models"
	"github.com/thomaspeugeot/phylotaxymusic/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __SpiralRhombus__dummysDeclaration__ models.SpiralRhombus
var __SpiralRhombus_time__dummyDeclaration time.Duration

var mutexSpiralRhombus sync.Mutex

// An SpiralRhombusID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSpiralRhombus updateSpiralRhombus deleteSpiralRhombus
type SpiralRhombusID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SpiralRhombusInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSpiralRhombus updateSpiralRhombus
type SpiralRhombusInput struct {
	// The SpiralRhombus to submit or modify
	// in: body
	SpiralRhombus *orm.SpiralRhombusAPI
}

// GetSpiralRhombuss
//
// swagger:route GET /spiralrhombuss spiralrhombuss getSpiralRhombuss
//
// # Get all spiralrhombuss
//
// Responses:
// default: genericError
//
//	200: spiralrhombusDBResponse
func (controller *Controller) GetSpiralRhombuss(c *gin.Context) {

	// source slice
	var spiralrhombusDBs []orm.SpiralRhombusDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralRhombuss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralRhombus.GetDB()

	query := db.Find(&spiralrhombusDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spiralrhombusAPIs := make([]orm.SpiralRhombusAPI, 0)

	// for each spiralrhombus, update fields from the database nullable fields
	for idx := range spiralrhombusDBs {
		spiralrhombusDB := &spiralrhombusDBs[idx]
		_ = spiralrhombusDB
		var spiralrhombusAPI orm.SpiralRhombusAPI

		// insertion point for updating fields
		spiralrhombusAPI.ID = spiralrhombusDB.ID
		spiralrhombusDB.CopyBasicFieldsToSpiralRhombus_WOP(&spiralrhombusAPI.SpiralRhombus_WOP)
		spiralrhombusAPI.SpiralRhombusPointersEncoding = spiralrhombusDB.SpiralRhombusPointersEncoding
		spiralrhombusAPIs = append(spiralrhombusAPIs, spiralrhombusAPI)
	}

	c.JSON(http.StatusOK, spiralrhombusAPIs)
}

// PostSpiralRhombus
//
// swagger:route POST /spiralrhombuss spiralrhombuss postSpiralRhombus
//
// Creates a spiralrhombus
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSpiralRhombus(c *gin.Context) {

	mutexSpiralRhombus.Lock()
	defer mutexSpiralRhombus.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSpiralRhombuss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralRhombus.GetDB()

	// Validate input
	var input orm.SpiralRhombusAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spiralrhombus
	spiralrhombusDB := orm.SpiralRhombusDB{}
	spiralrhombusDB.SpiralRhombusPointersEncoding = input.SpiralRhombusPointersEncoding
	spiralrhombusDB.CopyBasicFieldsFromSpiralRhombus_WOP(&input.SpiralRhombus_WOP)

	query := db.Create(&spiralrhombusDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSpiralRhombus.CheckoutPhaseOneInstance(&spiralrhombusDB)
	spiralrhombus := backRepo.BackRepoSpiralRhombus.Map_SpiralRhombusDBID_SpiralRhombusPtr[spiralrhombusDB.ID]

	if spiralrhombus != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spiralrhombus)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spiralrhombusDB)
}

// GetSpiralRhombus
//
// swagger:route GET /spiralrhombuss/{ID} spiralrhombuss getSpiralRhombus
//
// Gets the details for a spiralrhombus.
//
// Responses:
// default: genericError
//
//	200: spiralrhombusDBResponse
func (controller *Controller) GetSpiralRhombus(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralRhombus", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralRhombus.GetDB()

	// Get spiralrhombusDB in DB
	var spiralrhombusDB orm.SpiralRhombusDB
	if err := db.First(&spiralrhombusDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spiralrhombusAPI orm.SpiralRhombusAPI
	spiralrhombusAPI.ID = spiralrhombusDB.ID
	spiralrhombusAPI.SpiralRhombusPointersEncoding = spiralrhombusDB.SpiralRhombusPointersEncoding
	spiralrhombusDB.CopyBasicFieldsToSpiralRhombus_WOP(&spiralrhombusAPI.SpiralRhombus_WOP)

	c.JSON(http.StatusOK, spiralrhombusAPI)
}

// UpdateSpiralRhombus
//
// swagger:route PATCH /spiralrhombuss/{ID} spiralrhombuss updateSpiralRhombus
//
// # Update a spiralrhombus
//
// Responses:
// default: genericError
//
//	200: spiralrhombusDBResponse
func (controller *Controller) UpdateSpiralRhombus(c *gin.Context) {

	mutexSpiralRhombus.Lock()
	defer mutexSpiralRhombus.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSpiralRhombus", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralRhombus.GetDB()

	// Validate input
	var input orm.SpiralRhombusAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spiralrhombusDB orm.SpiralRhombusDB

	// fetch the spiralrhombus
	query := db.First(&spiralrhombusDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spiralrhombusDB.CopyBasicFieldsFromSpiralRhombus_WOP(&input.SpiralRhombus_WOP)
	spiralrhombusDB.SpiralRhombusPointersEncoding = input.SpiralRhombusPointersEncoding

	query = db.Model(&spiralrhombusDB).Updates(spiralrhombusDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spiralrhombusNew := new(models.SpiralRhombus)
	spiralrhombusDB.CopyBasicFieldsToSpiralRhombus(spiralrhombusNew)

	// redeem pointers
	spiralrhombusDB.DecodePointers(backRepo, spiralrhombusNew)

	// get stage instance from DB instance, and call callback function
	spiralrhombusOld := backRepo.BackRepoSpiralRhombus.Map_SpiralRhombusDBID_SpiralRhombusPtr[spiralrhombusDB.ID]
	if spiralrhombusOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), spiralrhombusOld, spiralrhombusNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spiralrhombusDB
	c.JSON(http.StatusOK, spiralrhombusDB)
}

// DeleteSpiralRhombus
//
// swagger:route DELETE /spiralrhombuss/{ID} spiralrhombuss deleteSpiralRhombus
//
// # Delete a spiralrhombus
//
// default: genericError
//
//	200: spiralrhombusDBResponse
func (controller *Controller) DeleteSpiralRhombus(c *gin.Context) {

	mutexSpiralRhombus.Lock()
	defer mutexSpiralRhombus.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSpiralRhombus", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralRhombus.GetDB()

	// Get model if exist
	var spiralrhombusDB orm.SpiralRhombusDB
	if err := db.First(&spiralrhombusDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&spiralrhombusDB)

	// get an instance (not staged) from DB instance, and call callback function
	spiralrhombusDeleted := new(models.SpiralRhombus)
	spiralrhombusDB.CopyBasicFieldsToSpiralRhombus(spiralrhombusDeleted)

	// get stage instance from DB instance, and call callback function
	spiralrhombusStaged := backRepo.BackRepoSpiralRhombus.Map_SpiralRhombusDBID_SpiralRhombusPtr[spiralrhombusDB.ID]
	if spiralrhombusStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spiralrhombusStaged, spiralrhombusDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
