
package controllers

type OptionsController struct {
	BaseController
}

func (c *OptionsController) Options() {
	c.Success()
}
