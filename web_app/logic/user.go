package logic

import (
	"hellogo/web_app/dao/mysql"
	"hellogo/web_app/models"
	"hellogo/web_app/pkg/jwt"
	"hellogo/web_app/pkg/snowflake"
)

//存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	//1.判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		//数据库查询出错
		return err
	}
	//2.生成UID
	userID := snowflake.GenID()
	//构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//3.保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	//生成jwt
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
