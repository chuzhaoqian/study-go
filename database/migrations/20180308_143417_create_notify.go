package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Notify_20180308_143417 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Notify_20180308_143417{}
	m.Created = "20180308_143417"

	migration.Register("Notify_20180308_143417", m)
}

// Run the migrations
func (m *Notify_20180308_143417) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `notify` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`plat` tinyint(1) DEFAULT '0' COMMENT '推送接受平台 0全部 1网站 1Android 2IOS'," +
		"`title` varchar(255) DEFAULT '' COMMENT '推送标题'," +
		"`content` text COMMENT '推送内容'," +
		"`is_cron` smallint(1) DEFAULT NULL COMMENT '是否是定时推送 0=>否 1=>是'," +
		"`receive_role` smallint(1) DEFAULT '0' COMMENT '接受角色 0全部 1学生 2老师'," +
		"`course_ids` varchar(255) DEFAULT '' COMMENT '课程id'," +
		"`category_ids` varchar(255) DEFAULT '' COMMENT '分类id'," +
		"`notify_time` datetime COMMENT '推送时间'," +
		"`notify_device` int(11) DEFAULT '0' COMMENT '通知设备数'," +
		"`status` tinyint(1) DEFAULT NULL COMMENT '推送状态 0未发送 1成功 2=>失败 3=>取消'," +
		"`type` tinyint(1) DEFAULT '0' COMMENT '消息类型 0 后台创建 1直播通知'," +
		"`user_id` int(11)  DEFAULT NULL COMMENT '创建者'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='信息表'")
}

// Reverse the migrations
func (m *Notify_20180308_143417) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `notify`")
}
