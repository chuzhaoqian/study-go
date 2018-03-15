package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Role_20180308_134538 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Role_20180308_134538{}
	m.Created = "20180308_134538"

	migration.Register("Role_20180308_134538", m)
}

// Run the migrations
func (m *Role_20180308_134538) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `role` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`name` varchar(20) NOT NULL COMMENT '角色名称'," +
		"`pid` int(11) DEFAULT '0' COMMENT '父角色ID'," +
		"`type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '类型 0 普通角色 1超级角色 2特殊'," +
		"`award` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0 不可授予 可授予 1'," +
		"`status` tinyint(1) DEFAULT '0' COMMENT '状态: 0=有效 1=禁用'," +
		"`remark` varchar(255) DEFAULT NULL COMMENT '备注'," +
		"`sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序字段'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)," +
		"KEY `parentId` (`pid`) USING BTREE," +
		"KEY `status` (`status`) USING BTREE" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表'")
}

// Reverse the migrations
func (m *Role_20180308_134538) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `role`")
}
