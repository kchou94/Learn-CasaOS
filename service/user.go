package service

import "Learn-CasaOS/pkg/config"

type UserService interface {
	SetUser(username, pwd, token, email, desc string) error
}

type user struct{}

func (c *user) SetUser(username, pwd, token, email, desc string) error {
	if len(username) > 0 {
		config.Cfg.Section("user").Key("UserName").SetValue(username)
		config.UserInfo.UserName = username
	}
	if len(pwd) > 0 {
		config.Cfg.Section("user").Key("PWD").SetValue(pwd)
		config.UserInfo.PWD = pwd
	}
	if len(token) > 0 {
		config.Cfg.Section("user").Key("Token").SetValue(token)
		config.UserInfo.Token = token
	}
	if len(email) > 0 {
		config.Cfg.Section("user").Key("Email").SetValue(email)
		config.UserInfo.Email = email
	}
	if len(desc) > 0 {
		config.Cfg.Section("user").Key("Description").SetValue(desc)
		config.UserInfo.Description = desc
	}
	config.Cfg.SaveTo("conf/conf.ini")
	return nil
}

func NewUserService() UserService {
	return &user{}
}
