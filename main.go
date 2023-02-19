package main

import (
	"shmily/model"
	"shmily/routers"
)

func main() {
	model.Database()

	r := routers.NewRouter()

	r.Run()
}
