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
var __SpiralRhombusGrid__dummysDeclaration__ models.SpiralRhombusGrid
var __SpiralRhombusGrid_time__dummyDeclaration time.Duration

var mutexSpiralRhombusGrid sync.Mutex

// An SpiralRhombusGridID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSpiralRhombusGrid updateSpiralRhombusGrid deleteSpiralRhombusGrid
type SpiralRhombusGridID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SpiralRhombusGridInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSpiralRhombusGrid updateSpiralRhombusGrid
type SpiralRhombusGridInput struct {
	// The SpiralRhombusGrid to submit or modify
	// in: body
	SpiralRhombusGrid *orm.SpiralRhombusGridAPI
}

// GetSpiralRhombusGrids
//
// swagger:route GET /spiralrhombusgrids spiralrhombusgrids getSpiralRhombusGrids
//
// # Get all spiralrhombusgrids
//
// Responses:
// default: genericError
//
//	200: spiralrhombusgridDBResponse
func (controller *Controller) GetSpiralRhombusGrids(c *gin.Context) {

	// source slice
	var spiralrhombusgridDBs []orm.SpiralRhombusGridDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralRhombusGrids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralRhombusGrid.GetDB()

	query := db.Find(&spiralrhombusgridDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spiralrhombusgridAPIs := make([]orm.SpiralRhombusGridAPI, 0)

	// for each spiralrhombusgrid, update fields from the database nullable fields
	for idx := range spiralrhombusgridDBs {
		spiralrhombusgridDB := &spiralrhombusgridDBs[idx]
		_ = spiralrhombusgridDB
		var spiralrhombusgridAPI orm.SpiralRhombusGridAPI

		// insertion point for updating fields
		spiralrhombusgridAPI.ID = spiralrhombusgridDB.ID
		spiralrhombusgridDB.CopyBasicFieldsToSpiralRhombusGrid_WOP(&spiralrhombusgridAPI.SpiralRhombusGrid_WOP)
		spiralrhombusgridAPI.SpiralRhombusGridPointersEncoding = spiralrhombusgridDB.SpiralRhombusGridPointersEncoding
		spiralrhombusgridAPIs = append(spiralrhombusgridAPIs, spiralrhombusgridAPI)
	}

	c.JSON(http.StatusOK, spiralrhombusgridAPIs)
}

// PostSpiralRhombusGrid
//
// swagger:route POST /spiralrhombusgrids spiralrhombusgrids postSpiralRhombusGrid
//
// Creates a spiralrhombusgrid
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSpiralRhombusGrid(c *gin.Context) {

	mutexSpiralRhombusGrid.Lock()
	defer mutexSpiralRhombusGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSpiralRhombusGrids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralRhombusGrid.GetDB()

	// Validate input
	var input orm.SpiralRhombusGridAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spiralrhombusgrid
	spiralrhombusgridDB := orm.SpiralRhombusGridDB{}
	spiralrhombusgridDB.SpiralRhombusGridPointersEncoding = input.SpiralRhombusGridPointersEncoding
	spiralrhombusgridDB.CopyBasicFieldsFromSpiralRhombusGrid_WOP(&input.SpiralRhombusGrid_WOP)

	query := db.Create(&spiralrhombusgridDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSpiralRhombusGrid.CheckoutPhaseOneInstance(&spiralrhombusgridDB)
	spiralrhombusgrid := backRepo.BackRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr[spiralrhombusgridDB.ID]

	if spiralrhombusgrid != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spiralrhombusgrid)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spiralrhombusgridDB)
}

