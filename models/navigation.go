package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Navigation struct {
	Id          int64     `orm:"column(id);auto" description:"导航ID"`
	SystemName  string    `orm:"column(system_name);size(30)" description:"系统默认名称，不可修改"`
	Name        string    `orm:"column(name);size(30)" description:"导航名称，可以自定义"`
	Url         string    `orm:"column(url);size(255)" description:"链接地址"`
	Sort        int64     `orm:"column(sort)" description:"显示顺序"`
	ParentId    int64     `orm:"column(parent_id)" description:"父导航ID"`
	Type        string    `orm:"column(type);size(30)" description:"类型"`
	IsOpen      uint8     `orm:"column(is_open)" description:"默认1开启 0关闭"`
	IsNewWindow uint8     `orm:"column(is_new_window)" description:"默认为1另开窗口"`
	CreateTime  time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间"`
	UpdateTime  time.Time `orm:"column(update_time);auto_now;type(datetime)" description:"更新时间"`
}

func (t *Navigation) TableName() string {
	return "navigation"
}

func init() {
	orm.RegisterModel(new(Navigation))
}

// AddNavigation insert a new Navigation into database and returns
// last inserted Id on success.
func AddNavigation(m *Navigation) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNavigationById retrieves Navigation by Id. Returns error if
// Id doesn't exist
func GetNavigationById(id int64) (v *Navigation, err error) {
	o := orm.NewOrm()
	v = &Navigation{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNavigation retrieves all Navigation matches certain condition. Returns empty list if
// no records exist
func GetAllNavigation(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Navigation))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Navigation
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateNavigation updates Navigation by Id and returns error if
// the record to be updated doesn't exist
func UpdateNavigationById(m *Navigation) (err error) {
	o := orm.NewOrm()
	v := Navigation{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNavigation deletes Navigation by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNavigation(id int64) (err error) {
	o := orm.NewOrm()
	v := Navigation{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Navigation{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
