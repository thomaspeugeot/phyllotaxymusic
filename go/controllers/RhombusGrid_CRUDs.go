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
var __RhombusGrid__dummysDeclaration__ models.RhombusGrid
var __RhombusGrid_time__dummyDeclaration time.Duration

var mutexRhombusGrid sync.Mutex

// An RhombusGridID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRhombusGrid updateRhombusGrid deleteRhombusGrid
type RhombusGridID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RhombusGridInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRhombusGrid updateRhombusGrid
type RhombusGridInput struct {
	// The RhombusGrid to submit or modify
	// in: body
	RhombusGrid *orm.RhombusGridAPI
}

// GetRhombusGrids
//
// swagger:route GET /rhombusgrids rhombusgrids getRhombusGrids
//
// # Get all rhombusgrids
//
// Responses:
// default: genericError
//
//	200: rhombusgridDBResponse
func (controller *Controller) GetRhombusGrids(c *gin.Context) {

	// source slice
	var rhombusgridDBs []orm.RhombusGridDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRhombusGrids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phyllotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRhombusGrid.GetDB()

	_, err := db.Find(&rhombusgridDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	rhombusgridAPIs := make([]orm.RhombusGridAPI, 0)

	// for each rhombusgrid, update fields from the database nullable fields
	for idx := range rhombusgridDBs {
		rhombusgridDB := &rhombusgridDBs[idx]
		_ = rhombusgridDB
		var rhombusgridAPI orm.RhombusGridAPI

		// insertion point for updating fields
		rhombusgridAPI.ID = rhombusgridDB.ID
		rhombusgridDB.CopyBasicFieldsToRhombusGrid_WOP(&rhombusgridAPI.RhombusGrid_WOP)
		rhombusgridAPI.RhombusGridPointersEncoding = rhombusgridDB.RhombusGridPointersEncoding
		rhombusgridAPIs = append(rhombusgridAPIs, rhombusgridAPI)
	}

	c.JSON(http.StatusOK, rhombusgridAPIs)
}

// PostRhombusGrid
//
// swagger:route POST /rhombusgrids rhombusgrids postRhombusGrid
//
// Creates a rhombusgrid
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRhombusGrid(c *gin.Context) {

	mutexRhombusGrid.Lock()
	defer mutexRhombusGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRhombusGrids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phyllotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRhombusGrid.GetDB()

	// Validate input
	var input orm.RhombusGridAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create rhombusgrid
	rhombusgridDB := orm.RhombusGridDB{}
	rhombusgridDB.RhombusGridPointersEncoding = input.RhombusGridPointersEncoding
	rhombusgridDB.CopyBasicFieldsFromRhombusGrid_WOP(&input.RhombusGrid_WOP)

	_, err = db.Create(&rhombusgridDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRhombusGrid.CheckoutPhaseOneInstance(&rhombusgridDB)
	rhombusgrid := backRepo.BackRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridPtr[rhombusgridDB.ID]

	if rhombusgrid != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), rhombusgrid)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, rhombusgridDB)
}

// GetRhombusGrid
//
// swagger:route GET /rhombusgrids/{ID} rhombusgrids getRhombusGrid
//
// Gets the details for a rhombusgrid.
//
// Responses:
// default: genericError
//
//	200: rhombusgridDBResponse
func (controller *Controller) GetRhombusGrid(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRhombusGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phyllotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRhombusGrid.GetDB()

	// Get rhombusgridDB in DB
	var rhombusgridDB orm.RhombusGridDB
	if _, err := db.First(&rhombusgridDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var rhombusgridAPI orm.RhombusGridAPI
	rhombusgridAPI.ID = rhombusgridDB.ID
	rhombusgridAPI.RhombusGridPointersEncoding = rhombusgridDB.RhombusGridPointersEncoding
	rhombusgridDB.CopyBasicFieldsToRhombusGrid_WOP(&rhombusgridAPI.RhombusGrid_WOP)

	c.JSON(http.StatusOK, rhombusgridAPI)
}

// UpdateRhombusGrid
//
// swagger:route PATCH /rhombusgrids/{ID} rhombusgrids updateRhombusGrid
//
// # Update a rhombusgrid
//
// Responses:
// default: genericError
//
//	200: rhombusgridDBResponse
func (controller *Controller) UpdateRhombusGrid(c *gin.Context) {

	mutexRhombusGrid.Lock()
	defer mutexRhombusGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRhombusGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phyllotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRhombusGrid.GetDB()

	// Validate input
	var input orm.RhombusGridAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rhombusgridDB orm.RhombusGridDB

	// fetch the rhombusgrid
	_, err := db.First(&rhombusgridDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rhombusgridDB.CopyBasicFieldsFromRhombusGrid_WOP(&input.RhombusGrid_WOP)
	rhombusgridDB.RhombusGridPointersEncoding = input.RhombusGridPointersEncoding

	db, _ = db.Model(&rhombusgridDB)
	_, err = db.Updates(&rhombusgridDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rhombusgridNew := new(models.RhombusGrid)
	rhombusgridDB.CopyBasicFieldsToRhombusGrid(rhombusgridNew)

	// redeem pointers
	rhombusgridDB.DecodePointers(backRepo, rhombusgridNew)

	// get stage instance from DB instance, and call callback function
	rhombusgridOld := backRepo.BackRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridPtr[rhombusgridDB.ID]
	if rhombusgridOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), rhombusgridOld, rhombusgridNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rhombusgridDB
	c.JSON(http.StatusOK, rhombusgridDB)
}

// DeleteRhombusGrid
//
// swagger:route DELETE /rhombusgrids/{ID} rhombusgrids deleteRhombusGrid
//
// # Delete a rhombusgrid
//
// default: genericError
//
//	200: rhombusgridDBResponse
func (controller *Controller) DeleteRhombusGrid(c *gin.Context) {

	mutexRhombusGrid.Lock()
	defer mutexRhombusGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRhombusGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phyllotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRhombusGrid.GetDB()

	// Get model if exist
	var rhombusgridDB orm.RhombusGridDB
	if _, err := db.First(&rhombusgridDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&rhombusgridDB)

	// get an instance (not staged) from DB instance, and call callback function
	rhombusgridDeleted := new(models.RhombusGrid)
	rhombusgridDB.CopyBasicFieldsToRhombusGrid(rhombusgridDeleted)

	// get stage instance from DB instance, and call callback function
	rhombusgridStaged := backRepo.BackRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridPtr[rhombusgridDB.ID]
	if rhombusgridStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), rhombusgridStaged, rhombusgridDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
