-- ChatBox Database Schema

-- 1. users table - 用户表
CREATE TABLE IF NOT EXISTS users (
    id            BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    phone         VARCHAR(20) UNIQUE NOT NULL COMMENT '手机号',
    nickname      VARCHAR(50) NOT NULL COMMENT '昵称',
    avatar        VARCHAR(255) COMMENT '头像URL',
    password_hash VARCHAR(255) NOT NULL COMMENT '密码哈希',
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at    DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 2. couples table - 情侣关系表
CREATE TABLE IF NOT EXISTS couples (
    id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user1_id    BIGINT UNSIGNED NOT NULL COMMENT '用户1 ID',
    user2_id    BIGINT UNSIGNED NOT NULL COMMENT '用户2 ID',
    invite_code VARCHAR(20) UNIQUE NOT NULL COMMENT '邀请码',
    status      TINYINT DEFAULT 1 COMMENT '状态: 1正常 0解绑',
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uk_users (user1_id, user2_id),
    FOREIGN KEY (user1_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (user2_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='情侣关系表';

-- 3. messages table - 聊天消息表
CREATE TABLE IF NOT EXISTS messages (
    id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    couple_id   BIGINT UNSIGNED NOT NULL COMMENT '情侣ID',
    sender_id   BIGINT UNSIGNED NOT NULL COMMENT '发送者ID',
    receiver_id BIGINT UNSIGNED NOT NULL COMMENT '接收者ID',
    msg_type    TINYINT NOT NULL COMMENT '类型: 1文字 2图片 3语音 4报备',
    content     TEXT COMMENT '文字内容/媒体URL',
    is_read     TINYINT DEFAULT 0 COMMENT '是否已读',
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_couple (couple_id),
    INDEX idx_created (created_at),
    INDEX idx_sender (sender_id),
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='聊天消息表';

-- 4. diaries table - 日记表
CREATE TABLE IF NOT EXISTS diaries (
    id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    couple_id  BIGINT UNSIGNED NOT NULL COMMENT '情侣ID',
    author_id  BIGINT UNSIGNED NOT NULL COMMENT '作者ID',
    title      VARCHAR(200) NOT NULL COMMENT '标题',
    content    TEXT NOT NULL COMMENT '内容',
    is_shared  TINYINT DEFAULT 1 COMMENT '是否共享给对方: 1是 0否',
    mood       VARCHAR(20) COMMENT '心情标签',
    weather    VARCHAR(20) COMMENT '天气',
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_couple (couple_id),
    INDEX idx_created (created_at),
    INDEX idx_author (author_id),
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日记表';

-- 5. diary_media table - 日记附件表
CREATE TABLE IF NOT EXISTS diary_media (
    id        BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    diary_id  BIGINT UNSIGNED NOT NULL COMMENT '日记ID',
    media_url VARCHAR(255) NOT NULL COMMENT '媒体URL',
    media_type VARCHAR(20) NOT NULL COMMENT '类型: image/audio',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_diary (diary_id),
    FOREIGN KEY (diary_id) REFERENCES diaries(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日记附件表';

-- 6. promises table - 承诺表
CREATE TABLE IF NOT EXISTS promises (
    id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    couple_id    BIGINT UNSIGNED NOT NULL COMMENT '情侣ID',
    creator_id   BIGINT UNSIGNED NOT NULL COMMENT '创建者ID',
    title        VARCHAR(200) NOT NULL COMMENT '标题',
    message_ids  JSON COMMENT '关联消息ID列表',
    color_tag    INT DEFAULT 0 COMMENT '颜色标签',
    is_completed BOOLEAN DEFAULT FALSE COMMENT '是否完成',
    completed_at DATETIME COMMENT '完成时间',
    created_at   DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_couple (couple_id),
    INDEX idx_creator (creator_id),
    INDEX idx_completed (is_completed),
    INDEX idx_created (created_at),
    FOREIGN KEY (creator_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='承诺表';
