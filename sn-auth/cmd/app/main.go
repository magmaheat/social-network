package main

import "github.com/magmaheat/social-network/sn-auth/internal/app"

const configPath = "config/local.yaml"

func main() {
	app.Run(configPath)
}
