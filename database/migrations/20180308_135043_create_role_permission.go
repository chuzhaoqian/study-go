package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type RolePermission_20180308_135043 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RolePermission_20180308_135043{}
	m.Created = "20180308_135043"

	migration.Register("RolePermission_20180308_135043", m)
}

// Run the migrations
func (m *RolePermission_20180308_135043) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `role_permission` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色id'," +
		"`permission_id` int(11) NOT NULL DEFAULT '0' COMMENT '权限规则id'," +
		"`permission_name` varchar(255) NOT NULL DEFAULT '' COMMENT '权限规则标识'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)," +
		"KEY `group_id` (`role_id`) USING BTREE," +
		"KEY `permission_id` (`permission_name`) USING BTREE" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色与权限关联表'")
}

// Reverse the migrations
func (m *RolePermission_20180308_135043) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `role_permission`")
}
