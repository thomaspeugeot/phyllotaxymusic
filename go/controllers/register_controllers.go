// generated code - do not edit
package controllers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/thomaspeugeot/phyllotaxymusic/go/orm"

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
	v1 := r.Group("/api/github.com/thomaspeugeot/phyllotaxymusic/go")
	{ // insertion point for registrations
		v1.GET("/v1/axiss", GetController().GetAxiss)
		v1.GET("/v1/axiss/:id", GetController().GetAxis)
		v1.POST("/v1/axiss", GetController().PostAxis)
		v1.PATCH("/v1/axiss/:id", GetController().UpdateAxis)
		v1.PUT("/v1/axiss/:id", GetController().UpdateAxis)
		v1.DELETE("/v1/axiss/:id", GetController().DeleteAxis)

		v1.GET("/v1/axisgrids", GetController().GetAxisGrids)
		v1.GET("/v1/axisgrids/:id", GetController().GetAxisGrid)
		v1.POST("/v1/axisgrids", GetController().PostAxisGrid)
		v1.PATCH("/v1/axisgrids/:id", GetController().UpdateAxisGrid)
		v1.PUT("/v1/axisgrids/:id", GetController().UpdateAxisGrid)
		v1.DELETE("/v1/axisgrids/:id", GetController().DeleteAxisGrid)

		v1.GET("/v1/beziers", GetController().GetBeziers)
		v1.GET("/v1/beziers/:id", GetController().GetBezier)
		v1.POST("/v1/beziers", GetController().PostBezier)
		v1.PATCH("/v1/beziers/:id", GetController().UpdateBezier)
		v1.PUT("/v1/beziers/:id", GetController().UpdateBezier)
		v1.DELETE("/v1/beziers/:id", GetController().DeleteBezier)

		v1.GET("/v1/beziergrids", GetController().GetBezierGrids)
		v1.GET("/v1/beziergrids/:id", GetController().GetBezierGrid)
		v1.POST("/v1/beziergrids", GetController().PostBezierGrid)
		v1.PATCH("/v1/beziergrids/:id", GetController().UpdateBezierGrid)
		v1.PUT("/v1/beziergrids/:id", GetController().UpdateBezierGrid)
		v1.DELETE("/v1/beziergrids/:id", GetController().DeleteBezierGrid)

		v1.GET("/v1/beziergridstacks", GetController().GetBezierGridStacks)
		v1.GET("/v1/beziergridstacks/:id", GetController().GetBezierGridStack)
		v1.POST("/v1/beziergridstacks", GetController().PostBezierGridStack)
		v1.PATCH("/v1/beziergridstacks/:id", GetController().UpdateBezierGridStack)
		v1.PUT("/v1/beziergridstacks/:id", GetController().UpdateBezierGridStack)
		v1.DELETE("/v1/beziergridstacks/:id", GetController().DeleteBezierGridStack)

		v1.GET("/v1/chapters", GetController().GetChapters)
		v1.GET("/v1/chapters/:id", GetController().GetChapter)
		v1.POST("/v1/chapters", GetController().PostChapter)
		v1.PATCH("/v1/chapters/:id", GetController().UpdateChapter)
		v1.PUT("/v1/chapters/:id", GetController().UpdateChapter)
		v1.DELETE("/v1/chapters/:id", GetController().DeleteChapter)

		v1.GET("/v1/circles", GetController().GetCircles)
		v1.GET("/v1/circles/:id", GetController().GetCircle)
		v1.POST("/v1/circles", GetController().PostCircle)
		v1.PATCH("/v1/circles/:id", GetController().UpdateCircle)
		v1.PUT("/v1/circles/:id", GetController().UpdateCircle)
		v1.DELETE("/v1/circles/:id", GetController().DeleteCircle)

		v1.GET("/v1/circlegrids", GetController().GetCircleGrids)
		v1.GET("/v1/circlegrids/:id", GetController().GetCircleGrid)
		v1.POST("/v1/circlegrids", GetController().PostCircleGrid)
		v1.PATCH("/v1/circlegrids/:id", GetController().UpdateCircleGrid)
		v1.PUT("/v1/circlegrids/:id", GetController().UpdateCircleGrid)
		v1.DELETE("/v1/circlegrids/:id", GetController().DeleteCircleGrid)

		v1.GET("/v1/contents", GetController().GetContents)
		v1.GET("/v1/contents/:id", GetController().GetContent)
		v1.POST("/v1/contents", GetController().PostContent)
		v1.PATCH("/v1/contents/:id", GetController().UpdateContent)
		v1.PUT("/v1/contents/:id", GetController().UpdateContent)
		v1.DELETE("/v1/contents/:id", GetController().DeleteContent)

		v1.GET("/v1/exporttomusicxmls", GetController().GetExportToMusicxmls)
		v1.GET("/v1/exporttomusicxmls/:id", GetController().GetExportToMusicxml)
		v1.POST("/v1/exporttomusicxmls", GetController().PostExportToMusicxml)
		v1.PATCH("/v1/exporttomusicxmls/:id", GetController().UpdateExportToMusicxml)
		v1.PUT("/v1/exporttomusicxmls/:id", GetController().UpdateExportToMusicxml)
		v1.DELETE("/v1/exporttomusicxmls/:id", GetController().DeleteExportToMusicxml)

		v1.GET("/v1/frontcurves", GetController().GetFrontCurves)
		v1.GET("/v1/frontcurves/:id", GetController().GetFrontCurve)
		v1.POST("/v1/frontcurves", GetController().PostFrontCurve)
		v1.PATCH("/v1/frontcurves/:id", GetController().UpdateFrontCurve)
		v1.PUT("/v1/frontcurves/:id", GetController().UpdateFrontCurve)
		v1.DELETE("/v1/frontcurves/:id", GetController().DeleteFrontCurve)

		v1.GET("/v1/frontcurvestacks", GetController().GetFrontCurveStacks)
		v1.GET("/v1/frontcurvestacks/:id", GetController().GetFrontCurveStack)
		v1.POST("/v1/frontcurvestacks", GetController().PostFrontCurveStack)
		v1.PATCH("/v1/frontcurvestacks/:id", GetController().UpdateFrontCurveStack)
		v1.PUT("/v1/frontcurvestacks/:id", GetController().UpdateFrontCurveStack)
		v1.DELETE("/v1/frontcurvestacks/:id", GetController().DeleteFrontCurveStack)

		v1.GET("/v1/horizontalaxiss", GetController().GetHorizontalAxiss)
		v1.GET("/v1/horizontalaxiss/:id", GetController().GetHorizontalAxis)
		v1.POST("/v1/horizontalaxiss", GetController().PostHorizontalAxis)
		v1.PATCH("/v1/horizontalaxiss/:id", GetController().UpdateHorizontalAxis)
		v1.PUT("/v1/horizontalaxiss/:id", GetController().UpdateHorizontalAxis)
		v1.DELETE("/v1/horizontalaxiss/:id", GetController().DeleteHorizontalAxis)

		v1.GET("/v1/keys", GetController().GetKeys)
		v1.GET("/v1/keys/:id", GetController().GetKey)
		v1.POST("/v1/keys", GetController().PostKey)
		v1.PATCH("/v1/keys/:id", GetController().UpdateKey)
		v1.PUT("/v1/keys/:id", GetController().UpdateKey)
		v1.DELETE("/v1/keys/:id", GetController().DeleteKey)

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

		v1.GET("/v1/shapecategorys", GetController().GetShapeCategorys)
		v1.GET("/v1/shapecategorys/:id", GetController().GetShapeCategory)
		v1.POST("/v1/shapecategorys", GetController().PostShapeCategory)
		v1.PATCH("/v1/shapecategorys/:id", GetController().UpdateShapeCategory)
		v1.PUT("/v1/shapecategorys/:id", GetController().UpdateShapeCategory)
		v1.DELETE("/v1/shapecategorys/:id", GetController().DeleteShapeCategory)

		v1.GET("/v1/spiralbeziers", GetController().GetSpiralBeziers)
		v1.GET("/v1/spiralbeziers/:id", GetController().GetSpiralBezier)
		v1.POST("/v1/spiralbeziers", GetController().PostSpiralBezier)
		v1.PATCH("/v1/spiralbeziers/:id", GetController().UpdateSpiralBezier)
		v1.PUT("/v1/spiralbeziers/:id", GetController().UpdateSpiralBezier)
		v1.DELETE("/v1/spiralbeziers/:id", GetController().DeleteSpiralBezier)

		v1.GET("/v1/spiralbeziergrids", GetController().GetSpiralBezierGrids)
		v1.GET("/v1/spiralbeziergrids/:id", GetController().GetSpiralBezierGrid)
		v1.POST("/v1/spiralbeziergrids", GetController().PostSpiralBezierGrid)
		v1.PATCH("/v1/spiralbeziergrids/:id", GetController().UpdateSpiralBezierGrid)
		v1.PUT("/v1/spiralbeziergrids/:id", GetController().UpdateSpiralBezierGrid)
		v1.DELETE("/v1/spiralbeziergrids/:id", GetController().DeleteSpiralBezierGrid)

		v1.GET("/v1/spiralcircles", GetController().GetSpiralCircles)
		v1.GET("/v1/spiralcircles/:id", GetController().GetSpiralCircle)
		v1.POST("/v1/spiralcircles", GetController().PostSpiralCircle)
		v1.PATCH("/v1/spiralcircles/:id", GetController().UpdateSpiralCircle)
		v1.PUT("/v1/spiralcircles/:id", GetController().UpdateSpiralCircle)
		v1.DELETE("/v1/spiralcircles/:id", GetController().DeleteSpiralCircle)

		v1.GET("/v1/spiralcirclegrids", GetController().GetSpiralCircleGrids)
		v1.GET("/v1/spiralcirclegrids/:id", GetController().GetSpiralCircleGrid)
		v1.POST("/v1/spiralcirclegrids", GetController().PostSpiralCircleGrid)
		v1.PATCH("/v1/spiralcirclegrids/:id", GetController().UpdateSpiralCircleGrid)
		v1.PUT("/v1/spiralcirclegrids/:id", GetController().UpdateSpiralCircleGrid)
		v1.DELETE("/v1/spiralcirclegrids/:id", GetController().DeleteSpiralCircleGrid)

		v1.GET("/v1/spirallines", GetController().GetSpiralLines)
		v1.GET("/v1/spirallines/:id", GetController().GetSpiralLine)
		v1.POST("/v1/spirallines", GetController().PostSpiralLine)
		v1.PATCH("/v1/spirallines/:id", GetController().UpdateSpiralLine)
		v1.PUT("/v1/spirallines/:id", GetController().UpdateSpiralLine)
		v1.DELETE("/v1/spirallines/:id", GetController().DeleteSpiralLine)

		v1.GET("/v1/spirallinegrids", GetController().GetSpiralLineGrids)
		v1.GET("/v1/spirallinegrids/:id", GetController().GetSpiralLineGrid)
		v1.POST("/v1/spirallinegrids", GetController().PostSpiralLineGrid)
		v1.PATCH("/v1/spirallinegrids/:id", GetController().UpdateSpiralLineGrid)
		v1.PUT("/v1/spirallinegrids/:id", GetController().UpdateSpiralLineGrid)
		v1.DELETE("/v1/spirallinegrids/:id", GetController().DeleteSpiralLineGrid)

		v1.GET("/v1/spiralorigins", GetController().GetSpiralOrigins)
		v1.GET("/v1/spiralorigins/:id", GetController().GetSpiralOrigin)
		v1.POST("/v1/spiralorigins", GetController().PostSpiralOrigin)
		v1.PATCH("/v1/spiralorigins/:id", GetController().UpdateSpiralOrigin)
		v1.PUT("/v1/spiralorigins/:id", GetController().UpdateSpiralOrigin)
		v1.DELETE("/v1/spiralorigins/:id", GetController().DeleteSpiralOrigin)

		v1.GET("/v1/spiralrhombuss", GetController().GetSpiralRhombuss)
		v1.GET("/v1/spiralrhombuss/:id", GetController().GetSpiralRhombus)
		v1.POST("/v1/spiralrhombuss", GetController().PostSpiralRhombus)
		v1.PATCH("/v1/spiralrhombuss/:id", GetController().UpdateSpiralRhombus)
		v1.PUT("/v1/spiralrhombuss/:id", GetController().UpdateSpiralRhombus)
		v1.DELETE("/v1/spiralrhombuss/:id", GetController().DeleteSpiralRhombus)

		v1.GET("/v1/spiralrhombusgrids", GetController().GetSpiralRhombusGrids)
		v1.GET("/v1/spiralrhombusgrids/:id", GetController().GetSpiralRhombusGrid)
		v1.POST("/v1/spiralrhombusgrids", GetController().PostSpiralRhombusGrid)
		v1.PATCH("/v1/spiralrhombusgrids/:id", GetController().UpdateSpiralRhombusGrid)
		v1.PUT("/v1/spiralrhombusgrids/:id", GetController().UpdateSpiralRhombusGrid)
		v1.DELETE("/v1/spiralrhombusgrids/:id", GetController().DeleteSpiralRhombusGrid)

		v1.GET("/v1/verticalaxiss", GetController().GetVerticalAxiss)
		v1.GET("/v1/verticalaxiss/:id", GetController().GetVerticalAxis)
		v1.POST("/v1/verticalaxiss", GetController().PostVerticalAxis)
		v1.PATCH("/v1/verticalaxiss/:id", GetController().UpdateVerticalAxis)
		v1.PUT("/v1/verticalaxiss/:id", GetController().UpdateVerticalAxis)
		v1.DELETE("/v1/verticalaxiss/:id", GetController().DeleteVerticalAxis)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)

		v1.GET("/v1/ws/stage", GetController().onWebSocketRequestForBackRepoContent)

		v1.GET("/v1/stacks", GetController().stacks)
	}
}

