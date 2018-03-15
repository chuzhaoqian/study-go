package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"math"
	"study-go/helper/response"
)

// Base operations for User
type BaseController struct {
	beego.Controller
}

var PageSize = 15
var Limit string
var Offset = 0
var Current = 1

type Paginate struct {
	Current  int // 当前页数
	Size     int // 每页数量
	PageNum  int // 总页数
	TotalNum int // 总数量
	IsAjax   int // 是否是ajax
}


func (c *BaseController) Success() {
	data := make(response.Data)
	c.Data["json"] = response.NewOKData(data)
	c.ServeJSON()
}

func (c *BaseController) SuccessPaginate(items []interface{}, paginate *Paginate) {

}

func (c *BaseController) Paginate(TotalNum int64) *Paginate {
	paginate := new(Paginate)

	if TotalNum <= 0 {
		return paginate
	}

	pageSize, err := c.GetInt("page", 0)
	if nil != err || pageSize == 0 {
		pageSize = PageSize
	}
	PageSize := pageSize

	paginate.Size = PageSize
	paginate.PageNum = int(math.Ceil(float64(TotalNum) / float64(pageSize)))
	page, err := c.GetInt("page", 0)
	paginate.Current = 1

	if nil != err || page == 0 {
		page = 1
		paginate.Current = page
	}

	if page > paginate.PageNum {
		paginate.Current = paginate.PageNum
	}

	Current = paginate.Current
	Offset = (paginate.Current - 1) * PageSize
	Limit = fmt.Sprintf("%d,%d", paginate.Size, Offset)

	return paginate
}

func (c *BaseController) SuccessData(data map[string]interface{}) {

}

func (c *BaseController) Error(id int) {

}

func (c *BaseController) ErrorInfo(err error) {

}

func (c *BaseController) ErrorValid(err error) {

}
