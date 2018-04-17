package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CourseLessonLearn_20180308_134359 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CourseLessonLearn_20180308_134359{}
	m.Created = "20180308_134359"

	migration.Register("CourseLessonLearn_20180308_134359", m)
}

// Run the migrations
func (m *CourseLessonLearn_20180308_134359) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `course_lesson_learn` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '课时学习记录ID'," +
		"`user_id` int(11) NOT NULL COMMENT '学员ID'," +
		"`course_id` int(11) NOT NULL COMMENT '课程id'," +
		"`course_lesson_id` int(10) NOT NULL DEFAULT '0' COMMENT '课时所属课时ID'," +
		"`status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '学习状态0 学习中, 1 学完'," +
		"`finished_time` datetime COMMENT '学习完成时间'," +
		"`learn_num` int(11) NOT NULL DEFAULT '0' COMMENT '次数'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)," +
		"KEY `IDX_USER_COURSE` (`user_id`,`course_id`) USING BTREE," +
		"KEY `IDX_COURSE_LESSON` (`course_id`,`course_lesson_id`) USING BTREE" +
		"KEY `IDX_CTIME` (`create_time`)" +
		"KEY `IDX_UTIME` (`update_time`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='课程学习表'")
}

// Reverse the migrations
func (m *CourseLessonLearn_20180308_134359) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `course_lesson_learn`")
}
