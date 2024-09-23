package main

import "github.com/magmaheat/social-network/sn-auth/internal/app"

const configPath = "configs/local.yaml"

func main() {
	app.Run(configPath)
}
