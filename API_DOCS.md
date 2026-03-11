# 情侣聊天+日记应用 API 文档

## 基本信息

- **Base URL**: `http://localhost:8080`
- **WebSocket URL**: `ws://localhost:8080/ws/chat`
- **数据格式**: JSON
- **字符编码**: UTF-8

## 认证方式

使用 JWT Token 进行认证，有两种方式传递 token：

1. **Header 方式**（推荐）:
   ```
   Authorization: Bearer {token}
   ```

2. **Query 参数方式**:
   ```
   /api/v1/user/profile?token={token}
   ```

Token 在登录和注册接口的响应中返回。

---

## 通用响应格式

### 成功响应
```json
{
  "code": 0,
  "message": "OK",
  "data": { ... }
}
```

### 错误响应
```json
{
  "code": 401,
  "message": "请先登录",
  "data": null
}
```

### 错误码说明
| Code | 说明 |
|------|------|
| 0 | 成功 |
| 401 | 未授权/Token无效 |
| 400 | 请求参数错误 |
| 500 | 服务器内部错误 |

---

## API 接口

### 1. 用户模块

#### 1.1 用户注册

**接口地址**: `POST /api/v1/user/register`

**是否需要认证**: 否

**请求参数**:
```json
{
  "phone": "13800138000",
  "password": "123456",
  "nickname": "小明"
}
```

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| phone | string | 是 | 手机号，11位数字 |
| password | string | 是 | 密码，6-32位 |
| nickname | string | 是 | 昵称 |

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "userId": 1,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "nickname": "小明"
  }
}
```

---

#### 1.2 用户登录

**接口地址**: `POST /api/v1/user/login`

**是否需要认证**: 否

**请求参数**:
```json
{
  "phone": "13800138000",
  "password": "123456"
}
```

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| phone | string | 是 | 手机号 |
| password | string | 是 | 密码 |

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "userId": 1,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "nickname": "小明",
    "avatar": ""
  }
}
```

---

#### 1.3 获取个人信息

**接口地址**: `GET /api/v1/user/profile`

**是否需要认证**: 是

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "userId": 1,
    "phone": "13800138000",
    "nickname": "小明",
    "avatar": "http://example.com/avatar.jpg",
    "createdAt": "2024-01-01T12:00:00+08:00"
  }
}
```

---

#### 1.4 更新个人信息

**接口地址**: `PUT /api/v1/user/profile`

**是否需要认证**: 是

**请求参数**:
```json
{
  "nickname": "新昵称",
  "avatar": "http://example.com/new-avatar.jpg"
}
```

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| nickname | string | 是 | 新昵称 |
| avatar | string | 否 | 头像URL |

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "userId": 1
  }
}
```

---

#### 1.5 生成情侣邀请码

**接口地址**: `POST /api/v1/user/generate-invite`

**是否需要认证**: 是

**说明**: 生成一个6位数字邀请码，有效期为24小时。需要先确保当前用户未绑定情侣关系。

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "inviteCode": "123456"
  }
}
```

---

#### 1.6 绑定情侣关系

**接口地址**: `POST /api/v1/user/bind-couple`

**是否需要认证**: 是

**请求参数**:
```json
{
  "inviteCode": "123456"
}
```

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| inviteCode | string | 是 | 对方生成的邀请码 |

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "coupleId": 1,
    "partnerId": 2
  }
}
```

---

#### 1.7 获取情侣信息

**接口地址**: `GET /api/v1/user/couple-info`

**是否需要认证**: 是

**说明**: 获取当前用户的情侣关系信息及伴侣的基本资料。

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "coupleId": 1,
    "userId": 1,
    "partnerId": 2,
    "nickname": "小红",
    "avatar": "http://example.com/partner-avatar.jpg"
  }
}
```

---

### 2. 聊天模块

#### 2.1 获取聊天历史

**接口地址**: `GET /api/v1/chat/history`

**是否需要认证**: 是

**请求参数**:
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| page | int64 | 否 | 页码，默认1 |
| pageSize | int64 | 否 | 每页数量，默认20，最大100 |
| lastId | uint64 | 否 | 获取此ID之前的消息（用于分页） |

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "list": [
      {
        "id": 1,
        "coupleId": 1,
        "senderId": 1,
        "receiverId": 2,
        "msgType": 1,
        "content": "你好呀",
        "isRead": true,
        "createdAt": "2024-01-01T12:00:00+08:00"
      }
    ],
    "hasMore": false,
    "unreadCount": 0
  }
}
```

