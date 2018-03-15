package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Course struct {
	Id           int64     `orm:"column(id);auto" description:"课程ID"`
	Title        string    `orm:"column(title);size(100)" description:"课程标题"`
	Status       int8      `orm:"column(status);null" description:"建课状态，0 未发布,1 已发布,2 已关闭"`
	Price        float32   `orm:"column(price)" description:"课程价格"`
	OriginalCost float32   `orm:"column(original_cost);null" description:"原价"`
	IsFree       int8      `orm:"column(is_free);null" description:"是否是收费课，1免费，0收费"`
	SmallPic     string    `orm:"column(small_pic);size(255)" description:"小图"`
	Pic          string    `orm:"column(pic);size(255)" description:"中图"`
	LargePic     string    `orm:"column(large_pic);size(255)" description:"大图"`
	Description  string    `orm:"column(description);null" description:"说明"`
	Degree       int8      `orm:"column(degree)" description:"难易程度  1 简单 2一般 3 困难"`
	Grade        float32   `orm:"column(grade);null" description:"课程评级"`
	StudentNum   int64     `orm:"column(student_num)" description:"学员数"`
	LessonNum    int64     `orm:"column(lesson_num)" description:"课时数"`
	PreviewPower int8      `orm:"column(preview_power)" description:"试读权限 0关闭试读  1试读第一章 2试读前2章 3试读前3章"`
	Type         int8      `orm:"column(type)" description:"课程类型/课程模板  0点播课 1直播课"`
	UserId       uint      `orm:"column(user_id);null" description:"创建者"`
	CreateTime   time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间"`
	UpdateTime   time.Time `orm:"column(update_time);auto_now;type(datetime)" description:"更新时间"`
}

func (t *Course) TableName() string {
	return "course"
}

func init() {
	orm.RegisterModel(new(Course))
}

// AddCourse insert a new Course into database and returns
// last inserted Id on success.
func AddCourse(m *Course) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCourseById retrieves Course by Id. Returns error if
// Id doesn't exist
func GetCourseById(id int64) (v *Course, err error) {
	o := orm.NewOrm()
	v = &Course{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCourse retrieves all Course matches certain condition. Returns empty list if
// no records exist
func GetAllCourse(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Course))
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

	var l []Course
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

// UpdateCourse updates Course by Id and returns error if
// the record to be updated doesn't exist
func UpdateCourseById(m *Course) (err error) {
	o := orm.NewOrm()
	v := Course{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCourse deletes Course by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCourse(id int64) (err error) {
	o := orm.NewOrm()
	v := Course{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Course{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
