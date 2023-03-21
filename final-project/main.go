package main

import (
	"io/ioutil"
	"net/http"

	"GLIM_Hacktiv8/golang-intermediate/final-project/client"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		content, err := ioutil.ReadFile("template/chat.html")
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "could not open html")
		}

		return ctx.HTML(http.StatusOK, string(content))
	})
	e.Static("/template", "template")

	e.Any("/ws", func(ctx echo.Context) error {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}
		currentGorillaConn, err := upgrader.Upgrade(ctx.Response().Writer, ctx.Request(), nil)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "Could not open websocket connection")
		}

		username := ctx.Request().URL.Query().Get("username")
		age := ctx.Request().URL.Query().Get("age")

		go client.HandleIO(&client.WebSocketConnection{
			Conn:     currentGorillaConn,
			Username: username,
			Age:      age,
		})
		return nil
	})

	e.Start(":8080")
}

/*
based on gorilla websocket repo
var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "template/home.html")
}

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
*/
