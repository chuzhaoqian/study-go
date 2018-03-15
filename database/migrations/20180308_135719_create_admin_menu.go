package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AdminMenu_20180308_135719 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AdminMenu_20180308_135719{}
	m.Created = "20180308_135719"

	migration.Register("AdminMenu_20180308_135719", m)
}

// Run the migrations
func (m *AdminMenu_20180308_135719) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `admin_menu` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`pid` int(11) NOT NULL DEFAULT '0'  COMMENT '父级id'," +
		"`model` varchar(30) NOT NULL DEFAULT '' COMMENT '控制器'," +
		"`status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态，1显示，0不显示'," +
		"`name` varchar(50) NOT NULL COMMENT '菜单名称'," +
		"`permission_name` varchar(50) NOT NULL COMMENT '权限标识'," +
		"`icon` varchar(50) DEFAULT NULL COMMENT '菜单图标'," +
		"`remark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注'," +
		"`sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序ID'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)," +
		"KEY `status` (`status`) USING BTREE," +
		"KEY `pid` (`pid`) USING BTREE" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台菜单表'")
}

// Reverse the migrations
func (m *AdminMenu_20180308_135719) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `admin_menu`")
}
