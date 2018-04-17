package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UserProfile_20180306_084325 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserProfile_20180306_084325{}
	m.Created = "20180306_084325"

	migration.Register("UserProfile_20180306_084325", m)
}

// Run the migrations
func (m *UserProfile_20180306_084325) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `user_profile` (" +
		"`id` int(11) NOT NULL COMMENT '用户ID'," +
		"`gender` tinyint(1) DEFAULT '0' COMMENT '性别 ：0女  1男'," +
		"`nickname` varchar(60) DEFAULT '' COMMENT '昵称'," +
		"`avatar` varchar(255) DEFAULT '' COMMENT '头像'," +
		"`description` varchar(500) DEFAULT '' COMMENT '备注'," +
		"`create_time` datetime NOT NULL COMMENT '注册时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)" +
		"KEY `IDX_CTIME` (`create_time`)" +
		"KEY `IDX_UTIME` (`update_time`)" +
		")ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息扩展表'")
}

// Reverse the migrations
func (m *UserProfile_20180306_084325) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user_profile`")
}
