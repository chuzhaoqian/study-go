package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CourseCategory_20180308_134039 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CourseCategory_20180308_134039{}
	m.Created = "20180308_134039"

	migration.Register("CourseCategory_20180308_134039", m)
}

// Run the migrations
func (m *CourseCategory_20180308_134039) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `course_category` (" +
		"`id` int(11)  NOT NULL AUTO_INCREMENT COMMENT '分类ID'," +
		"`name` varchar(100) DEFAULT '' COMMENT '分类名称'," +
		"`weight` int(11)  NOT NULL DEFAULT '0' COMMENT '分类权重'," +
		"`group_id` int(11)  NOT NULL COMMENT '分类组ID'," +
		"`parent_id` int(11)  NOT NULL DEFAULT '0' COMMENT '父分类ID'," +
		"`level` tinyint(1)  DEFAULT '0' COMMENT '层级：顶级=0 次级=1 重级=2 依次类推'," +
		"`is_node` tinyint(1)  DEFAULT '0' COMMENT '是否叶结点: 0=>非叶结点  1=>叶结点'," +
		"`is_recommend` tinyint(1)  DEFAULT '0' COMMENT ' 是否是推荐分类 1. 推荐 0. 不推荐'," +
		"`recommend_name` varchar(36) DEFAULT '' COMMENT '推荐名字'," +
		"`is_default` tinyint(1) DEFAULT '0' COMMENT '是否为默认分类 0：不是 1：是'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '最后更新时间'," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='分类表'")
}

// Reverse the migrations
func (m *CourseCategory_20180308_134039) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `course_category`")
}
