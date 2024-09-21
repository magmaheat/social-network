package main

import "github.com/magmaheat/social-network/sn-post/internal/app"

const pathConfig = "config/local.yaml"

func main() {
	app.Run(pathConfig)
}
