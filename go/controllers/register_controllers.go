// generated code - do not edit
package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thomaspeugeot/phylotaxymusic/go/orm"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/thomaspeugeot/phylotaxymusic/go")
	{ // insertion point for registrations
		v1.GET("/v1/circles", GetController().GetCircles)
		v1.GET("/v1/circles/:id", GetController().GetCircle)
		v1.POST("/v1/circles", GetController().PostCircle)
		v1.PATCH("/v1/circles/:id", GetController().UpdateCircle)
		v1.PUT("/v1/circles/:id", GetController().UpdateCircle)
		v1.DELETE("/v1/circles/:id", GetController().DeleteCircle)

		v1.GET("/v1/horizontalaxiss", GetController().GetHorizontalAxiss)
		v1.GET("/v1/horizontalaxiss/:id", GetController().GetHorizontalAxis)
		v1.POST("/v1/horizontalaxiss", GetController().PostHorizontalAxis)
		v1.PATCH("/v1/horizontalaxiss/:id", GetController().UpdateHorizontalAxis)
		v1.PUT("/v1/horizontalaxiss/:id", GetController().UpdateHorizontalAxis)
		v1.DELETE("/v1/horizontalaxiss/:id", GetController().DeleteHorizontalAxis)

		v1.GET("/v1/initialaxiss", GetController().GetInitialAxiss)
		v1.GET("/v1/initialaxiss/:id", GetController().GetInitialAxis)
		v1.POST("/v1/initialaxiss", GetController().PostInitialAxis)
		v1.PATCH("/v1/initialaxiss/:id", GetController().UpdateInitialAxis)
		v1.PUT("/v1/initialaxiss/:id", GetController().UpdateInitialAxis)
		v1.DELETE("/v1/initialaxiss/:id", GetController().DeleteInitialAxis)

		v1.GET("/v1/lines", GetController().GetLines)
		v1.GET("/v1/lines/:id", GetController().GetLine)
		v1.POST("/v1/lines", GetController().PostLine)
		v1.PATCH("/v1/lines/:id", GetController().UpdateLine)
		v1.PUT("/v1/lines/:id", GetController().UpdateLine)
		v1.DELETE("/v1/lines/:id", GetController().DeleteLine)

		v1.GET("/v1/parameters", GetController().GetParameters)
		v1.GET("/v1/parameters/:id", GetController().GetParameter)
		v1.POST("/v1/parameters", GetController().PostParameter)
		v1.PATCH("/v1/parameters/:id", GetController().UpdateParameter)
		v1.PUT("/v1/parameters/:id", GetController().UpdateParameter)
		v1.DELETE("/v1/parameters/:id", GetController().DeleteParameter)

		v1.GET("/v1/rhombuss", GetController().GetRhombuss)
		v1.GET("/v1/rhombuss/:id", GetController().GetRhombus)
		v1.POST("/v1/rhombuss", GetController().PostRhombus)
		v1.PATCH("/v1/rhombuss/:id", GetController().UpdateRhombus)
		v1.PUT("/v1/rhombuss/:id", GetController().UpdateRhombus)
		v1.DELETE("/v1/rhombuss/:id", GetController().DeleteRhombus)

		v1.GET("/v1/rhombusgrids", GetController().GetRhombusGrids)
		v1.GET("/v1/rhombusgrids/:id", GetController().GetRhombusGrid)
		v1.POST("/v1/rhombusgrids", GetController().PostRhombusGrid)
		v1.PATCH("/v1/rhombusgrids/:id", GetController().UpdateRhombusGrid)
		v1.PUT("/v1/rhombusgrids/:id", GetController().UpdateRhombusGrid)
		v1.DELETE("/v1/rhombusgrids/:id", GetController().DeleteRhombusGrid)

		v1.GET("/v1/verticalaxiss", GetController().GetVerticalAxiss)
		v1.GET("/v1/verticalaxiss/:id", GetController().GetVerticalAxis)
		v1.POST("/v1/verticalaxiss", GetController().PostVerticalAxis)
		v1.PATCH("/v1/verticalaxiss/:id", GetController().UpdateVerticalAxis)
		v1.PUT("/v1/verticalaxiss/:id", GetController().UpdateVerticalAxis)
		v1.DELETE("/v1/verticalaxiss/:id", GetController().DeleteVerticalAxis)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)

		v1.GET("/v1/ws/commitfrombacknb", GetController().onWebSocketRequestForCommitFromBackNb)
		v1.GET("/v1/ws/stage", GetController().onWebSocketRequestForBackRepoContent)
	}
}

// onWebSocketRequestForCommitFromBackNb is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForCommitFromBackNb(c *gin.Context) {

	// log.Println("Stack github.com/thomaspeugeot/phylotaxymusic/go, onWebSocketRequestForCommitFromBackNb")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:8080" || origin == "http://localhost:4200"
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb()

	for nbCommitBackRepo := range updateCommitBackRepoNbChannel {

		// Send elapsed time as a string over the WebSocket connection
		err = wsConnection.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", nbCommitBackRepo)))
		if err != nil {
			log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
			fmt.Println(err)
			return
		}
	}
}

// onWebSocketRequestForBackRepoContent is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForBackRepoContent(c *gin.Context) {

	// log.Println("Stack github.com/thomaspeugeot/phylotaxymusic/go, onWebSocketRequestForBackRepoContent")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:8080" || origin == "http://localhost:4200"
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb()

	backRepoData := new(orm.BackRepoData)
	orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

	err = wsConnection.WriteJSON(backRepoData)
	// log.Println("Stack github.com/thomaspeugeot/phylotaxymusic/go, onWebSocketRequestForBackRepoContent, first sent back repo of", stackPath)
	if err != nil {
		log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	}

	for nbCommitBackRepo := range updateCommitBackRepoNbChannel {
		_ = nbCommitBackRepo

		backRepoData := new(orm.BackRepoData)
		orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

		// Send backRepo data
		err = wsConnection.WriteJSON(backRepoData)

		// log.Println("Stack github.com/thomaspeugeot/phylotaxymusic/go, onWebSocketRequestForBackRepoContent, sent back repo of", stackPath)

		if err != nil {
			log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
			fmt.Println(err)
			return
		}
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/phylotaxymusic/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