func (controller *Controller) stacks(c *gin.Context) {

	var res []string

	for k := range controller.Map_BackRepos {
		res = append(res, k)
	}

	c.JSON(http.StatusOK, res)
}

// onWebSocketRequestForBackRepoContent is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForBackRepoContent(c *gin.Context) {

	// log.Println("Stack github.com/thomaspeugeot/phyllotaxymusic/go, onWebSocketRequestForBackRepoContent")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.

	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			if origin == "" {
				log.Printf("CheckOrigin: Rejected - Origin header is empty. Request from: %s", r.RemoteAddr)
				return false // Or handle as per your security policy
			}

			u, err := url.Parse(origin)
			if err != nil {
				log.Printf("CheckOrigin: Rejected - Invalid Origin URL '%s'. Error: %v", origin, err)
				return false // Invalid URL
			}

			portStr := u.Port()

			if portStr == "" {
				// If no port is specified, it might be using default HTTP/HTTPS ports.
				// For this specific request, we'll assume a port must be present.
				log.Printf("CheckOrigin: Rejected - No port specified in Origin URL '%s'", origin)
				return false
			}

			port, err := strconv.Atoi(portStr)
			if err != nil {
				log.Printf("CheckOrigin: Rejected - Port '%s' in Origin URL '%s' is not a valid number. Error: %v", portStr, origin, err)
				return false // Port is not a valid number
			}

			// Check if the port is 4200 OR in the range 8000-9000
			allowed := port == 4200 || (port >= 8000 && port <= 9000)
			if !allowed {
				log.Printf("CheckOrigin: Rejected - Port %d from Origin '%s' is not in the allowed list (4200 or 8000-9000)", port, origin)
				return false
			}

			// log.Printf("CheckOrigin: Accepted - Origin '%s' with port %d", origin, port)
			return true
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	// Create a context that is canceled when the connection is closed
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "Name", stackPath)
		}
	}

	index := controller.listenerIndex
	controller.listenerIndex++

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "WebSocket Connect, Stack github.com/thomaspeugeot/phyllotaxymusic/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb(ctx)

	// Start a goroutine to read from the WebSocket to detect disconnection
	go func() {
		for {
			// ReadMessage is used to detect client disconnection
			_, _, err := wsConnection.ReadMessage()
			if err != nil {
				log.Println("github.com/thomaspeugeot/phyllotaxymusic/go", stackPath, "WS client disconnected:", err)
				cancel() // Cancel the context
				return
			}
		}
	}()

	backRepoData := new(orm.BackRepoData)
	orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)
	backRepoData.GONG__Index = index

	refresh := 0
	// Marshal the data to JSON first to be able to get its size
	jsonData, err := json.Marshal(backRepoData)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return
	}

	// Get the size of the JSON data in bytes
	jsonSize := len(jsonData)

    // Calculate the full SHA-256 hash
    fullHash := sha256.Sum256(jsonData)

    // Use the first 12 characters for a shorter, yet highly unique, signature
    shortHash := hex.EncodeToString(fullHash[:])[0:12]

	// Use WriteMessage to send the pre-marshaled JSON data.
	// websocket.TextMessage is typically what WriteJSON uses.
	err = wsConnection.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Println("github.com/thomaspeugeot/phyllotaxymusic/go:\n",
			"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	} else {
		// 1. Extract the component name from the long path for cleaner logs
		// For example, "github.com/fullstack-lang/gong/lib/table/go" becomes "table"
		parts := strings.Split("github.com/thomaspeugeot/phyllotaxymusic/go", "/") // Assuming goFilePath holds the path
		component := "unknown"
		if len(parts) > 2 {
			component = parts[len(parts)-2]
		}

		// 2. Use a single, formatted log line
		log.Printf(
			"%-12s | %-85s | Idx: %d | Size: %-9s | Hash: %s",
			component,
			stackPath,
			index,
			formatBytes(jsonSize),
			shortHash,
		)
	}
	for {
		select {
		case <-ctx.Done():
			// Context canceled, exit the loop
			return
		default:
			for nbCommitBackRepo := range updateCommitBackRepoNbChannel {
				_ = nbCommitBackRepo
				refresh += 1

				backRepoData := new(orm.BackRepoData)
				orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)
				backRepoData.GONG__Index = index

				// Set write deadline to prevent blocking indefinitely
				wsConnection.SetWriteDeadline(time.Now().Add(10 * time.Second))

				// Send backRepo data
				// Marshal the data to JSON first to be able to get its size
				jsonData, err := json.Marshal(backRepoData)
				if err != nil {
					log.Printf("Error marshaling JSON: %v", err)
					return
				}

				// Get the size of the JSON data in bytes
				jsonSize := len(jsonData)

				// Calculate the full SHA-256 hash
				fullHash := sha256.Sum256(jsonData)

				// Use the first 12 characters for a shorter, yet highly unique, signature
				shortHash := hex.EncodeToString(fullHash[:])[0:12]

				// Use WriteMessage to send the pre-marshaled JSON data.
				// websocket.TextMessage is typically what WriteJSON uses.
				err = wsConnection.WriteMessage(websocket.TextMessage, jsonData)
				if err != nil {
					log.Println("github.com/thomaspeugeot/phyllotaxymusic/go:\n",
						"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
					fmt.Println(err)
					return
				} else {
					// 1. Extract the component name from the long path for cleaner logs
					// For example, "github.com/fullstack-lang/gong/lib/table/go" becomes "table"
					parts := strings.Split("github.com/thomaspeugeot/phyllotaxymusic/go", "/") // Assuming goFilePath holds the path
					component := "unknown"
					if len(parts) > 2 {
						component = parts[len(parts)-2]
					}

					// 2. Use a single, formatted log line
					log.Printf(
						"%-12s | %-85s | Idx: %d | Size: %-9s | Hash: %s",
						component,
						stackPath,
						index,
						formatBytes(jsonSize),
						shortHash,
					)
				}
			}
		}
	}
}

// formatBytes converts a size in bytes to a human-readable string (KB, MB, GB).
func formatBytes(size int) string {
    if size < 1024 {
        return fmt.Sprintf("%d B", size)
    }
    sizeInKB := float64(size) / 1024.0
    if sizeInKB < 1024.0 {
        // For KB, show one decimal place if it's not a whole number
        if math.Mod(sizeInKB, 1.0) == 0 {
            return fmt.Sprintf("%.0f KB", sizeInKB)
        }
        return fmt.Sprintf("%.1f KB", sizeInKB)
    }
    sizeInMB := sizeInKB / 1024.0
    return fmt.Sprintf("%.2f MB", sizeInMB)
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "Name", stackPath)
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
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "Name", stackPath)
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
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
