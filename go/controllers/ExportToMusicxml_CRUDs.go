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
var __ExportToMusicxml__dummysDeclaration__ models.ExportToMusicxml
var __ExportToMusicxml_time__dummyDeclaration time.Duration

var mutexExportToMusicxml sync.Mutex

// An ExportToMusicxmlID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getExportToMusicxml updateExportToMusicxml deleteExportToMusicxml
type ExportToMusicxmlID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ExportToMusicxmlInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postExportToMusicxml updateExportToMusicxml
type ExportToMusicxmlInput struct {
	// The ExportToMusicxml to submit or modify
	// in: body
	ExportToMusicxml *orm.ExportToMusicxmlAPI
}

// GetExportToMusicxmls
//
// swagger:route GET /exporttomusicxmls exporttomusicxmls getExportToMusicxmls
//
// # Get all exporttomusicxmls
//
// Responses:
// default: genericError
//
//	200: exporttomusicxmlDBResponse
func (controller *Controller) GetExportToMusicxmls(c *gin.Context) {

	// source slice
	var exporttomusicxmlDBs []orm.ExportToMusicxmlDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetExportToMusicxmls", "Name", stackPath)
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
	db := backRepo.BackRepoExportToMusicxml.GetDB()

	_, err := db.Find(&exporttomusicxmlDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	exporttomusicxmlAPIs := make([]orm.ExportToMusicxmlAPI, 0)

	// for each exporttomusicxml, update fields from the database nullable fields
	for idx := range exporttomusicxmlDBs {
		exporttomusicxmlDB := &exporttomusicxmlDBs[idx]
		_ = exporttomusicxmlDB
		var exporttomusicxmlAPI orm.ExportToMusicxmlAPI

		// insertion point for updating fields
		exporttomusicxmlAPI.ID = exporttomusicxmlDB.ID
		exporttomusicxmlDB.CopyBasicFieldsToExportToMusicxml_WOP(&exporttomusicxmlAPI.ExportToMusicxml_WOP)
		exporttomusicxmlAPI.ExportToMusicxmlPointersEncoding = exporttomusicxmlDB.ExportToMusicxmlPointersEncoding
		exporttomusicxmlAPIs = append(exporttomusicxmlAPIs, exporttomusicxmlAPI)
	}

	c.JSON(http.StatusOK, exporttomusicxmlAPIs)
}

// PostExportToMusicxml
//
// swagger:route POST /exporttomusicxmls exporttomusicxmls postExportToMusicxml
//
// Creates a exporttomusicxml
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostExportToMusicxml(c *gin.Context) {

	mutexExportToMusicxml.Lock()
	defer mutexExportToMusicxml.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostExportToMusicxmls", "Name", stackPath)
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
	db := backRepo.BackRepoExportToMusicxml.GetDB()

	// Validate input
	var input orm.ExportToMusicxmlAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create exporttomusicxml
	exporttomusicxmlDB := orm.ExportToMusicxmlDB{}
	exporttomusicxmlDB.ExportToMusicxmlPointersEncoding = input.ExportToMusicxmlPointersEncoding
	exporttomusicxmlDB.CopyBasicFieldsFromExportToMusicxml_WOP(&input.ExportToMusicxml_WOP)

	_, err = db.Create(&exporttomusicxmlDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoExportToMusicxml.CheckoutPhaseOneInstance(&exporttomusicxmlDB)
	exporttomusicxml := backRepo.BackRepoExportToMusicxml.Map_ExportToMusicxmlDBID_ExportToMusicxmlPtr[exporttomusicxmlDB.ID]

	if exporttomusicxml != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), exporttomusicxml)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, exporttomusicxmlDB)
}

