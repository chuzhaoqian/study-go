package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type IniUser_20180306_084623 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &IniUser_20180306_084623{}
	m.Created = "20180306_084623"

	migration.Register("IniUser_20180306_084623", m)
}

// Run the migrations
func (m *IniUser_20180306_084623) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("insert into `user` ( `register_type`, `email`, `mobile`, `password`, `salt`, `locked`, `ip`, `create_time`, `update_time`) values ( 0, '1028290810@qq.com', 111111111111, '4999f135df7be44a43a1afe60c6027c4', '2279046668356600329', 0, '127.0.0.1', '2018-03-06 08:00:00', '2018-03-06 08:00:00')")
	m.SQL("insert into `user_profile` ( `id`, `gender`, `nickname`, `avatar`,  `description`, `create_time`, `update_time`) values ( 1, 1, 'ZhaoQian',  '', '', '2018-03-06 08:00:00', '2018-03-06 08:00:00')")
}

// Reverse the migrations
func (m *IniUser_20180306_084623) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("truncate table `user`")
	m.SQL("truncate table `user_profile`")
}
