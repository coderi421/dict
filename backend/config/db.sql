CREATE TABLE dictionary (
                            id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
                            chinese VARCHAR(100) NOT NULL COMMENT '中文词语',
                            chinese_explanation TEXT NOT NULL COMMENT '中文解释',
                            english VARCHAR(100) NOT NULL COMMENT '英文词语',
                            english_explanation TEXT NOT NULL COMMENT '英文解释',
                            category_id INT UNSIGNED NOT NULL COMMENT '分类（政策话语、经济产业、文旅资源、历史文化、社会民生等）',
                            source VARCHAR(255) COMMENT '出处',
                            remark TEXT COMMENT '备注',
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            updated_by INT COMMENT '更新人',
                            -- 添加 ngram 全文索引
                            FULLTEXT INDEX idx_ngram_chinese (chinese) WITH PARSER ngram,
                            FULLTEXT INDEX idx_ngram_english (english) WITH PARSER ngram,
                            FULLTEXT INDEX idx_ngram_chinese_explanation (chinese_explanation) WITH PARSER ngram,
                            FULLTEXT INDEX idx_ngram_english_explanation (english_explanation) WITH PARSER ngram
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '词典表';

CREATE TABLE users (
                       id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                       username VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       role INT NOT NULL COMMENT '1-管理员，2-普通用户',
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       updated_by INT COMMENT '更新人'
);

CREATE TABLE categories (
                            id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                            name VARCHAR(255) NOT NULL UNIQUE,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE search_hot_keywords (
                                     id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                                     keyword VARCHAR(255) NOT NULL,
                                     search_count INT UNSIGNED NOT NULL DEFAULT 0,
                                     last_searched_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                     UNIQUE INDEX idx_keyword (keyword)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '热点词条表';

-- 测试数据
-- INSERT INTO dictionary (chinese, chinese_explanation, english, english_explanation, category_id, source, remark) VALUES
--                                                                                                                   ('科技', '科学技术的简称，指人类利用知识和工具改造自然的能力。', 'technology', 'The application of scientific knowledge for practical purposes.', 1, '现代汉语词典', '常用词'),
--                                                                                                                   ('山川', '山和河流的总称，常用于文学描述自然景观。', 'mountains and rivers', 'A general term for mountains and rivers, often used in literature.', 2, '古诗词', '文学意象'),
--                                                                                                                   ('应用', '使用某种技术或方法解决问题。', 'application', 'The act of applying something to solve a problem.', 1, '辞海', NULL);
--
-- INSERT INTO users (username, password, role) VALUES
--                                                  ('admin', 'hashed_password_here', 1),
--                                                  ('user1', 'hashed_password_here', 2);


-- my.cnf 相关配置
-- [mysqld]
-- ngram_token_size=1  # 支持单字分词