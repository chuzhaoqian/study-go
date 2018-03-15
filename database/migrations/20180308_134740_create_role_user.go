package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type RoleUser_20180308_134740 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RoleUser_20180308_134740{}
	m.Created = "20180308_134740"

	migration.Register("RoleUser_20180308_134740", m)
}

// Run the migrations
func (m *RoleUser_20180308_134740) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `role_user` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`role_id` int(11) DEFAULT '0' COMMENT '角色 id'," +
		"`user_id` int(11) DEFAULT '0' COMMENT '用户id'," +
		"`type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '类型 0 普通角色 1超级角色 2特殊'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)," +
		"KEY `group_id` (`role_id`) USING BTREE," +
		"KEY `user_id` (`user_id`) USING BTREE" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色与管理员关联表'")
}

// Reverse the migrations
func (m *RoleUser_20180308_134740) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `role_user`")
}