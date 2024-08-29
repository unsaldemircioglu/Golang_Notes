package main

import (
	"fmt"

	"github.com/twpayne/go-geom"
)

func main() {
	// println("Hello World!")
	fmt.Println("Hello,World")

	geom.NewMultiPolygon(geom.XY)
}

/*
go version
go run main.go

go mod init oneLessonGo

go mod tidy


go get github.com/twpayne/go-geom
*/
