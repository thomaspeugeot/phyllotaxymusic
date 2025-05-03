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
var __SpiralCircleGrid__dummysDeclaration__ models.SpiralCircleGrid
var __SpiralCircleGrid_time__dummyDeclaration time.Duration

var mutexSpiralCircleGrid sync.Mutex

// An SpiralCircleGridID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSpiralCircleGrid updateSpiralCircleGrid deleteSpiralCircleGrid
type SpiralCircleGridID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SpiralCircleGridInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSpiralCircleGrid updateSpiralCircleGrid
type SpiralCircleGridInput struct {
	// The SpiralCircleGrid to submit or modify
	// in: body
	SpiralCircleGrid *orm.SpiralCircleGridAPI
}

// GetSpiralCircleGrids
//
// swagger:route GET /spiralcirclegrids spiralcirclegrids getSpiralCircleGrids
//
// # Get all spiralcirclegrids
//
// Responses:
// default: genericError
//
//	200: spiralcirclegridDBResponse
func (controller *Controller) GetSpiralCircleGrids(c *gin.Context) {

	// source slice
	var spiralcirclegridDBs []orm.SpiralCircleGridDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralCircleGrids", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSpiralCircleGrid.GetDB()

	_, err := db.Find(&spiralcirclegridDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spiralcirclegridAPIs := make([]orm.SpiralCircleGridAPI, 0)

	// for each spiralcirclegrid, update fields from the database nullable fields
	for idx := range spiralcirclegridDBs {
		spiralcirclegridDB := &spiralcirclegridDBs[idx]
		_ = spiralcirclegridDB
		var spiralcirclegridAPI orm.SpiralCircleGridAPI

		// insertion point for updating fields
		spiralcirclegridAPI.ID = spiralcirclegridDB.ID
		spiralcirclegridDB.CopyBasicFieldsToSpiralCircleGrid_WOP(&spiralcirclegridAPI.SpiralCircleGrid_WOP)
		spiralcirclegridAPI.SpiralCircleGridPointersEncoding = spiralcirclegridDB.SpiralCircleGridPointersEncoding
		spiralcirclegridAPIs = append(spiralcirclegridAPIs, spiralcirclegridAPI)
	}

	c.JSON(http.StatusOK, spiralcirclegridAPIs)
}

// PostSpiralCircleGrid
//
// swagger:route POST /spiralcirclegrids spiralcirclegrids postSpiralCircleGrid
//
// Creates a spiralcirclegrid
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSpiralCircleGrid(c *gin.Context) {

	mutexSpiralCircleGrid.Lock()
	defer mutexSpiralCircleGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSpiralCircleGrids", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSpiralCircleGrid.GetDB()

	// Validate input
	var input orm.SpiralCircleGridAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spiralcirclegrid
	spiralcirclegridDB := orm.SpiralCircleGridDB{}
	spiralcirclegridDB.SpiralCircleGridPointersEncoding = input.SpiralCircleGridPointersEncoding
	spiralcirclegridDB.CopyBasicFieldsFromSpiralCircleGrid_WOP(&input.SpiralCircleGrid_WOP)

	_, err = db.Create(&spiralcirclegridDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSpiralCircleGrid.CheckoutPhaseOneInstance(&spiralcirclegridDB)
	spiralcirclegrid := backRepo.BackRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridPtr[spiralcirclegridDB.ID]

	if spiralcirclegrid != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spiralcirclegrid)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spiralcirclegridDB)
}

// GetSpiralCircleGrid
//
// swagger:route GET /spiralcirclegrids/{ID} spiralcirclegrids getSpiralCircleGrid
//
// Gets the details for a spiralcirclegrid.
//
// Responses:
// default: genericError
//
//	200: spiralcirclegridDBResponse
func (controller *Controller) GetSpiralCircleGrid(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralCircleGrid", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSpiralCircleGrid.GetDB()

	// Get spiralcirclegridDB in DB
	var spiralcirclegridDB orm.SpiralCircleGridDB
	if _, err := db.First(&spiralcirclegridDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spiralcirclegridAPI orm.SpiralCircleGridAPI
	spiralcirclegridAPI.ID = spiralcirclegridDB.ID
	spiralcirclegridAPI.SpiralCircleGridPointersEncoding = spiralcirclegridDB.SpiralCircleGridPointersEncoding
	spiralcirclegridDB.CopyBasicFieldsToSpiralCircleGrid_WOP(&spiralcirclegridAPI.SpiralCircleGrid_WOP)

	c.JSON(http.StatusOK, spiralcirclegridAPI)
}

// UpdateSpiralCircleGrid
//
// swagger:route PATCH /spiralcirclegrids/{ID} spiralcirclegrids updateSpiralCircleGrid
//
// # Update a spiralcirclegrid
//
// Responses:
// default: genericError
//
//	200: spiralcirclegridDBResponse
func (controller *Controller) UpdateSpiralCircleGrid(c *gin.Context) {

	mutexSpiralCircleGrid.Lock()
	defer mutexSpiralCircleGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSpiralCircleGrid", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSpiralCircleGrid.GetDB()

	// Validate input
	var input orm.SpiralCircleGridAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spiralcirclegridDB orm.SpiralCircleGridDB

	// fetch the spiralcirclegrid
	_, err := db.First(&spiralcirclegridDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spiralcirclegridDB.CopyBasicFieldsFromSpiralCircleGrid_WOP(&input.SpiralCircleGrid_WOP)
	spiralcirclegridDB.SpiralCircleGridPointersEncoding = input.SpiralCircleGridPointersEncoding

	db, _ = db.Model(&spiralcirclegridDB)
	_, err = db.Updates(&spiralcirclegridDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spiralcirclegridNew := new(models.SpiralCircleGrid)
	spiralcirclegridDB.CopyBasicFieldsToSpiralCircleGrid(spiralcirclegridNew)

	// redeem pointers
	spiralcirclegridDB.DecodePointers(backRepo, spiralcirclegridNew)

	// get stage instance from DB instance, and call callback function
	spiralcirclegridOld := backRepo.BackRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridPtr[spiralcirclegridDB.ID]
	if spiralcirclegridOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), spiralcirclegridOld, spiralcirclegridNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spiralcirclegridDB
	c.JSON(http.StatusOK, spiralcirclegridDB)
}

// DeleteSpiralCircleGrid
//
// swagger:route DELETE /spiralcirclegrids/{ID} spiralcirclegrids deleteSpiralCircleGrid
//
// # Delete a spiralcirclegrid
//
// default: genericError
//
//	200: spiralcirclegridDBResponse
func (controller *Controller) DeleteSpiralCircleGrid(c *gin.Context) {

	mutexSpiralCircleGrid.Lock()
	defer mutexSpiralCircleGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSpiralCircleGrid", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSpiralCircleGrid.GetDB()

	// Get model if exist
	var spiralcirclegridDB orm.SpiralCircleGridDB
	if _, err := db.First(&spiralcirclegridDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&spiralcirclegridDB)

	// get an instance (not staged) from DB instance, and call callback function
	spiralcirclegridDeleted := new(models.SpiralCircleGrid)
	spiralcirclegridDB.CopyBasicFieldsToSpiralCircleGrid(spiralcirclegridDeleted)

	// get stage instance from DB instance, and call callback function
	spiralcirclegridStaged := backRepo.BackRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridPtr[spiralcirclegridDB.ID]
	if spiralcirclegridStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spiralcirclegridStaged, spiralcirclegridDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
