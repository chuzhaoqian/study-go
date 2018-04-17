package main

import (
	"github.com/astaxie/beego/migration"
)

// Generation command
// 生成命令
// bee generate migration user -driver=mysql -fields="id:auto,register_type:bool,email:string:50,mobile:string:12,password:string:32,salt:string:32,locked:bool,ip:string:15,create_time:datetime,update_time:datetime"
// 会报错
// [E] execute error: Error 1075: Incorrect table definition; there can be only one auto column and it must be defined as a key

// DO NOT MODIFY
type User_20180305_182251 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20180305_182251{}
	m.Created = "20180305_182251"

	migration.Register("User_20180305_182251", m)
}

// Run the migrations
func (m *User_20180305_182251) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `user` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`register_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '注册方式 0网站 1Android 2IOS'," +
		"`email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱'," +
		"`mobile` varchar(12) NOT NULL DEFAULT '' COMMENT '手机'," +
		"`password` varchar(32) NOT NULL COMMENT '密码'," +
		"`salt` varchar(32) NOT NULL COMMENT '密码SALT'," +
		"`locked` tinyint(1) NOT NULL DEFAULT '0' COMMENT '用户封号状态 0默认，1封号'," +
		"`ip` varchar(15) NOT NULL DEFAULT '' COMMENT '注册ip'," +
		"`create_time` datetime NOT NULL COMMENT '注册时间'," +
		"`update_time` datetime NOT NULL COMMENT '修改时间'," +
		"PRIMARY KEY (`id`)" +
		"KEY `IDX_CTIME` (`create_time`)" +
		"KEY `IDX_UTIME` (`update_time`)" +
		")ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户信息主表'")
}

// Reverse the migrations
func (m *User_20180305_182251) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user`")
}
