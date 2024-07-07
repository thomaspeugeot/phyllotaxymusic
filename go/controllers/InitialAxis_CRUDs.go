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
var __InitialAxis__dummysDeclaration__ models.InitialAxis
var __InitialAxis_time__dummyDeclaration time.Duration

var mutexInitialAxis sync.Mutex

// An InitialAxisID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getInitialAxis updateInitialAxis deleteInitialAxis
type InitialAxisID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// InitialAxisInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postInitialAxis updateInitialAxis
type InitialAxisInput struct {
	// The InitialAxis to submit or modify
	// in: body
	InitialAxis *orm.InitialAxisAPI
}

// GetInitialAxiss
//
// swagger:route GET /initialaxiss initialaxiss getInitialAxiss
//
// # Get all initialaxiss
//
// Responses:
// default: genericError
//
//	200: initialaxisDBResponse
func (controller *Controller) GetInitialAxiss(c *gin.Context) {

	// source slice
	var initialaxisDBs []orm.InitialAxisDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetInitialAxiss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInitialAxis.GetDB()

	query := db.Find(&initialaxisDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	initialaxisAPIs := make([]orm.InitialAxisAPI, 0)

	// for each initialaxis, update fields from the database nullable fields
	for idx := range initialaxisDBs {
		initialaxisDB := &initialaxisDBs[idx]
		_ = initialaxisDB
		var initialaxisAPI orm.InitialAxisAPI

		// insertion point for updating fields
		initialaxisAPI.ID = initialaxisDB.ID
		initialaxisDB.CopyBasicFieldsToInitialAxis_WOP(&initialaxisAPI.InitialAxis_WOP)
		initialaxisAPI.InitialAxisPointersEncoding = initialaxisDB.InitialAxisPointersEncoding
		initialaxisAPIs = append(initialaxisAPIs, initialaxisAPI)
	}

	c.JSON(http.StatusOK, initialaxisAPIs)
}

// PostInitialAxis
//
// swagger:route POST /initialaxiss initialaxiss postInitialAxis
//
// Creates a initialaxis
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostInitialAxis(c *gin.Context) {

	mutexInitialAxis.Lock()
	defer mutexInitialAxis.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostInitialAxiss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInitialAxis.GetDB()

	// Validate input
	var input orm.InitialAxisAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create initialaxis
	initialaxisDB := orm.InitialAxisDB{}
	initialaxisDB.InitialAxisPointersEncoding = input.InitialAxisPointersEncoding
	initialaxisDB.CopyBasicFieldsFromInitialAxis_WOP(&input.InitialAxis_WOP)

	query := db.Create(&initialaxisDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoInitialAxis.CheckoutPhaseOneInstance(&initialaxisDB)
	initialaxis := backRepo.BackRepoInitialAxis.Map_InitialAxisDBID_InitialAxisPtr[initialaxisDB.ID]

	if initialaxis != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), initialaxis)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, initialaxisDB)
}

// GetInitialAxis
//
// swagger:route GET /initialaxiss/{ID} initialaxiss getInitialAxis
//
// Gets the details for a initialaxis.
//
// Responses:
// default: genericError
//
//	200: initialaxisDBResponse
func (controller *Controller) GetInitialAxis(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetInitialAxis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInitialAxis.GetDB()

	// Get initialaxisDB in DB
	var initialaxisDB orm.InitialAxisDB
	if err := db.First(&initialaxisDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var initialaxisAPI orm.InitialAxisAPI
	initialaxisAPI.ID = initialaxisDB.ID
	initialaxisAPI.InitialAxisPointersEncoding = initialaxisDB.InitialAxisPointersEncoding
	initialaxisDB.CopyBasicFieldsToInitialAxis_WOP(&initialaxisAPI.InitialAxis_WOP)

	c.JSON(http.StatusOK, initialaxisAPI)
}

// UpdateInitialAxis
//
// swagger:route PATCH /initialaxiss/{ID} initialaxiss updateInitialAxis
//
// # Update a initialaxis
//
// Responses:
// default: genericError
//
//	200: initialaxisDBResponse
func (controller *Controller) UpdateInitialAxis(c *gin.Context) {

	mutexInitialAxis.Lock()
	defer mutexInitialAxis.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateInitialAxis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInitialAxis.GetDB()

	// Validate input
	var input orm.InitialAxisAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var initialaxisDB orm.InitialAxisDB

	// fetch the initialaxis
	query := db.First(&initialaxisDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	initialaxisDB.CopyBasicFieldsFromInitialAxis_WOP(&input.InitialAxis_WOP)
	initialaxisDB.InitialAxisPointersEncoding = input.InitialAxisPointersEncoding

	query = db.Model(&initialaxisDB).Updates(initialaxisDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	initialaxisNew := new(models.InitialAxis)
	initialaxisDB.CopyBasicFieldsToInitialAxis(initialaxisNew)

	// redeem pointers
	initialaxisDB.DecodePointers(backRepo, initialaxisNew)

	// get stage instance from DB instance, and call callback function
	initialaxisOld := backRepo.BackRepoInitialAxis.Map_InitialAxisDBID_InitialAxisPtr[initialaxisDB.ID]
	if initialaxisOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), initialaxisOld, initialaxisNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the initialaxisDB
	c.JSON(http.StatusOK, initialaxisDB)
}

// DeleteInitialAxis
//
// swagger:route DELETE /initialaxiss/{ID} initialaxiss deleteInitialAxis
//
// # Delete a initialaxis
//
// default: genericError
//
//	200: initialaxisDBResponse
func (controller *Controller) DeleteInitialAxis(c *gin.Context) {

	mutexInitialAxis.Lock()
	defer mutexInitialAxis.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteInitialAxis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInitialAxis.GetDB()

	// Get model if exist
	var initialaxisDB orm.InitialAxisDB
	if err := db.First(&initialaxisDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&initialaxisDB)

	// get an instance (not staged) from DB instance, and call callback function
	initialaxisDeleted := new(models.InitialAxis)
	initialaxisDB.CopyBasicFieldsToInitialAxis(initialaxisDeleted)

	// get stage instance from DB instance, and call callback function
	initialaxisStaged := backRepo.BackRepoInitialAxis.Map_InitialAxisDBID_InitialAxisPtr[initialaxisDB.ID]
	if initialaxisStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), initialaxisStaged, initialaxisDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
