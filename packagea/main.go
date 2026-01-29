package packagea

import "github.com/shengshunyan/go-monorepo/packagecommon"

func Common() {
	packagecommon.Common()
	println("This is package a 1")
}
