package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Setting_20180308_140640 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Setting_20180308_140640{}
	m.Created = "20180308_140640"

	migration.Register("Setting_20180308_140640", m)
}

// Run the migrations
func (m *Setting_20180308_140640) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `setting` (" +
		"`id` int(10) NOT NULL AUTO_INCREMENT COMMENT '系统设置ID'," +
		"`name` varchar(64) NOT NULL DEFAULT '' COMMENT '系统设置名'," +
		"`value` text COMMENT '系统设置值'," +
		"`user_id` int(11) DEFAULT NULL COMMENT '创建者'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)," +
		"KEY `IDX_NAME` (`name`) USING BTREE" +
		"KEY `IDX_CTIME` (`create_time`)" +
		"KEY `IDX_UTIME` (`update_time`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='站点信息设置表'")
}

// Reverse the migrations
func (m *Setting_20180308_140640) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `setting`")
}
