package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CourseMember_20180308_133648 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CourseMember_20180308_133648{}
	m.Created = "20180308_133648"

	migration.Register("CourseMember_20180308_133648", m)
}

// Run the migrations
func (m *CourseMember_20180308_133648) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `course_member` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`course_id` int(11) NOT NULL COMMENT '课程id'," +
		"`user_id` int(11) NOT NULL COMMENT '用户id'," +
		"`join_type` tinyint(1) DEFAULT '0' COMMENT '加入方式  ： 0、购买加入  1 、后台导入'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '最后更新时间'," +
		"PRIMARY KEY (`id`)," +
		"KEY `IDX_COURSE` (`course_id`) USING BTREE" +
		"KEY `IDX_CTIME` (`create_time`)" +
		"KEY `IDX_UTIME` (`update_time`)" +
		")  ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='课程授课班学生表'")
}

// Reverse the migrations
func (m *CourseMember_20180308_133648) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `course_member`")
}
