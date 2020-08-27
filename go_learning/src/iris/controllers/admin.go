package controllers

type AdminController struct {}

func (c *AdminController) Any() interface{} {
	return map[string]interface{}{
		"code": 0,
		"msg": "ok",
	}
}