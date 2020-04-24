package main

import (
	"fmt"
	"gallery/routes"
	"gallery/services"
)

func main() {
	err := services.ConnectDB("root:dylinh952001@tcp(127.0.0.1:3306)/galleries?parseTime=true")
	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println("Connected !")
	err = services.CreateLogger()
	if err != nil {
		panic(err)
	}
	defer services.CloseLogger()
	g := routes.Create()
	g.Run("127.0.0.1:3000")

}
