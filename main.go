package main

import (
	"fmt"
	"race-proj/router"
	"race-proj/setting"
)

func main() {
	router.InitRouter().Run(fmt.Sprintf(":%d", setting.Http_port))
}