**消息类型 (msgType)**:
| 值 | 说明 |
|----|------|
| 1 | 文字消息 |
| 2 | 图片消息 |
| 3 | 语音消息 |

---

#### 2.2 标记消息已读

**接口地址**: `PUT /api/v1/chat/read`

**是否需要认证**: 是

**请求参数**:
```json
{
  "messageId": 1
}
```

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "success": true
  }
}
```

---

#### 2.3 上传聊天文件

**接口地址**: `POST /api/v1/chat/upload`

**是否需要认证**: 是

**请求类型**: `multipart/form-data`

**请求参数**:
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| file | File | 是 | 图片或语音文件 |

**文件限制**:
- 图片支持：jpg, png, gif, webp
- 语音支持：mp3, amr, mp4, wav
- 最大文件大小：10MB

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "url": "/resource/uploads/2024/01/01/xxx.jpg",
    "filename": "photo.jpg",
    "size": 102400
  }
}
```

---

### 3. 日记模块

#### 3.1 创建日记

**接口地址**: `POST /api/v1/diary`

**是否需要认证**: 是

**请求参数**:
```json
{
  "title": "美好的一天",
  "content": "今天天气真好，和喜欢的人一起出去玩了...",
  "isShared": 1,
  "mood": "开心",
  "weather": "晴",
  "mediaIds": []
}
```

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| title | string | 是 | 日记标题 |
| content | string | 是 | 日记内容 |
| isShared | int8 | 否 | 是否共享给对方，1是/0否，默认1 |
| mood | string | 否 | 心情标签 |
| weather | string | 否 | 天气 |
| mediaIds | array | 否 | 附件ID列表 |

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "diaryId": 1
  }
}
```

---

#### 3.2 获取日记列表

**接口地址**: `GET /api/v1/diary/list`

**是否需要认证**: 是

**请求参数**:
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| page | int64 | 否 | 页码，默认1 |
| pageSize | int64 | 否 | 每页数量，默认20，最大100 |

**说明**: 返回自己创建的日记 + 对方共享给自己的日记

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "list": [
      {
        "id": 1,
        "title": "美好的一天",
        "content": "今天天气真好...",
        "isShared": true,
        "mood": "开心",
        "weather": "晴",
        "createdAt": "2024-01-01T12:00:00+08:00",
        "updatedAt": "2024-01-01T12:00:00+08:00",
        "media": []
      }
    ],
    "total": 10,
    "page": 1,
    "pageSize": 20
  }
}
```

---

#### 3.3 获取日记详情

**接口地址**: `GET /api/v1/diary/{id}`

**是否需要认证**: 是

**路径参数**:
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | uint64 | 是 | 日记ID |

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "id": 1,
    "title": "美好的一天",
    "content": "今天天气真好...",
    "isShared": true,
    "mood": "开心",
    "weather": "晴",
    "createdAt": "2024-01-01T12:00:00+08:00",
    "updatedAt": "2024-01-01T12:00:00+08:00",
    "media": []
  }
}
```

---

#### 3.4 更新日记

**接口地址**: `PUT /api/v1/diary/{id}`

**是否需要认证**: 是

**路径参数**:
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | uint64 | 是 | 日记ID |

**请求参数**:
```json
{
  "title": "美好的一天（已更新）",
  "content": "更新后的内容...",
  "isShared": 0,
  "mood": "非常开心",
  "weather": "晴转多云",
  "mediaIds": []
}
```

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "success": true
  }
}
```

---

#### 3.5 删除日记

**接口地址**: `DELETE /api/v1/diary/{id}`

**是否需要认证**: 是

**路径参数**:
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | uint64 | 是 | 日记ID |

