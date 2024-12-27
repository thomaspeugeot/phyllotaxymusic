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
var __Parameter__dummysDeclaration__ models.Parameter
var __Parameter_time__dummyDeclaration time.Duration

var mutexParameter sync.Mutex

// An ParameterID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getParameter updateParameter deleteParameter
type ParameterID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ParameterInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postParameter updateParameter
type ParameterInput struct {
	// The Parameter to submit or modify
	// in: body
	Parameter *orm.ParameterAPI
}

// GetParameters
//
// swagger:route GET /parameters parameters getParameters
//
// # Get all parameters
//
// Responses:
// default: genericError
//
//	200: parameterDBResponse
func (controller *Controller) GetParameters(c *gin.Context) {

	// source slice
	var parameterDBs []orm.ParameterDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetParameters", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phyllotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParameter.GetDB()

	_, err := db.Find(&parameterDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	parameterAPIs := make([]orm.ParameterAPI, 0)

	// for each parameter, update fields from the database nullable fields
	for idx := range parameterDBs {
		parameterDB := &parameterDBs[idx]
		_ = parameterDB
		var parameterAPI orm.ParameterAPI

		// insertion point for updating fields
		parameterAPI.ID = parameterDB.ID
		parameterDB.CopyBasicFieldsToParameter_WOP(&parameterAPI.Parameter_WOP)
		parameterAPI.ParameterPointersEncoding = parameterDB.ParameterPointersEncoding
		parameterAPIs = append(parameterAPIs, parameterAPI)
	}

	c.JSON(http.StatusOK, parameterAPIs)
}

// PostParameter
//
// swagger:route POST /parameters parameters postParameter
//
// Creates a parameter
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostParameter(c *gin.Context) {

	mutexParameter.Lock()
	defer mutexParameter.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostParameters", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phyllotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParameter.GetDB()

	// Validate input
	var input orm.ParameterAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create parameter
	parameterDB := orm.ParameterDB{}
	parameterDB.ParameterPointersEncoding = input.ParameterPointersEncoding
	parameterDB.CopyBasicFieldsFromParameter_WOP(&input.Parameter_WOP)

	_, err = db.Create(&parameterDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoParameter.CheckoutPhaseOneInstance(&parameterDB)
	parameter := backRepo.BackRepoParameter.Map_ParameterDBID_ParameterPtr[parameterDB.ID]

	if parameter != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), parameter)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, parameterDB)
}

// GetParameter
//
// swagger:route GET /parameters/{ID} parameters getParameter
//
// Gets the details for a parameter.
//
// Responses:
// default: genericError
//
//	200: parameterDBResponse
func (controller *Controller) GetParameter(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetParameter", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phyllotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParameter.GetDB()

	// Get parameterDB in DB
	var parameterDB orm.ParameterDB
	if _, err := db.First(&parameterDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var parameterAPI orm.ParameterAPI
	parameterAPI.ID = parameterDB.ID
	parameterAPI.ParameterPointersEncoding = parameterDB.ParameterPointersEncoding
	parameterDB.CopyBasicFieldsToParameter_WOP(&parameterAPI.Parameter_WOP)

	c.JSON(http.StatusOK, parameterAPI)
}

// UpdateParameter
//
// swagger:route PATCH /parameters/{ID} parameters updateParameter
//
// # Update a parameter
//
// Responses:
// default: genericError
//
//	200: parameterDBResponse
func (controller *Controller) UpdateParameter(c *gin.Context) {

	mutexParameter.Lock()
	defer mutexParameter.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateParameter", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phyllotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParameter.GetDB()

	// Validate input
	var input orm.ParameterAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var parameterDB orm.ParameterDB

	// fetch the parameter
	_, err := db.First(&parameterDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	parameterDB.CopyBasicFieldsFromParameter_WOP(&input.Parameter_WOP)
	parameterDB.ParameterPointersEncoding = input.ParameterPointersEncoding

	db, _ = db.Model(&parameterDB)
	_, err = db.Updates(&parameterDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	parameterNew := new(models.Parameter)
	parameterDB.CopyBasicFieldsToParameter(parameterNew)

	// redeem pointers
	parameterDB.DecodePointers(backRepo, parameterNew)

	// get stage instance from DB instance, and call callback function
	parameterOld := backRepo.BackRepoParameter.Map_ParameterDBID_ParameterPtr[parameterDB.ID]
	if parameterOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), parameterOld, parameterNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the parameterDB
	c.JSON(http.StatusOK, parameterDB)
}

// DeleteParameter
//
// swagger:route DELETE /parameters/{ID} parameters deleteParameter
//
// # Delete a parameter
//
// default: genericError
//
//	200: parameterDBResponse
func (controller *Controller) DeleteParameter(c *gin.Context) {

	mutexParameter.Lock()
	defer mutexParameter.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteParameter", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phyllotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParameter.GetDB()

	// Get model if exist
	var parameterDB orm.ParameterDB
	if _, err := db.First(&parameterDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&parameterDB)

	// get an instance (not staged) from DB instance, and call callback function
	parameterDeleted := new(models.Parameter)
	parameterDB.CopyBasicFieldsToParameter(parameterDeleted)

	// get stage instance from DB instance, and call callback function
	parameterStaged := backRepo.BackRepoParameter.Map_ParameterDBID_ParameterPtr[parameterDB.ID]
	if parameterStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), parameterStaged, parameterDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
