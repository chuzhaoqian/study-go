package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CourseCategory struct {
	Id            int64     `orm:"column(id);auto" description:"分类ID"`
	Name          string    `orm:"column(name);size(100);null" description:"分类名称"`
	Weight        int64     `orm:"column(weight)" description:"分类权重"`
	GroupId       int64     `orm:"column(group_id)" description:"分类组ID"`
	ParentId      int64     `orm:"column(parent_id)" description:"父分类ID"`
	Level         int8      `orm:"column(level);null" description:"层级：顶级=0 次级=1 重级=2 依次类推"`
	IsNode        int8      `orm:"column(is_node);null" description:"是否叶结点: 0=>非叶结点  1=>叶结点"`
	IsRecommend   int8      `orm:"column(is_recommend);null" description:" 是否是推荐分类 1. 推荐 0. 不推荐"`
	RecommendName string    `orm:"column(recommend_name);size(36);null" description:"推荐名字"`
	IsDefault     int8      `orm:"column(is_default);null" description:"是否为默认分类 0：不是 1：是"`
	CreateTime    time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间"`
	UpdateTime    time.Time `orm:"column(update_time);auto_now;type(datetime)" description:"更新时间"`
}

func (t *CourseCategory) TableName() string {
	return "course_category"
}

func init() {
	orm.RegisterModel(new(CourseCategory))
}

// AddCourseCategory insert a new CourseCategory into database and returns
// last inserted Id on success.
func AddCourseCategory(m *CourseCategory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCourseCategoryById retrieves CourseCategory by Id. Returns error if
// Id doesn't exist
func GetCourseCategoryById(id int64) (v *CourseCategory, err error) {
	o := orm.NewOrm()
	v = &CourseCategory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCourseCategory retrieves all CourseCategory matches certain condition. Returns empty list if
// no records exist
func GetAllCourseCategory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CourseCategory))
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

	var l []CourseCategory
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

// UpdateCourseCategory updates CourseCategory by Id and returns error if
// the record to be updated doesn't exist
func UpdateCourseCategoryById(m *CourseCategory) (err error) {
	o := orm.NewOrm()
	v := CourseCategory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCourseCategory deletes CourseCategory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCourseCategory(id int64) (err error) {
	o := orm.NewOrm()
	v := CourseCategory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CourseCategory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
