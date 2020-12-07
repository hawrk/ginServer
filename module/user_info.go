package module

import (
	"ginServer/dao"
	"log"
)

// 以下为一个表名的demo，当需要操作多个表时， 一个表必须单独一个文件出来

//CREATE TABLE `t_user_info` (
//`id` int unsigned NOT NULL AUTO_INCREMENT,
//`name` varchar(64) DEFAULT NULL,
//`age` int DEFAULT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

// 表字段结构体定义
type UserInfo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}


func (u *UserInfo) Insert()(int64,error) {
	res, err := dao.SqlDB.Exec("insert into t_user_info(name, age) " +
		"values (?,?",u.Name, u.Age)
	if err != nil  {
		log.Printf("exec sql fail:%v", err)
	}
	id, err := res.LastInsertId()
	log.Printf("insert success , insert id:%d", id)
	return id, err
}

func (u *UserInfo) Select() (info []*UserInfo, err error) {
	info = make([]*UserInfo,0)
	rows, err := dao.SqlDB.Query("select name, age from t_user_info")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var data UserInfo
		rows.Scan(&data.Name, &data.Age)
		info = append(info, &data)
	}
	log.Printf("out:%v", info)
	return
}