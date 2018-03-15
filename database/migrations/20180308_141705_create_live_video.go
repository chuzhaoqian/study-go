package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type LiveVideo_20180308_141705 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &LiveVideo_20180308_141705{}
	m.Created = "20180308_141705"

	migration.Register("LiveVideo_20180308_141705", m)
}

// Run the migrations
func (m *LiveVideo_20180308_141705) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `live_video` (" +
		"`id` int(11)  NOT NULL AUTO_INCREMENT COMMENT '录播视频id'," +
		"`title` varchar(255) NOT NULL DEFAULT '' COMMENT '视频名称'," +
		"`live_id` int(11) NOT NULL COMMENT '直播id'," +
		"`video_id` varchar(100) NOT NULL COMMENT '视频id'," +
		"`video_detail` varchar(300) NOT NULL DEFAULT '' COMMENT '视频详情'," +
		"`video_size` varchar(11) DEFAULT '0' COMMENT '视频时长'," +
		"`sort` int(11) DEFAULT NULL DEFAULT '0' COMMENT '排序'," +
		"`size` int(11) NOT NULL DEFAULT '0' COMMENT '大小'," +
		"`user_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者id'," +
		"`start_time` int(10) NOT NULL DEFAULT '0' COMMENT '开始录制时间'," +
		"`end_time` int(10) NOT NULL DEFAULT '0' COMMENT '录制结束时间'," +
		"`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'," +
		"`update_time` datetime NOT NULL COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)," +
		"KEY `live_video` (`live_id`,`video_id`) USING BTREE," +
		"KEY `live` (`live_id`) USING BTREE" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='直播视频录像'")
}

// Reverse the migrations
func (m *LiveVideo_20180308_141705) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `live_video`")
}
