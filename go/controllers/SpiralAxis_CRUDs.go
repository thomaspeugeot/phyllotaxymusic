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
var __SpiralAxis__dummysDeclaration__ models.SpiralAxis
var __SpiralAxis_time__dummyDeclaration time.Duration

var mutexSpiralAxis sync.Mutex

// An SpiralAxisID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSpiralAxis updateSpiralAxis deleteSpiralAxis
type SpiralAxisID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SpiralAxisInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSpiralAxis updateSpiralAxis
type SpiralAxisInput struct {
	// The SpiralAxis to submit or modify
	// in: body
	SpiralAxis *orm.SpiralAxisAPI
}

// GetSpiralAxiss
//
// swagger:route GET /spiralaxiss spiralaxiss getSpiralAxiss
//
// # Get all spiralaxiss
//
// Responses:
// default: genericError
//
//	200: spiralaxisDBResponse
func (controller *Controller) GetSpiralAxiss(c *gin.Context) {

	// source slice
	var spiralaxisDBs []orm.SpiralAxisDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralAxiss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralAxis.GetDB()

	query := db.Find(&spiralaxisDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spiralaxisAPIs := make([]orm.SpiralAxisAPI, 0)

	// for each spiralaxis, update fields from the database nullable fields
	for idx := range spiralaxisDBs {
		spiralaxisDB := &spiralaxisDBs[idx]
		_ = spiralaxisDB
		var spiralaxisAPI orm.SpiralAxisAPI

		// insertion point for updating fields
		spiralaxisAPI.ID = spiralaxisDB.ID
		spiralaxisDB.CopyBasicFieldsToSpiralAxis_WOP(&spiralaxisAPI.SpiralAxis_WOP)
		spiralaxisAPI.SpiralAxisPointersEncoding = spiralaxisDB.SpiralAxisPointersEncoding
		spiralaxisAPIs = append(spiralaxisAPIs, spiralaxisAPI)
	}

	c.JSON(http.StatusOK, spiralaxisAPIs)
}

// PostSpiralAxis
//
// swagger:route POST /spiralaxiss spiralaxiss postSpiralAxis
//
// Creates a spiralaxis
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSpiralAxis(c *gin.Context) {

	mutexSpiralAxis.Lock()
	defer mutexSpiralAxis.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSpiralAxiss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralAxis.GetDB()

	// Validate input
	var input orm.SpiralAxisAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spiralaxis
	spiralaxisDB := orm.SpiralAxisDB{}
	spiralaxisDB.SpiralAxisPointersEncoding = input.SpiralAxisPointersEncoding
	spiralaxisDB.CopyBasicFieldsFromSpiralAxis_WOP(&input.SpiralAxis_WOP)

	query := db.Create(&spiralaxisDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSpiralAxis.CheckoutPhaseOneInstance(&spiralaxisDB)
	spiralaxis := backRepo.BackRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisPtr[spiralaxisDB.ID]

	if spiralaxis != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spiralaxis)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spiralaxisDB)
}

// GetSpiralAxis
//
// swagger:route GET /spiralaxiss/{ID} spiralaxiss getSpiralAxis
//
// Gets the details for a spiralaxis.
//
// Responses:
// default: genericError
//
//	200: spiralaxisDBResponse
func (controller *Controller) GetSpiralAxis(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralAxis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralAxis.GetDB()

	// Get spiralaxisDB in DB
	var spiralaxisDB orm.SpiralAxisDB
	if err := db.First(&spiralaxisDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spiralaxisAPI orm.SpiralAxisAPI
	spiralaxisAPI.ID = spiralaxisDB.ID
	spiralaxisAPI.SpiralAxisPointersEncoding = spiralaxisDB.SpiralAxisPointersEncoding
	spiralaxisDB.CopyBasicFieldsToSpiralAxis_WOP(&spiralaxisAPI.SpiralAxis_WOP)

	c.JSON(http.StatusOK, spiralaxisAPI)
}

// UpdateSpiralAxis
//
// swagger:route PATCH /spiralaxiss/{ID} spiralaxiss updateSpiralAxis
//
// # Update a spiralaxis
//
// Responses:
// default: genericError
//
//	200: spiralaxisDBResponse
func (controller *Controller) UpdateSpiralAxis(c *gin.Context) {

	mutexSpiralAxis.Lock()
	defer mutexSpiralAxis.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSpiralAxis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralAxis.GetDB()

	// Validate input
	var input orm.SpiralAxisAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spiralaxisDB orm.SpiralAxisDB

	// fetch the spiralaxis
	query := db.First(&spiralaxisDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spiralaxisDB.CopyBasicFieldsFromSpiralAxis_WOP(&input.SpiralAxis_WOP)
	spiralaxisDB.SpiralAxisPointersEncoding = input.SpiralAxisPointersEncoding

	query = db.Model(&spiralaxisDB).Updates(spiralaxisDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spiralaxisNew := new(models.SpiralAxis)
	spiralaxisDB.CopyBasicFieldsToSpiralAxis(spiralaxisNew)

	// redeem pointers
	spiralaxisDB.DecodePointers(backRepo, spiralaxisNew)

	// get stage instance from DB instance, and call callback function
	spiralaxisOld := backRepo.BackRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisPtr[spiralaxisDB.ID]
	if spiralaxisOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), spiralaxisOld, spiralaxisNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spiralaxisDB
	c.JSON(http.StatusOK, spiralaxisDB)
}

// DeleteSpiralAxis
//
// swagger:route DELETE /spiralaxiss/{ID} spiralaxiss deleteSpiralAxis
//
// # Delete a spiralaxis
//
// default: genericError
//
//	200: spiralaxisDBResponse
func (controller *Controller) DeleteSpiralAxis(c *gin.Context) {

	mutexSpiralAxis.Lock()
	defer mutexSpiralAxis.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSpiralAxis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralAxis.GetDB()

	// Get model if exist
	var spiralaxisDB orm.SpiralAxisDB
	if err := db.First(&spiralaxisDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&spiralaxisDB)

	// get an instance (not staged) from DB instance, and call callback function
	spiralaxisDeleted := new(models.SpiralAxis)
	spiralaxisDB.CopyBasicFieldsToSpiralAxis(spiralaxisDeleted)

	// get stage instance from DB instance, and call callback function
	spiralaxisStaged := backRepo.BackRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisPtr[spiralaxisDB.ID]
	if spiralaxisStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spiralaxisStaged, spiralaxisDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
