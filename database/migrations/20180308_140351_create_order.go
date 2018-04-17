package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Order_20180308_140351 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Order_20180308_140351{}
	m.Created = "20180308_140351"

	migration.Register("Order_20180308_140351", m)
}

// Run the migrations
func (m *Order_20180308_140351) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `order` (" +
		"`id` int(11)  NOT NULL AUTO_INCREMENT COMMENT '订单ID'," +
		"`sn` varchar(32) NOT NULL COMMENT '订单编号'," +
		"`payment_type` tinyint(1)  NOT NULL DEFAULT '0' COMMENT '支付方式: 1支付宝 2微信 0其他'," +
		"`payment_sn` varchar(32) DEFAULT '' COMMENT '支付流水号'," +
		"`payment_source` tinyint(1)  DEFAULT '0' COMMENT '支付订单来源: 0=Web 1=Android 2=IOS'," +
		"`status` tinyint(1)  NOT NULL DEFAULT '0' COMMENT '订单状态: 0未付款(created) 1已付款(paid) 2已过期(overdue) 3已退款(refunded) 4已取消(cancelled)'," +
		"`amount` float(10,2)  NOT NULL DEFAULT '0.00' COMMENT '订单实付金额'," +
		"`total_price` float(10,2)  NOT NULL DEFAULT '0.00' COMMENT '订单总价'," +
		"`user_id` int(11)  NOT NULL COMMENT '订单购买人'," +
		"`course_id` int(11)  NOT NULL DEFAULT '0' COMMENT '商品id'," +
		"`create_source` tinyint(1)  NOT NULL DEFAULT '0' COMMENT '产生订单来源: 0未知 1Web 2Android 3IOS'," +
		"`payment_time` datetime NULL COMMENT '支付时间 完成时间'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)," +
		"KEY `sn` (`sn`) USING BTREE," +
		"KEY `user_id` (`user_id`) USING BTREE" +
		"KEY `IDX_CTIME` (`create_time`)" +
		"KEY `IDX_UTIME` (`update_time`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单信息表'")
}

// Reverse the migrations
func (m *Order_20180308_140351) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `order`")
}
