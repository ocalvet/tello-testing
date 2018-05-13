package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	drone := tello.NewDriver("8888")
	master := gobot.NewMaster()
	a := api.NewAPI(master)
	a.Start()
	// work := func() {
	// 	drone.TakeOff()

	// 	gobot.After(5*time.Second, func() {
	// 		drone.FrontFlip()
	// 	})

	// 	gobot.After(10*time.Second, func() {
	// 		drone.BackFlip()
	// 	})

	// 	gobot.After(15*time.Second, func() {
	// 		drone.Land()
	// 	})
	// }

	work := func() {
		drone.TakeOff()
		r := gin.Default()
		r.Use(cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowCredentials: true,
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"*"},
		}))
		r.Use()
		r.GET("/tello/takeoff", func(c *gin.Context) {
			drone.TakeOff()
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/land", func(c *gin.Context) {
			drone.Land()
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/left", func(c *gin.Context) {
			drone.Left(1)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/right", func(c *gin.Context) {
			drone.Right(1)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/forward", func(c *gin.Context) {
			drone.Forward(1)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/backward", func(c *gin.Context) {
			drone.Backward(1)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.Run(":9081")
	}
	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)
	master.AddRobot(robot)
	master.Start()
}
