package main

import "sn-post/internal/app"

const pathConfig = "config/local.yaml"

func main() {
	app.Run(pathConfig)
}
