package main

import (
    "github.com/gorilla/websocket"
    "net/http"
    "time"
)

var upGrader = websocket.Upgrader{}

func main() {
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        http.ServeFile(writer, request, "index.html")
    })

    http.HandleFunc("/v1/ws", func(writer http.ResponseWriter, request *http.Request) {
        conn, err := upGrader.Upgrade(writer, request, nil)
        if err != nil {
            return
        }
        go func(conn *websocket.Conn) {
            for {
                mType, msg, _ := conn.ReadMessage()
                conn.WriteMessage(mType, msg)
            }
        }(conn)
    })

    http.HandleFunc("/v2/ws", func(writer http.ResponseWriter, request *http.Request) {
        conn, _ := upGrader.Upgrade(writer, request, nil)
        go func(conn *websocket.Conn) {
            _, msg, _ := conn.ReadMessage()
            println(string(msg))
        }(conn)
    })

    http.HandleFunc("/v3/ws", func(writer http.ResponseWriter, request *http.Request) {
        conn, _ := upGrader.Upgrade(writer, request, nil)
        go func(conn *websocket.Conn) {
            ch := time.Tick(5 * time.Second)
            for range ch {
                conn.WriteJSON(User{
                    Username: "mark",
                    Age:      24,
                })
            }
        }(conn)
    })

    http.HandleFunc("/v4/ws", func(writer http.ResponseWriter, request *http.Request) {
        conn, _ := upGrader.Upgrade(writer, request, nil)
        go func(conn *websocket.Conn) {
            for {
                _, _, err := conn.ReadMessage()
                if err != nil {
                    conn.Close()
                }
            }
        }(conn)
        go func(conn *websocket.Conn) {
            ch := time.Tick(5 * time.Second)
            for range ch {
                conn.WriteJSON(User{
                    Username: "mark",
                    Age:      24,
                })
            }
        }(conn)
    })

    http.ListenAndServe(":3000", nil)
}

type User struct {
    Username string `json:"username"`
    Age      int    `json:"age"`
}
