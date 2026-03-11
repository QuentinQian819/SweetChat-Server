# 🚀 承诺墙API后端部署 - 完整指南

## 📦 当前状态

### ✅ 已完成
1. **后端代码** - Promise API已添加到 `/Users/quentinqian/chatBox`
2. **编译完成** - 二进制文件已生成：`/Users/quentinqian/chatBox/main`
3. **Android更新** - 应用已支持服务器同步

### 📋 需要你执行的部署步骤

---

## 方式1：一键部署（最简单）

### 在你的Mac上执行：

```bash
cd /Users/quentinqian/chatBox
./upload_and_deploy.sh
```

这个脚本会自动：
1. 上传编译好的二进制文件到服务器
2. 上传部署脚本和SQL文件
3. 给你清晰的后续步骤指引

---

## 方式2：手动部署（如果一键脚本有问题）

### 步骤1：上传文件到服务器

```bash
# 上传新的二进制文件
scp /Users/quentinqian/chatBox/main root@43.138.154.39:/root/chatBox.new

# 上传SQL文件（创建promises表）
scp /Users/quentinqian/chatBox/deploy_sql/add_promises_table.sql root@43.138.154.39:/root/
```

### 步骤2：SSH到服务器

```bash
ssh root@43.138.154.39
```

**如果SSH连接不上，可能的原因：**
- 防火墙阻止了22端口
- SSH服务未启动
- 密码不对

### 步骤3：在服务器上创建promises表

```bash
mysql -u root -p chatbox < /root/add_promises_table.sql
```

或者手动执行SQL（连接MySQL后）：

```sql
USE chatbox;

CREATE TABLE IF NOT EXISTS `promises` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `couple_id` BIGINT UNSIGNED NOT NULL,
    `creator_id` BIGINT UNSIGNED NOT NULL,
    `title` VARCHAR(200) NOT NULL,
    `message_ids` JSON NOT NULL,
    `color_tag` INT DEFAULT 0,
    `is_completed` BOOLEAN DEFAULT FALSE,
    `completed_at` DATETIME DEFAULT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_couple` (`couple_id`),
    INDEX `idx_creator` (`creator_id`),
    INDEX `idx_completed` (`is_completed`),
    INDEX `idx_created` (`created_at`),
    FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='承诺表';
```

### 步骤4：停止旧服务

```bash
# 方式1：如果使用systemd
systemctl stop chatbox

# 方式2：手动停止进程
ps aux | grep chatBox
kill <PID>
```

### 步骤5：部署新版本

```bash
cd /root
mv chatBox.new chatBox
chmod +x chatBox
```

### 步骤6：启动服务

```bash
# 后台运行
nohup ./chatBox > chatBox.log 2>&1 &

# 检查是否启动成功
sleep 3
ps aux | grep chatBox

# 查看日志
tail -f chatbox.log
```

---

## ✅ 验证部署

### 测试API（在服务器上）

```bash
# 测试承诺列表API
curl http://localhost:8080/api/v1/promise/list
```

应该返回JSON格式的承诺列表。

### 测试Android应用

1. 重启Android应用
2. 进入承诺墙
3. 应该能看到从服务器同步的承诺数据

---

## 🔧 故障排查

### 问题1：SSH连接超时

**解决方案：**
- 检查网络连接
- 确认服务器IP地址正确：43.138.154.39
- 尝试ping：`ping 43.138.154.39`

### 问题2：API返回404

**原因：**promises表未创建

**解决：**
```bash
# 在服务器上执行
mysql -u root -p chatbox -e "SHOW TABLES LIKE 'promises';"

# 如果表不存在，执行SQL创建
mysql -u root -p chatbox < /root/add_promises_table.sql
```

### 问题3：服务启动失败

**查看日志：**
```bash
tail -100 /root/chatbox.log
```

**常见错误：**
- 端口被占用：`lsof -i:8080`
- 数据库连接失败：检查MySQL是否运行
- 配置文件错误：检查 `manifest/config/config.yaml`

### 问题4：Android应用无法同步

**检查Android日志：**
```bash
adb logcat | grep PromiseViewModel
adb logcat | grep "ApiClient"
```

**检查Token是否有效：**
- 重新登录应用获取新Token
- 确认后端API地址正确：`43.138.154.39:8080`

---

## 📊 API端点列表

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/v1/promise | 创建承诺 |
| GET | /api/v1/promise/list | 获取承诺列表 |
| GET | /api/v1/promise/:id | 获取承诺详情 |
| PUT | /api/v1/promise/:id | 更新承诺 |
| DELETE | /api/v1/promise/:id | 删除承诺 |
| PUT | /api/v1/promise/:id/complete | 切换完成状态 |

---

## 📞 需要帮助？

如果遇到问题，请提供以下信息：

1. **错误日志**
   - 服务器日志：`tail -100 /root/chatbox.log`
   - Android日志：`adb logcat | grep PromiseViewModel`

2. **配置信息**
   - 服务器IP和端口
   - MySQL版本：`mysql --version`
   - Go版本：`go version`

3. **网络测试结果**
   - `ping 43.138.154.39`
   - `telnet 43.138.154.39 8080`

---

## 🎉 部署成功后

你的承诺墙功能将：
- ✅ 数据永久保存在服务器
- ✅ 多设备同步
- ✅ 应用重装后数据恢复
- ✅ 实时同步创建/删除/完成状态

开始部署吧！💪
