package service

import (
	"aws-case/db"
	"aws-case/entity"
	"aws-case/log"
	"fmt"
	"mime/multipart"
)

type AccountService struct {
	accountRepository db.AccountRepository
	s3Repo            db.S3Repo
}

func (this *AccountService) Register(username string, password string, fileName string, file *multipart.File) (*entity.Account, error) {
	account, err := this.accountRepository.FindByUserName(username)
	if err != nil {
		err := fmt.Errorf("Register Exception:%v", err)
		log.Error("Register Exception", err)
		return nil, err
	}

	if account != nil {
		err := fmt.Errorf("username:%v already exists", username)
		log.Error("register error: ", err)
		return nil, err
	}
	url := this.s3Repo.PutFile(username, fileName, *file)
	return this.accountRepository.Save(&entity.Account{
		UserName: username,
		Password: password,
		Avatar:   url,
	})
}

// Login 登陆用户
func (this *AccountService) Login(username string, password string) (*entity.Account, error) {
	account, err := this.accountRepository.FindByUserName(username)
	if err != nil {
		err = fmt.Errorf("Login Exception:%v", err)
		log.Error("获取账户信息错误", err)
		return nil, err
	}
	if account == nil {
		err = fmt.Errorf("Incorrect username or password")
		log.Error("账户不存在")
		return nil, err
	}
	pwd := account.Password
	if pwd != password {
		err = fmt.Errorf("Incorrect username or password")
		log.Error("用户名密码不正确")
		return nil, err
	}
	return account, nil
}

// UpdateAvatar 更新用户头像
func (this *AccountService) UpdateAvatar(username string, fileName string, file *multipart.File) (*entity.Account, error) {
	account, err := this.accountRepository.FindByUserName(username)
	if account == nil || err != nil {
		log.Error("username does not exist:%v", username)
		return nil, err
	}

	url := this.s3Repo.PutFile(username, fileName, *file)

	err = this.accountRepository.UpdateAvatar(username, url)
	if err != nil {
		log.Error("Modification failed:%v", err)
		return nil, err
	}
	accountDb, _ := this.accountRepository.FindByUserName(username)
	return accountDb, nil
}

func NewAccountService() AccountService {
	return AccountService{
		accountRepository: db.NewDynamoDBRepository(),
		s3Repo:            db.NewS3Repo(),
	}
}