// GetSpiralRhombusGrid
//
// swagger:route GET /spiralrhombusgrids/{ID} spiralrhombusgrids getSpiralRhombusGrid
//
// Gets the details for a spiralrhombusgrid.
//
// Responses:
// default: genericError
//
//	200: spiralrhombusgridDBResponse
func (controller *Controller) GetSpiralRhombusGrid(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralRhombusGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralRhombusGrid.GetDB()

	// Get spiralrhombusgridDB in DB
	var spiralrhombusgridDB orm.SpiralRhombusGridDB
	if err := db.First(&spiralrhombusgridDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spiralrhombusgridAPI orm.SpiralRhombusGridAPI
	spiralrhombusgridAPI.ID = spiralrhombusgridDB.ID
	spiralrhombusgridAPI.SpiralRhombusGridPointersEncoding = spiralrhombusgridDB.SpiralRhombusGridPointersEncoding
	spiralrhombusgridDB.CopyBasicFieldsToSpiralRhombusGrid_WOP(&spiralrhombusgridAPI.SpiralRhombusGrid_WOP)

	c.JSON(http.StatusOK, spiralrhombusgridAPI)
}

// UpdateSpiralRhombusGrid
//
// swagger:route PATCH /spiralrhombusgrids/{ID} spiralrhombusgrids updateSpiralRhombusGrid
//
// # Update a spiralrhombusgrid
//
// Responses:
// default: genericError
//
//	200: spiralrhombusgridDBResponse
func (controller *Controller) UpdateSpiralRhombusGrid(c *gin.Context) {

	mutexSpiralRhombusGrid.Lock()
	defer mutexSpiralRhombusGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSpiralRhombusGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralRhombusGrid.GetDB()

	// Validate input
	var input orm.SpiralRhombusGridAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spiralrhombusgridDB orm.SpiralRhombusGridDB

	// fetch the spiralrhombusgrid
	query := db.First(&spiralrhombusgridDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spiralrhombusgridDB.CopyBasicFieldsFromSpiralRhombusGrid_WOP(&input.SpiralRhombusGrid_WOP)
	spiralrhombusgridDB.SpiralRhombusGridPointersEncoding = input.SpiralRhombusGridPointersEncoding

	query = db.Model(&spiralrhombusgridDB).Updates(spiralrhombusgridDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spiralrhombusgridNew := new(models.SpiralRhombusGrid)
	spiralrhombusgridDB.CopyBasicFieldsToSpiralRhombusGrid(spiralrhombusgridNew)

	// redeem pointers
	spiralrhombusgridDB.DecodePointers(backRepo, spiralrhombusgridNew)

	// get stage instance from DB instance, and call callback function
	spiralrhombusgridOld := backRepo.BackRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr[spiralrhombusgridDB.ID]
	if spiralrhombusgridOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), spiralrhombusgridOld, spiralrhombusgridNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spiralrhombusgridDB
	c.JSON(http.StatusOK, spiralrhombusgridDB)
}

// DeleteSpiralRhombusGrid
//
// swagger:route DELETE /spiralrhombusgrids/{ID} spiralrhombusgrids deleteSpiralRhombusGrid
//
// # Delete a spiralrhombusgrid
//
// default: genericError
//
//	200: spiralrhombusgridDBResponse
func (controller *Controller) DeleteSpiralRhombusGrid(c *gin.Context) {

	mutexSpiralRhombusGrid.Lock()
	defer mutexSpiralRhombusGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSpiralRhombusGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralRhombusGrid.GetDB()

	// Get model if exist
	var spiralrhombusgridDB orm.SpiralRhombusGridDB
	if err := db.First(&spiralrhombusgridDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&spiralrhombusgridDB)

	// get an instance (not staged) from DB instance, and call callback function
	spiralrhombusgridDeleted := new(models.SpiralRhombusGrid)
	spiralrhombusgridDB.CopyBasicFieldsToSpiralRhombusGrid(spiralrhombusgridDeleted)

	// get stage instance from DB instance, and call callback function
	spiralrhombusgridStaged := backRepo.BackRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr[spiralrhombusgridDB.ID]
	if spiralrhombusgridStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spiralrhombusgridStaged, spiralrhombusgridDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
