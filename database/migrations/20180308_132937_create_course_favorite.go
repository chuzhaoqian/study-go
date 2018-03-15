package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CourseFavorite_20180308_132937 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CourseFavorite_20180308_132937{}
	m.Created = "20180308_132937"

	migration.Register("CourseFavorite_20180308_132937", m)
}

// Run the migrations
func (m *CourseFavorite_20180308_132937) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `course_favorite` (" +
		"`id` int(11)  NOT NULL AUTO_INCREMENT COMMENT '收藏ID'," +
		"`user_id` int(11)  NOT NULL COMMENT '学员ID'," +
		"`course_id` int(11)  NOT NULL COMMENT '课程id'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '最后更新时间'," +
		"PRIMARY KEY (`id`)," +
		"KEY `user_id` (`user_id`) USING BTREE," +
		"KEY `course_id` (`course_id`) USING BTREE" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='课程收藏表'")
}

// Reverse the migrations
func (m *CourseFavorite_20180308_132937) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `course_favorite`")
}
