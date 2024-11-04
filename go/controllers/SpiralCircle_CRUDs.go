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
var __SpiralCircle__dummysDeclaration__ models.SpiralCircle
var __SpiralCircle_time__dummyDeclaration time.Duration

var mutexSpiralCircle sync.Mutex

// An SpiralCircleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSpiralCircle updateSpiralCircle deleteSpiralCircle
type SpiralCircleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SpiralCircleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSpiralCircle updateSpiralCircle
type SpiralCircleInput struct {
	// The SpiralCircle to submit or modify
	// in: body
	SpiralCircle *orm.SpiralCircleAPI
}

// GetSpiralCircles
//
// swagger:route GET /spiralcircles spiralcircles getSpiralCircles
//
// # Get all spiralcircles
//
// Responses:
// default: genericError
//
//	200: spiralcircleDBResponse
func (controller *Controller) GetSpiralCircles(c *gin.Context) {

	// source slice
	var spiralcircleDBs []orm.SpiralCircleDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralCircles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralCircle.GetDB()

	_, err := db.Find(&spiralcircleDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spiralcircleAPIs := make([]orm.SpiralCircleAPI, 0)

	// for each spiralcircle, update fields from the database nullable fields
	for idx := range spiralcircleDBs {
		spiralcircleDB := &spiralcircleDBs[idx]
		_ = spiralcircleDB
		var spiralcircleAPI orm.SpiralCircleAPI

		// insertion point for updating fields
		spiralcircleAPI.ID = spiralcircleDB.ID
		spiralcircleDB.CopyBasicFieldsToSpiralCircle_WOP(&spiralcircleAPI.SpiralCircle_WOP)
		spiralcircleAPI.SpiralCirclePointersEncoding = spiralcircleDB.SpiralCirclePointersEncoding
		spiralcircleAPIs = append(spiralcircleAPIs, spiralcircleAPI)
	}

	c.JSON(http.StatusOK, spiralcircleAPIs)
}

// PostSpiralCircle
//
// swagger:route POST /spiralcircles spiralcircles postSpiralCircle
//
// Creates a spiralcircle
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSpiralCircle(c *gin.Context) {

	mutexSpiralCircle.Lock()
	defer mutexSpiralCircle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSpiralCircles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralCircle.GetDB()

	// Validate input
	var input orm.SpiralCircleAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spiralcircle
	spiralcircleDB := orm.SpiralCircleDB{}
	spiralcircleDB.SpiralCirclePointersEncoding = input.SpiralCirclePointersEncoding
	spiralcircleDB.CopyBasicFieldsFromSpiralCircle_WOP(&input.SpiralCircle_WOP)

	_, err = db.Create(&spiralcircleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSpiralCircle.CheckoutPhaseOneInstance(&spiralcircleDB)
	spiralcircle := backRepo.BackRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCirclePtr[spiralcircleDB.ID]

	if spiralcircle != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spiralcircle)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spiralcircleDB)
}

// GetSpiralCircle
//
// swagger:route GET /spiralcircles/{ID} spiralcircles getSpiralCircle
//
// Gets the details for a spiralcircle.
//
// Responses:
// default: genericError
//
//	200: spiralcircleDBResponse
func (controller *Controller) GetSpiralCircle(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralCircle", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralCircle.GetDB()

	// Get spiralcircleDB in DB
	var spiralcircleDB orm.SpiralCircleDB
	if _, err := db.First(&spiralcircleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spiralcircleAPI orm.SpiralCircleAPI
	spiralcircleAPI.ID = spiralcircleDB.ID
	spiralcircleAPI.SpiralCirclePointersEncoding = spiralcircleDB.SpiralCirclePointersEncoding
	spiralcircleDB.CopyBasicFieldsToSpiralCircle_WOP(&spiralcircleAPI.SpiralCircle_WOP)

	c.JSON(http.StatusOK, spiralcircleAPI)
}

// UpdateSpiralCircle
//
// swagger:route PATCH /spiralcircles/{ID} spiralcircles updateSpiralCircle
//
// # Update a spiralcircle
//
// Responses:
// default: genericError
//
//	200: spiralcircleDBResponse
func (controller *Controller) UpdateSpiralCircle(c *gin.Context) {

	mutexSpiralCircle.Lock()
	defer mutexSpiralCircle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSpiralCircle", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralCircle.GetDB()

	// Validate input
	var input orm.SpiralCircleAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spiralcircleDB orm.SpiralCircleDB

	// fetch the spiralcircle
	_, err := db.First(&spiralcircleDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spiralcircleDB.CopyBasicFieldsFromSpiralCircle_WOP(&input.SpiralCircle_WOP)
	spiralcircleDB.SpiralCirclePointersEncoding = input.SpiralCirclePointersEncoding

	db, _ = db.Model(&spiralcircleDB)
	_, err = db.Updates(&spiralcircleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spiralcircleNew := new(models.SpiralCircle)
	spiralcircleDB.CopyBasicFieldsToSpiralCircle(spiralcircleNew)

	// redeem pointers
	spiralcircleDB.DecodePointers(backRepo, spiralcircleNew)

	// get stage instance from DB instance, and call callback function
	spiralcircleOld := backRepo.BackRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCirclePtr[spiralcircleDB.ID]
	if spiralcircleOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), spiralcircleOld, spiralcircleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spiralcircleDB
	c.JSON(http.StatusOK, spiralcircleDB)
}

// DeleteSpiralCircle
//
// swagger:route DELETE /spiralcircles/{ID} spiralcircles deleteSpiralCircle
//
// # Delete a spiralcircle
//
// default: genericError
//
//	200: spiralcircleDBResponse
func (controller *Controller) DeleteSpiralCircle(c *gin.Context) {

	mutexSpiralCircle.Lock()
	defer mutexSpiralCircle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSpiralCircle", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralCircle.GetDB()

	// Get model if exist
	var spiralcircleDB orm.SpiralCircleDB
	if _, err := db.First(&spiralcircleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&spiralcircleDB)

	// get an instance (not staged) from DB instance, and call callback function
	spiralcircleDeleted := new(models.SpiralCircle)
	spiralcircleDB.CopyBasicFieldsToSpiralCircle(spiralcircleDeleted)

	// get stage instance from DB instance, and call callback function
	spiralcircleStaged := backRepo.BackRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCirclePtr[spiralcircleDB.ID]
	if spiralcircleStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spiralcircleStaged, spiralcircleDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
