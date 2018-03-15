package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type LiveCourse_20180308_141250 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &LiveCourse_20180308_141250{}
	m.Created = "20180308_141250"

	migration.Register("LiveCourse_20180308_141250", m)
}

// Run the migrations
func (m *LiveCourse_20180308_141250) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `live_course` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`live_id` int(11) NOT NULL COMMENT '直播id'," +
		"`course_id` int(11) NOT NULL COMMENT '课程id'," +
		"`user_id` int(11) DEFAULT NULL COMMENT '创建者'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='直播关联课程'")
}

// Reverse the migrations
func (m *LiveCourse_20180308_141250) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `live_course`")
}
