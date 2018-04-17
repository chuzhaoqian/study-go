package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Navigation_20180308_140008 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Navigation_20180308_140008{}
	m.Created = "20180308_140008"

	migration.Register("Navigation_20180308_140008", m)
}

// Run the migrations
func (m *Navigation_20180308_140008) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `navigation` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '导航ID'," +
		"`system_name` varchar(30) NOT NULL DEFAULT '' COMMENT '系统默认名称，不可修改'," +
		"`name` varchar(30) NOT NULL DEFAULT '' COMMENT '导航名称，可以自定义'," +
		"`url` varchar(255) NOT NULL DEFAULT '' COMMENT '链接地址'," +
		"`sort` int(11) NOT NULL DEFAULT '0' COMMENT '显示顺序'," +
		"`parent_id` int(10) NOT NULL DEFAULT '0' COMMENT '父导航ID'," +
		"`type` varchar(30) NOT NULL DEFAULT '0' COMMENT '类型'," +
		"`is_open` tinyint(1) NOT NULL DEFAULT '0' COMMENT '默认1开启 0关闭'," +
		"`is_new_window` tinyint(1) NOT NULL DEFAULT '0' COMMENT '默认为1另开窗口'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)" +
		"KEY `IDX_CTIME` (`create_time`)" +
		"KEY `IDX_UTIME` (`update_time`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='导航数据表'")
}

// Reverse the migrations
func (m *Navigation_20180308_140008) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `navigation`")
}
