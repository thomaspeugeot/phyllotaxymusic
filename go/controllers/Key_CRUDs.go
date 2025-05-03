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
var __Key__dummysDeclaration__ models.Key
var __Key_time__dummyDeclaration time.Duration

var mutexKey sync.Mutex

// An KeyID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getKey updateKey deleteKey
type KeyID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// KeyInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postKey updateKey
type KeyInput struct {
	// The Key to submit or modify
	// in: body
	Key *orm.KeyAPI
}

// GetKeys
//
// swagger:route GET /keys keys getKeys
//
// # Get all keys
//
// Responses:
// default: genericError
//
//	200: keyDBResponse
func (controller *Controller) GetKeys(c *gin.Context) {

	// source slice
	var keyDBs []orm.KeyDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKeys", "Name", stackPath)
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
	db := backRepo.BackRepoKey.GetDB()

	_, err := db.Find(&keyDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	keyAPIs := make([]orm.KeyAPI, 0)

	// for each key, update fields from the database nullable fields
	for idx := range keyDBs {
		keyDB := &keyDBs[idx]
		_ = keyDB
		var keyAPI orm.KeyAPI

		// insertion point for updating fields
		keyAPI.ID = keyDB.ID
		keyDB.CopyBasicFieldsToKey_WOP(&keyAPI.Key_WOP)
		keyAPI.KeyPointersEncoding = keyDB.KeyPointersEncoding
		keyAPIs = append(keyAPIs, keyAPI)
	}

	c.JSON(http.StatusOK, keyAPIs)
}

// PostKey
//
// swagger:route POST /keys keys postKey
//
// Creates a key
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostKey(c *gin.Context) {

	mutexKey.Lock()
	defer mutexKey.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostKeys", "Name", stackPath)
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
	db := backRepo.BackRepoKey.GetDB()

	// Validate input
	var input orm.KeyAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create key
	keyDB := orm.KeyDB{}
	keyDB.KeyPointersEncoding = input.KeyPointersEncoding
	keyDB.CopyBasicFieldsFromKey_WOP(&input.Key_WOP)

	_, err = db.Create(&keyDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoKey.CheckoutPhaseOneInstance(&keyDB)
	key := backRepo.BackRepoKey.Map_KeyDBID_KeyPtr[keyDB.ID]

	if key != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), key)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, keyDB)
}

// GetKey
//
// swagger:route GET /keys/{ID} keys getKey
//
// Gets the details for a key.
//
// Responses:
// default: genericError
//
//	200: keyDBResponse
func (controller *Controller) GetKey(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKey", "Name", stackPath)
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
	db := backRepo.BackRepoKey.GetDB()

	// Get keyDB in DB
	var keyDB orm.KeyDB
	if _, err := db.First(&keyDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var keyAPI orm.KeyAPI
	keyAPI.ID = keyDB.ID
	keyAPI.KeyPointersEncoding = keyDB.KeyPointersEncoding
	keyDB.CopyBasicFieldsToKey_WOP(&keyAPI.Key_WOP)

	c.JSON(http.StatusOK, keyAPI)
}

// UpdateKey
//
// swagger:route PATCH /keys/{ID} keys updateKey
//
// # Update a key
//
// Responses:
// default: genericError
//
//	200: keyDBResponse
func (controller *Controller) UpdateKey(c *gin.Context) {

	mutexKey.Lock()
	defer mutexKey.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateKey", "Name", stackPath)
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
	db := backRepo.BackRepoKey.GetDB()

	// Validate input
	var input orm.KeyAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var keyDB orm.KeyDB

	// fetch the key
	_, err := db.First(&keyDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	keyDB.CopyBasicFieldsFromKey_WOP(&input.Key_WOP)
	keyDB.KeyPointersEncoding = input.KeyPointersEncoding

	db, _ = db.Model(&keyDB)
	_, err = db.Updates(&keyDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	keyNew := new(models.Key)
	keyDB.CopyBasicFieldsToKey(keyNew)

	// redeem pointers
	keyDB.DecodePointers(backRepo, keyNew)

	// get stage instance from DB instance, and call callback function
	keyOld := backRepo.BackRepoKey.Map_KeyDBID_KeyPtr[keyDB.ID]
	if keyOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), keyOld, keyNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the keyDB
	c.JSON(http.StatusOK, keyDB)
}

// DeleteKey
//
// swagger:route DELETE /keys/{ID} keys deleteKey
//
// # Delete a key
//
// default: genericError
//
//	200: keyDBResponse
func (controller *Controller) DeleteKey(c *gin.Context) {

	mutexKey.Lock()
	defer mutexKey.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteKey", "Name", stackPath)
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
	db := backRepo.BackRepoKey.GetDB()

	// Get model if exist
	var keyDB orm.KeyDB
	if _, err := db.First(&keyDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&keyDB)

	// get an instance (not staged) from DB instance, and call callback function
	keyDeleted := new(models.Key)
	keyDB.CopyBasicFieldsToKey(keyDeleted)

	// get stage instance from DB instance, and call callback function
	keyStaged := backRepo.BackRepoKey.Map_KeyDBID_KeyPtr[keyDB.ID]
	if keyStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), keyStaged, keyDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
