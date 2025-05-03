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
var __SpiralBezierGrid__dummysDeclaration__ models.SpiralBezierGrid
var __SpiralBezierGrid_time__dummyDeclaration time.Duration

var mutexSpiralBezierGrid sync.Mutex

// An SpiralBezierGridID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSpiralBezierGrid updateSpiralBezierGrid deleteSpiralBezierGrid
type SpiralBezierGridID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SpiralBezierGridInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSpiralBezierGrid updateSpiralBezierGrid
type SpiralBezierGridInput struct {
	// The SpiralBezierGrid to submit or modify
	// in: body
	SpiralBezierGrid *orm.SpiralBezierGridAPI
}

// GetSpiralBezierGrids
//
// swagger:route GET /spiralbeziergrids spiralbeziergrids getSpiralBezierGrids
//
// # Get all spiralbeziergrids
//
// Responses:
// default: genericError
//
//	200: spiralbeziergridDBResponse
func (controller *Controller) GetSpiralBezierGrids(c *gin.Context) {

	// source slice
	var spiralbeziergridDBs []orm.SpiralBezierGridDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralBezierGrids", "Name", stackPath)
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
	db := backRepo.BackRepoSpiralBezierGrid.GetDB()

	_, err := db.Find(&spiralbeziergridDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spiralbeziergridAPIs := make([]orm.SpiralBezierGridAPI, 0)

	// for each spiralbeziergrid, update fields from the database nullable fields
	for idx := range spiralbeziergridDBs {
		spiralbeziergridDB := &spiralbeziergridDBs[idx]
		_ = spiralbeziergridDB
		var spiralbeziergridAPI orm.SpiralBezierGridAPI

		// insertion point for updating fields
		spiralbeziergridAPI.ID = spiralbeziergridDB.ID
		spiralbeziergridDB.CopyBasicFieldsToSpiralBezierGrid_WOP(&spiralbeziergridAPI.SpiralBezierGrid_WOP)
		spiralbeziergridAPI.SpiralBezierGridPointersEncoding = spiralbeziergridDB.SpiralBezierGridPointersEncoding
		spiralbeziergridAPIs = append(spiralbeziergridAPIs, spiralbeziergridAPI)
	}

	c.JSON(http.StatusOK, spiralbeziergridAPIs)
}

// PostSpiralBezierGrid
//
// swagger:route POST /spiralbeziergrids spiralbeziergrids postSpiralBezierGrid
//
// Creates a spiralbeziergrid
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSpiralBezierGrid(c *gin.Context) {

	mutexSpiralBezierGrid.Lock()
	defer mutexSpiralBezierGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSpiralBezierGrids", "Name", stackPath)
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
	db := backRepo.BackRepoSpiralBezierGrid.GetDB()

	// Validate input
	var input orm.SpiralBezierGridAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spiralbeziergrid
	spiralbeziergridDB := orm.SpiralBezierGridDB{}
	spiralbeziergridDB.SpiralBezierGridPointersEncoding = input.SpiralBezierGridPointersEncoding
	spiralbeziergridDB.CopyBasicFieldsFromSpiralBezierGrid_WOP(&input.SpiralBezierGrid_WOP)

	_, err = db.Create(&spiralbeziergridDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSpiralBezierGrid.CheckoutPhaseOneInstance(&spiralbeziergridDB)
	spiralbeziergrid := backRepo.BackRepoSpiralBezierGrid.Map_SpiralBezierGridDBID_SpiralBezierGridPtr[spiralbeziergridDB.ID]

	if spiralbeziergrid != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spiralbeziergrid)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spiralbeziergridDB)
}

