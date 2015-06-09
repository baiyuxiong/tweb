package controllers

import (
	"github.com/revel/revel"
)

// init is called when the first request into the controller is made
func init() {
	revel.InterceptMethod((*BaseController).Before, revel.BEFORE)
}

type BaseController struct {
	*revel.Controller
}

// Before is called prior to the controller method
func (c *BaseController) Before() revel.Result {
	return nil
}