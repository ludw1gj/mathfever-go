package router

import "github.com/oxtoacart/bpool"

func init() {
	loadTemplates()
	tmplBufpool = bpool.NewBufferPool(64)
}
