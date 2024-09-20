package main

import "sn-auth/internal/app"

const configPath = "config/local.yaml"

func main() {
	app.Run(configPath)
}
