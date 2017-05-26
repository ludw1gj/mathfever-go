package router

import "github.com/oxtoacart/bpool"

func init() {
	loadTemplates()
	tmplBufPool = bpool.NewBufferPool(64)
}
