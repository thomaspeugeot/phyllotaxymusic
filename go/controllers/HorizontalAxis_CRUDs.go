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
var __HorizontalAxis__dummysDeclaration__ models.HorizontalAxis
var __HorizontalAxis_time__dummyDeclaration time.Duration

var mutexHorizontalAxis sync.Mutex

// An HorizontalAxisID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getHorizontalAxis updateHorizontalAxis deleteHorizontalAxis
type HorizontalAxisID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// HorizontalAxisInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postHorizontalAxis updateHorizontalAxis
type HorizontalAxisInput struct {
	// The HorizontalAxis to submit or modify
	// in: body
	HorizontalAxis *orm.HorizontalAxisAPI
}

// GetHorizontalAxiss
//
// swagger:route GET /horizontalaxiss horizontalaxiss getHorizontalAxiss
//
// # Get all horizontalaxiss
//
// Responses:
// default: genericError
//
//	200: horizontalaxisDBResponse
func (controller *Controller) GetHorizontalAxiss(c *gin.Context) {

	// source slice
	var horizontalaxisDBs []orm.HorizontalAxisDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHorizontalAxiss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHorizontalAxis.GetDB()

	_, err := db.Find(&horizontalaxisDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	horizontalaxisAPIs := make([]orm.HorizontalAxisAPI, 0)

	// for each horizontalaxis, update fields from the database nullable fields
	for idx := range horizontalaxisDBs {
		horizontalaxisDB := &horizontalaxisDBs[idx]
		_ = horizontalaxisDB
		var horizontalaxisAPI orm.HorizontalAxisAPI

		// insertion point for updating fields
		horizontalaxisAPI.ID = horizontalaxisDB.ID
		horizontalaxisDB.CopyBasicFieldsToHorizontalAxis_WOP(&horizontalaxisAPI.HorizontalAxis_WOP)
		horizontalaxisAPI.HorizontalAxisPointersEncoding = horizontalaxisDB.HorizontalAxisPointersEncoding
		horizontalaxisAPIs = append(horizontalaxisAPIs, horizontalaxisAPI)
	}

	c.JSON(http.StatusOK, horizontalaxisAPIs)
}

// PostHorizontalAxis
//
// swagger:route POST /horizontalaxiss horizontalaxiss postHorizontalAxis
//
// Creates a horizontalaxis
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostHorizontalAxis(c *gin.Context) {

	mutexHorizontalAxis.Lock()
	defer mutexHorizontalAxis.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostHorizontalAxiss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHorizontalAxis.GetDB()

	// Validate input
	var input orm.HorizontalAxisAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create horizontalaxis
	horizontalaxisDB := orm.HorizontalAxisDB{}
	horizontalaxisDB.HorizontalAxisPointersEncoding = input.HorizontalAxisPointersEncoding
	horizontalaxisDB.CopyBasicFieldsFromHorizontalAxis_WOP(&input.HorizontalAxis_WOP)

	_, err = db.Create(&horizontalaxisDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoHorizontalAxis.CheckoutPhaseOneInstance(&horizontalaxisDB)
	horizontalaxis := backRepo.BackRepoHorizontalAxis.Map_HorizontalAxisDBID_HorizontalAxisPtr[horizontalaxisDB.ID]

	if horizontalaxis != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), horizontalaxis)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, horizontalaxisDB)
}

// GetHorizontalAxis
//
// swagger:route GET /horizontalaxiss/{ID} horizontalaxiss getHorizontalAxis
//
// Gets the details for a horizontalaxis.
//
// Responses:
// default: genericError
//
//	200: horizontalaxisDBResponse
func (controller *Controller) GetHorizontalAxis(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHorizontalAxis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHorizontalAxis.GetDB()

	// Get horizontalaxisDB in DB
	var horizontalaxisDB orm.HorizontalAxisDB
	if _, err := db.First(&horizontalaxisDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var horizontalaxisAPI orm.HorizontalAxisAPI
	horizontalaxisAPI.ID = horizontalaxisDB.ID
	horizontalaxisAPI.HorizontalAxisPointersEncoding = horizontalaxisDB.HorizontalAxisPointersEncoding
	horizontalaxisDB.CopyBasicFieldsToHorizontalAxis_WOP(&horizontalaxisAPI.HorizontalAxis_WOP)

	c.JSON(http.StatusOK, horizontalaxisAPI)
}

// UpdateHorizontalAxis
//
// swagger:route PATCH /horizontalaxiss/{ID} horizontalaxiss updateHorizontalAxis
//
// # Update a horizontalaxis
//
// Responses:
// default: genericError
//
//	200: horizontalaxisDBResponse
func (controller *Controller) UpdateHorizontalAxis(c *gin.Context) {

	mutexHorizontalAxis.Lock()
	defer mutexHorizontalAxis.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateHorizontalAxis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHorizontalAxis.GetDB()

	// Validate input
	var input orm.HorizontalAxisAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var horizontalaxisDB orm.HorizontalAxisDB

	// fetch the horizontalaxis
	_, err := db.First(&horizontalaxisDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	horizontalaxisDB.CopyBasicFieldsFromHorizontalAxis_WOP(&input.HorizontalAxis_WOP)
	horizontalaxisDB.HorizontalAxisPointersEncoding = input.HorizontalAxisPointersEncoding

	db, _ = db.Model(&horizontalaxisDB)
	_, err = db.Updates(horizontalaxisDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	horizontalaxisNew := new(models.HorizontalAxis)
	horizontalaxisDB.CopyBasicFieldsToHorizontalAxis(horizontalaxisNew)

	// redeem pointers
	horizontalaxisDB.DecodePointers(backRepo, horizontalaxisNew)

	// get stage instance from DB instance, and call callback function
	horizontalaxisOld := backRepo.BackRepoHorizontalAxis.Map_HorizontalAxisDBID_HorizontalAxisPtr[horizontalaxisDB.ID]
	if horizontalaxisOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), horizontalaxisOld, horizontalaxisNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the horizontalaxisDB
	c.JSON(http.StatusOK, horizontalaxisDB)
}

// DeleteHorizontalAxis
//
// swagger:route DELETE /horizontalaxiss/{ID} horizontalaxiss deleteHorizontalAxis
//
// # Delete a horizontalaxis
//
// default: genericError
//
//	200: horizontalaxisDBResponse
func (controller *Controller) DeleteHorizontalAxis(c *gin.Context) {

	mutexHorizontalAxis.Lock()
	defer mutexHorizontalAxis.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteHorizontalAxis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHorizontalAxis.GetDB()

	// Get model if exist
	var horizontalaxisDB orm.HorizontalAxisDB
	if _, err := db.First(&horizontalaxisDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&horizontalaxisDB)

	// get an instance (not staged) from DB instance, and call callback function
	horizontalaxisDeleted := new(models.HorizontalAxis)
	horizontalaxisDB.CopyBasicFieldsToHorizontalAxis(horizontalaxisDeleted)

	// get stage instance from DB instance, and call callback function
	horizontalaxisStaged := backRepo.BackRepoHorizontalAxis.Map_HorizontalAxisDBID_HorizontalAxisPtr[horizontalaxisDB.ID]
	if horizontalaxisStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), horizontalaxisStaged, horizontalaxisDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
