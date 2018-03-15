package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Order struct {
	Id            int64     `orm:"column(id);auto" description:"订单ID"`
	Sn            string    `orm:"column(sn);size(32)" description:"订单编号"`
	PaymentType   int8      `orm:"column(payment_type)" description:"支付方式: 1支付宝 2微信 0其他"`
	PaymentSn     string    `orm:"column(payment_sn);size(32);null" description:"支付流水号"`
	PaymentSource int8      `orm:"column(payment_source);null" description:"支付订单来源: 0=Web 1=Android 2=IOS"`
	Status        int8      `orm:"column(status)" description:"订单状态: 0未付款(created) 1已付款(paid) 2已过期(overdue) 3已退款(refunded) 4已取消(cancelled)"`
	Amount        float32   `orm:"column(amount)" description:"订单实付金额"`
	TotalPrice    float32   `orm:"column(total_price)" description:"订单总价"`
	UserId        int64     `orm:"column(user_id)" description:"订单购买人"`
	CourseId      int64     `orm:"column(course_id)" description:"商品id"`
	CreateSource  int8      `orm:"column(create_source)" description:"产生订单来源: 0未知 1Web 2Android 3IOS"`
	PaymentTime   time.Time `orm:"column(payment_time);type(datetime);null" description:"支付时间 完成时间"`
	CreateTime    time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间"`
	UpdateTime    time.Time `orm:"column(update_time);auto_now;type(datetime)" description:"更新时间"`
}

func (t *Order) TableName() string {
	return "order"
}

func init() {
	orm.RegisterModel(new(Order))
}

// AddOrder insert a new Order into database and returns
// last inserted Id on success.
func AddOrder(m *Order) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrderById retrieves Order by Id. Returns error if
// Id doesn't exist
func GetOrderById(id int64) (v *Order, err error) {
	o := orm.NewOrm()
	v = &Order{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrder retrieves all Order matches certain condition. Returns empty list if
// no records exist
func GetAllOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Order))
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

	var l []Order
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

// UpdateOrder updates Order by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrderById(m *Order) (err error) {
	o := orm.NewOrm()
	v := Order{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrder deletes Order by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrder(id int64) (err error) {
	o := orm.NewOrm()
	v := Order{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Order{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
