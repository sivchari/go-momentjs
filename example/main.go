package main

import (
	"log"

	"github.com/sivchari/go-momentjs"
)

func main() {
	m, err := momentjs.New()
	if err != nil {
		panic(err)
	}
	const format = "YYYY/MM/DD HH:mm:ss dddd"
	m.Global.Set("fmt", m.QjsCtx.String(format))
	defer m.ClearGlobal("fmt")
	result, err := m.QjsCtx.Eval("moment().format(fmt)")
	if err != nil {
		panic(err)
	}
	resultDate := result.String()
	result.Free()
	log.Println(resultDate)
}
