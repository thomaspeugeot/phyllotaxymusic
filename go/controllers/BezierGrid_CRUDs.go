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
var __BezierGrid__dummysDeclaration__ models.BezierGrid
var __BezierGrid_time__dummyDeclaration time.Duration

var mutexBezierGrid sync.Mutex

// An BezierGridID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBezierGrid updateBezierGrid deleteBezierGrid
type BezierGridID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BezierGridInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBezierGrid updateBezierGrid
type BezierGridInput struct {
	// The BezierGrid to submit or modify
	// in: body
	BezierGrid *orm.BezierGridAPI
}

// GetBezierGrids
//
// swagger:route GET /beziergrids beziergrids getBezierGrids
//
// # Get all beziergrids
//
// Responses:
// default: genericError
//
//	200: beziergridDBResponse
func (controller *Controller) GetBezierGrids(c *gin.Context) {

	// source slice
	var beziergridDBs []orm.BezierGridDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBezierGrids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoBezierGrid.GetDB()

	_, err := db.Find(&beziergridDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	beziergridAPIs := make([]orm.BezierGridAPI, 0)

	// for each beziergrid, update fields from the database nullable fields
	for idx := range beziergridDBs {
		beziergridDB := &beziergridDBs[idx]
		_ = beziergridDB
		var beziergridAPI orm.BezierGridAPI

		// insertion point for updating fields
		beziergridAPI.ID = beziergridDB.ID
		beziergridDB.CopyBasicFieldsToBezierGrid_WOP(&beziergridAPI.BezierGrid_WOP)
		beziergridAPI.BezierGridPointersEncoding = beziergridDB.BezierGridPointersEncoding
		beziergridAPIs = append(beziergridAPIs, beziergridAPI)
	}

	c.JSON(http.StatusOK, beziergridAPIs)
}

// PostBezierGrid
//
// swagger:route POST /beziergrids beziergrids postBezierGrid
//
// Creates a beziergrid
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBezierGrid(c *gin.Context) {

	mutexBezierGrid.Lock()
	defer mutexBezierGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBezierGrids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoBezierGrid.GetDB()

	// Validate input
	var input orm.BezierGridAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create beziergrid
	beziergridDB := orm.BezierGridDB{}
	beziergridDB.BezierGridPointersEncoding = input.BezierGridPointersEncoding
	beziergridDB.CopyBasicFieldsFromBezierGrid_WOP(&input.BezierGrid_WOP)

	_, err = db.Create(&beziergridDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBezierGrid.CheckoutPhaseOneInstance(&beziergridDB)
	beziergrid := backRepo.BackRepoBezierGrid.Map_BezierGridDBID_BezierGridPtr[beziergridDB.ID]

	if beziergrid != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), beziergrid)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, beziergridDB)
}

// GetBezierGrid
//
// swagger:route GET /beziergrids/{ID} beziergrids getBezierGrid
//
// Gets the details for a beziergrid.
//
// Responses:
// default: genericError
//
//	200: beziergridDBResponse
func (controller *Controller) GetBezierGrid(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBezierGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoBezierGrid.GetDB()

	// Get beziergridDB in DB
	var beziergridDB orm.BezierGridDB
	if _, err := db.First(&beziergridDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var beziergridAPI orm.BezierGridAPI
	beziergridAPI.ID = beziergridDB.ID
	beziergridAPI.BezierGridPointersEncoding = beziergridDB.BezierGridPointersEncoding
	beziergridDB.CopyBasicFieldsToBezierGrid_WOP(&beziergridAPI.BezierGrid_WOP)

	c.JSON(http.StatusOK, beziergridAPI)
}

// UpdateBezierGrid
//
// swagger:route PATCH /beziergrids/{ID} beziergrids updateBezierGrid
//
// # Update a beziergrid
//
// Responses:
// default: genericError
//
//	200: beziergridDBResponse
func (controller *Controller) UpdateBezierGrid(c *gin.Context) {

	mutexBezierGrid.Lock()
	defer mutexBezierGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBezierGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoBezierGrid.GetDB()

	// Validate input
	var input orm.BezierGridAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var beziergridDB orm.BezierGridDB

	// fetch the beziergrid
	_, err := db.First(&beziergridDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	beziergridDB.CopyBasicFieldsFromBezierGrid_WOP(&input.BezierGrid_WOP)
	beziergridDB.BezierGridPointersEncoding = input.BezierGridPointersEncoding

	db, _ = db.Model(&beziergridDB)
	_, err = db.Updates(&beziergridDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	beziergridNew := new(models.BezierGrid)
	beziergridDB.CopyBasicFieldsToBezierGrid(beziergridNew)

	// redeem pointers
	beziergridDB.DecodePointers(backRepo, beziergridNew)

	// get stage instance from DB instance, and call callback function
	beziergridOld := backRepo.BackRepoBezierGrid.Map_BezierGridDBID_BezierGridPtr[beziergridDB.ID]
	if beziergridOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), beziergridOld, beziergridNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the beziergridDB
	c.JSON(http.StatusOK, beziergridDB)
}

// DeleteBezierGrid
//
// swagger:route DELETE /beziergrids/{ID} beziergrids deleteBezierGrid
//
// # Delete a beziergrid
//
// default: genericError
//
//	200: beziergridDBResponse
func (controller *Controller) DeleteBezierGrid(c *gin.Context) {

	mutexBezierGrid.Lock()
	defer mutexBezierGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBezierGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoBezierGrid.GetDB()

	// Get model if exist
	var beziergridDB orm.BezierGridDB
	if _, err := db.First(&beziergridDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&beziergridDB)

	// get an instance (not staged) from DB instance, and call callback function
	beziergridDeleted := new(models.BezierGrid)
	beziergridDB.CopyBasicFieldsToBezierGrid(beziergridDeleted)

	// get stage instance from DB instance, and call callback function
	beziergridStaged := backRepo.BackRepoBezierGrid.Map_BezierGridDBID_BezierGridPtr[beziergridDB.ID]
	if beziergridStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), beziergridStaged, beziergridDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
