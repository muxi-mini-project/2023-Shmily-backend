package main

import (
	"shmily/conf"
	"shmily/routers"
)

func main() {
	conf.Init()

	r := routers.NewRouter()

	r.Run()
}
