# Novel-Go 项目

这是一个基于 Go 语言开发的后端项目，使用了 Gin Web 框架构建 RESTful API，并集成了 Swag 自动化接口文档生成工具。项目还使用 Redis 实现了验证码的存储和数据缓存功能。

## 🚀 技术栈

- **Gin**：轻量级的 Go Web 框架，处理路由和 HTTP 请求。
- **Swag**：基于注解生成 Swagger 文档，方便接口调试和文档管理。
- **Redis**：作为缓存中间件，用于存储图片验证码、提升系统性能。
- **base64Captcha**：用于生成图形验证码。

## 🔧 功能模块

### ✅ 用户端接口（/api/front）

- `GET /resource/img_verify_code`：生成图片验证码，返回 Base64 图片。
- `POST /resource/image`：图片上传接口（需登录）。

### ✅ 管理后台接口（规划中）

## 📖 接口文档

启动项目后访问：

