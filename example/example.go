package main

import (
	"fmt"

	"github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"

	echoSocket "github.com/partyzanex/echo-socket.io"
)

func main() {
	e := echo.New()

	e.Any("/socket.io/", socketIOWrapper().HandlerFunc)

	e.Logger.Fatal(e.Start(":8080"))
}

func socketIOWrapper() *echoSocket.Wrapper {
	wrapper, err := echoSocket.NewWrapper(nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	wrapper.OnConnect("/", func(context echo.Context, conn socketio.Conn) error {
		conn.SetContext("")
		fmt.Println("connected:", conn.ID())
		return nil
	})
	wrapper.OnError("/", func(context echo.Context, e error) {
		fmt.Println("meet error:", e)
	})
	wrapper.OnDisconnect("/", func(context echo.Context, conn socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})

	wrapper.OnEvent("/", "test", func(context echo.Context, conn socketio.Conn, msg string) {
		conn.SetContext(msg)
		fmt.Println("notice:", msg)
		conn.Emit("test", msg)
	})

	go wrapper.Serve()

	return wrapper
}
