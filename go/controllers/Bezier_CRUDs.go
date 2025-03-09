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
var __Bezier__dummysDeclaration__ models.Bezier
var __Bezier_time__dummyDeclaration time.Duration

var mutexBezier sync.Mutex

// An BezierID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBezier updateBezier deleteBezier
type BezierID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BezierInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBezier updateBezier
type BezierInput struct {
	// The Bezier to submit or modify
	// in: body
	Bezier *orm.BezierAPI
}

// GetBeziers
//
// swagger:route GET /beziers beziers getBeziers
//
// # Get all beziers
//
// Responses:
// default: genericError
//
//	200: bezierDBResponse
func (controller *Controller) GetBeziers(c *gin.Context) {

	// source slice
	var bezierDBs []orm.BezierDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBeziers", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoBezier.GetDB()

	_, err := db.Find(&bezierDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	bezierAPIs := make([]orm.BezierAPI, 0)

	// for each bezier, update fields from the database nullable fields
	for idx := range bezierDBs {
		bezierDB := &bezierDBs[idx]
		_ = bezierDB
		var bezierAPI orm.BezierAPI

		// insertion point for updating fields
		bezierAPI.ID = bezierDB.ID
		bezierDB.CopyBasicFieldsToBezier_WOP(&bezierAPI.Bezier_WOP)
		bezierAPI.BezierPointersEncoding = bezierDB.BezierPointersEncoding
		bezierAPIs = append(bezierAPIs, bezierAPI)
	}

	c.JSON(http.StatusOK, bezierAPIs)
}

// PostBezier
//
// swagger:route POST /beziers beziers postBezier
//
// Creates a bezier
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBezier(c *gin.Context) {

	mutexBezier.Lock()
	defer mutexBezier.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBeziers", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoBezier.GetDB()

	// Validate input
	var input orm.BezierAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create bezier
	bezierDB := orm.BezierDB{}
	bezierDB.BezierPointersEncoding = input.BezierPointersEncoding
	bezierDB.CopyBasicFieldsFromBezier_WOP(&input.Bezier_WOP)

	_, err = db.Create(&bezierDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBezier.CheckoutPhaseOneInstance(&bezierDB)
	bezier := backRepo.BackRepoBezier.Map_BezierDBID_BezierPtr[bezierDB.ID]

	if bezier != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), bezier)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, bezierDB)
}

// GetBezier
//
// swagger:route GET /beziers/{ID} beziers getBezier
//
// Gets the details for a bezier.
//
// Responses:
// default: genericError
//
//	200: bezierDBResponse
func (controller *Controller) GetBezier(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBezier", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoBezier.GetDB()

	// Get bezierDB in DB
	var bezierDB orm.BezierDB
	if _, err := db.First(&bezierDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var bezierAPI orm.BezierAPI
	bezierAPI.ID = bezierDB.ID
	bezierAPI.BezierPointersEncoding = bezierDB.BezierPointersEncoding
	bezierDB.CopyBasicFieldsToBezier_WOP(&bezierAPI.Bezier_WOP)

	c.JSON(http.StatusOK, bezierAPI)
}

// UpdateBezier
//
// swagger:route PATCH /beziers/{ID} beziers updateBezier
//
// # Update a bezier
//
// Responses:
// default: genericError
//
//	200: bezierDBResponse
func (controller *Controller) UpdateBezier(c *gin.Context) {

	mutexBezier.Lock()
	defer mutexBezier.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBezier", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoBezier.GetDB()

	// Validate input
	var input orm.BezierAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var bezierDB orm.BezierDB

	// fetch the bezier
	_, err := db.First(&bezierDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	bezierDB.CopyBasicFieldsFromBezier_WOP(&input.Bezier_WOP)
	bezierDB.BezierPointersEncoding = input.BezierPointersEncoding

	db, _ = db.Model(&bezierDB)
	_, err = db.Updates(&bezierDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	bezierNew := new(models.Bezier)
	bezierDB.CopyBasicFieldsToBezier(bezierNew)

	// redeem pointers
	bezierDB.DecodePointers(backRepo, bezierNew)

	// get stage instance from DB instance, and call callback function
	bezierOld := backRepo.BackRepoBezier.Map_BezierDBID_BezierPtr[bezierDB.ID]
	if bezierOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), bezierOld, bezierNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the bezierDB
	c.JSON(http.StatusOK, bezierDB)
}

// DeleteBezier
//
// swagger:route DELETE /beziers/{ID} beziers deleteBezier
//
// # Delete a bezier
//
// default: genericError
//
//	200: bezierDBResponse
func (controller *Controller) DeleteBezier(c *gin.Context) {

	mutexBezier.Lock()
	defer mutexBezier.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBezier", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoBezier.GetDB()

	// Get model if exist
	var bezierDB orm.BezierDB
	if _, err := db.First(&bezierDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&bezierDB)

	// get an instance (not staged) from DB instance, and call callback function
	bezierDeleted := new(models.Bezier)
	bezierDB.CopyBasicFieldsToBezier(bezierDeleted)

	// get stage instance from DB instance, and call callback function
	bezierStaged := backRepo.BackRepoBezier.Map_BezierDBID_BezierPtr[bezierDB.ID]
	if bezierStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), bezierStaged, bezierDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