// GetSpiralBezierGrid
//
// swagger:route GET /spiralbeziergrids/{ID} spiralbeziergrids getSpiralBezierGrid
//
// Gets the details for a spiralbeziergrid.
//
// Responses:
// default: genericError
//
//	200: spiralbeziergridDBResponse
func (controller *Controller) GetSpiralBezierGrid(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralBezierGrid", "Name", stackPath)
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
	db := backRepo.BackRepoSpiralBezierGrid.GetDB()

	// Get spiralbeziergridDB in DB
	var spiralbeziergridDB orm.SpiralBezierGridDB
	if _, err := db.First(&spiralbeziergridDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spiralbeziergridAPI orm.SpiralBezierGridAPI
	spiralbeziergridAPI.ID = spiralbeziergridDB.ID
	spiralbeziergridAPI.SpiralBezierGridPointersEncoding = spiralbeziergridDB.SpiralBezierGridPointersEncoding
	spiralbeziergridDB.CopyBasicFieldsToSpiralBezierGrid_WOP(&spiralbeziergridAPI.SpiralBezierGrid_WOP)

	c.JSON(http.StatusOK, spiralbeziergridAPI)
}

// UpdateSpiralBezierGrid
//
// swagger:route PATCH /spiralbeziergrids/{ID} spiralbeziergrids updateSpiralBezierGrid
//
// # Update a spiralbeziergrid
//
// Responses:
// default: genericError
//
//	200: spiralbeziergridDBResponse
func (controller *Controller) UpdateSpiralBezierGrid(c *gin.Context) {

	mutexSpiralBezierGrid.Lock()
	defer mutexSpiralBezierGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSpiralBezierGrid", "Name", stackPath)
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
	db := backRepo.BackRepoSpiralBezierGrid.GetDB()

	// Validate input
	var input orm.SpiralBezierGridAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spiralbeziergridDB orm.SpiralBezierGridDB

	// fetch the spiralbeziergrid
	_, err := db.First(&spiralbeziergridDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spiralbeziergridDB.CopyBasicFieldsFromSpiralBezierGrid_WOP(&input.SpiralBezierGrid_WOP)
	spiralbeziergridDB.SpiralBezierGridPointersEncoding = input.SpiralBezierGridPointersEncoding

	db, _ = db.Model(&spiralbeziergridDB)
	_, err = db.Updates(&spiralbeziergridDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spiralbeziergridNew := new(models.SpiralBezierGrid)
	spiralbeziergridDB.CopyBasicFieldsToSpiralBezierGrid(spiralbeziergridNew)

	// redeem pointers
	spiralbeziergridDB.DecodePointers(backRepo, spiralbeziergridNew)

	// get stage instance from DB instance, and call callback function
	spiralbeziergridOld := backRepo.BackRepoSpiralBezierGrid.Map_SpiralBezierGridDBID_SpiralBezierGridPtr[spiralbeziergridDB.ID]
	if spiralbeziergridOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), spiralbeziergridOld, spiralbeziergridNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spiralbeziergridDB
	c.JSON(http.StatusOK, spiralbeziergridDB)
}

// DeleteSpiralBezierGrid
//
// swagger:route DELETE /spiralbeziergrids/{ID} spiralbeziergrids deleteSpiralBezierGrid
//
// # Delete a spiralbeziergrid
//
// default: genericError
//
//	200: spiralbeziergridDBResponse
func (controller *Controller) DeleteSpiralBezierGrid(c *gin.Context) {

	mutexSpiralBezierGrid.Lock()
	defer mutexSpiralBezierGrid.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSpiralBezierGrid", "Name", stackPath)
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
	db := backRepo.BackRepoSpiralBezierGrid.GetDB()

	// Get model if exist
	var spiralbeziergridDB orm.SpiralBezierGridDB
	if _, err := db.First(&spiralbeziergridDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&spiralbeziergridDB)

	// get an instance (not staged) from DB instance, and call callback function
	spiralbeziergridDeleted := new(models.SpiralBezierGrid)
	spiralbeziergridDB.CopyBasicFieldsToSpiralBezierGrid(spiralbeziergridDeleted)

	// get stage instance from DB instance, and call callback function
	spiralbeziergridStaged := backRepo.BackRepoSpiralBezierGrid.Map_SpiralBezierGridDBID_SpiralBezierGridPtr[spiralbeziergridDB.ID]
	if spiralbeziergridStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spiralbeziergridStaged, spiralbeziergridDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
