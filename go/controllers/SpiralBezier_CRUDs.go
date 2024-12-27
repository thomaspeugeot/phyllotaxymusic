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
var __SpiralBezier__dummysDeclaration__ models.SpiralBezier
var __SpiralBezier_time__dummyDeclaration time.Duration

var mutexSpiralBezier sync.Mutex

// An SpiralBezierID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSpiralBezier updateSpiralBezier deleteSpiralBezier
type SpiralBezierID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SpiralBezierInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSpiralBezier updateSpiralBezier
type SpiralBezierInput struct {
	// The SpiralBezier to submit or modify
	// in: body
	SpiralBezier *orm.SpiralBezierAPI
}

// GetSpiralBeziers
//
// swagger:route GET /spiralbeziers spiralbeziers getSpiralBeziers
//
// # Get all spiralbeziers
//
// Responses:
// default: genericError
//
//	200: spiralbezierDBResponse
func (controller *Controller) GetSpiralBeziers(c *gin.Context) {

	// source slice
	var spiralbezierDBs []orm.SpiralBezierDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralBeziers", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phyllotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralBezier.GetDB()

	_, err := db.Find(&spiralbezierDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spiralbezierAPIs := make([]orm.SpiralBezierAPI, 0)

	// for each spiralbezier, update fields from the database nullable fields
	for idx := range spiralbezierDBs {
		spiralbezierDB := &spiralbezierDBs[idx]
		_ = spiralbezierDB
		var spiralbezierAPI orm.SpiralBezierAPI

		// insertion point for updating fields
		spiralbezierAPI.ID = spiralbezierDB.ID
		spiralbezierDB.CopyBasicFieldsToSpiralBezier_WOP(&spiralbezierAPI.SpiralBezier_WOP)
		spiralbezierAPI.SpiralBezierPointersEncoding = spiralbezierDB.SpiralBezierPointersEncoding
		spiralbezierAPIs = append(spiralbezierAPIs, spiralbezierAPI)
	}

	c.JSON(http.StatusOK, spiralbezierAPIs)
}

// PostSpiralBezier
//
// swagger:route POST /spiralbeziers spiralbeziers postSpiralBezier
//
// Creates a spiralbezier
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSpiralBezier(c *gin.Context) {

	mutexSpiralBezier.Lock()
	defer mutexSpiralBezier.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSpiralBeziers", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phyllotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralBezier.GetDB()

	// Validate input
	var input orm.SpiralBezierAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spiralbezier
	spiralbezierDB := orm.SpiralBezierDB{}
	spiralbezierDB.SpiralBezierPointersEncoding = input.SpiralBezierPointersEncoding
	spiralbezierDB.CopyBasicFieldsFromSpiralBezier_WOP(&input.SpiralBezier_WOP)

	_, err = db.Create(&spiralbezierDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSpiralBezier.CheckoutPhaseOneInstance(&spiralbezierDB)
	spiralbezier := backRepo.BackRepoSpiralBezier.Map_SpiralBezierDBID_SpiralBezierPtr[spiralbezierDB.ID]

	if spiralbezier != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spiralbezier)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spiralbezierDB)
}

// GetSpiralBezier
//
// swagger:route GET /spiralbeziers/{ID} spiralbeziers getSpiralBezier
//
// Gets the details for a spiralbezier.
//
// Responses:
// default: genericError
//
//	200: spiralbezierDBResponse
func (controller *Controller) GetSpiralBezier(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSpiralBezier", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phyllotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralBezier.GetDB()

	// Get spiralbezierDB in DB
	var spiralbezierDB orm.SpiralBezierDB
	if _, err := db.First(&spiralbezierDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spiralbezierAPI orm.SpiralBezierAPI
	spiralbezierAPI.ID = spiralbezierDB.ID
	spiralbezierAPI.SpiralBezierPointersEncoding = spiralbezierDB.SpiralBezierPointersEncoding
	spiralbezierDB.CopyBasicFieldsToSpiralBezier_WOP(&spiralbezierAPI.SpiralBezier_WOP)

	c.JSON(http.StatusOK, spiralbezierAPI)
}

// UpdateSpiralBezier
//
// swagger:route PATCH /spiralbeziers/{ID} spiralbeziers updateSpiralBezier
//
// # Update a spiralbezier
//
// Responses:
// default: genericError
//
//	200: spiralbezierDBResponse
func (controller *Controller) UpdateSpiralBezier(c *gin.Context) {

	mutexSpiralBezier.Lock()
	defer mutexSpiralBezier.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSpiralBezier", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phyllotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralBezier.GetDB()

	// Validate input
	var input orm.SpiralBezierAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spiralbezierDB orm.SpiralBezierDB

	// fetch the spiralbezier
	_, err := db.First(&spiralbezierDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spiralbezierDB.CopyBasicFieldsFromSpiralBezier_WOP(&input.SpiralBezier_WOP)
	spiralbezierDB.SpiralBezierPointersEncoding = input.SpiralBezierPointersEncoding

	db, _ = db.Model(&spiralbezierDB)
	_, err = db.Updates(&spiralbezierDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spiralbezierNew := new(models.SpiralBezier)
	spiralbezierDB.CopyBasicFieldsToSpiralBezier(spiralbezierNew)

	// redeem pointers
	spiralbezierDB.DecodePointers(backRepo, spiralbezierNew)

	// get stage instance from DB instance, and call callback function
	spiralbezierOld := backRepo.BackRepoSpiralBezier.Map_SpiralBezierDBID_SpiralBezierPtr[spiralbezierDB.ID]
	if spiralbezierOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), spiralbezierOld, spiralbezierNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spiralbezierDB
	c.JSON(http.StatusOK, spiralbezierDB)
}

// DeleteSpiralBezier
//
// swagger:route DELETE /spiralbeziers/{ID} spiralbeziers deleteSpiralBezier
//
// # Delete a spiralbezier
//
// default: genericError
//
//	200: spiralbezierDBResponse
func (controller *Controller) DeleteSpiralBezier(c *gin.Context) {

	mutexSpiralBezier.Lock()
	defer mutexSpiralBezier.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSpiralBezier", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phyllotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSpiralBezier.GetDB()

	// Get model if exist
	var spiralbezierDB orm.SpiralBezierDB
	if _, err := db.First(&spiralbezierDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&spiralbezierDB)

	// get an instance (not staged) from DB instance, and call callback function
	spiralbezierDeleted := new(models.SpiralBezier)
	spiralbezierDB.CopyBasicFieldsToSpiralBezier(spiralbezierDeleted)

	// get stage instance from DB instance, and call callback function
	spiralbezierStaged := backRepo.BackRepoSpiralBezier.Map_SpiralBezierDBID_SpiralBezierPtr[spiralbezierDB.ID]
	if spiralbezierStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spiralbezierStaged, spiralbezierDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
