package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CourseLessonLearn struct {
	Id             int64     `orm:"column(id);auto" description:"课时学习记录ID"`
	UserId         int64     `orm:"column(user_id)" description:"学员ID"`
	CourseId       int64     `orm:"column(course_id)" description:"课程id"`
	CourseLessonId int64     `orm:"column(course_lesson_id)" description:"课时所属课时ID"`
	Status         uint8     `orm:"column(status)" description:"学习状态0 学习中, 1 学完"`
	FinishedTime   time.Time `orm:"column(finished_time);type(datetime);null" description:"学习完成时间"`
	LearnNum       int64     `orm:"column(learn_num)" description:"次数"`
	CreateTime     time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间"`
	UpdateTime     time.Time `orm:"column(update_time);auto_now;type(datetime)" description:"更新时间"`
}

func (t *CourseLessonLearn) TableName() string {
	return "course_lesson_learn"
}

func init() {
	orm.RegisterModel(new(CourseLessonLearn))
}

// AddCourseLessonLearn insert a new CourseLessonLearn into database and returns
// last inserted Id on success.
func AddCourseLessonLearn(m *CourseLessonLearn) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCourseLessonLearnById retrieves CourseLessonLearn by Id. Returns error if
// Id doesn't exist
func GetCourseLessonLearnById(id int64) (v *CourseLessonLearn, err error) {
	o := orm.NewOrm()
	v = &CourseLessonLearn{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCourseLessonLearn retrieves all CourseLessonLearn matches certain condition. Returns empty list if
// no records exist
func GetAllCourseLessonLearn(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CourseLessonLearn))
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

	var l []CourseLessonLearn
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

// UpdateCourseLessonLearn updates CourseLessonLearn by Id and returns error if
// the record to be updated doesn't exist
func UpdateCourseLessonLearnById(m *CourseLessonLearn) (err error) {
	o := orm.NewOrm()
	v := CourseLessonLearn{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCourseLessonLearn deletes CourseLessonLearn by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCourseLessonLearn(id int64) (err error) {
	o := orm.NewOrm()
	v := CourseLessonLearn{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CourseLessonLearn{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
