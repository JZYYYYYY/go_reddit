package logic

import (
	"hellogo/web_app/dao/mysql"
	"hellogo/web_app/models"
)

func GetCommunityList() ([]*models.Community, error) {
	//查数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
