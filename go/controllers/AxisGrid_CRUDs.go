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
var __AxisGrid__dummysDeclaration__ models.AxisGrid
var __AxisGrid_time__dummyDeclaration time.Duration

var mutexAxisGrid sync.Mutex

// An AxisGridID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAxisGrid updateAxisGrid deleteAxisGrid
type AxisGridID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AxisGridInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAxisGrid updateAxisGrid
type AxisGridInput struct {
	// The AxisGrid to submit or modify
	// in: body
	AxisGrid *orm.AxisGridAPI
}

// GetAxisGrids
//
// swagger:route GET /axisgrids axisgrids getAxisGrids
//
// # Get all axisgrids
//
// Responses:
// default: genericError
//
//	200: axisgridDBResponse
func (controller *Controller) GetAxisGrids(c *gin.Context) {

	// source slice
	var axisgridDBs []orm.AxisGridDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAxisGrids", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoAxisGrid.GetDB()

	_, err := db.Find(&axisgridDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	axisgridAPIs := make([]orm.AxisGridAPI, 0)

	// for each axisgrid, update fields from the database nullable fields
	for idx := range axisgridDBs {
		axisgridDB := &axisgridDBs[idx]
		_ = axisgridDB
		var axisgridAPI orm.AxisGridAPI

		// insertion point for updating fields
		axisgridAPI.ID = axisgridDB.ID
		axisgridDB.CopyBasicFieldsToAxisGrid_WOP(&axisgridAPI.AxisGrid_WOP)
		axisgridAPI.AxisGridPointersEncoding = axisgridDB.AxisGridPointersEncoding
		axisgridAPIs = append(axisgridAPIs, axisgridAPI)
	}

	c.JSON(http.StatusOK, axisgridAPIs)
}

// PostAxisGrid
//
// swagger:route POST /axisgrids axisgrids postAxisGrid
//
// Creates a axisgrid
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAxisGrid(c *gin.Context) {

	mutexAxisGrid.Lock()
	defer mutexAxisGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAxisGrids", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoAxisGrid.GetDB()

	// Validate input
	var input orm.AxisGridAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create axisgrid
	axisgridDB := orm.AxisGridDB{}
	axisgridDB.AxisGridPointersEncoding = input.AxisGridPointersEncoding
	axisgridDB.CopyBasicFieldsFromAxisGrid_WOP(&input.AxisGrid_WOP)

	_, err = db.Create(&axisgridDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAxisGrid.CheckoutPhaseOneInstance(&axisgridDB)
	axisgrid := backRepo.BackRepoAxisGrid.Map_AxisGridDBID_AxisGridPtr[axisgridDB.ID]

	if axisgrid != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), axisgrid)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, axisgridDB)
}

// GetAxisGrid
//
// swagger:route GET /axisgrids/{ID} axisgrids getAxisGrid
//
// Gets the details for a axisgrid.
//
// Responses:
// default: genericError
//
//	200: axisgridDBResponse
func (controller *Controller) GetAxisGrid(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAxisGrid", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoAxisGrid.GetDB()

	// Get axisgridDB in DB
	var axisgridDB orm.AxisGridDB
	if _, err := db.First(&axisgridDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var axisgridAPI orm.AxisGridAPI
	axisgridAPI.ID = axisgridDB.ID
	axisgridAPI.AxisGridPointersEncoding = axisgridDB.AxisGridPointersEncoding
	axisgridDB.CopyBasicFieldsToAxisGrid_WOP(&axisgridAPI.AxisGrid_WOP)

	c.JSON(http.StatusOK, axisgridAPI)
}

// UpdateAxisGrid
//
// swagger:route PATCH /axisgrids/{ID} axisgrids updateAxisGrid
//
// # Update a axisgrid
//
// Responses:
// default: genericError
//
//	200: axisgridDBResponse
func (controller *Controller) UpdateAxisGrid(c *gin.Context) {

	mutexAxisGrid.Lock()
	defer mutexAxisGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAxisGrid", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoAxisGrid.GetDB()

	// Validate input
	var input orm.AxisGridAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var axisgridDB orm.AxisGridDB

	// fetch the axisgrid
	_, err := db.First(&axisgridDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	axisgridDB.CopyBasicFieldsFromAxisGrid_WOP(&input.AxisGrid_WOP)
	axisgridDB.AxisGridPointersEncoding = input.AxisGridPointersEncoding

	db, _ = db.Model(&axisgridDB)
	_, err = db.Updates(&axisgridDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	axisgridNew := new(models.AxisGrid)
	axisgridDB.CopyBasicFieldsToAxisGrid(axisgridNew)

	// redeem pointers
	axisgridDB.DecodePointers(backRepo, axisgridNew)

	// get stage instance from DB instance, and call callback function
	axisgridOld := backRepo.BackRepoAxisGrid.Map_AxisGridDBID_AxisGridPtr[axisgridDB.ID]
	if axisgridOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), axisgridOld, axisgridNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the axisgridDB
	c.JSON(http.StatusOK, axisgridDB)
}

// DeleteAxisGrid
//
// swagger:route DELETE /axisgrids/{ID} axisgrids deleteAxisGrid
//
// # Delete a axisgrid
//
// default: genericError
//
//	200: axisgridDBResponse
func (controller *Controller) DeleteAxisGrid(c *gin.Context) {

	mutexAxisGrid.Lock()
	defer mutexAxisGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAxisGrid", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoAxisGrid.GetDB()

	// Get model if exist
	var axisgridDB orm.AxisGridDB
	if _, err := db.First(&axisgridDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&axisgridDB)

	// get an instance (not staged) from DB instance, and call callback function
	axisgridDeleted := new(models.AxisGrid)
	axisgridDB.CopyBasicFieldsToAxisGrid(axisgridDeleted)

	// get stage instance from DB instance, and call callback function
	axisgridStaged := backRepo.BackRepoAxisGrid.Map_AxisGridDBID_AxisGridPtr[axisgridDB.ID]
	if axisgridStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), axisgridStaged, axisgridDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
