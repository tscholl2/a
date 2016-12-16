package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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
	r.GET("/game", func(c *gin.Context) {
		state.Lock()
		c.JSON(http.StatusOK, state.board)
		state.Unlock()
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
	state.Lock()
	state.board.MakeWorld(10)
	state.Unlock()
	for {
		state.Lock()
		state.board.Step()
		bJSON, _ := json.MarshalIndent(state.board, "", " ")
		ioutil.WriteFile("board.json", bJSON, 0644)
		state.Unlock()
		time.Sleep(1 * time.Second)
	}
}
