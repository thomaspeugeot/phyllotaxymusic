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
var __NoteInfo__dummysDeclaration__ models.NoteInfo
var __NoteInfo_time__dummyDeclaration time.Duration

var mutexNoteInfo sync.Mutex

// An NoteInfoID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNoteInfo updateNoteInfo deleteNoteInfo
type NoteInfoID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// NoteInfoInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNoteInfo updateNoteInfo
type NoteInfoInput struct {
	// The NoteInfo to submit or modify
	// in: body
	NoteInfo *orm.NoteInfoAPI
}

// GetNoteInfos
//
// swagger:route GET /noteinfos noteinfos getNoteInfos
//
// # Get all noteinfos
//
// Responses:
// default: genericError
//
//	200: noteinfoDBResponse
func (controller *Controller) GetNoteInfos(c *gin.Context) {

	// source slice
	var noteinfoDBs []orm.NoteInfoDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNoteInfos", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNoteInfo.GetDB()

	_, err := db.Find(&noteinfoDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	noteinfoAPIs := make([]orm.NoteInfoAPI, 0)

	// for each noteinfo, update fields from the database nullable fields
	for idx := range noteinfoDBs {
		noteinfoDB := &noteinfoDBs[idx]
		_ = noteinfoDB
		var noteinfoAPI orm.NoteInfoAPI

		// insertion point for updating fields
		noteinfoAPI.ID = noteinfoDB.ID
		noteinfoDB.CopyBasicFieldsToNoteInfo_WOP(&noteinfoAPI.NoteInfo_WOP)
		noteinfoAPI.NoteInfoPointersEncoding = noteinfoDB.NoteInfoPointersEncoding
		noteinfoAPIs = append(noteinfoAPIs, noteinfoAPI)
	}

	c.JSON(http.StatusOK, noteinfoAPIs)
}

// PostNoteInfo
//
// swagger:route POST /noteinfos noteinfos postNoteInfo
//
// Creates a noteinfo
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostNoteInfo(c *gin.Context) {

	mutexNoteInfo.Lock()
	defer mutexNoteInfo.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostNoteInfos", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNoteInfo.GetDB()

	// Validate input
	var input orm.NoteInfoAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create noteinfo
	noteinfoDB := orm.NoteInfoDB{}
	noteinfoDB.NoteInfoPointersEncoding = input.NoteInfoPointersEncoding
	noteinfoDB.CopyBasicFieldsFromNoteInfo_WOP(&input.NoteInfo_WOP)

	_, err = db.Create(&noteinfoDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoNoteInfo.CheckoutPhaseOneInstance(&noteinfoDB)
	noteinfo := backRepo.BackRepoNoteInfo.Map_NoteInfoDBID_NoteInfoPtr[noteinfoDB.ID]

	if noteinfo != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), noteinfo)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, noteinfoDB)
}

// GetNoteInfo
//
// swagger:route GET /noteinfos/{ID} noteinfos getNoteInfo
//
// Gets the details for a noteinfo.
//
// Responses:
// default: genericError
//
//	200: noteinfoDBResponse
func (controller *Controller) GetNoteInfo(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNoteInfo", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNoteInfo.GetDB()

	// Get noteinfoDB in DB
	var noteinfoDB orm.NoteInfoDB
	if _, err := db.First(&noteinfoDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var noteinfoAPI orm.NoteInfoAPI
	noteinfoAPI.ID = noteinfoDB.ID
	noteinfoAPI.NoteInfoPointersEncoding = noteinfoDB.NoteInfoPointersEncoding
	noteinfoDB.CopyBasicFieldsToNoteInfo_WOP(&noteinfoAPI.NoteInfo_WOP)

	c.JSON(http.StatusOK, noteinfoAPI)
}

// UpdateNoteInfo
//
// swagger:route PATCH /noteinfos/{ID} noteinfos updateNoteInfo
//
// # Update a noteinfo
//
// Responses:
// default: genericError
//
//	200: noteinfoDBResponse
func (controller *Controller) UpdateNoteInfo(c *gin.Context) {

	mutexNoteInfo.Lock()
	defer mutexNoteInfo.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateNoteInfo", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNoteInfo.GetDB()

	// Validate input
	var input orm.NoteInfoAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var noteinfoDB orm.NoteInfoDB

	// fetch the noteinfo
	_, err := db.First(&noteinfoDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	noteinfoDB.CopyBasicFieldsFromNoteInfo_WOP(&input.NoteInfo_WOP)
	noteinfoDB.NoteInfoPointersEncoding = input.NoteInfoPointersEncoding

	db, _ = db.Model(&noteinfoDB)
	_, err = db.Updates(noteinfoDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	noteinfoNew := new(models.NoteInfo)
	noteinfoDB.CopyBasicFieldsToNoteInfo(noteinfoNew)

	// redeem pointers
	noteinfoDB.DecodePointers(backRepo, noteinfoNew)

	// get stage instance from DB instance, and call callback function
	noteinfoOld := backRepo.BackRepoNoteInfo.Map_NoteInfoDBID_NoteInfoPtr[noteinfoDB.ID]
	if noteinfoOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), noteinfoOld, noteinfoNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the noteinfoDB
	c.JSON(http.StatusOK, noteinfoDB)
}

// DeleteNoteInfo
//
// swagger:route DELETE /noteinfos/{ID} noteinfos deleteNoteInfo
//
// # Delete a noteinfo
//
// default: genericError
//
//	200: noteinfoDBResponse
func (controller *Controller) DeleteNoteInfo(c *gin.Context) {

	mutexNoteInfo.Lock()
	defer mutexNoteInfo.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteNoteInfo", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNoteInfo.GetDB()

	// Get model if exist
	var noteinfoDB orm.NoteInfoDB
	if _, err := db.First(&noteinfoDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&noteinfoDB)

	// get an instance (not staged) from DB instance, and call callback function
	noteinfoDeleted := new(models.NoteInfo)
	noteinfoDB.CopyBasicFieldsToNoteInfo(noteinfoDeleted)

	// get stage instance from DB instance, and call callback function
	noteinfoStaged := backRepo.BackRepoNoteInfo.Map_NoteInfoDBID_NoteInfoPtr[noteinfoDB.ID]
	if noteinfoStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), noteinfoStaged, noteinfoDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
