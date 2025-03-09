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
var __SpiralOrigin__dummysDeclaration__ models.SpiralOrigin
var __SpiralOrigin_time__dummyDeclaration time.Duration

var mutexSpiralOrigin sync.Mutex

// An SpiralOriginID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSpiralOrigin updateSpiralOrigin deleteSpiralOrigin
type SpiralOriginID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SpiralOriginInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSpiralOrigin updateSpiralOrigin
type SpiralOriginInput struct {
	// The SpiralOrigin to submit or modify
	// in: body
	SpiralOrigin *orm.SpiralOriginAPI
}

// GetSpiralOrigins
//
// swagger:route GET /spiralorigins spiralorigins getSpiralOrigins
//
// # Get all spiralorigins
//
// Responses:
// default: genericError
//
//	200: spiraloriginDBResponse
func (controller *Controller) GetSpiralOrigins(c *gin.Context) {

	// source slice
	var spiraloriginDBs []orm.SpiralOriginDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralOrigins", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoSpiralOrigin.GetDB()

	_, err := db.Find(&spiraloriginDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spiraloriginAPIs := make([]orm.SpiralOriginAPI, 0)

	// for each spiralorigin, update fields from the database nullable fields
	for idx := range spiraloriginDBs {
		spiraloriginDB := &spiraloriginDBs[idx]
		_ = spiraloriginDB
		var spiraloriginAPI orm.SpiralOriginAPI

		// insertion point for updating fields
		spiraloriginAPI.ID = spiraloriginDB.ID
		spiraloriginDB.CopyBasicFieldsToSpiralOrigin_WOP(&spiraloriginAPI.SpiralOrigin_WOP)
		spiraloriginAPI.SpiralOriginPointersEncoding = spiraloriginDB.SpiralOriginPointersEncoding
		spiraloriginAPIs = append(spiraloriginAPIs, spiraloriginAPI)
	}

	c.JSON(http.StatusOK, spiraloriginAPIs)
}

// PostSpiralOrigin
//
// swagger:route POST /spiralorigins spiralorigins postSpiralOrigin
//
// Creates a spiralorigin
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSpiralOrigin(c *gin.Context) {

	mutexSpiralOrigin.Lock()
	defer mutexSpiralOrigin.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSpiralOrigins", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoSpiralOrigin.GetDB()

	// Validate input
	var input orm.SpiralOriginAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spiralorigin
	spiraloriginDB := orm.SpiralOriginDB{}
	spiraloriginDB.SpiralOriginPointersEncoding = input.SpiralOriginPointersEncoding
	spiraloriginDB.CopyBasicFieldsFromSpiralOrigin_WOP(&input.SpiralOrigin_WOP)

	_, err = db.Create(&spiraloriginDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSpiralOrigin.CheckoutPhaseOneInstance(&spiraloriginDB)
	spiralorigin := backRepo.BackRepoSpiralOrigin.Map_SpiralOriginDBID_SpiralOriginPtr[spiraloriginDB.ID]

	if spiralorigin != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spiralorigin)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spiraloriginDB)
}

// GetSpiralOrigin
//
// swagger:route GET /spiralorigins/{ID} spiralorigins getSpiralOrigin
//
// Gets the details for a spiralorigin.
//
// Responses:
// default: genericError
//
//	200: spiraloriginDBResponse
func (controller *Controller) GetSpiralOrigin(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralOrigin", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoSpiralOrigin.GetDB()

	// Get spiraloriginDB in DB
	var spiraloriginDB orm.SpiralOriginDB
	if _, err := db.First(&spiraloriginDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spiraloriginAPI orm.SpiralOriginAPI
	spiraloriginAPI.ID = spiraloriginDB.ID
	spiraloriginAPI.SpiralOriginPointersEncoding = spiraloriginDB.SpiralOriginPointersEncoding
	spiraloriginDB.CopyBasicFieldsToSpiralOrigin_WOP(&spiraloriginAPI.SpiralOrigin_WOP)

	c.JSON(http.StatusOK, spiraloriginAPI)
}

// UpdateSpiralOrigin
//
// swagger:route PATCH /spiralorigins/{ID} spiralorigins updateSpiralOrigin
//
// # Update a spiralorigin
//
// Responses:
// default: genericError
//
//	200: spiraloriginDBResponse
func (controller *Controller) UpdateSpiralOrigin(c *gin.Context) {

	mutexSpiralOrigin.Lock()
	defer mutexSpiralOrigin.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSpiralOrigin", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoSpiralOrigin.GetDB()

	// Validate input
	var input orm.SpiralOriginAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spiraloriginDB orm.SpiralOriginDB

	// fetch the spiralorigin
	_, err := db.First(&spiraloriginDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spiraloriginDB.CopyBasicFieldsFromSpiralOrigin_WOP(&input.SpiralOrigin_WOP)
	spiraloriginDB.SpiralOriginPointersEncoding = input.SpiralOriginPointersEncoding

	db, _ = db.Model(&spiraloriginDB)
	_, err = db.Updates(&spiraloriginDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spiraloriginNew := new(models.SpiralOrigin)
	spiraloriginDB.CopyBasicFieldsToSpiralOrigin(spiraloriginNew)

	// redeem pointers
	spiraloriginDB.DecodePointers(backRepo, spiraloriginNew)

	// get stage instance from DB instance, and call callback function
	spiraloriginOld := backRepo.BackRepoSpiralOrigin.Map_SpiralOriginDBID_SpiralOriginPtr[spiraloriginDB.ID]
	if spiraloriginOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), spiraloriginOld, spiraloriginNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spiraloriginDB
	c.JSON(http.StatusOK, spiraloriginDB)
}

// DeleteSpiralOrigin
//
// swagger:route DELETE /spiralorigins/{ID} spiralorigins deleteSpiralOrigin
//
// # Delete a spiralorigin
//
// default: genericError
//
//	200: spiraloriginDBResponse
func (controller *Controller) DeleteSpiralOrigin(c *gin.Context) {

	mutexSpiralOrigin.Lock()
	defer mutexSpiralOrigin.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSpiralOrigin", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoSpiralOrigin.GetDB()

	// Get model if exist
	var spiraloriginDB orm.SpiralOriginDB
	if _, err := db.First(&spiraloriginDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&spiraloriginDB)

	// get an instance (not staged) from DB instance, and call callback function
	spiraloriginDeleted := new(models.SpiralOrigin)
	spiraloriginDB.CopyBasicFieldsToSpiralOrigin(spiraloriginDeleted)

	// get stage instance from DB instance, and call callback function
	spiraloriginStaged := backRepo.BackRepoSpiralOrigin.Map_SpiralOriginDBID_SpiralOriginPtr[spiraloriginDB.ID]
	if spiraloriginStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spiraloriginStaged, spiraloriginDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
