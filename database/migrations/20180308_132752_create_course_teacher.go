package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CourseTeacher_20180308_132752 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CourseTeacher_20180308_132752{}
	m.Created = "20180308_132752"

	migration.Register("CourseTeacher_20180308_132752", m)
}

// Run the migrations
func (m *CourseTeacher_20180308_132752) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `course_teacher`(" +
		"`id` int(11)  NOT NULL AUTO_INCREMENT," +
		"`course_id` int(11)  NOT NULL DEFAULT '0' COMMENT '课程id'," +
		"`user_id` int(11)  NOT NULL DEFAULT '0' COMMENT '用户id'," +
		"`type` tinyint(1)  NOT NULL DEFAULT '0' COMMENT '0.讲师  1.管理老师'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '最后更新时间'," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 CHARSET=utf8 COMMENT='课程老师表'")
}

// Reverse the migrations
func (m *CourseTeacher_20180308_132752) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `course_teacher`")
}
