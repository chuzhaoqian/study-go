package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type AdminMenu struct {
	Id             int64     `orm:"column(id);auto"`
	Pid            int64     `orm:"column(pid)" description:"父级id"`
	Model          string    `orm:"column(model);size(30)" description:"控制器"`
	Status         int8      `orm:"column(status)" description:"状态，1显示，0不显示"`
	Name           string    `orm:"column(name);size(50)" description:"菜单名称"`
	PermissionName string    `orm:"column(permission_name);size(50)" description:"权限标识"`
	Icon           string    `orm:"column(icon);size(50);null" description:"菜单图标"`
	Remark         string    `orm:"column(remark);size(200)" description:"备注"`
	Sort           int64     `orm:"column(sort)" description:"排序ID"`
	CreateTime     time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间"`
	UpdateTime     time.Time `orm:"column(update_time);auto_now;type(datetime)" description:"更新时间"`
}

func (t *AdminMenu) TableName() string {
	return "admin_menu"
}

func init() {
	orm.RegisterModel(new(AdminMenu))
}

// AddAdminMenu insert a new AdminMenu into database and returns
// last inserted Id on success.
func AddAdminMenu(m *AdminMenu) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAdminMenuById retrieves AdminMenu by Id. Returns error if
// Id doesn't exist
func GetAdminMenuById(id int64) (v *AdminMenu, err error) {
	o := orm.NewOrm()
	v = &AdminMenu{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAdminMenu retrieves all AdminMenu matches certain condition. Returns empty list if
// no records exist
func GetAllAdminMenu(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AdminMenu))
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

	var l []AdminMenu
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

// UpdateAdminMenu updates AdminMenu by Id and returns error if
// the record to be updated doesn't exist
func UpdateAdminMenuById(m *AdminMenu) (err error) {
	o := orm.NewOrm()
	v := AdminMenu{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAdminMenu deletes AdminMenu by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAdminMenu(id int64) (err error) {
	o := orm.NewOrm()
	v := AdminMenu{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AdminMenu{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
