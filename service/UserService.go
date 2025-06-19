package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"novel-go/config"
	"novel-go/constants"
	"novel-go/model/pojo"
	"novel-go/model/req"
	"novel-go/model/resp"
	"novel-go/utils"
	"time"
)

// UserService 用户服务接口
type UserService interface {
	Register(req *req.UserRegisterReqDto) (*resp.UserRegisterRespDto, error)

	Login(r *req.UserLoginReqDto) (*resp.UserLoginRespDto, error)
}

// UserServiceImpl 用户服务接口实现
type UserServiceImpl struct {
}

// NewUserServiceImpl 构造函数
func NewUserServiceImpl() UserService {
	return &UserServiceImpl{}
}

// Register 用户注册实现
func (s *UserServiceImpl) Register(req *req.UserRegisterReqDto) (*resp.UserRegisterRespDto, error) {
	redisKey := constants.RedisKeyCaptcha + req.SessionId

	// 1. 校验验证码
	dbCode, err := config.RedisClient.Get(config.Ctx, redisKey).Result()
	if err != nil {
		return nil, errors.New("验证码不存在或已过期")
	}
	if dbCode != req.VelCode {
		return nil, errors.New("验证码错误")
	}

	// 校验成功后再删除 Redis 验证码，避免误删
	defer config.RedisClient.Del(config.Ctx, redisKey)

	// 2. 检查用户名是否已存在
	var count int64
	if err := config.DB.Model(&pojo.UserInfo{}).
		Where("username = ?", req.Username).
		Count(&count).Error; err != nil {
		return nil, fmt.Errorf("数据库查询失败: %w", err)
	}
	if count > 0 {
		return nil, errors.New("用户名已存在")
	}

	// 3. 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("密码加密失败: %w", err)
	}

	// 4. 创建用户
	user := pojo.UserInfo{
		Username:   req.Username,
		Password:   string(hashedPassword),
		NickName:   req.Username,
		Status:     0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("用户创建失败: %w", err)
	}

	// 5. 构造返回信息
	respDto := &resp.UserRegisterRespDto{
		UID:      int64(user.ID),
		NickName: user.NickName,
		Avatar:   user.UserPhoto,
	}

	return respDto, nil
}

// Login 用户登录实现
func (s *UserServiceImpl) Login(r *req.UserLoginReqDto) (*resp.UserLoginRespDto, error) {
	username := r.Username
	password := r.Password

	// 1. 查询用户信息
	var user pojo.UserInfo
	if err := config.DB.Where("username = ?", username).Limit(1).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, fmt.Errorf("数据库查询失败: %w", err)
	}

	// 2. 校验密码（bcrypt）
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("密码错误")
	}

	// 3. 生成 JWT Token
	token, err := utils.GenerateToken(user.ID, constants.JwtSecret)
	if err != nil {
		return nil, fmt.Errorf("Token 生成失败: %w", err)
	}

	// 4. 构造返回对象
	respDto := &resp.UserLoginRespDto{
		UID:      user.ID,
		NickName: user.NickName,
		Avatar:   user.UserPhoto,
		Token:    token,
	}

	return respDto, nil
}
