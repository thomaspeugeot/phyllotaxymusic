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
var __CircleGrid__dummysDeclaration__ models.CircleGrid
var __CircleGrid_time__dummyDeclaration time.Duration

var mutexCircleGrid sync.Mutex

// An CircleGridID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCircleGrid updateCircleGrid deleteCircleGrid
type CircleGridID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CircleGridInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCircleGrid updateCircleGrid
type CircleGridInput struct {
	// The CircleGrid to submit or modify
	// in: body
	CircleGrid *orm.CircleGridAPI
}

// GetCircleGrids
//
// swagger:route GET /circlegrids circlegrids getCircleGrids
//
// # Get all circlegrids
//
// Responses:
// default: genericError
//
//	200: circlegridDBResponse
func (controller *Controller) GetCircleGrids(c *gin.Context) {

	// source slice
	var circlegridDBs []orm.CircleGridDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCircleGrids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCircleGrid.GetDB()

	_, err := db.Find(&circlegridDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	circlegridAPIs := make([]orm.CircleGridAPI, 0)

	// for each circlegrid, update fields from the database nullable fields
	for idx := range circlegridDBs {
		circlegridDB := &circlegridDBs[idx]
		_ = circlegridDB
		var circlegridAPI orm.CircleGridAPI

		// insertion point for updating fields
		circlegridAPI.ID = circlegridDB.ID
		circlegridDB.CopyBasicFieldsToCircleGrid_WOP(&circlegridAPI.CircleGrid_WOP)
		circlegridAPI.CircleGridPointersEncoding = circlegridDB.CircleGridPointersEncoding
		circlegridAPIs = append(circlegridAPIs, circlegridAPI)
	}

	c.JSON(http.StatusOK, circlegridAPIs)
}

// PostCircleGrid
//
// swagger:route POST /circlegrids circlegrids postCircleGrid
//
// Creates a circlegrid
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCircleGrid(c *gin.Context) {

	mutexCircleGrid.Lock()
	defer mutexCircleGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCircleGrids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCircleGrid.GetDB()

	// Validate input
	var input orm.CircleGridAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create circlegrid
	circlegridDB := orm.CircleGridDB{}
	circlegridDB.CircleGridPointersEncoding = input.CircleGridPointersEncoding
	circlegridDB.CopyBasicFieldsFromCircleGrid_WOP(&input.CircleGrid_WOP)

	_, err = db.Create(&circlegridDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCircleGrid.CheckoutPhaseOneInstance(&circlegridDB)
	circlegrid := backRepo.BackRepoCircleGrid.Map_CircleGridDBID_CircleGridPtr[circlegridDB.ID]

	if circlegrid != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), circlegrid)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, circlegridDB)
}

// GetCircleGrid
//
// swagger:route GET /circlegrids/{ID} circlegrids getCircleGrid
//
// Gets the details for a circlegrid.
//
// Responses:
// default: genericError
//
//	200: circlegridDBResponse
func (controller *Controller) GetCircleGrid(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCircleGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCircleGrid.GetDB()

	// Get circlegridDB in DB
	var circlegridDB orm.CircleGridDB
	if _, err := db.First(&circlegridDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var circlegridAPI orm.CircleGridAPI
	circlegridAPI.ID = circlegridDB.ID
	circlegridAPI.CircleGridPointersEncoding = circlegridDB.CircleGridPointersEncoding
	circlegridDB.CopyBasicFieldsToCircleGrid_WOP(&circlegridAPI.CircleGrid_WOP)

	c.JSON(http.StatusOK, circlegridAPI)
}

// UpdateCircleGrid
//
// swagger:route PATCH /circlegrids/{ID} circlegrids updateCircleGrid
//
// # Update a circlegrid
//
// Responses:
// default: genericError
//
//	200: circlegridDBResponse
func (controller *Controller) UpdateCircleGrid(c *gin.Context) {

	mutexCircleGrid.Lock()
	defer mutexCircleGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCircleGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCircleGrid.GetDB()

	// Validate input
	var input orm.CircleGridAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var circlegridDB orm.CircleGridDB

	// fetch the circlegrid
	_, err := db.First(&circlegridDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	circlegridDB.CopyBasicFieldsFromCircleGrid_WOP(&input.CircleGrid_WOP)
	circlegridDB.CircleGridPointersEncoding = input.CircleGridPointersEncoding

	db, _ = db.Model(&circlegridDB)
	_, err = db.Updates(circlegridDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	circlegridNew := new(models.CircleGrid)
	circlegridDB.CopyBasicFieldsToCircleGrid(circlegridNew)

	// redeem pointers
	circlegridDB.DecodePointers(backRepo, circlegridNew)

	// get stage instance from DB instance, and call callback function
	circlegridOld := backRepo.BackRepoCircleGrid.Map_CircleGridDBID_CircleGridPtr[circlegridDB.ID]
	if circlegridOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), circlegridOld, circlegridNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the circlegridDB
	c.JSON(http.StatusOK, circlegridDB)
}

// DeleteCircleGrid
//
// swagger:route DELETE /circlegrids/{ID} circlegrids deleteCircleGrid
//
// # Delete a circlegrid
//
// default: genericError
//
//	200: circlegridDBResponse
func (controller *Controller) DeleteCircleGrid(c *gin.Context) {

	mutexCircleGrid.Lock()
	defer mutexCircleGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCircleGrid", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCircleGrid.GetDB()

	// Get model if exist
	var circlegridDB orm.CircleGridDB
	if _, err := db.First(&circlegridDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&circlegridDB)

	// get an instance (not staged) from DB instance, and call callback function
	circlegridDeleted := new(models.CircleGrid)
	circlegridDB.CopyBasicFieldsToCircleGrid(circlegridDeleted)

	// get stage instance from DB instance, and call callback function
	circlegridStaged := backRepo.BackRepoCircleGrid.Map_CircleGridDBID_CircleGridPtr[circlegridDB.ID]
	if circlegridStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), circlegridStaged, circlegridDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
