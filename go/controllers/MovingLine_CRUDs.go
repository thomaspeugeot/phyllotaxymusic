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
var __MovingLine__dummysDeclaration__ models.MovingLine
var __MovingLine_time__dummyDeclaration time.Duration

var mutexMovingLine sync.Mutex

// An MovingLineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMovingLine updateMovingLine deleteMovingLine
type MovingLineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MovingLineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMovingLine updateMovingLine
type MovingLineInput struct {
	// The MovingLine to submit or modify
	// in: body
	MovingLine *orm.MovingLineAPI
}

// GetMovingLines
//
// swagger:route GET /movinglines movinglines getMovingLines
//
// # Get all movinglines
//
// Responses:
// default: genericError
//
//	200: movinglineDBResponse
func (controller *Controller) GetMovingLines(c *gin.Context) {

	// source slice
	var movinglineDBs []orm.MovingLineDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMovingLines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMovingLine.GetDB()

	_, err := db.Find(&movinglineDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	movinglineAPIs := make([]orm.MovingLineAPI, 0)

	// for each movingline, update fields from the database nullable fields
	for idx := range movinglineDBs {
		movinglineDB := &movinglineDBs[idx]
		_ = movinglineDB
		var movinglineAPI orm.MovingLineAPI

		// insertion point for updating fields
		movinglineAPI.ID = movinglineDB.ID
		movinglineDB.CopyBasicFieldsToMovingLine_WOP(&movinglineAPI.MovingLine_WOP)
		movinglineAPI.MovingLinePointersEncoding = movinglineDB.MovingLinePointersEncoding
		movinglineAPIs = append(movinglineAPIs, movinglineAPI)
	}

	c.JSON(http.StatusOK, movinglineAPIs)
}

// PostMovingLine
//
// swagger:route POST /movinglines movinglines postMovingLine
//
// Creates a movingline
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMovingLine(c *gin.Context) {

	mutexMovingLine.Lock()
	defer mutexMovingLine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMovingLines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMovingLine.GetDB()

	// Validate input
	var input orm.MovingLineAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create movingline
	movinglineDB := orm.MovingLineDB{}
	movinglineDB.MovingLinePointersEncoding = input.MovingLinePointersEncoding
	movinglineDB.CopyBasicFieldsFromMovingLine_WOP(&input.MovingLine_WOP)

	_, err = db.Create(&movinglineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMovingLine.CheckoutPhaseOneInstance(&movinglineDB)
	movingline := backRepo.BackRepoMovingLine.Map_MovingLineDBID_MovingLinePtr[movinglineDB.ID]

	if movingline != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), movingline)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, movinglineDB)
}

// GetMovingLine
//
// swagger:route GET /movinglines/{ID} movinglines getMovingLine
//
// Gets the details for a movingline.
//
// Responses:
// default: genericError
//
//	200: movinglineDBResponse
func (controller *Controller) GetMovingLine(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMovingLine", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMovingLine.GetDB()

	// Get movinglineDB in DB
	var movinglineDB orm.MovingLineDB
	if _, err := db.First(&movinglineDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var movinglineAPI orm.MovingLineAPI
	movinglineAPI.ID = movinglineDB.ID
	movinglineAPI.MovingLinePointersEncoding = movinglineDB.MovingLinePointersEncoding
	movinglineDB.CopyBasicFieldsToMovingLine_WOP(&movinglineAPI.MovingLine_WOP)

	c.JSON(http.StatusOK, movinglineAPI)
}

// UpdateMovingLine
//
// swagger:route PATCH /movinglines/{ID} movinglines updateMovingLine
//
// # Update a movingline
//
// Responses:
// default: genericError
//
//	200: movinglineDBResponse
func (controller *Controller) UpdateMovingLine(c *gin.Context) {

	mutexMovingLine.Lock()
	defer mutexMovingLine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMovingLine", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMovingLine.GetDB()

	// Validate input
	var input orm.MovingLineAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var movinglineDB orm.MovingLineDB

	// fetch the movingline
	_, err := db.First(&movinglineDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	movinglineDB.CopyBasicFieldsFromMovingLine_WOP(&input.MovingLine_WOP)
	movinglineDB.MovingLinePointersEncoding = input.MovingLinePointersEncoding

	db, _ = db.Model(&movinglineDB)
	_, err = db.Updates(&movinglineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	movinglineNew := new(models.MovingLine)
	movinglineDB.CopyBasicFieldsToMovingLine(movinglineNew)

	// redeem pointers
	movinglineDB.DecodePointers(backRepo, movinglineNew)

	// get stage instance from DB instance, and call callback function
	movinglineOld := backRepo.BackRepoMovingLine.Map_MovingLineDBID_MovingLinePtr[movinglineDB.ID]
	if movinglineOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), movinglineOld, movinglineNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the movinglineDB
	c.JSON(http.StatusOK, movinglineDB)
}

// DeleteMovingLine
//
// swagger:route DELETE /movinglines/{ID} movinglines deleteMovingLine
//
// # Delete a movingline
//
// default: genericError
//
//	200: movinglineDBResponse
func (controller *Controller) DeleteMovingLine(c *gin.Context) {

	mutexMovingLine.Lock()
	defer mutexMovingLine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMovingLine", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMovingLine.GetDB()

	// Get model if exist
	var movinglineDB orm.MovingLineDB
	if _, err := db.First(&movinglineDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&movinglineDB)

	// get an instance (not staged) from DB instance, and call callback function
	movinglineDeleted := new(models.MovingLine)
	movinglineDB.CopyBasicFieldsToMovingLine(movinglineDeleted)

	// get stage instance from DB instance, and call callback function
	movinglineStaged := backRepo.BackRepoMovingLine.Map_MovingLineDBID_MovingLinePtr[movinglineDB.ID]
	if movinglineStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), movinglineStaged, movinglineDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
