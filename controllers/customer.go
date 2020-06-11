package controllers

import (
	_ "beego-vuejs/configuration"
	"github.com/astaxie/beego"
)

// CustomerController operations for CustomerController
type CustomerController struct {
	beego.Controller
}

type Customer interface {
	Info() string

}

// URLMapping ...
func (c *CustomerController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create CustomerController
// @Param	body		body 	models.CustomerController	true		"body for CustomerController content"
// @Success 201 {object} models.CustomerController
// @Failure 403 body is empty
// @router / [post]
func (c *CustomerController) Post() {


}

// GetOne ...
// @Title GetOne
// @Description get CustomerController by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.CustomerController
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CustomerController) GetOne() {
	//_ := configuration.Client("numbers")
	//col.InsertOne({name: ""})


}

// GetAll ...
// @Title GetAll
// @Description get CustomerController
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.CustomerController
// @Failure 403
// @router / [get]
func (c *CustomerController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the CustomerController
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.CustomerController	true		"body for CustomerController content"
// @Success 200 {object} models.CustomerController
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CustomerController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the CustomerController
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CustomerController) Delete() {

}
