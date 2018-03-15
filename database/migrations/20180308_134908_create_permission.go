package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Permission_20180308_134908 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Permission_20180308_134908{}
	m.Created = "20180308_134908"

	migration.Register("Permission_20180308_134908", m)
}

// Run the migrations
func (m *Permission_20180308_134908) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `permission` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '规则id,自增主键'," +
		"`module` varchar(20) NOT NULL COMMENT '规则所属module'," +
		"`type` varchar(30) NOT NULL DEFAULT '1' COMMENT '权限规则分类，请加应用前缀,如admin_'," +
		"`name` varchar(30) NOT NULL DEFAULT '' COMMENT '规则唯一英文标识,全小写'," +
		"`param` varchar(255) DEFAULT NULL COMMENT '额外url参数'," +
		"`title` varchar(20) NOT NULL DEFAULT '' COMMENT '规则中文描述'," +
		"`status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否有效(0:无效,1:有效)'," +
		"`condition` varchar(300) NOT NULL DEFAULT '' COMMENT '规则附加条件(为以后的页面视图权限做准备)'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)," +
		"KEY `module` (`module`,`status`,`type`) USING BTREE" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限规则表'")
}

// Reverse the migrations
func (m *Permission_20180308_134908) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `permission`")
}
