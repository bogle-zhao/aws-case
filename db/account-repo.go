package db

import "aws-case/entity"

type AccountRepository interface {
	// Save 保存用户
	Save(account *entity.Account) (*entity.Account, error)
	// FindByUserName 根据用户名获取账户信息
	FindByUserName(userName string) (*entity.Account, error)
	// UpdateAvatar 更新用户头像
	UpdateAvatar(username string, avatar string) error
}