**说明**: 只能删除自己创建的日记

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "success": true
  }
}
```

---

#### 3.6 上传日记图片

**接口地址**: `POST /api/v1/diary/upload`

**是否需要认证**: 是

**请求类型**: `multipart/form-data`

**请求参数**:
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| file | File | 是 | 图片文件 |

**文件限制**:
- 支持格式：jpg, png, gif, webp
- 最大文件大小：10MB

**响应示例**:
```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "url": "/resource/uploads/2024/01/01/xxx.jpg",
    "filename": "photo.jpg",
    "size": 102400
  }
}
```

---

## WebSocket 接口

### 连接地址

```
ws://localhost:8080/ws/chat
```

### 消息格式

所有消息均为 JSON 格式：

```json
{
  "type": "消息类型",
  "data": { ... }
}
```

### 消息类型

#### 1. 认证消息

客户端发送（连接后首先发送）:
```json
{
  "type": "auth",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

服务端响应:
```json
{
  "type": "auth",
  "success": true
}
```

---

#### 2. 心跳消息

客户端发送:
```json
{
  "type": "ping"
}
```

服务端响应:
```json
{
  "type": "pong"
}
```

**建议**: 客户端每30秒发送一次ping保持连接。

---

#### 3. 聊天消息

客户端发送:
```json
{
  "type": "message",
  "msgType": 1,
  "content": "你好呀"
}
```

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| msgType | int | 是 | 1文字/2图片/3语音 |
| content | string | 是 | 消息内容或URL |

服务端广播给双方:
```json
{
  "type": "message",
  "data": {
    "id": 1,
    "coupleId": 1,
    "senderId": 1,
    "receiverId": 2,
    "msgType": 1,
    "content": "你好呀"
  }
}
```

---

#### 4. 错误消息

服务端发送:
```json
{
  "type": "error",
  "message": "错误描述"
}
```

**常见错误**:
- `认证失败` - Token无效或过期
- `请先认证` - 未发送认证消息就发送其他消息
- `未绑定情侣关系` - 未绑定情侣关系时发送聊天消息
- `消息发送失败` - 保存消息到数据库失败

---

### WebSocket 连接流程

1. 客户端连接到 `ws://localhost:8080/ws/chat`
2. 发送认证消息 `{"type": "auth", "token": "..."}`
3. 收到认证成功响应 `{"type": "auth", "success": true}`
4. 可以开始发送聊天消息
5. 定期发送ping保持连接

---

## 静态资源

上传的文件通过静态资源服务访问：

```
http://localhost:8080/resource/uploads/{日期}/{文件名}
```

例如：
```
http://localhost:8080/resource/uploads/2024/01/01/abc123.jpg
```

---

## 错误处理

### 常见错误响应

#### 1. 未授权
```json
{
  "code": 401,
  "message": "请先登录",
  "data": null
}
```

#### 2. Token过期
```json
{
  "code": 401,
  "message": "登录已过期，请重新登录",
  "data": null
}
```

#### 3. 参数验证失败
```json
{
  "code": 400,
  "message": "请输入手机号",
  "data": null
}
```

#### 4. 业务错误
```json
{
  "code": 500,
  "message": "该手机号已注册",
  "data": null
}
```

---

## 开发建议

### Token 管理
- Token 有效期为 7 天
- 建议在客户端本地存储 Token
- Token 过期后需要重新登录获取新 Token

### 文件上传
- 图片上传前建议压缩
- 大文件建议分片上传
- 上传后使用返回的 URL 进行展示

### WebSocket 连接
- 实现断线重连机制
- 心跳间隔建议 30 秒
- 收到消息后需要确认是否自己发送的消息

### 分页加载
- 聊天历史建议使用 `lastId` 进行分页
- 日记列表使用传统 page/pageSize 分页
- 每次请求获取 20-50 条数据

---

## 联调工具

- **Swagger UI**: http://localhost:8080/swagger/
- **OpenAPI Spec**: http://localhost:8080/api.json

---

## 更新日志

| 日期 | 版本 | 说明 |
|------|------|------|
| 2024-02-22 | v1.0 | 初始版本，实现基础功能 |
