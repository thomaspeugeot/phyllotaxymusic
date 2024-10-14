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
var __SpiralLineGrid__dummysDeclaration__ models.SpiralLineGrid
var __SpiralLineGrid_time__dummyDeclaration time.Duration

var mutexSpiralLineGrid sync.Mutex

// An SpiralLineGridID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSpiralLineGrid updateSpiralLineGrid deleteSpiralLineGrid
type SpiralLineGridID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SpiralLineGridInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSpiralLineGrid updateSpiralLineGrid
type SpiralLineGridInput struct {
	// The SpiralLineGrid to submit or modify
	// in: body
	SpiralLineGrid *orm.SpiralLineGridAPI
}

// GetSpiralLineGrids
//
// swagger:route GET /spirallinegrids spirallinegrids getSpiralLineGrids
//
// # Get all spirallinegrids
//
// Responses:
// default: genericError
//
//	200: spirallinegridDBResponse
func (controller *Controller) GetSpiralLineGrids(c *gin.Context) {

	// source slice
	var spirallinegridDBs []orm.SpiralLineGridDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralLineGrids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralLineGrid.GetDB()

	query := db.Find(&spirallinegridDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spirallinegridAPIs := make([]orm.SpiralLineGridAPI, 0)

	// for each spirallinegrid, update fields from the database nullable fields
	for idx := range spirallinegridDBs {
		spirallinegridDB := &spirallinegridDBs[idx]
		_ = spirallinegridDB
		var spirallinegridAPI orm.SpiralLineGridAPI

		// insertion point for updating fields
		spirallinegridAPI.ID = spirallinegridDB.ID
		spirallinegridDB.CopyBasicFieldsToSpiralLineGrid_WOP(&spirallinegridAPI.SpiralLineGrid_WOP)
		spirallinegridAPI.SpiralLineGridPointersEncoding = spirallinegridDB.SpiralLineGridPointersEncoding
		spirallinegridAPIs = append(spirallinegridAPIs, spirallinegridAPI)
	}

	c.JSON(http.StatusOK, spirallinegridAPIs)
}

// PostSpiralLineGrid
//
// swagger:route POST /spirallinegrids spirallinegrids postSpiralLineGrid
//
// Creates a spirallinegrid
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSpiralLineGrid(c *gin.Context) {

	mutexSpiralLineGrid.Lock()
	defer mutexSpiralLineGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSpiralLineGrids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralLineGrid.GetDB()

	// Validate input
	var input orm.SpiralLineGridAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spirallinegrid
	spirallinegridDB := orm.SpiralLineGridDB{}
	spirallinegridDB.SpiralLineGridPointersEncoding = input.SpiralLineGridPointersEncoding
	spirallinegridDB.CopyBasicFieldsFromSpiralLineGrid_WOP(&input.SpiralLineGrid_WOP)

	query := db.Create(&spirallinegridDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSpiralLineGrid.CheckoutPhaseOneInstance(&spirallinegridDB)
	spirallinegrid := backRepo.BackRepoSpiralLineGrid.Map_SpiralLineGridDBID_SpiralLineGridPtr[spirallinegridDB.ID]

	if spirallinegrid != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spirallinegrid)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spirallinegridDB)
}

// GetSpiralLineGrid
//
// swagger:route GET /spirallinegrids/{ID} spirallinegrids getSpiralLineGrid
//
// Gets the details for a spirallinegrid.
//
// Responses:
// default: genericError
//
//	200: spirallinegridDBResponse
func (controller *Controller) GetSpiralLineGrid(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralLineGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralLineGrid.GetDB()

	// Get spirallinegridDB in DB
	var spirallinegridDB orm.SpiralLineGridDB
	if err := db.First(&spirallinegridDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spirallinegridAPI orm.SpiralLineGridAPI
	spirallinegridAPI.ID = spirallinegridDB.ID
	spirallinegridAPI.SpiralLineGridPointersEncoding = spirallinegridDB.SpiralLineGridPointersEncoding
	spirallinegridDB.CopyBasicFieldsToSpiralLineGrid_WOP(&spirallinegridAPI.SpiralLineGrid_WOP)

	c.JSON(http.StatusOK, spirallinegridAPI)
}

// UpdateSpiralLineGrid
//
// swagger:route PATCH /spirallinegrids/{ID} spirallinegrids updateSpiralLineGrid
//
// # Update a spirallinegrid
//
// Responses:
// default: genericError
//
//	200: spirallinegridDBResponse
func (controller *Controller) UpdateSpiralLineGrid(c *gin.Context) {

	mutexSpiralLineGrid.Lock()
	defer mutexSpiralLineGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSpiralLineGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralLineGrid.GetDB()

	// Validate input
	var input orm.SpiralLineGridAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spirallinegridDB orm.SpiralLineGridDB

	// fetch the spirallinegrid
	query := db.First(&spirallinegridDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spirallinegridDB.CopyBasicFieldsFromSpiralLineGrid_WOP(&input.SpiralLineGrid_WOP)
	spirallinegridDB.SpiralLineGridPointersEncoding = input.SpiralLineGridPointersEncoding

	query = db.Model(&spirallinegridDB).Updates(spirallinegridDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spirallinegridNew := new(models.SpiralLineGrid)
	spirallinegridDB.CopyBasicFieldsToSpiralLineGrid(spirallinegridNew)

	// redeem pointers
	spirallinegridDB.DecodePointers(backRepo, spirallinegridNew)

	// get stage instance from DB instance, and call callback function
	spirallinegridOld := backRepo.BackRepoSpiralLineGrid.Map_SpiralLineGridDBID_SpiralLineGridPtr[spirallinegridDB.ID]
	if spirallinegridOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), spirallinegridOld, spirallinegridNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spirallinegridDB
	c.JSON(http.StatusOK, spirallinegridDB)
}

// DeleteSpiralLineGrid
//
// swagger:route DELETE /spirallinegrids/{ID} spirallinegrids deleteSpiralLineGrid
//
// # Delete a spirallinegrid
//
// default: genericError
//
//	200: spirallinegridDBResponse
func (controller *Controller) DeleteSpiralLineGrid(c *gin.Context) {

	mutexSpiralLineGrid.Lock()
	defer mutexSpiralLineGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSpiralLineGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralLineGrid.GetDB()

	// Get model if exist
	var spirallinegridDB orm.SpiralLineGridDB
	if err := db.First(&spirallinegridDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&spirallinegridDB)

	// get an instance (not staged) from DB instance, and call callback function
	spirallinegridDeleted := new(models.SpiralLineGrid)
	spirallinegridDB.CopyBasicFieldsToSpiralLineGrid(spirallinegridDeleted)

	// get stage instance from DB instance, and call callback function
	spirallinegridStaged := backRepo.BackRepoSpiralLineGrid.Map_SpiralLineGridDBID_SpiralLineGridPtr[spirallinegridDB.ID]
	if spirallinegridStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spirallinegridStaged, spirallinegridDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
