package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Course_20180306_123556 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Course_20180306_123556{}
	m.Created = "20180306_123556"

	migration.Register("Course_20180306_123556", m)
}

// Run the migrations
func (m *Course_20180306_123556) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `course` (" +
		"`id` int(11)  NOT NULL AUTO_INCREMENT COMMENT '课程ID'," +
		"`title` varchar(100) NOT NULL COMMENT '课程标题'," +
		"`status` tinyint(1) DEFAULT '1' COMMENT '建课状态，0 未发布,1 已发布,2 已关闭'," +
		"`price` float(10,2) NOT NULL DEFAULT '0.00' COMMENT '课程价格'," +
		"`original_cost` float(10,2)  DEFAULT '0.00' COMMENT '原价'," +
		"`is_free` tinyint(4)  DEFAULT '0' COMMENT '是否是收费课，1免费，0收费'," +
		"`small_pic` varchar(255) NOT NULL DEFAULT '' COMMENT '小图'," +
		"`pic` varchar(255) NOT NULL DEFAULT '' COMMENT '中图'," +
		"`large_pic` varchar(255) NOT NULL DEFAULT '' COMMENT '大图'," +
		"`description` text COMMENT '说明'," +
		"`degree` tinyint(1) NOT NULL DEFAULT '1' COMMENT '难易程度  1 简单 2一般 3 困难'," +
		"`grade` float(4,2) DEFAULT '4.00' COMMENT '课程评级'," +
		"`student_num` int(10)  NOT NULL DEFAULT '0' COMMENT '学员数'," +
		"`lesson_num` int(10)  NOT NULL DEFAULT '0' COMMENT '课时数'," +
		"`preview_power` tinyint(1) NOT NULL DEFAULT '0' COMMENT '试读权限 0关闭试读  1试读第一章 2试读前2章 3试读前3章'," +
		"`type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '课程类型/课程模板  0点播课 1直播课'," +
		"`user_id` int(11)  DEFAULT NULL COMMENT '创建者'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)" +
		")ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='课程表'")
}

// Reverse the migrations
func (m *Course_20180306_123556) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `course`")
}
