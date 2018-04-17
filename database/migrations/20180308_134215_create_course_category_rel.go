package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CourseCategoryRel_20180308_134215 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CourseCategoryRel_20180308_134215{}
	m.Created = "20180308_134215"

	migration.Register("CourseCategoryRel_20180308_134215", m)
}

// Run the migrations
func (m *CourseCategoryRel_20180308_134215) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `course_category_rel` (" +
		"`id` int(10) NOT NULL AUTO_INCREMENT," +
		"`course_id` int(11) DEFAULT NULL COMMENT '课程id'," +
		"`category_id` int(11) DEFAULT NULL COMMENT '课程分类id'," +
		"`user_id` int(11) DEFAULT NULL COMMENT '创建者'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '最后更新时间'," +
		"PRIMARY KEY (`id`)," +
		"KEY `IDX_COURSE` (`course_id`) USING BTREE" +
		"KEY `IDX_CTIME` (`create_time`)" +
		"KEY `IDX_UTIME` (`update_time`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='课程分类关联表'")
}

// Reverse the migrations
func (m *CourseCategoryRel_20180308_134215) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `course_category_rel`")
}
