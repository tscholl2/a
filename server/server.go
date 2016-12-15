package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tscholl2/a/entity"
	"github.com/tscholl2/a/game"
)

var Port string

var state = struct {
	sync.RWMutex
	board *game.Game
}{board: new(game.Game)}

func main() {
	flag.StringVar(&Port, "port", "8072", "port to run this server on (default: 8072)")
	flag.Parse()

	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	r.POST("/new", func(c *gin.Context) {
		var json entity.Attributes
		err := c.BindJSON(&json)
		if err == nil {
			fmt.Println(json)
		} else {
			log.Println(err)
		}
		c.String(http.StatusOK, "recieved")
	})

	fmt.Println("Running on 127.0.0.1:" + Port)

	fmt.Println("Starting game...")
	go startGame()
	r.Run(":" + Port)
}

func startGame() {
	state.board.MakeWorld(10)
	state.Lock()
	for {
		state.Unlock()
		state.board.Step()
		state.Lock()
		time.Sleep(1)
	}
}
