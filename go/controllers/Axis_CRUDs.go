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
var __Axis__dummysDeclaration__ models.Axis
var __Axis_time__dummyDeclaration time.Duration

var mutexAxis sync.Mutex

// An AxisID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAxis updateAxis deleteAxis
type AxisID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AxisInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAxis updateAxis
type AxisInput struct {
	// The Axis to submit or modify
	// in: body
	Axis *orm.AxisAPI
}

// GetAxiss
//
// swagger:route GET /axiss axiss getAxiss
//
// # Get all axiss
//
// Responses:
// default: genericError
//
//	200: axisDBResponse
func (controller *Controller) GetAxiss(c *gin.Context) {

	// source slice
	var axisDBs []orm.AxisDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAxiss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAxis.GetDB()

	query := db.Find(&axisDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	axisAPIs := make([]orm.AxisAPI, 0)

	// for each axis, update fields from the database nullable fields
	for idx := range axisDBs {
		axisDB := &axisDBs[idx]
		_ = axisDB
		var axisAPI orm.AxisAPI

		// insertion point for updating fields
		axisAPI.ID = axisDB.ID
		axisDB.CopyBasicFieldsToAxis_WOP(&axisAPI.Axis_WOP)
		axisAPI.AxisPointersEncoding = axisDB.AxisPointersEncoding
		axisAPIs = append(axisAPIs, axisAPI)
	}

	c.JSON(http.StatusOK, axisAPIs)
}

// PostAxis
//
// swagger:route POST /axiss axiss postAxis
//
// Creates a axis
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAxis(c *gin.Context) {

	mutexAxis.Lock()
	defer mutexAxis.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAxiss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAxis.GetDB()

	// Validate input
	var input orm.AxisAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create axis
	axisDB := orm.AxisDB{}
	axisDB.AxisPointersEncoding = input.AxisPointersEncoding
	axisDB.CopyBasicFieldsFromAxis_WOP(&input.Axis_WOP)

	query := db.Create(&axisDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAxis.CheckoutPhaseOneInstance(&axisDB)
	axis := backRepo.BackRepoAxis.Map_AxisDBID_AxisPtr[axisDB.ID]

	if axis != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), axis)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, axisDB)
}

// GetAxis
//
// swagger:route GET /axiss/{ID} axiss getAxis
//
// Gets the details for a axis.
//
// Responses:
// default: genericError
//
//	200: axisDBResponse
func (controller *Controller) GetAxis(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAxis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAxis.GetDB()

	// Get axisDB in DB
	var axisDB orm.AxisDB
	if err := db.First(&axisDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var axisAPI orm.AxisAPI
	axisAPI.ID = axisDB.ID
	axisAPI.AxisPointersEncoding = axisDB.AxisPointersEncoding
	axisDB.CopyBasicFieldsToAxis_WOP(&axisAPI.Axis_WOP)

	c.JSON(http.StatusOK, axisAPI)
}

// UpdateAxis
//
// swagger:route PATCH /axiss/{ID} axiss updateAxis
//
// # Update a axis
//
// Responses:
// default: genericError
//
//	200: axisDBResponse
func (controller *Controller) UpdateAxis(c *gin.Context) {

	mutexAxis.Lock()
	defer mutexAxis.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAxis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAxis.GetDB()

	// Validate input
	var input orm.AxisAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var axisDB orm.AxisDB

	// fetch the axis
	query := db.First(&axisDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	axisDB.CopyBasicFieldsFromAxis_WOP(&input.Axis_WOP)
	axisDB.AxisPointersEncoding = input.AxisPointersEncoding

	query = db.Model(&axisDB).Updates(axisDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	axisNew := new(models.Axis)
	axisDB.CopyBasicFieldsToAxis(axisNew)

	// redeem pointers
	axisDB.DecodePointers(backRepo, axisNew)

	// get stage instance from DB instance, and call callback function
	axisOld := backRepo.BackRepoAxis.Map_AxisDBID_AxisPtr[axisDB.ID]
	if axisOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), axisOld, axisNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the axisDB
	c.JSON(http.StatusOK, axisDB)
}

// DeleteAxis
//
// swagger:route DELETE /axiss/{ID} axiss deleteAxis
//
// # Delete a axis
//
// default: genericError
//
//	200: axisDBResponse
func (controller *Controller) DeleteAxis(c *gin.Context) {

	mutexAxis.Lock()
	defer mutexAxis.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAxis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAxis.GetDB()

	// Get model if exist
	var axisDB orm.AxisDB
	if err := db.First(&axisDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&axisDB)

	// get an instance (not staged) from DB instance, and call callback function
	axisDeleted := new(models.Axis)
	axisDB.CopyBasicFieldsToAxis(axisDeleted)

	// get stage instance from DB instance, and call callback function
	axisStaged := backRepo.BackRepoAxis.Map_AxisDBID_AxisPtr[axisDB.ID]
	if axisStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), axisStaged, axisDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}