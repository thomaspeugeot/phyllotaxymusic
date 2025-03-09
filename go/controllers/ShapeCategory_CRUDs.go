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
var __ShapeCategory__dummysDeclaration__ models.ShapeCategory
var __ShapeCategory_time__dummyDeclaration time.Duration

var mutexShapeCategory sync.Mutex

// An ShapeCategoryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getShapeCategory updateShapeCategory deleteShapeCategory
type ShapeCategoryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ShapeCategoryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postShapeCategory updateShapeCategory
type ShapeCategoryInput struct {
	// The ShapeCategory to submit or modify
	// in: body
	ShapeCategory *orm.ShapeCategoryAPI
}

// GetShapeCategorys
//
// swagger:route GET /shapecategorys shapecategorys getShapeCategorys
//
// # Get all shapecategorys
//
// Responses:
// default: genericError
//
//	200: shapecategoryDBResponse
func (controller *Controller) GetShapeCategorys(c *gin.Context) {

	// source slice
	var shapecategoryDBs []orm.ShapeCategoryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetShapeCategorys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoShapeCategory.GetDB()

	_, err := db.Find(&shapecategoryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	shapecategoryAPIs := make([]orm.ShapeCategoryAPI, 0)

	// for each shapecategory, update fields from the database nullable fields
	for idx := range shapecategoryDBs {
		shapecategoryDB := &shapecategoryDBs[idx]
		_ = shapecategoryDB
		var shapecategoryAPI orm.ShapeCategoryAPI

		// insertion point for updating fields
		shapecategoryAPI.ID = shapecategoryDB.ID
		shapecategoryDB.CopyBasicFieldsToShapeCategory_WOP(&shapecategoryAPI.ShapeCategory_WOP)
		shapecategoryAPI.ShapeCategoryPointersEncoding = shapecategoryDB.ShapeCategoryPointersEncoding
		shapecategoryAPIs = append(shapecategoryAPIs, shapecategoryAPI)
	}

	c.JSON(http.StatusOK, shapecategoryAPIs)
}

// PostShapeCategory
//
// swagger:route POST /shapecategorys shapecategorys postShapeCategory
//
// Creates a shapecategory
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostShapeCategory(c *gin.Context) {

	mutexShapeCategory.Lock()
	defer mutexShapeCategory.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostShapeCategorys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoShapeCategory.GetDB()

	// Validate input
	var input orm.ShapeCategoryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create shapecategory
	shapecategoryDB := orm.ShapeCategoryDB{}
	shapecategoryDB.ShapeCategoryPointersEncoding = input.ShapeCategoryPointersEncoding
	shapecategoryDB.CopyBasicFieldsFromShapeCategory_WOP(&input.ShapeCategory_WOP)

	_, err = db.Create(&shapecategoryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoShapeCategory.CheckoutPhaseOneInstance(&shapecategoryDB)
	shapecategory := backRepo.BackRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr[shapecategoryDB.ID]

	if shapecategory != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), shapecategory)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, shapecategoryDB)
}

// GetShapeCategory
//
// swagger:route GET /shapecategorys/{ID} shapecategorys getShapeCategory
//
// Gets the details for a shapecategory.
//
// Responses:
// default: genericError
//
//	200: shapecategoryDBResponse
func (controller *Controller) GetShapeCategory(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetShapeCategory", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoShapeCategory.GetDB()

	// Get shapecategoryDB in DB
	var shapecategoryDB orm.ShapeCategoryDB
	if _, err := db.First(&shapecategoryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var shapecategoryAPI orm.ShapeCategoryAPI
	shapecategoryAPI.ID = shapecategoryDB.ID
	shapecategoryAPI.ShapeCategoryPointersEncoding = shapecategoryDB.ShapeCategoryPointersEncoding
	shapecategoryDB.CopyBasicFieldsToShapeCategory_WOP(&shapecategoryAPI.ShapeCategory_WOP)

	c.JSON(http.StatusOK, shapecategoryAPI)
}

// UpdateShapeCategory
//
// swagger:route PATCH /shapecategorys/{ID} shapecategorys updateShapeCategory
//
// # Update a shapecategory
//
// Responses:
// default: genericError
//
//	200: shapecategoryDBResponse
func (controller *Controller) UpdateShapeCategory(c *gin.Context) {

	mutexShapeCategory.Lock()
	defer mutexShapeCategory.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateShapeCategory", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoShapeCategory.GetDB()

	// Validate input
	var input orm.ShapeCategoryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var shapecategoryDB orm.ShapeCategoryDB

	// fetch the shapecategory
	_, err := db.First(&shapecategoryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	shapecategoryDB.CopyBasicFieldsFromShapeCategory_WOP(&input.ShapeCategory_WOP)
	shapecategoryDB.ShapeCategoryPointersEncoding = input.ShapeCategoryPointersEncoding

	db, _ = db.Model(&shapecategoryDB)
	_, err = db.Updates(&shapecategoryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	shapecategoryNew := new(models.ShapeCategory)
	shapecategoryDB.CopyBasicFieldsToShapeCategory(shapecategoryNew)

	// redeem pointers
	shapecategoryDB.DecodePointers(backRepo, shapecategoryNew)

	// get stage instance from DB instance, and call callback function
	shapecategoryOld := backRepo.BackRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr[shapecategoryDB.ID]
	if shapecategoryOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), shapecategoryOld, shapecategoryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the shapecategoryDB
	c.JSON(http.StatusOK, shapecategoryDB)
}

// DeleteShapeCategory
//
// swagger:route DELETE /shapecategorys/{ID} shapecategorys deleteShapeCategory
//
// # Delete a shapecategory
//
// default: genericError
//
//	200: shapecategoryDBResponse
func (controller *Controller) DeleteShapeCategory(c *gin.Context) {

	mutexShapeCategory.Lock()
	defer mutexShapeCategory.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteShapeCategory", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoShapeCategory.GetDB()

	// Get model if exist
	var shapecategoryDB orm.ShapeCategoryDB
	if _, err := db.First(&shapecategoryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&shapecategoryDB)

	// get an instance (not staged) from DB instance, and call callback function
	shapecategoryDeleted := new(models.ShapeCategory)
	shapecategoryDB.CopyBasicFieldsToShapeCategory(shapecategoryDeleted)

	// get stage instance from DB instance, and call callback function
	shapecategoryStaged := backRepo.BackRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr[shapecategoryDB.ID]
	if shapecategoryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), shapecategoryStaged, shapecategoryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
