package main

import (
	"github.com/shengshunyan/go-monorepo/packagea"
	"github.com/shengshunyan/go-monorepo/packagecommon"
)

func main() {
	packagecommon.Common()
	packagea.Common()
	println("hello service a")
}
