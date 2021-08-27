package main

import (
	"gee"
)

func main() {
	r := gee.New()
	r.Static("/assets", "./static_asset")
	r.Run(":9999")
}
