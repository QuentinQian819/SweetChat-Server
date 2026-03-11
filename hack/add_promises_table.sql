-- Add promises table to ChatBox database
-- Run this to create the promises table in an existing database

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
