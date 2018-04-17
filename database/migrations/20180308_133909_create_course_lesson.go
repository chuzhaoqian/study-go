package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CourseLesson_20180308_133909 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CourseLesson_20180308_133909{}
	m.Created = "20180308_133909"

	migration.Register("CourseLesson_20180308_133909", m)
}

// Run the migrations
func (m *CourseLesson_20180308_133909) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `course_lesson` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '课时ID'," +
		"`course_id` int(11) NOT NULL COMMENT '课程id'," +
		"`sort` int(11) NOT NULL COMMENT '总排序序号'," +
		"`status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '课时状态0 未发布, 1发布'," +
		"`title` varchar(100) NOT NULL COMMENT '课时标题'," +
		"`description` varchar(200) NOT NULL DEFAULT '' COMMENT '课时摘要'," +
		"`type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '课时类型 0章 1视频课时 2文档课时'," +
		"`user_id` int(11) DEFAULT NULL COMMENT '创建者'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '最后更新时间'," +
		"PRIMARY KEY (`id`)" +
		"KEY `IDX_CTIME` (`create_time`)" +
		"KEY `IDX_UTIME` (`update_time`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='课程课时内容表'")
}

// Reverse the migrations
func (m *CourseLesson_20180308_133909) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `course_lesson`")
}
