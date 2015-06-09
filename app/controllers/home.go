package controllers

import (
	"github.com/revel/revel"
)

type Home struct {
	BaseController
}

func (c Home) Index() revel.Result {
	data := map[string]interface{}{
		"title":   "跟踪您的工作进度-随时随地",
	}
	return c.Render(data)
}
