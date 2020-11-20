package main

import (
	_ "github.com/zhaoyunxing92/letter/public"
	"github.com/zhaoyunxing92/letter/router"
)

func main() {

	engine := router.InitRouter()

	if err := engine.Run();err!=nil{
		panic("gin error")
	}
}
