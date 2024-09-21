package main

import "github.com/magmaheat/social-network/tree/main/sn-post/internal/app"

const pathConfig = "configs/local.yaml"

func main() {
	app.Run(pathConfig)
}
