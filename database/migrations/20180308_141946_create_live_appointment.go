package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type LiveAppointment_20180308_141946 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &LiveAppointment_20180308_141946{}
	m.Created = "20180308_141946"

	migration.Register("LiveAppointment_20180308_141946", m)
}

// Run the migrations
func (m *LiveAppointment_20180308_141946) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `live_appointment` (" +
		"`id` int(11)  NOT NULL AUTO_INCREMENT," +
		"`live_id` int(11)  NOT NULL DEFAULT '0' COMMENT '直播id'," +
		"`user_id` int(11)  NOT NULL DEFAULT '0' COMMENT '用户id'," +
		"`send_time` datetime NULL COMMENT '发送时间'," +
		"`status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态 0未发送  1已发送'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)," +
		"KEY `IDX_USER` (`user_id`) USING BTREE," +
		"KEY `IDX_LIVE_USER` (`live_id`,`user_id`) USING BTREE" +
		"KEY `IDX_CTIME` (`create_time`)" +
		"KEY `IDX_UTIME` (`update_time`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='直播预约'")
}

// Reverse the migrations
func (m *LiveAppointment_20180308_141946) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `live_appointment`")
}
