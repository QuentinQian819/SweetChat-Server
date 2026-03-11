-- ChatBox 数据库表创建脚本
-- 数据库: chatbox

-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS chatbox DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE chatbox;

-- 用户表
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `phone` varchar(20) NOT NULL COMMENT '手机号',
  `nickname` varchar(50) NOT NULL COMMENT '昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像URL',
  `password_hash` varchar(255) NOT NULL COMMENT '密码哈希',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 情侣表
DROP TABLE IF EXISTS `couples`;
CREATE TABLE `couples` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '情侣关系ID',
  `user1_id` bigint unsigned NOT NULL COMMENT '用户1 ID',
  `user2_id` bigint unsigned NOT NULL COMMENT '用户2 ID',
  `invite_code` varchar(20) NOT NULL COMMENT '邀请码',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态: 0-待确认, 1-已绑定',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_invite_code` (`invite_code`),
  KEY `idx_user1_id` (`user1_id`),
  KEY `idx_user2_id` (`user2_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='情侣表';

-- 消息表
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '消息ID',
  `couple_id` bigint unsigned NOT NULL COMMENT '情侣关系ID',
  `sender_id` bigint unsigned NOT NULL COMMENT '发送者ID',
  `receiver_id` bigint unsigned NOT NULL COMMENT '接收者ID',
  `msg_type` tinyint NOT NULL COMMENT '消息类型: 1-文本, 2-图片, 3-语音, 4-视频',
  `content` text NOT NULL COMMENT '消息内容',
  `is_read` tinyint NOT NULL DEFAULT '0' COMMENT '是否已读: 0-未读, 1-已读',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_couple_id` (`couple_id`),
  KEY `idx_sender_id` (`sender_id`),
  KEY `idx_receiver_id` (`receiver_id`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='消息表';

-- 承诺/约定表
DROP TABLE IF EXISTS `promises`;
CREATE TABLE `promises` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '承诺ID',
  `couple_id` bigint unsigned NOT NULL COMMENT '情侣关系ID',
  `creator_id` bigint unsigned NOT NULL COMMENT '创建者ID',
  `title` varchar(255) NOT NULL COMMENT '承诺标题',
  `message_ids` text DEFAULT NULL COMMENT '关联的消息ID列表',
  `color_tag` int NOT NULL DEFAULT '0' COMMENT '颜色标签',
  `is_completed` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否完成',
  `completed_at` datetime DEFAULT NULL COMMENT '完成时间',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_couple_id` (`couple_id`),
  KEY `idx_creator_id` (`creator_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='承诺表';

-- 日记表
DROP TABLE IF EXISTS `diaries`;
CREATE TABLE `diaries` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '日记ID',
  `couple_id` bigint unsigned NOT NULL COMMENT '情侣关系ID',
  `author_id` bigint unsigned NOT NULL COMMENT '作者ID',
  `title` varchar(255) NOT NULL COMMENT '日记标题',
  `content` text NOT NULL COMMENT '日记内容',
  `is_shared` tinyint NOT NULL DEFAULT '0' COMMENT '是否共享: 0-私密, 1-共享',
  `mood` varchar(50) DEFAULT NULL COMMENT '心情',
  `weather` varchar(50) DEFAULT NULL COMMENT '天气',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_couple_id` (`couple_id`),
  KEY `idx_author_id` (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='日记表';

-- 日记媒体表
DROP TABLE IF EXISTS `diary_media`;
CREATE TABLE `diary_media` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '媒体ID',
  `diary_id` bigint unsigned NOT NULL COMMENT '日记ID',
  `media_url` varchar(500) NOT NULL COMMENT '媒体文件URL',
  `media_type` varchar(50) NOT NULL COMMENT '媒体类型: image, video等',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_diary_id` (`diary_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='日记媒体表';
