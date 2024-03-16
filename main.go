package main

import "belajar-gin/routers"

func main() {
	var PORT = ":8080"

	routers.StarServer().Run(PORT)
}
