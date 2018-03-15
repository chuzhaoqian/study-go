package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Notify struct {
	Id           int64     `orm:"column(id);auto"`
	Plat         int8      `orm:"column(plat);null" description:"推送接受平台 0全部 1网站 1Android 2IOS"`
	Title        string    `orm:"column(title);size(255);null" description:"推送标题"`
	Content      string    `orm:"column(content);null" description:"推送内容"`
	IsCron       int8      `orm:"column(is_cron);null" description:"是否是定时推送 0=>否 1=>是"`
	ReceiveRole  int8      `orm:"column(receive_role);null" description:"接受角色 0全部 1学生 2老师"`
	CourseIds    string    `orm:"column(course_ids);size(255);null" description:"课程id"`
	CategoryIds  string    `orm:"column(category_ids);size(255);null" description:"分类id"`
	NotifyTime   time.Time `orm:"column(notify_time);type(datetime);null" description:"推送时间"`
	NotifyDevice int64     `orm:"column(notify_device);null" description:"通知设备数"`
	Status       int8      `orm:"column(status);null" description:"推送状态 0未发送 1成功 2=>失败 3=>取消"`
	Type         int8      `orm:"column(type);null" description:"消息类型 0 后台创建 1直播通知"`
	UserId       int64     `orm:"column(user_id);null" description:"创建者"`
	CreateTime   time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间"`
	UpdateTime   time.Time `orm:"column(update_time);auto_now;type(datetime)" description:"更新时间"`
}

func (t *Notify) TableName() string {
	return "notify"
}

func init() {
	orm.RegisterModel(new(Notify))
}

// AddNotify insert a new Notify into database and returns
// last inserted Id on success.
func AddNotify(m *Notify) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNotifyById retrieves Notify by Id. Returns error if
// Id doesn't exist
func GetNotifyById(id int64) (v *Notify, err error) {
	o := orm.NewOrm()
	v = &Notify{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNotify retrieves all Notify matches certain condition. Returns empty list if
// no records exist
func GetAllNotify(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Notify))
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

	var l []Notify
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

// UpdateNotify updates Notify by Id and returns error if
// the record to be updated doesn't exist
func UpdateNotifyById(m *Notify) (err error) {
	o := orm.NewOrm()
	v := Notify{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNotify deletes Notify by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNotify(id int64) (err error) {
	o := orm.NewOrm()
	v := Notify{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Notify{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
