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
var __VerticalAxis__dummysDeclaration__ models.VerticalAxis
var __VerticalAxis_time__dummyDeclaration time.Duration

var mutexVerticalAxis sync.Mutex

// An VerticalAxisID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getVerticalAxis updateVerticalAxis deleteVerticalAxis
type VerticalAxisID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// VerticalAxisInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postVerticalAxis updateVerticalAxis
type VerticalAxisInput struct {
	// The VerticalAxis to submit or modify
	// in: body
	VerticalAxis *orm.VerticalAxisAPI
}

// GetVerticalAxiss
//
// swagger:route GET /verticalaxiss verticalaxiss getVerticalAxiss
//
// # Get all verticalaxiss
//
// Responses:
// default: genericError
//
//	200: verticalaxisDBResponse
func (controller *Controller) GetVerticalAxiss(c *gin.Context) {

	// source slice
	var verticalaxisDBs []orm.VerticalAxisDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVerticalAxiss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVerticalAxis.GetDB()

	query := db.Find(&verticalaxisDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	verticalaxisAPIs := make([]orm.VerticalAxisAPI, 0)

	// for each verticalaxis, update fields from the database nullable fields
	for idx := range verticalaxisDBs {
		verticalaxisDB := &verticalaxisDBs[idx]
		_ = verticalaxisDB
		var verticalaxisAPI orm.VerticalAxisAPI

		// insertion point for updating fields
		verticalaxisAPI.ID = verticalaxisDB.ID
		verticalaxisDB.CopyBasicFieldsToVerticalAxis_WOP(&verticalaxisAPI.VerticalAxis_WOP)
		verticalaxisAPI.VerticalAxisPointersEncoding = verticalaxisDB.VerticalAxisPointersEncoding
		verticalaxisAPIs = append(verticalaxisAPIs, verticalaxisAPI)
	}

	c.JSON(http.StatusOK, verticalaxisAPIs)
}

// PostVerticalAxis
//
// swagger:route POST /verticalaxiss verticalaxiss postVerticalAxis
//
// Creates a verticalaxis
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostVerticalAxis(c *gin.Context) {

	mutexVerticalAxis.Lock()
	defer mutexVerticalAxis.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostVerticalAxiss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVerticalAxis.GetDB()

	// Validate input
	var input orm.VerticalAxisAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create verticalaxis
	verticalaxisDB := orm.VerticalAxisDB{}
	verticalaxisDB.VerticalAxisPointersEncoding = input.VerticalAxisPointersEncoding
	verticalaxisDB.CopyBasicFieldsFromVerticalAxis_WOP(&input.VerticalAxis_WOP)

	query := db.Create(&verticalaxisDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoVerticalAxis.CheckoutPhaseOneInstance(&verticalaxisDB)
	verticalaxis := backRepo.BackRepoVerticalAxis.Map_VerticalAxisDBID_VerticalAxisPtr[verticalaxisDB.ID]

	if verticalaxis != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), verticalaxis)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, verticalaxisDB)
}

// GetVerticalAxis
//
// swagger:route GET /verticalaxiss/{ID} verticalaxiss getVerticalAxis
//
// Gets the details for a verticalaxis.
//
// Responses:
// default: genericError
//
//	200: verticalaxisDBResponse
func (controller *Controller) GetVerticalAxis(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVerticalAxis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVerticalAxis.GetDB()

	// Get verticalaxisDB in DB
	var verticalaxisDB orm.VerticalAxisDB
	if err := db.First(&verticalaxisDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var verticalaxisAPI orm.VerticalAxisAPI
	verticalaxisAPI.ID = verticalaxisDB.ID
	verticalaxisAPI.VerticalAxisPointersEncoding = verticalaxisDB.VerticalAxisPointersEncoding
	verticalaxisDB.CopyBasicFieldsToVerticalAxis_WOP(&verticalaxisAPI.VerticalAxis_WOP)

	c.JSON(http.StatusOK, verticalaxisAPI)
}

// UpdateVerticalAxis
//
// swagger:route PATCH /verticalaxiss/{ID} verticalaxiss updateVerticalAxis
//
// # Update a verticalaxis
//
// Responses:
// default: genericError
//
//	200: verticalaxisDBResponse
func (controller *Controller) UpdateVerticalAxis(c *gin.Context) {

	mutexVerticalAxis.Lock()
	defer mutexVerticalAxis.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateVerticalAxis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVerticalAxis.GetDB()

	// Validate input
	var input orm.VerticalAxisAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var verticalaxisDB orm.VerticalAxisDB

	// fetch the verticalaxis
	query := db.First(&verticalaxisDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	verticalaxisDB.CopyBasicFieldsFromVerticalAxis_WOP(&input.VerticalAxis_WOP)
	verticalaxisDB.VerticalAxisPointersEncoding = input.VerticalAxisPointersEncoding

	query = db.Model(&verticalaxisDB).Updates(verticalaxisDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	verticalaxisNew := new(models.VerticalAxis)
	verticalaxisDB.CopyBasicFieldsToVerticalAxis(verticalaxisNew)

	// redeem pointers
	verticalaxisDB.DecodePointers(backRepo, verticalaxisNew)

	// get stage instance from DB instance, and call callback function
	verticalaxisOld := backRepo.BackRepoVerticalAxis.Map_VerticalAxisDBID_VerticalAxisPtr[verticalaxisDB.ID]
	if verticalaxisOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), verticalaxisOld, verticalaxisNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the verticalaxisDB
	c.JSON(http.StatusOK, verticalaxisDB)
}

// DeleteVerticalAxis
//
// swagger:route DELETE /verticalaxiss/{ID} verticalaxiss deleteVerticalAxis
//
// # Delete a verticalaxis
//
// default: genericError
//
//	200: verticalaxisDBResponse
func (controller *Controller) DeleteVerticalAxis(c *gin.Context) {

	mutexVerticalAxis.Lock()
	defer mutexVerticalAxis.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteVerticalAxis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVerticalAxis.GetDB()

	// Get model if exist
	var verticalaxisDB orm.VerticalAxisDB
	if err := db.First(&verticalaxisDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&verticalaxisDB)

	// get an instance (not staged) from DB instance, and call callback function
	verticalaxisDeleted := new(models.VerticalAxis)
	verticalaxisDB.CopyBasicFieldsToVerticalAxis(verticalaxisDeleted)

	// get stage instance from DB instance, and call callback function
	verticalaxisStaged := backRepo.BackRepoVerticalAxis.Map_VerticalAxisDBID_VerticalAxisPtr[verticalaxisDB.ID]
	if verticalaxisStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), verticalaxisStaged, verticalaxisDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
