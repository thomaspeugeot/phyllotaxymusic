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
var __Rhombus__dummysDeclaration__ models.Rhombus
var __Rhombus_time__dummyDeclaration time.Duration

var mutexRhombus sync.Mutex

// An RhombusID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRhombus updateRhombus deleteRhombus
type RhombusID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RhombusInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRhombus updateRhombus
type RhombusInput struct {
	// The Rhombus to submit or modify
	// in: body
	Rhombus *orm.RhombusAPI
}

// GetRhombuss
//
// swagger:route GET /rhombuss rhombuss getRhombuss
//
// # Get all rhombuss
//
// Responses:
// default: genericError
//
//	200: rhombusDBResponse
func (controller *Controller) GetRhombuss(c *gin.Context) {

	// source slice
	var rhombusDBs []orm.RhombusDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRhombuss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoRhombus.GetDB()

	_, err := db.Find(&rhombusDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	rhombusAPIs := make([]orm.RhombusAPI, 0)

	// for each rhombus, update fields from the database nullable fields
	for idx := range rhombusDBs {
		rhombusDB := &rhombusDBs[idx]
		_ = rhombusDB
		var rhombusAPI orm.RhombusAPI

		// insertion point for updating fields
		rhombusAPI.ID = rhombusDB.ID
		rhombusDB.CopyBasicFieldsToRhombus_WOP(&rhombusAPI.Rhombus_WOP)
		rhombusAPI.RhombusPointersEncoding = rhombusDB.RhombusPointersEncoding
		rhombusAPIs = append(rhombusAPIs, rhombusAPI)
	}

	c.JSON(http.StatusOK, rhombusAPIs)
}

// PostRhombus
//
// swagger:route POST /rhombuss rhombuss postRhombus
//
// Creates a rhombus
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRhombus(c *gin.Context) {

	mutexRhombus.Lock()
	defer mutexRhombus.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRhombuss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoRhombus.GetDB()

	// Validate input
	var input orm.RhombusAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create rhombus
	rhombusDB := orm.RhombusDB{}
	rhombusDB.RhombusPointersEncoding = input.RhombusPointersEncoding
	rhombusDB.CopyBasicFieldsFromRhombus_WOP(&input.Rhombus_WOP)

	_, err = db.Create(&rhombusDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRhombus.CheckoutPhaseOneInstance(&rhombusDB)
	rhombus := backRepo.BackRepoRhombus.Map_RhombusDBID_RhombusPtr[rhombusDB.ID]

	if rhombus != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), rhombus)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, rhombusDB)
}

// GetRhombus
//
// swagger:route GET /rhombuss/{ID} rhombuss getRhombus
//
// Gets the details for a rhombus.
//
// Responses:
// default: genericError
//
//	200: rhombusDBResponse
func (controller *Controller) GetRhombus(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRhombus", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoRhombus.GetDB()

	// Get rhombusDB in DB
	var rhombusDB orm.RhombusDB
	if _, err := db.First(&rhombusDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var rhombusAPI orm.RhombusAPI
	rhombusAPI.ID = rhombusDB.ID
	rhombusAPI.RhombusPointersEncoding = rhombusDB.RhombusPointersEncoding
	rhombusDB.CopyBasicFieldsToRhombus_WOP(&rhombusAPI.Rhombus_WOP)

	c.JSON(http.StatusOK, rhombusAPI)
}

// UpdateRhombus
//
// swagger:route PATCH /rhombuss/{ID} rhombuss updateRhombus
//
// # Update a rhombus
//
// Responses:
// default: genericError
//
//	200: rhombusDBResponse
func (controller *Controller) UpdateRhombus(c *gin.Context) {

	mutexRhombus.Lock()
	defer mutexRhombus.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRhombus", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoRhombus.GetDB()

	// Validate input
	var input orm.RhombusAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rhombusDB orm.RhombusDB

	// fetch the rhombus
	_, err := db.First(&rhombusDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rhombusDB.CopyBasicFieldsFromRhombus_WOP(&input.Rhombus_WOP)
	rhombusDB.RhombusPointersEncoding = input.RhombusPointersEncoding

	db, _ = db.Model(&rhombusDB)
	_, err = db.Updates(&rhombusDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rhombusNew := new(models.Rhombus)
	rhombusDB.CopyBasicFieldsToRhombus(rhombusNew)

	// redeem pointers
	rhombusDB.DecodePointers(backRepo, rhombusNew)

	// get stage instance from DB instance, and call callback function
	rhombusOld := backRepo.BackRepoRhombus.Map_RhombusDBID_RhombusPtr[rhombusDB.ID]
	if rhombusOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), rhombusOld, rhombusNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rhombusDB
	c.JSON(http.StatusOK, rhombusDB)
}

// DeleteRhombus
//
// swagger:route DELETE /rhombuss/{ID} rhombuss deleteRhombus
//
// # Delete a rhombus
//
// default: genericError
//
//	200: rhombusDBResponse
func (controller *Controller) DeleteRhombus(c *gin.Context) {

	mutexRhombus.Lock()
	defer mutexRhombus.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRhombus", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoRhombus.GetDB()

	// Get model if exist
	var rhombusDB orm.RhombusDB
	if _, err := db.First(&rhombusDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&rhombusDB)

	// get an instance (not staged) from DB instance, and call callback function
	rhombusDeleted := new(models.Rhombus)
	rhombusDB.CopyBasicFieldsToRhombus(rhombusDeleted)

	// get stage instance from DB instance, and call callback function
	rhombusStaged := backRepo.BackRepoRhombus.Map_RhombusDBID_RhombusPtr[rhombusDB.ID]
	if rhombusStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), rhombusStaged, rhombusDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
