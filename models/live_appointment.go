package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type LiveAppointment struct {
	Id         int64     `orm:"column(id);auto"`
	LiveId     int64     `orm:"column(live_id)" description:"直播id"`
	UserId     int64     `orm:"column(user_id)" description:"用户id"`
	SendTime   time.Time `orm:"column(send_time);type(datetime);null" description:"发送时间"`
	Status     int8      `orm:"column(status)" description:"状态 0未发送  1已发送"`
	CreateTime time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间"`
	UpdateTime time.Time `orm:"column(update_time);auto_now;type(datetime)" description:"更新时间"`
}

func (t *LiveAppointment) TableName() string {
	return "live_appointment"
}

func init() {
	orm.RegisterModel(new(LiveAppointment))
}

// AddLiveAppointment insert a new LiveAppointment into database and returns
// last inserted Id on success.
func AddLiveAppointment(m *LiveAppointment) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLiveAppointmentById retrieves LiveAppointment by Id. Returns error if
// Id doesn't exist
func GetLiveAppointmentById(id int64) (v *LiveAppointment, err error) {
	o := orm.NewOrm()
	v = &LiveAppointment{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLiveAppointment retrieves all LiveAppointment matches certain condition. Returns empty list if
// no records exist
func GetAllLiveAppointment(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LiveAppointment))
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

	var l []LiveAppointment
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

// UpdateLiveAppointment updates LiveAppointment by Id and returns error if
// the record to be updated doesn't exist
func UpdateLiveAppointmentById(m *LiveAppointment) (err error) {
	o := orm.NewOrm()
	v := LiveAppointment{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLiveAppointment deletes LiveAppointment by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLiveAppointment(id int64) (err error) {
	o := orm.NewOrm()
	v := LiveAppointment{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LiveAppointment{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
