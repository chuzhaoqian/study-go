package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UserNotify_20180308_143318 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserNotify_20180308_143318{}
	m.Created = "20180308_143318"

	migration.Register("UserNotify_20180308_143318", m)
}

// Run the migrations
func (m *UserNotify_20180308_143318) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `user_notify` (" +
		"`id` int(11)  NOT NULL AUTO_INCREMENT," +
		"`user_id` int(11)  DEFAULT NULL COMMENT '用户id'," +
		"`message_id` int(11)  DEFAULT NULL COMMENT '消息id'," +
		"`status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0未读 1已读'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)" +
		"KEY `IDX_CTIME` (`create_time`)" +
		"KEY `IDX_UTIME` (`update_time`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户信息关联表'")
}

// Reverse the migrations
func (m *UserNotify_20180308_143318) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user_notify`")
}
