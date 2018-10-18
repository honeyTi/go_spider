package main

import (
	"single_spider/engine"
	"single_spider/zhenai/parser"
)

func main() {
	engine.Run(engine.Resquest{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}