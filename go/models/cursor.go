package models

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

type Cursor struct {
	Name string

	// in degrees
	AngleDegree float64
	Length      float64

	CenterX, CenterY float64

	StartX, EndX    float64
	DurationSeconds float64
	IsMoving        bool

	Presentation
}

func (movingline *Cursor) Draw(gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	parameter *Parameter,
) {
	themeWidth := parameter.RotatedAxis.Length
	_ = themeWidth

	line := new(gongsvg_models.Line)
	line.Name = movingline.Name
	layer.Lines = append(layer.Lines, line)

	angleRad := DegreesToRadians(movingline.AngleDegree)

	line.X1 = parameter.OriginX + movingline.CenterX
	line.Y1 = parameter.OriginY - movingline.CenterY

	line.X2 = line.X1 + movingline.Length*math.Cos(angleRad)
	line.Y2 = line.Y1 - movingline.Length*math.Sin(angleRad)

}

// OnWebSocketConnection is the callback from the client
func (cursor *Cursor) OnWebSocketConnection(c *gin.Context) {
	log.Println("Cursor cursorStart, OnWebSocketConnection")

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

	// Create a context that is canceled when the connection is closed
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}

	// Start a goroutine to read from the WebSocket to detect disconnection
	go func() {
		for {
			// ReadMessage is used to detect client disconnection
			_, _, err := wsConnection.ReadMessage()
			if err != nil {
				log.Println("github.com/thomaspeugeot/phylotaxymusic/go", stackPath, "WS client disconnected:", err)
				cancel() // Cancel the context
				return
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			// Context canceled, exit the loop
			return
		default:
			// message := ""

			// // Send backRepo data
			// err = wsConnection.WriteJSON(message)
			// if err != nil {
			// 	log.Println("cursor start\n", stackPath,
			// 		"client no longer receiver web socket message,closing websocket handler")
			// 	fmt.Println(err)
			// 	cancel() // Cancel the context
			// 	return
			// } else {
			// 	log.Println(time.Now().Format("2006-01-02 15:04:05.000000"),
			// 		"sent cursor start")
			// }

		}
	}
}
