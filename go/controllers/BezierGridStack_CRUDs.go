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
var __BezierGridStack__dummysDeclaration__ models.BezierGridStack
var __BezierGridStack_time__dummyDeclaration time.Duration

var mutexBezierGridStack sync.Mutex

// An BezierGridStackID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBezierGridStack updateBezierGridStack deleteBezierGridStack
type BezierGridStackID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BezierGridStackInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBezierGridStack updateBezierGridStack
type BezierGridStackInput struct {
	// The BezierGridStack to submit or modify
	// in: body
	BezierGridStack *orm.BezierGridStackAPI
}

// GetBezierGridStacks
//
// swagger:route GET /beziergridstacks beziergridstacks getBezierGridStacks
//
// # Get all beziergridstacks
//
// Responses:
// default: genericError
//
//	200: beziergridstackDBResponse
func (controller *Controller) GetBezierGridStacks(c *gin.Context) {

	// source slice
	var beziergridstackDBs []orm.BezierGridStackDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBezierGridStacks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBezierGridStack.GetDB()

	query := db.Find(&beziergridstackDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	beziergridstackAPIs := make([]orm.BezierGridStackAPI, 0)

	// for each beziergridstack, update fields from the database nullable fields
	for idx := range beziergridstackDBs {
		beziergridstackDB := &beziergridstackDBs[idx]
		_ = beziergridstackDB
		var beziergridstackAPI orm.BezierGridStackAPI

		// insertion point for updating fields
		beziergridstackAPI.ID = beziergridstackDB.ID
		beziergridstackDB.CopyBasicFieldsToBezierGridStack_WOP(&beziergridstackAPI.BezierGridStack_WOP)
		beziergridstackAPI.BezierGridStackPointersEncoding = beziergridstackDB.BezierGridStackPointersEncoding
		beziergridstackAPIs = append(beziergridstackAPIs, beziergridstackAPI)
	}

	c.JSON(http.StatusOK, beziergridstackAPIs)
}

// PostBezierGridStack
//
// swagger:route POST /beziergridstacks beziergridstacks postBezierGridStack
//
// Creates a beziergridstack
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBezierGridStack(c *gin.Context) {

	mutexBezierGridStack.Lock()
	defer mutexBezierGridStack.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBezierGridStacks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBezierGridStack.GetDB()

	// Validate input
	var input orm.BezierGridStackAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create beziergridstack
	beziergridstackDB := orm.BezierGridStackDB{}
	beziergridstackDB.BezierGridStackPointersEncoding = input.BezierGridStackPointersEncoding
	beziergridstackDB.CopyBasicFieldsFromBezierGridStack_WOP(&input.BezierGridStack_WOP)

	query := db.Create(&beziergridstackDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBezierGridStack.CheckoutPhaseOneInstance(&beziergridstackDB)
	beziergridstack := backRepo.BackRepoBezierGridStack.Map_BezierGridStackDBID_BezierGridStackPtr[beziergridstackDB.ID]

	if beziergridstack != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), beziergridstack)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, beziergridstackDB)
}

// GetBezierGridStack
//
// swagger:route GET /beziergridstacks/{ID} beziergridstacks getBezierGridStack
//
// Gets the details for a beziergridstack.
//
// Responses:
// default: genericError
//
//	200: beziergridstackDBResponse
func (controller *Controller) GetBezierGridStack(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBezierGridStack", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBezierGridStack.GetDB()

	// Get beziergridstackDB in DB
	var beziergridstackDB orm.BezierGridStackDB
	if err := db.First(&beziergridstackDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var beziergridstackAPI orm.BezierGridStackAPI
	beziergridstackAPI.ID = beziergridstackDB.ID
	beziergridstackAPI.BezierGridStackPointersEncoding = beziergridstackDB.BezierGridStackPointersEncoding
	beziergridstackDB.CopyBasicFieldsToBezierGridStack_WOP(&beziergridstackAPI.BezierGridStack_WOP)

	c.JSON(http.StatusOK, beziergridstackAPI)
}

// UpdateBezierGridStack
//
// swagger:route PATCH /beziergridstacks/{ID} beziergridstacks updateBezierGridStack
//
// # Update a beziergridstack
//
// Responses:
// default: genericError
//
//	200: beziergridstackDBResponse
func (controller *Controller) UpdateBezierGridStack(c *gin.Context) {

	mutexBezierGridStack.Lock()
	defer mutexBezierGridStack.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBezierGridStack", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBezierGridStack.GetDB()

	// Validate input
	var input orm.BezierGridStackAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var beziergridstackDB orm.BezierGridStackDB

	// fetch the beziergridstack
	query := db.First(&beziergridstackDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	beziergridstackDB.CopyBasicFieldsFromBezierGridStack_WOP(&input.BezierGridStack_WOP)
	beziergridstackDB.BezierGridStackPointersEncoding = input.BezierGridStackPointersEncoding

	query = db.Model(&beziergridstackDB).Updates(beziergridstackDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	beziergridstackNew := new(models.BezierGridStack)
	beziergridstackDB.CopyBasicFieldsToBezierGridStack(beziergridstackNew)

	// redeem pointers
	beziergridstackDB.DecodePointers(backRepo, beziergridstackNew)

	// get stage instance from DB instance, and call callback function
	beziergridstackOld := backRepo.BackRepoBezierGridStack.Map_BezierGridStackDBID_BezierGridStackPtr[beziergridstackDB.ID]
	if beziergridstackOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), beziergridstackOld, beziergridstackNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the beziergridstackDB
	c.JSON(http.StatusOK, beziergridstackDB)
}

// DeleteBezierGridStack
//
// swagger:route DELETE /beziergridstacks/{ID} beziergridstacks deleteBezierGridStack
//
// # Delete a beziergridstack
//
// default: genericError
//
//	200: beziergridstackDBResponse
func (controller *Controller) DeleteBezierGridStack(c *gin.Context) {

	mutexBezierGridStack.Lock()
	defer mutexBezierGridStack.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBezierGridStack", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBezierGridStack.GetDB()

	// Get model if exist
	var beziergridstackDB orm.BezierGridStackDB
	if err := db.First(&beziergridstackDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&beziergridstackDB)

	// get an instance (not staged) from DB instance, and call callback function
	beziergridstackDeleted := new(models.BezierGridStack)
	beziergridstackDB.CopyBasicFieldsToBezierGridStack(beziergridstackDeleted)

	// get stage instance from DB instance, and call callback function
	beziergridstackStaged := backRepo.BackRepoBezierGridStack.Map_BezierGridStackDBID_BezierGridStackPtr[beziergridstackDB.ID]
	if beziergridstackStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), beziergridstackStaged, beziergridstackDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
