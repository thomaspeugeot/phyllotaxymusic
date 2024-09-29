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
var __SpiralAxisGrid__dummysDeclaration__ models.SpiralAxisGrid
var __SpiralAxisGrid_time__dummyDeclaration time.Duration

var mutexSpiralAxisGrid sync.Mutex

// An SpiralAxisGridID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSpiralAxisGrid updateSpiralAxisGrid deleteSpiralAxisGrid
type SpiralAxisGridID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SpiralAxisGridInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSpiralAxisGrid updateSpiralAxisGrid
type SpiralAxisGridInput struct {
	// The SpiralAxisGrid to submit or modify
	// in: body
	SpiralAxisGrid *orm.SpiralAxisGridAPI
}

// GetSpiralAxisGrids
//
// swagger:route GET /spiralaxisgrids spiralaxisgrids getSpiralAxisGrids
//
// # Get all spiralaxisgrids
//
// Responses:
// default: genericError
//
//	200: spiralaxisgridDBResponse
func (controller *Controller) GetSpiralAxisGrids(c *gin.Context) {

	// source slice
	var spiralaxisgridDBs []orm.SpiralAxisGridDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralAxisGrids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralAxisGrid.GetDB()

	query := db.Find(&spiralaxisgridDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spiralaxisgridAPIs := make([]orm.SpiralAxisGridAPI, 0)

	// for each spiralaxisgrid, update fields from the database nullable fields
	for idx := range spiralaxisgridDBs {
		spiralaxisgridDB := &spiralaxisgridDBs[idx]
		_ = spiralaxisgridDB
		var spiralaxisgridAPI orm.SpiralAxisGridAPI

		// insertion point for updating fields
		spiralaxisgridAPI.ID = spiralaxisgridDB.ID
		spiralaxisgridDB.CopyBasicFieldsToSpiralAxisGrid_WOP(&spiralaxisgridAPI.SpiralAxisGrid_WOP)
		spiralaxisgridAPI.SpiralAxisGridPointersEncoding = spiralaxisgridDB.SpiralAxisGridPointersEncoding
		spiralaxisgridAPIs = append(spiralaxisgridAPIs, spiralaxisgridAPI)
	}

	c.JSON(http.StatusOK, spiralaxisgridAPIs)
}

// PostSpiralAxisGrid
//
// swagger:route POST /spiralaxisgrids spiralaxisgrids postSpiralAxisGrid
//
// Creates a spiralaxisgrid
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSpiralAxisGrid(c *gin.Context) {

	mutexSpiralAxisGrid.Lock()
	defer mutexSpiralAxisGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSpiralAxisGrids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralAxisGrid.GetDB()

	// Validate input
	var input orm.SpiralAxisGridAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spiralaxisgrid
	spiralaxisgridDB := orm.SpiralAxisGridDB{}
	spiralaxisgridDB.SpiralAxisGridPointersEncoding = input.SpiralAxisGridPointersEncoding
	spiralaxisgridDB.CopyBasicFieldsFromSpiralAxisGrid_WOP(&input.SpiralAxisGrid_WOP)

	query := db.Create(&spiralaxisgridDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSpiralAxisGrid.CheckoutPhaseOneInstance(&spiralaxisgridDB)
	spiralaxisgrid := backRepo.BackRepoSpiralAxisGrid.Map_SpiralAxisGridDBID_SpiralAxisGridPtr[spiralaxisgridDB.ID]

	if spiralaxisgrid != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spiralaxisgrid)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spiralaxisgridDB)
}

// GetSpiralAxisGrid
//
// swagger:route GET /spiralaxisgrids/{ID} spiralaxisgrids getSpiralAxisGrid
//
// Gets the details for a spiralaxisgrid.
//
// Responses:
// default: genericError
//
//	200: spiralaxisgridDBResponse
func (controller *Controller) GetSpiralAxisGrid(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralAxisGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralAxisGrid.GetDB()

	// Get spiralaxisgridDB in DB
	var spiralaxisgridDB orm.SpiralAxisGridDB
	if err := db.First(&spiralaxisgridDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spiralaxisgridAPI orm.SpiralAxisGridAPI
	spiralaxisgridAPI.ID = spiralaxisgridDB.ID
	spiralaxisgridAPI.SpiralAxisGridPointersEncoding = spiralaxisgridDB.SpiralAxisGridPointersEncoding
	spiralaxisgridDB.CopyBasicFieldsToSpiralAxisGrid_WOP(&spiralaxisgridAPI.SpiralAxisGrid_WOP)

	c.JSON(http.StatusOK, spiralaxisgridAPI)
}

// UpdateSpiralAxisGrid
//
// swagger:route PATCH /spiralaxisgrids/{ID} spiralaxisgrids updateSpiralAxisGrid
//
// # Update a spiralaxisgrid
//
// Responses:
// default: genericError
//
//	200: spiralaxisgridDBResponse
func (controller *Controller) UpdateSpiralAxisGrid(c *gin.Context) {

	mutexSpiralAxisGrid.Lock()
	defer mutexSpiralAxisGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSpiralAxisGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralAxisGrid.GetDB()

	// Validate input
	var input orm.SpiralAxisGridAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spiralaxisgridDB orm.SpiralAxisGridDB

	// fetch the spiralaxisgrid
	query := db.First(&spiralaxisgridDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spiralaxisgridDB.CopyBasicFieldsFromSpiralAxisGrid_WOP(&input.SpiralAxisGrid_WOP)
	spiralaxisgridDB.SpiralAxisGridPointersEncoding = input.SpiralAxisGridPointersEncoding

	query = db.Model(&spiralaxisgridDB).Updates(spiralaxisgridDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spiralaxisgridNew := new(models.SpiralAxisGrid)
	spiralaxisgridDB.CopyBasicFieldsToSpiralAxisGrid(spiralaxisgridNew)

	// redeem pointers
	spiralaxisgridDB.DecodePointers(backRepo, spiralaxisgridNew)

	// get stage instance from DB instance, and call callback function
	spiralaxisgridOld := backRepo.BackRepoSpiralAxisGrid.Map_SpiralAxisGridDBID_SpiralAxisGridPtr[spiralaxisgridDB.ID]
	if spiralaxisgridOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), spiralaxisgridOld, spiralaxisgridNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spiralaxisgridDB
	c.JSON(http.StatusOK, spiralaxisgridDB)
}

// DeleteSpiralAxisGrid
//
// swagger:route DELETE /spiralaxisgrids/{ID} spiralaxisgrids deleteSpiralAxisGrid
//
// # Delete a spiralaxisgrid
//
// default: genericError
//
//	200: spiralaxisgridDBResponse
func (controller *Controller) DeleteSpiralAxisGrid(c *gin.Context) {

	mutexSpiralAxisGrid.Lock()
	defer mutexSpiralAxisGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSpiralAxisGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralAxisGrid.GetDB()

	// Get model if exist
	var spiralaxisgridDB orm.SpiralAxisGridDB
	if err := db.First(&spiralaxisgridDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&spiralaxisgridDB)

	// get an instance (not staged) from DB instance, and call callback function
	spiralaxisgridDeleted := new(models.SpiralAxisGrid)
	spiralaxisgridDB.CopyBasicFieldsToSpiralAxisGrid(spiralaxisgridDeleted)

	// get stage instance from DB instance, and call callback function
	spiralaxisgridStaged := backRepo.BackRepoSpiralAxisGrid.Map_SpiralAxisGridDBID_SpiralAxisGridPtr[spiralaxisgridDB.ID]
	if spiralaxisgridStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spiralaxisgridStaged, spiralaxisgridDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
