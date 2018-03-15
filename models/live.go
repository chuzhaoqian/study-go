package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Live struct {
	Id                  int64     `orm:"column(id);auto" description:"直播id"`
	LiveTitle           string    `orm:"column(live_title);size(40)" description:"直播主题"`
	StartTime           time.Time `orm:"column(start_time);type(datetime);null" description:"开始时间"`
	EndTime             time.Time `orm:"column(end_time);type(datetime);null" description:"结束时间"`
	RealStartTime       int64     `orm:"column(real_start_time)" description:"实际直播开始时间"`
	RealEndTime         int64     `orm:"column(real_end_time)" description:"实际直播结束时间"`
	Description         string    `orm:"column(description);null" description:"直播介绍"`
	Picture             string    `orm:"column(picture);size(255);null" description:"直播封面图片"`
	LiveStatus          int8      `orm:"column(live_status)" description:"直播状态 0=>未开始 1=>直播中 2=>已结束 3=>已关闭 "`
	IsDesignatedCourses int8      `orm:"column(is_designated_courses)" description:"发送对象(1=>指定课程,0=>全网开放)"`
	IsNotification      int8      `orm:"column(is_notification)" description:"是否发送通知(1=>是,0=>否)"`
	IsPlayback          int8      `orm:"column(is_playback)" description:"是否回放(1=>是,0=>否)"`
	IsClosed            int8      `orm:"column(is_closed)" description:"是否关闭直播（0=>未关闭,1=>已关闭）"`
	Playing             int8      `orm:"column(playing)" description:"0 未播放 1正播放"`
	SpeakerId           int64     `orm:"column(speaker_id)" description:"主讲人id"`
	UserId              int64     `orm:"column(user_id);null" description:"创建者"`
	CreateTime          time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间"`
	UpdateTime          time.Time `orm:"column(update_time);auto_now;type(datetime)" description:"更新时间"`
}

func (t *Live) TableName() string {
	return "live"
}

func init() {
	orm.RegisterModel(new(Live))
}

// AddLive insert a new Live into database and returns
// last inserted Id on success.
func AddLive(m *Live) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLiveById retrieves Live by Id. Returns error if
// Id doesn't exist
func GetLiveById(id int64) (v *Live, err error) {
	o := orm.NewOrm()
	v = &Live{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLive retrieves all Live matches certain condition. Returns empty list if
// no records exist
func GetAllLive(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Live))
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

	var l []Live
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

// UpdateLive updates Live by Id and returns error if
// the record to be updated doesn't exist
func UpdateLiveById(m *Live) (err error) {
	o := orm.NewOrm()
	v := Live{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLive deletes Live by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLive(id int64) (err error) {
	o := orm.NewOrm()
	v := Live{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Live{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
