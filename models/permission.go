package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Permission struct {
	Id         int64     `orm:"column(id);auto" description:"规则id,自增主键"`
	Module     string    `orm:"column(module);size(20)" description:"规则所属module"`
	Type       string    `orm:"column(type);size(30)" description:"权限规则分类，请加应用前缀,如admin_"`
	Name       string    `orm:"column(name);size(30)" description:"规则唯一英文标识,全小写"`
	Param      string    `orm:"column(param);size(255);null" description:"额外url参数"`
	Title      string    `orm:"column(title);size(20)" description:"规则中文描述"`
	Status     int8      `orm:"column(status)" description:"是否有效(0:无效,1:有效)"`
	Condition  string    `orm:"column(condition);size(300)" description:"规则附加条件(为以后的页面视图权限做准备)"`
	CreateTime time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间"`
	UpdateTime time.Time `orm:"column(update_time);auto_now;type(datetime)" description:"更新时间"`
}

func (t *Permission) TableName() string {
	return "permission"
}

func init() {
	orm.RegisterModel(new(Permission))
}

// AddPermission insert a new Permission into database and returns
// last inserted Id on success.
func AddPermission(m *Permission) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPermissionById retrieves Permission by Id. Returns error if
// Id doesn't exist
func GetPermissionById(id int64) (v *Permission, err error) {
	o := orm.NewOrm()
	v = &Permission{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPermission retrieves all Permission matches certain condition. Returns empty list if
// no records exist
func GetAllPermission(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Permission))
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

	var l []Permission
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

// UpdatePermission updates Permission by Id and returns error if
// the record to be updated doesn't exist
func UpdatePermissionById(m *Permission) (err error) {
	o := orm.NewOrm()
	v := Permission{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePermission deletes Permission by Id and returns error if
// the record to be deleted doesn't exist
func DeletePermission(id int64) (err error) {
	o := orm.NewOrm()
	v := Permission{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Permission{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
