package main

import Routers "belajar-gin/Routers"

func main() {

	var PORT = ":8080"

	Routers.StartServer().Run(PORT)
}