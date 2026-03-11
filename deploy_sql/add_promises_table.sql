-- 承诺表创建SQL
-- 执行方式: mysql -u root -p chatbox < add_promises_table.sql

USE chatbox;

-- 创建承诺表
CREATE TABLE IF NOT EXISTS `promises` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `couple_id` BIGINT UNSIGNED NOT NULL COMMENT '情侣ID',
    `creator_id` BIGINT UNSIGNED NOT NULL COMMENT '创建者ID',
    `title` VARCHAR(200) NOT NULL COMMENT '承诺标题',
    `message_ids` JSON NOT NULL COMMENT '关联消息ID列表',
    `color_tag` INT DEFAULT 0 COMMENT '颜色标签: 0=粉色 1=蓝色 2=绿色 3=橙色',
    `is_completed` BOOLEAN DEFAULT FALSE COMMENT '是否完成',
    `completed_at` DATETIME DEFAULT NULL COMMENT '完成时间',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX `idx_couple` (`couple_id`),
    INDEX `idx_creator` (`creator_id`),
    INDEX `idx_completed` (`is_completed`),
    INDEX `idx_created` (`created_at`),
    FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='承诺表';

-- 显示表结构
DESCRIBE promises;

-- 完成提示
SELECT 'Promises table created successfully!' as Status;