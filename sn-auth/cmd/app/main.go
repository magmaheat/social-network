package main

import "sn-auth/internal/app"

const configPath = "configs/local.yaml"

func main() {
	app.Run(configPath)
}
