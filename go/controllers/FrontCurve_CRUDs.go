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
var __FrontCurve__dummysDeclaration__ models.FrontCurve
var __FrontCurve_time__dummyDeclaration time.Duration

var mutexFrontCurve sync.Mutex

// An FrontCurveID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFrontCurve updateFrontCurve deleteFrontCurve
type FrontCurveID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FrontCurveInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFrontCurve updateFrontCurve
type FrontCurveInput struct {
	// The FrontCurve to submit or modify
	// in: body
	FrontCurve *orm.FrontCurveAPI
}

// GetFrontCurves
//
// swagger:route GET /frontcurves frontcurves getFrontCurves
//
// # Get all frontcurves
//
// Responses:
// default: genericError
//
//	200: frontcurveDBResponse
func (controller *Controller) GetFrontCurves(c *gin.Context) {

	// source slice
	var frontcurveDBs []orm.FrontCurveDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFrontCurves", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoFrontCurve.GetDB()

	_, err := db.Find(&frontcurveDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	frontcurveAPIs := make([]orm.FrontCurveAPI, 0)

	// for each frontcurve, update fields from the database nullable fields
	for idx := range frontcurveDBs {
		frontcurveDB := &frontcurveDBs[idx]
		_ = frontcurveDB
		var frontcurveAPI orm.FrontCurveAPI

		// insertion point for updating fields
		frontcurveAPI.ID = frontcurveDB.ID
		frontcurveDB.CopyBasicFieldsToFrontCurve_WOP(&frontcurveAPI.FrontCurve_WOP)
		frontcurveAPI.FrontCurvePointersEncoding = frontcurveDB.FrontCurvePointersEncoding
		frontcurveAPIs = append(frontcurveAPIs, frontcurveAPI)
	}

	c.JSON(http.StatusOK, frontcurveAPIs)
}

// PostFrontCurve
//
// swagger:route POST /frontcurves frontcurves postFrontCurve
//
// Creates a frontcurve
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFrontCurve(c *gin.Context) {

	mutexFrontCurve.Lock()
	defer mutexFrontCurve.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFrontCurves", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoFrontCurve.GetDB()

	// Validate input
	var input orm.FrontCurveAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create frontcurve
	frontcurveDB := orm.FrontCurveDB{}
	frontcurveDB.FrontCurvePointersEncoding = input.FrontCurvePointersEncoding
	frontcurveDB.CopyBasicFieldsFromFrontCurve_WOP(&input.FrontCurve_WOP)

	_, err = db.Create(&frontcurveDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFrontCurve.CheckoutPhaseOneInstance(&frontcurveDB)
	frontcurve := backRepo.BackRepoFrontCurve.Map_FrontCurveDBID_FrontCurvePtr[frontcurveDB.ID]

	if frontcurve != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), frontcurve)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, frontcurveDB)
}

// GetFrontCurve
//
// swagger:route GET /frontcurves/{ID} frontcurves getFrontCurve
//
// Gets the details for a frontcurve.
//
// Responses:
// default: genericError
//
//	200: frontcurveDBResponse
func (controller *Controller) GetFrontCurve(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFrontCurve", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoFrontCurve.GetDB()

	// Get frontcurveDB in DB
	var frontcurveDB orm.FrontCurveDB
	if _, err := db.First(&frontcurveDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var frontcurveAPI orm.FrontCurveAPI
	frontcurveAPI.ID = frontcurveDB.ID
	frontcurveAPI.FrontCurvePointersEncoding = frontcurveDB.FrontCurvePointersEncoding
	frontcurveDB.CopyBasicFieldsToFrontCurve_WOP(&frontcurveAPI.FrontCurve_WOP)

	c.JSON(http.StatusOK, frontcurveAPI)
}

// UpdateFrontCurve
//
// swagger:route PATCH /frontcurves/{ID} frontcurves updateFrontCurve
//
// # Update a frontcurve
//
// Responses:
// default: genericError
//
//	200: frontcurveDBResponse
func (controller *Controller) UpdateFrontCurve(c *gin.Context) {

	mutexFrontCurve.Lock()
	defer mutexFrontCurve.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFrontCurve", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoFrontCurve.GetDB()

	// Validate input
	var input orm.FrontCurveAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var frontcurveDB orm.FrontCurveDB

	// fetch the frontcurve
	_, err := db.First(&frontcurveDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	frontcurveDB.CopyBasicFieldsFromFrontCurve_WOP(&input.FrontCurve_WOP)
	frontcurveDB.FrontCurvePointersEncoding = input.FrontCurvePointersEncoding

	db, _ = db.Model(&frontcurveDB)
	_, err = db.Updates(&frontcurveDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	frontcurveNew := new(models.FrontCurve)
	frontcurveDB.CopyBasicFieldsToFrontCurve(frontcurveNew)

	// redeem pointers
	frontcurveDB.DecodePointers(backRepo, frontcurveNew)

	// get stage instance from DB instance, and call callback function
	frontcurveOld := backRepo.BackRepoFrontCurve.Map_FrontCurveDBID_FrontCurvePtr[frontcurveDB.ID]
	if frontcurveOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), frontcurveOld, frontcurveNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the frontcurveDB
	c.JSON(http.StatusOK, frontcurveDB)
}

// DeleteFrontCurve
//
// swagger:route DELETE /frontcurves/{ID} frontcurves deleteFrontCurve
//
// # Delete a frontcurve
//
// default: genericError
//
//	200: frontcurveDBResponse
func (controller *Controller) DeleteFrontCurve(c *gin.Context) {

	mutexFrontCurve.Lock()
	defer mutexFrontCurve.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFrontCurve", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoFrontCurve.GetDB()

	// Get model if exist
	var frontcurveDB orm.FrontCurveDB
	if _, err := db.First(&frontcurveDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&frontcurveDB)

	// get an instance (not staged) from DB instance, and call callback function
	frontcurveDeleted := new(models.FrontCurve)
	frontcurveDB.CopyBasicFieldsToFrontCurve(frontcurveDeleted)

	// get stage instance from DB instance, and call callback function
	frontcurveStaged := backRepo.BackRepoFrontCurve.Map_FrontCurveDBID_FrontCurvePtr[frontcurveDB.ID]
	if frontcurveStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), frontcurveStaged, frontcurveDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
