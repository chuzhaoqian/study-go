package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Live_20180308_140844 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Live_20180308_140844{}
	m.Created = "20180308_140844"

	migration.Register("Live_20180308_140844", m)
}

// Run the migrations
func (m *Live_20180308_140844) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `live` (" +
		"`id` int(11)  NOT NULL AUTO_INCREMENT COMMENT '直播id'," +
		"`live_title` varchar(40) NOT NULL COMMENT '直播主题'," +
		"`start_time` datetime NULL COMMENT '直播开始时间'," +
		"`end_time` datetime NULL COMMENT '直播结束时间'," +
		"`real_start_time` int(11) NOT NULL DEFAULT '0' COMMENT '实际直播开始时间'," +
		"`real_end_time` int(11) NOT NULL DEFAULT '0' COMMENT '实际直播结束时间'," +
		"`description` text COMMENT '直播介绍'," +
		"`picture` varchar(255) DEFAULT '' COMMENT '直播封面图片'," +
		"`live_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '直播状态 0=>未开始 1=>直播中 2=>已结束 3=>已关闭 '," +
		"`is_designated_courses` tinyint(1) NOT NULL DEFAULT '0' COMMENT '发送对象(1=>指定课程,0=>全网开放)'," +
		"`is_notification` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否发送通知(1=>是,0=>否)'," +
		"`is_playback` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否回放(1=>是,0=>否)'," +
		"`is_closed` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否关闭直播（0=>未关闭,1=>已关闭）'," +
		"`playing` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0 未播放 1正播放'," +
		"`speaker_id` int(11)  NOT NULL DEFAULT '2' COMMENT '主讲人id'," +
		"`user_id` int(11)  DEFAULT NULL COMMENT '创建者'," +
		"`create_time` datetime NOT NULL COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)" +
		"KEY `IDX_CTIME` (`create_time`)" +
		"KEY `IDX_UTIME` (`update_time`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='直播表'")
}

// Reverse the migrations
func (m *Live_20180308_140844) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `live`")
}