// GetExportToMusicxml
//
// swagger:route GET /exporttomusicxmls/{ID} exporttomusicxmls getExportToMusicxml
//
// Gets the details for a exporttomusicxml.
//
// Responses:
// default: genericError
//
//	200: exporttomusicxmlDBResponse
func (controller *Controller) GetExportToMusicxml(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetExportToMusicxml", "Name", stackPath)
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
	db := backRepo.BackRepoExportToMusicxml.GetDB()

	// Get exporttomusicxmlDB in DB
	var exporttomusicxmlDB orm.ExportToMusicxmlDB
	if _, err := db.First(&exporttomusicxmlDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var exporttomusicxmlAPI orm.ExportToMusicxmlAPI
	exporttomusicxmlAPI.ID = exporttomusicxmlDB.ID
	exporttomusicxmlAPI.ExportToMusicxmlPointersEncoding = exporttomusicxmlDB.ExportToMusicxmlPointersEncoding
	exporttomusicxmlDB.CopyBasicFieldsToExportToMusicxml_WOP(&exporttomusicxmlAPI.ExportToMusicxml_WOP)

	c.JSON(http.StatusOK, exporttomusicxmlAPI)
}

// UpdateExportToMusicxml
//
// swagger:route PATCH /exporttomusicxmls/{ID} exporttomusicxmls updateExportToMusicxml
//
// # Update a exporttomusicxml
//
// Responses:
// default: genericError
//
//	200: exporttomusicxmlDBResponse
func (controller *Controller) UpdateExportToMusicxml(c *gin.Context) {

	mutexExportToMusicxml.Lock()
	defer mutexExportToMusicxml.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
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
	db := backRepo.BackRepoExportToMusicxml.GetDB()

	// Validate input
	var input orm.ExportToMusicxmlAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var exporttomusicxmlDB orm.ExportToMusicxmlDB

	// fetch the exporttomusicxml
	_, err := db.First(&exporttomusicxmlDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	exporttomusicxmlDB.CopyBasicFieldsFromExportToMusicxml_WOP(&input.ExportToMusicxml_WOP)
	exporttomusicxmlDB.ExportToMusicxmlPointersEncoding = input.ExportToMusicxmlPointersEncoding

	db, _ = db.Model(&exporttomusicxmlDB)
	_, err = db.Updates(&exporttomusicxmlDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	exporttomusicxmlNew := new(models.ExportToMusicxml)
	exporttomusicxmlDB.CopyBasicFieldsToExportToMusicxml(exporttomusicxmlNew)

	// redeem pointers
	exporttomusicxmlDB.DecodePointers(backRepo, exporttomusicxmlNew)

	// get stage instance from DB instance, and call callback function
	exporttomusicxmlOld := backRepo.BackRepoExportToMusicxml.Map_ExportToMusicxmlDBID_ExportToMusicxmlPtr[exporttomusicxmlDB.ID]
	if exporttomusicxmlOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), exporttomusicxmlOld, exporttomusicxmlNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the exporttomusicxmlDB
	c.JSON(http.StatusOK, exporttomusicxmlDB)
}

// DeleteExportToMusicxml
//
// swagger:route DELETE /exporttomusicxmls/{ID} exporttomusicxmls deleteExportToMusicxml
//
// # Delete a exporttomusicxml
//
// default: genericError
//
//	200: exporttomusicxmlDBResponse
func (controller *Controller) DeleteExportToMusicxml(c *gin.Context) {

	mutexExportToMusicxml.Lock()
	defer mutexExportToMusicxml.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteExportToMusicxml", "Name", stackPath)
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
	db := backRepo.BackRepoExportToMusicxml.GetDB()

	// Get model if exist
	var exporttomusicxmlDB orm.ExportToMusicxmlDB
	if _, err := db.First(&exporttomusicxmlDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&exporttomusicxmlDB)

	// get an instance (not staged) from DB instance, and call callback function
	exporttomusicxmlDeleted := new(models.ExportToMusicxml)
	exporttomusicxmlDB.CopyBasicFieldsToExportToMusicxml(exporttomusicxmlDeleted)

	// get stage instance from DB instance, and call callback function
	exporttomusicxmlStaged := backRepo.BackRepoExportToMusicxml.Map_ExportToMusicxmlDBID_ExportToMusicxmlPtr[exporttomusicxmlDB.ID]
	if exporttomusicxmlStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), exporttomusicxmlStaged, exporttomusicxmlDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
