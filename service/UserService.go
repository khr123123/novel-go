package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"novel-go/config"
	"novel-go/constants"
	"novel-go/model/pojo"
	"novel-go/model/req"
	"novel-go/model/resp"
	"time"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Register(req *req.UserRegisterReqDto) (*resp.UserRegisterRespDto, error) {
	// 1. 校验参数
	if req.Username == "" || req.Password == "" || req.VelCode == "" || req.SessionId == "" {
		return nil, fmt.Errorf("缺少必要的注册参数")
	}

	// 2. 从 Redis 读取验证码
	redisKey := constants.RedisKeyCaptcha + req.SessionId
	dbCode, err := config.RedisClient.Get(config.Ctx, redisKey).Result()
	if err != nil {
		return nil, fmt.Errorf("验证码不存在或已过期")
	}

	// 3. 校验验证码是否匹配
	if dbCode != req.VelCode {
		return nil, fmt.Errorf("验证码错误")
	}

	// 4. 检查用户名是否已存在
	var count int64
	err = config.DB.Model(&pojo.UserInfo{}).Where("username = ?", req.Username).Count(&count).Error
	if err != nil {
		return nil, fmt.Errorf("数据库查询错误: %v", err)
	}
	if count > 0 {
		return nil, fmt.Errorf("用户名已存在")
	}

	// 5. 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("密码加密失败: %v", err)
	}

	// 6. 创建用户实体
	user := pojo.UserInfo{
		Username: req.Username,
		Password: string(hashedPassword),
		// 可根据需求初始化默认值，比如昵称、状态等
		NickName:   req.Username, // 默认昵称为用户名，也可后续修改
		Status:     0,            // 默认正常状态
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	// 7. 入库
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("用户创建失败: %v", err)
	}

	// 8. 构造返回数据（示例，这里生成一个简单的 token，生产环境请替换为JWT等）
	//token, err := s.generateToken(user.ID, user.Username)
	//if err != nil {
	//	return nil, fmt.Errorf("token生成失败: %v", err)
	//}

	respDto := &resp.UserRegisterRespDto{
		UID:      int64(user.ID),
		Token:    "",
		NickName: user.NickName,
		Avatar:   user.UserPhoto, // 头像可为空，后续用户可上传
	}
	return respDto, nil
}
