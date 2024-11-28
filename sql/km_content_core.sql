CREATE TABLE `content_label` (
                                 `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '标签id',
                                 `content_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '内容id',
                                 `content_type` smallint(5) NOT NULL DEFAULT '0' COMMENT '内容类型 笔记：1',
                                 `label_type` smallint(5) NOT NULL DEFAULT '0' COMMENT '标签类型：内容标签：1，笔记标签：2，年龄：3，性别：4，用户类型：5',
                                 `label_value` varchar(255) NOT NULL DEFAULT '' COMMENT '标签值',
                                 `label_name` varchar(255) NOT NULL DEFAULT '' COMMENT '标签名',
                                 `group_name` varchar(255) NOT NULL DEFAULT '' COMMENT '标签组名称',
                                 `group_value` varchar(255) NOT NULL DEFAULT '' COMMENT '标签组值',
                                 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                 PRIMARY KEY (`id`) USING BTREE,
                                 KEY `idx_label` (`content_id`,`content_type`,`label_type`) USING BTREE COMMENT '标签查询',
                                 KEY `idx_content` (`label_value`,`label_type`) USING BTREE COMMENT '内容查询'
) ENGINE=InnoDB AUTO_INCREMENT=5782831 DEFAULT CHARSET=utf8mb4 COMMENT='内容标签表' ;

CREATE TABLE `KAEKZK__________content_work_old` (
                                                    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                                    `title` varchar(256) NOT NULL DEFAULT '' COMMENT '作品名称',
                                                    `artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '横版封面',
                                                    `tab_artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '竖版封面',
                                                    `description_keypoint` varchar(1024) NOT NULL DEFAULT '' COMMENT '简介',
                                                    `description_text` varchar(16383) NOT NULL DEFAULT '' COMMENT '内容详情',
                                                    `updated_section_count` smallint(5) NOT NULL DEFAULT '0' COMMENT '已更新总集数',
                                                    `expect_section_count` smallint(5) NOT NULL DEFAULT '0' COMMENT '总集数',
                                                    `leading_role` varchar(512) NOT NULL DEFAULT '' COMMENT '主角名称',
                                                    `categories_type` int(10) NOT NULL DEFAULT '0' COMMENT '内容分类0故事1知识',
                                                    `space_type` int(10) NOT NULL DEFAULT '0' COMMENT '是否为长篇1：长篇，0：短篇',
                                                    `expected_word_count` int(32) NOT NULL DEFAULT '0' COMMENT '预计作品字数（万）',
                                                    `product_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '商品类型',
                                                    `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                                    `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                                    `updater` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
                                                    `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                                    `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                                    `cell_type` int(10) NOT NULL DEFAULT '0' COMMENT 'Cell 类型 对应视频、音频、图文',
                                                    `word_total` bigint(20) NOT NULL DEFAULT '0' COMMENT '作品下小节总字数',
                                                    `sku_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'sku_id',
                                                    `audit_status` varchar(32) NOT NULL DEFAULT '' COMMENT '审核状态',
                                                    `clue_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '线索id',
                                                    `source_from` varchar(16) NOT NULL DEFAULT '' COMMENT '来源',
                                                    `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                                    `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                                    `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '资源id',
                                                    `extra` text COMMENT '额外信息',
                                                    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                                    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1766084292613169153 DEFAULT CHARSET=utf8 COMMENT='内容库作品表';

CREATE TABLE `content_work` (
                                `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                `title` varchar(256) NOT NULL DEFAULT '' COMMENT '作品名称',
                                `artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '横版封面',
                                `tab_artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '竖版封面',
                                `description_keypoint` varchar(1024) NOT NULL DEFAULT '' COMMENT '简介',
                                `description_text` varchar(16383) NOT NULL DEFAULT '' COMMENT '内容详情',
                                `updated_section_count` smallint(5) NOT NULL DEFAULT '0' COMMENT '已更新总集数',
                                `expect_section_count` smallint(5) NOT NULL DEFAULT '0' COMMENT '总集数',
                                `leading_role` varchar(512) NOT NULL DEFAULT '' COMMENT '主角名称',
                                `categories_type` int(10) NOT NULL DEFAULT '0' COMMENT '内容分类0故事1知识',
                                `space_type` int(10) NOT NULL DEFAULT '0' COMMENT '是否为长篇1：长篇，0：短篇',
                                `expected_word_count` int(32) NOT NULL DEFAULT '0' COMMENT '预计作品字数（万）',
                                `product_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '商品类型',
                                `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                `updater` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
                                `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                `cell_type` int(10) NOT NULL DEFAULT '0' COMMENT 'Cell 类型 对应视频、音频、图文',
                                `word_total` bigint(20) NOT NULL DEFAULT '0' COMMENT '作品下小节总字数',
                                `sku_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'sku_id',
                                `audit_status` varchar(32) NOT NULL DEFAULT '' COMMENT '审核状态',
                                `clue_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '线索id',
                                `series_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '系列ID',
                                `source_from` varchar(16) NOT NULL DEFAULT '' COMMENT '来源',
                                `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '资源id',
                                `extra` text COMMENT '额外信息',
                                `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1771238188012670977 DEFAULT CHARSET=utf8 COMMENT='内容库作品表';

CREATE TABLE `BW0DDP__________content_work_old` (
                                                    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                                    `title` varchar(256) NOT NULL DEFAULT '' COMMENT '作品名称',
                                                    `artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '横版封面',
                                                    `tab_artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '竖版封面',
                                                    `description_keypoint` varchar(1024) NOT NULL DEFAULT '' COMMENT '简介',
                                                    `description_text` varchar(16383) NOT NULL DEFAULT '' COMMENT '内容详情',
                                                    `leading_role` varchar(512) NOT NULL DEFAULT '' COMMENT '主角名称',
                                                    `categories_type` int(10) NOT NULL DEFAULT '0' COMMENT '内容分类0故事1知识',
                                                    `space_type` int(10) NOT NULL DEFAULT '0' COMMENT '是否为长篇1：长篇，0：短篇',
                                                    `expected_word_count` int(32) NOT NULL DEFAULT '0' COMMENT '预计作品字数（万）',
                                                    `product_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '商品类型',
                                                    `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                                    `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                                    `updater` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
                                                    `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                                    `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                                    `cell_type` int(10) NOT NULL DEFAULT '0' COMMENT 'Cell 类型 对应视频、音频、图文',
                                                    `word_total` bigint(20) NOT NULL DEFAULT '0' COMMENT '作品下小节总字数',
                                                    `sku_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'sku_id',
                                                    `audit_status` varchar(32) NOT NULL DEFAULT '' COMMENT '审核状态',
                                                    `clue_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '线索id',
                                                    `source_from` varchar(16) NOT NULL DEFAULT '' COMMENT '来源',
                                                    `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                                    `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                                    `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '资源id',
                                                    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                                    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1755601580013842433 DEFAULT CHARSET=utf8 COMMENT='内容库作品表';

CREATE TABLE `9Z60X0__________content_work_old` (
                                                    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                                    `title` varchar(256) NOT NULL DEFAULT '' COMMENT '作品名称',
                                                    `artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '横版封面',
                                                    `tab_artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '竖版封面',
                                                    `description_keypoint` varchar(1024) NOT NULL DEFAULT '' COMMENT '简介',
                                                    `description_text` varchar(16383) NOT NULL DEFAULT '' COMMENT '内容详情',
                                                    `leading_role` varchar(512) NOT NULL DEFAULT '' COMMENT '主角名称',
                                                    `categories_type` int(10) NOT NULL DEFAULT '0' COMMENT '内容分类0故事1知识',
                                                    `space_type` int(10) NOT NULL DEFAULT '0' COMMENT '是否为长篇1：长篇，0：短篇',
                                                    `expected_word_count` int(32) NOT NULL DEFAULT '0' COMMENT '预计作品字数（万）',
                                                    `product_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '商品类型',
                                                    `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                                    `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                                    `updater` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
                                                    `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                                    `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                                    `cell_type` int(10) NOT NULL DEFAULT '0' COMMENT 'Cell 类型 对应视频、音频、图文',
                                                    `sku_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'sku_id',
                                                    `audit_status` varchar(32) NOT NULL DEFAULT '' COMMENT '审核状态',
                                                    `clue_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '线索id',
                                                    `source_from` varchar(16) NOT NULL DEFAULT '' COMMENT '来源',
                                                    `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                                    `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                                    `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '资源id',
                                                    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                                    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1746489434034147329 DEFAULT CHARSET=utf8 COMMENT='内容库作品表';

CREATE TABLE `work_section` (
                                `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                `work_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '作品id',
                                `sku_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '商品 id',
                                `title` varchar(256) NOT NULL DEFAULT '' COMMENT '小节名称',
                                `section_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '小节类型正文、番外',
                                `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '文本资源id',
                                `audit_status` varchar(32) NOT NULL DEFAULT '' COMMENT '审稿状态',
                                `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                `is_free` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否免费, 0 付费, 1 免费',
                                `global_idx` int(11) NOT NULL DEFAULT '0' COMMENT '全局索引',
                                `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                PRIMARY KEY (`id`),
                                KEY `idx_work_id` (`work_id`),
                                KEY `idx_section_resource_id` (`resource_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1771238188012670977 DEFAULT CHARSET=utf8 COMMENT='作品小节关系表';

CREATE TABLE `_______work_section_old` (
                                           `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                           `work_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '作品id',
                                           `title` varchar(256) NOT NULL DEFAULT '' COMMENT '小节名称',
                                           `section_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '小节类型正文、番外',
                                           `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '文本资源id',
                                           `audit_status` varchar(32) NOT NULL DEFAULT '' COMMENT '审稿状态',
                                           `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                           `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                           `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                           `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                           `global_idx` int(11) NOT NULL DEFAULT '0' COMMENT '全局索引',
                                           `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                           `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                           `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                           `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                           PRIMARY KEY (`id`),
                                           KEY `idx_work_id` (`work_id`),
                                           KEY `idx_section_resource_id` (`resource_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1753431975753601025 DEFAULT CHARSET=utf8 COMMENT='作品小节关系表';

CREATE TABLE `______work_section_old` (
                                          `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                          `work_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '作品id',
                                          `title` varchar(256) NOT NULL DEFAULT '' COMMENT '小节名称',
                                          `section_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '小节类型正文、番外',
                                          `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '文本资源id',
                                          `audit_status` varchar(32) NOT NULL DEFAULT '' COMMENT '审稿状态',
                                          `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                          `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                          `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                          `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                          `global_idx` int(11) NOT NULL DEFAULT '0' COMMENT '全局索引',
                                          `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                          `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                          `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                          `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                          PRIMARY KEY (`id`),
                                          KEY `idx_work_id` (`work_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1746570494093066241 DEFAULT CHARSET=utf8 COMMENT='作品小节关系表';

CREATE TABLE `_________content_work_old` (
                                             `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                             `title` varchar(256) NOT NULL DEFAULT '' COMMENT '作品名称',
                                             `artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '横版封面',
                                             `tab_artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '竖版封面',
                                             `description_keypoint` varchar(1024) NOT NULL DEFAULT '' COMMENT '简介',
                                             `description_text` varchar(16383) NOT NULL DEFAULT '' COMMENT '内容详情',
                                             `leading_role` varchar(512) NOT NULL DEFAULT '' COMMENT '主角名称',
                                             `categories_type` int(10) NOT NULL DEFAULT '0' COMMENT '内容分类0故事1知识',
                                             `space_type` int(10) NOT NULL DEFAULT '0' COMMENT '是否为长篇1：长篇，0：短篇',
                                             `expected_word_count` int(32) NOT NULL DEFAULT '0' COMMENT '预计作品字数（万）',
                                             `product_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '商品类型',
                                             `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                             `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                             `updater` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
                                             `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                             `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                             `cell_type` int(10) NOT NULL DEFAULT '0' COMMENT 'Cell 类型 对应视频、音频、图文',
                                             `audit_status` varchar(32) NOT NULL DEFAULT '' COMMENT '审核状态',
                                             `clue_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '线索id',
                                             `source_from` varchar(16) NOT NULL DEFAULT '' COMMENT '来源',
                                             `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                             `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                             `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '资源id',
                                             `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                             `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                             PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1725957659025547266 DEFAULT CHARSET=utf8 COMMENT='内容库作品表';

CREATE TABLE `operate_log` (
                               `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                               `bussiness_id` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '业务ID',
                               `bussiness_name` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '业务方名称',
                               `old_value` varchar(2048) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '数据快照',
                               `new_value` varchar(2048) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '更新后的数据快照',
                               `operator` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
                               `operate_mark` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '操作日志信息(核心)',
                               `operate_type` int(11) NOT NULL DEFAULT '0' COMMENT '操作类型',
                               `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`id`),
                               KEY `idx_operate_type_bussiness_id` (`operate_type`,`bussiness_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2805478 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='日志表';

CREATE TABLE `work_author_relation` (
                                        `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                        `work_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '作品id',
                                        `author_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '作者id',
                                        `author_type` int(10) NOT NULL DEFAULT '0' COMMENT '作者类型, 0 普通作者, 1 短剧原著作者, 2 短剧创作者 导演, 3 短剧创作者 编剧',
                                        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                        `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                        PRIMARY KEY (`id`),
                                        KEY `idx_work_id` (`work_id`)
) ENGINE=InnoDB AUTO_INCREMENT=918416 DEFAULT CHARSET=utf8 COMMENT='作品作者关系表';

CREATE TABLE `vip_pin` (
                           `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                           `member_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '发布人 id',
                           `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'pin标题',
                           `public_status` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '上下架状态 0 未上架 1上架 2已下架',
                           `audit_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '审核状态',
                           `is_delete` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '删除状态 0 未删除 1已删除',
                           `owner_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '角色类型',
                           `resource_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '资源类型',
                           `extra` varchar(5000) NOT NULL DEFAULT '' COMMENT '额外数据，如笔记首图以及对应宽高',
                           `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人 id',
                           `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人 id',
                           `publish_at` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '自动发布时间戳',
                           `expire_at` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '自动隐藏时间戳',
                           `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           `pin_type` int(8) NOT NULL DEFAULT '0' COMMENT '笔记类型，0默认类型，1系统笔记',
                           PRIMARY KEY (`id`),
                           KEY `idx_poster_id_is_delete` (`member_id`,`is_delete`),
                           KEY `idx_publish_at` (`publish_at`),
                           KEY `idx_expire_at` (`expire_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1771237852811890689 DEFAULT CHARSET=utf8mb4 COMMENT='想法笔记表';

CREATE TABLE `_work_author_relation_old` (
                                             `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                             `work_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '作品id',
                                             `author_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '作者id',
                                             `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                             `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                             PRIMARY KEY (`id`),
                                             KEY `idx_work_id` (`work_id`)
) ENGINE=InnoDB AUTO_INCREMENT=697044 DEFAULT CHARSET=utf8 COMMENT='作品作者关系表';

CREATE TABLE `pin_image` (
                             `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                             `pin_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '资源 id',
                             `img_hash` varchar(255) NOT NULL DEFAULT '' COMMENT '图片hash',
                             `img_meta` varchar(1024) NOT NULL DEFAULT '' COMMENT '图片元数据，如宽高',
                             `sort` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '图片排序',
                             `resource_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '资源id',
                             `resource_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '资源类型',
                             `relation_content_token` varchar(255) NOT NULL DEFAULT '' COMMENT '关联作品id，如长短篇、书单、活动页',
                             `relation_content_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '关联作品类型 short_story/long_story/book_list/h5_activity',
                             `relation_content_extra` text COMMENT '扩展资源额外数据',
                             `extra` text COMMENT '资源扩展数据',
                             `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                             `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                             PRIMARY KEY (`id`),
                             UNIQUE KEY `uniq_resource_order` (`pin_id`,`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=95736 DEFAULT CHARSET=utf8mb4 COMMENT='笔记相册表';

CREATE TABLE `____work_section_old` (
                                        `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                        `work_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '作品id',
                                        `title` varchar(256) NOT NULL DEFAULT '' COMMENT '小节名称',
                                        `section_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '小节类型正文、番外',
                                        `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '文本资源id',
                                        `audit_status` varchar(32) NOT NULL DEFAULT '' COMMENT '审稿状态',
                                        `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                        `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                        `global_idx` int(11) NOT NULL DEFAULT '0' COMMENT '全局索引',
                                        `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                        `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                        `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                        PRIMARY KEY (`id`),
                                        KEY `idx_work_id` (`work_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1717921164721131521 DEFAULT CHARSET=utf8 COMMENT='作品小节关系表';

CREATE TABLE `________vip_pin_old` (
                                       `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                       `member_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '发布人 id',
                                       `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'pin标题',
                                       `public_status` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '上下架状态 0 未上架 1上架 2已下架',
                                       `is_delete` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '删除状态 0 未删除 1已删除',
                                       `owner_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '角色类型',
                                       `resource_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '资源类型',
                                       `extra` varchar(5000) NOT NULL DEFAULT '' COMMENT '额外数据，如笔记首图以及对应宽高',
                                       `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人 id',
                                       `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人 id',
                                       `publish_at` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '自动发布时间戳',
                                       `expire_at` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '自动隐藏时间戳',
                                       `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                       `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                       `pin_type` int(8) NOT NULL DEFAULT '0' COMMENT '笔记类型，0默认类型，1系统笔记',
                                       PRIMARY KEY (`id`),
                                       KEY `idx_poster_id_is_delete` (`member_id`,`is_delete`),
                                       KEY `idx_publish_at` (`publish_at`),
                                       KEY `idx_expire_at` (`expire_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1754457865673101313 DEFAULT CHARSET=utf8mb4 COMMENT='想法笔记表';

CREATE TABLE `_____content_work_old` (
                                         `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                         `title` varchar(256) NOT NULL DEFAULT '' COMMENT '作品名称',
                                         `artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '横版封面',
                                         `tab_artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '竖版封面',
                                         `description_keypoint` varchar(256) NOT NULL DEFAULT '' COMMENT '一句话简介',
                                         `description_text` varchar(16383) NOT NULL DEFAULT '' COMMENT '内容详情',
                                         `leading_role` varchar(512) NOT NULL DEFAULT '' COMMENT '主角名称',
                                         `categories_type` int(10) NOT NULL DEFAULT '0' COMMENT '内容分类1故事2知识',
                                         `space_type` int(10) NOT NULL DEFAULT '0' COMMENT '是否为长篇1：长篇，2：短篇',
                                         `expected_word_count` int(32) NOT NULL DEFAULT '0' COMMENT '预计作品字数（万）',
                                         `product_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '商品类型',
                                         `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                         `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                         `updater` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
                                         `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                         `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                         `cell_type` int(10) NOT NULL DEFAULT '0' COMMENT 'Cell 类型 对应视频、音频、图文',
                                         `clue_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '线索id',
                                         `source_from` varchar(16) NOT NULL DEFAULT '' COMMENT '来源',
                                         `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                         `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                         `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '资源id',
                                         `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                         `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1717921107473076225 DEFAULT CHARSET=utf8 COMMENT='内容库作品表';

CREATE TABLE `_____work_section_old` (
                                         `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                         `work_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '作品id',
                                         `title` varchar(256) NOT NULL DEFAULT '' COMMENT '小节名称',
                                         `section_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '小节类型正文、番外',
                                         `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '文本资源id',
                                         `audit_status` varchar(32) NOT NULL DEFAULT '' COMMENT '审稿状态',
                                         `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                         `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                         `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                         `global_idx` int(11) NOT NULL DEFAULT '0' COMMENT '全局索引',
                                         `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                         `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                         `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                         `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                         PRIMARY KEY (`id`),
                                         KEY `idx_work_id` (`work_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1717921164721131521 DEFAULT CHARSET=utf8 COMMENT='作品小节关系表';

CREATE TABLE `_______vip_pin_old` (
                                      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                      `member_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '发布人 id',
                                      `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'pin标题',
                                      `public_status` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '上下架状态 0 未上架 1上架 2已下架',
                                      `is_delete` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '删除状态 0 未删除 1已删除',
                                      `owner_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '角色类型',
                                      `resource_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '资源类型',
                                      `extra` varchar(5000) NOT NULL DEFAULT '' COMMENT '额外数据，如笔记首图以及对应宽高',
                                      `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人 id',
                                      `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人 id',
                                      `publish_at` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '自动发布时间戳',
                                      `expire_at` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '自动隐藏时间戳',
                                      `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                      `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                      `pin_type` int(8) NOT NULL DEFAULT '0' COMMENT '笔记类型，0默认类型，1系统笔记',
                                      PRIMARY KEY (`id`),
                                      KEY `idx_poster_id_is_delete` (`member_id`,`is_delete`),
                                      KEY `idx_publish_at` (`publish_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1744301904581292033 DEFAULT CHARSET=utf8mb4 COMMENT='想法笔记表';

CREATE TABLE `_____vip_pin_old` (
                                    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                    `member_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '发布人 id',
                                    `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'pin标题',
                                    `public_status` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '上下架状态 0 未上架 1上架 2已下架',
                                    `is_delete` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '删除状态 0 未删除 1已删除',
                                    `owner_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '角色类型',
                                    `resource_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '资源类型',
                                    `extra` varchar(5000) NOT NULL DEFAULT '' COMMENT '额外数据，如笔记首图以及对应宽高',
                                    `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人 id',
                                    `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人 id',
                                    `publish_at` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '自动发布时间戳',
                                    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                    `pin_type` int(8) NOT NULL DEFAULT '0' COMMENT '笔记类型，0默认类型，1系统笔记',
                                    PRIMARY KEY (`id`),
                                    KEY `idx_poster_id_is_delete` (`member_id`,`is_delete`)
) ENGINE=InnoDB AUTO_INCREMENT=1744301904581292033 DEFAULT CHARSET=utf8mb4 COMMENT='想法笔记表';

CREATE TABLE `______vip_pin_old` (
                                     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                     `member_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '发布人 id',
                                     `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'pin标题',
                                     `public_status` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '上下架状态 0 未上架 1上架 2已下架',
                                     `is_delete` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '删除状态 0 未删除 1已删除',
                                     `owner_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '角色类型',
                                     `resource_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '资源类型',
                                     `extra` varchar(5000) NOT NULL DEFAULT '' COMMENT '额外数据，如笔记首图以及对应宽高',
                                     `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人 id',
                                     `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人 id',
                                     `publish_at` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '自动发布时间戳',
                                     `expire_at` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '自动隐藏时间戳',
                                     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                     `pin_type` int(8) NOT NULL DEFAULT '0' COMMENT '笔记类型，0默认类型，1系统笔记',
                                     PRIMARY KEY (`id`),
                                     KEY `idx_poster_id_is_delete` (`member_id`,`is_delete`)
) ENGINE=InnoDB AUTO_INCREMENT=1744301904581292033 DEFAULT CHARSET=utf8mb4 COMMENT='想法笔记表';

CREATE TABLE `________content_work_old` (
                                            `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                            `title` varchar(256) NOT NULL DEFAULT '' COMMENT '作品名称',
                                            `artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '横版封面',
                                            `tab_artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '竖版封面',
                                            `description_keypoint` varchar(256) NOT NULL DEFAULT '' COMMENT '一句话简介',
                                            `description_text` varchar(16383) NOT NULL DEFAULT '' COMMENT '内容详情',
                                            `leading_role` varchar(512) NOT NULL DEFAULT '' COMMENT '主角名称',
                                            `categories_type` int(10) NOT NULL DEFAULT '0' COMMENT '内容分类0故事1知识',
                                            `space_type` int(10) NOT NULL DEFAULT '0' COMMENT '是否为长篇1：长篇，0：短篇',
                                            `expected_word_count` int(32) NOT NULL DEFAULT '0' COMMENT '预计作品字数（万）',
                                            `product_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '商品类型',
                                            `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                            `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                            `updater` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
                                            `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                            `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                            `cell_type` int(10) NOT NULL DEFAULT '0' COMMENT 'Cell 类型 对应视频、音频、图文',
                                            `audit_status` varchar(32) NOT NULL DEFAULT '' COMMENT '审核状态',
                                            `clue_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '线索id',
                                            `source_from` varchar(16) NOT NULL DEFAULT '' COMMENT '来源',
                                            `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                            `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                            `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '资源id',
                                            `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                            `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1718637197719363585 DEFAULT CHARSET=utf8 COMMENT='内容库作品表';

CREATE TABLE `_______content_work_old` (
                                           `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                           `title` varchar(256) NOT NULL DEFAULT '' COMMENT '作品名称',
                                           `artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '横版封面',
                                           `tab_artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '竖版封面',
                                           `description_keypoint` varchar(256) NOT NULL DEFAULT '' COMMENT '一句话简介',
                                           `description_text` varchar(16383) NOT NULL DEFAULT '' COMMENT '内容详情',
                                           `leading_role` varchar(512) NOT NULL DEFAULT '' COMMENT '主角名称',
                                           `categories_type` int(10) NOT NULL DEFAULT '0' COMMENT '内容分类0故事1知识',
                                           `space_type` int(10) NOT NULL DEFAULT '0' COMMENT '是否为长篇1：长篇，0：短篇',
                                           `expected_word_count` int(32) NOT NULL DEFAULT '0' COMMENT '预计作品字数（万）',
                                           `product_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '商品类型',
                                           `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                           `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                           `updater` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
                                           `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                           `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                           `cell_type` int(10) NOT NULL DEFAULT '0' COMMENT 'Cell 类型 对应视频、音频、图文',
                                           `clue_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '线索id',
                                           `source_from` varchar(16) NOT NULL DEFAULT '' COMMENT '来源',
                                           `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                           `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                           `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '资源id',
                                           `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                           `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                           PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1717921107473076225 DEFAULT CHARSET=utf8 COMMENT='内容库作品表';

CREATE TABLE `______content_work_old` (
                                          `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                          `title` varchar(256) NOT NULL DEFAULT '' COMMENT '作品名称',
                                          `artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '横版封面',
                                          `tab_artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '竖版封面',
                                          `description_keypoint` varchar(256) NOT NULL DEFAULT '' COMMENT '一句话简介',
                                          `description_text` varchar(16383) NOT NULL DEFAULT '' COMMENT '内容详情',
                                          `leading_role` varchar(512) NOT NULL DEFAULT '' COMMENT '主角名称',
                                          `categories_type` int(10) NOT NULL DEFAULT '0' COMMENT '内容分类0故事1知识',
                                          `space_type` int(10) NOT NULL DEFAULT '0' COMMENT '是否为长篇1：长篇，2：短篇',
                                          `expected_word_count` int(32) NOT NULL DEFAULT '0' COMMENT '预计作品字数（万）',
                                          `product_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '商品类型',
                                          `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                          `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                          `updater` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
                                          `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                          `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                          `cell_type` int(10) NOT NULL DEFAULT '0' COMMENT 'Cell 类型 对应视频、音频、图文',
                                          `clue_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '线索id',
                                          `source_from` varchar(16) NOT NULL DEFAULT '' COMMENT '来源',
                                          `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                          `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                          `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '资源id',
                                          `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                          `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1717921107473076225 DEFAULT CHARSET=utf8 COMMENT='内容库作品表';

CREATE TABLE `____vip_pin_old` (
                                   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                   `member_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '发布人 id',
                                   `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'pin标题',
                                   `public_status` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '上下架状态 0 未上架 1上架 2已下架',
                                   `is_delete` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '删除状态 0 未删除 1已删除',
                                   `owner_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '角色类型',
                                   `resource_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '资源类型',
                                   `extra` varchar(5000) NOT NULL DEFAULT '' COMMENT '额外数据，如笔记首图以及对应宽高',
                                   `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人 id',
                                   `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人 id',
                                   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                   `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                   `pin_type` int(8) NOT NULL DEFAULT '0' COMMENT '笔记类型，0默认类型，1系统笔记',
                                   PRIMARY KEY (`id`),
                                   KEY `idx_poster_id_is_delete` (`member_id`,`is_delete`)
) ENGINE=InnoDB AUTO_INCREMENT=1744301904581292033 DEFAULT CHARSET=utf8mb4 COMMENT='想法笔记表';

CREATE TABLE `_work_section_old` (
                                     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                     `work_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '作品id',
                                     `title` varchar(256) NOT NULL DEFAULT '' COMMENT '小节名称',
                                     `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '文本资源id',
                                     `audit_status` varchar(32) NOT NULL DEFAULT '' COMMENT '审稿状态',
                                     `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                     `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                     `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                     PRIMARY KEY (`id`),
                                     KEY `idx_work_id` (`work_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1713219037151105025 DEFAULT CHARSET=utf8 COMMENT='作品小节关系表';

CREATE TABLE `__work_section_old` (
                                      `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                      `work_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '作品id',
                                      `title` varchar(256) NOT NULL DEFAULT '' COMMENT '小节名称',
                                      `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '文本资源id',
                                      `audit_status` varchar(32) NOT NULL DEFAULT '' COMMENT '审稿状态',
                                      `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                      `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                      `global_idx` int(11) NOT NULL DEFAULT '0' COMMENT '全局索引',
                                      `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                      `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                      `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                      PRIMARY KEY (`id`),
                                      KEY `idx_work_id` (`work_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1716047657112289282 DEFAULT CHARSET=utf8 COMMENT='作品小节关系表';

CREATE TABLE `____pin_image_old` (
                                     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                     `pin_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '资源 id',
                                     `img_hash` varchar(255) NOT NULL DEFAULT '' COMMENT '图片hash',
                                     `img_meta` varchar(1024) NOT NULL DEFAULT '' COMMENT '图片元数据，如宽高',
                                     `sort` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '图片排序',
                                     `resource_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '资源id',
                                     `resource_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '资源类型',
                                     `relation_content_token` varchar(255) NOT NULL DEFAULT '' COMMENT '关联作品id，如长短篇、书单、活动页',
                                     `relation_content_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '关联作品类型 short_story/long_story/book_list/h5_activity',
                                     `relation_content_extra` text COMMENT '扩展资源额外数据',
                                     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                     PRIMARY KEY (`id`),
                                     UNIQUE KEY `uniq_resource_order` (`pin_id`,`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=25572 DEFAULT CHARSET=utf8mb4 COMMENT='笔记相册表';

CREATE TABLE `___pin_image_old` (
                                    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                    `pin_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '资源 id',
                                    `img_hash` varchar(255) NOT NULL DEFAULT '' COMMENT '图片hash',
                                    `img_meta` varchar(1024) NOT NULL DEFAULT '' COMMENT '图片元数据，如宽高',
                                    `sort` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '图片排序',
                                    `resource_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '资源id',
                                    `resource_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '资源类型',
                                    `relation_content_token` varchar(255) NOT NULL DEFAULT '' COMMENT '关联作品id，如长短篇、书单、活动页',
                                    `relation_content_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '关联作品类型 short_story/long_story/book_list/h5_activity',
                                    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                    PRIMARY KEY (`id`),
                                    UNIQUE KEY `uniq_resource_order` (`pin_id`,`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=25572 DEFAULT CHARSET=utf8mb4 COMMENT='笔记相册表';

CREATE TABLE `___work_section_old` (
                                       `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                       `work_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '作品id',
                                       `title` varchar(256) NOT NULL DEFAULT '' COMMENT '小节名称',
                                       `section_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '小节类型正文、番外',
                                       `resource_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '文本资源id',
                                       `audit_status` varchar(32) NOT NULL DEFAULT '' COMMENT '审稿状态',
                                       `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                       `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                       `global_idx` int(11) NOT NULL DEFAULT '0' COMMENT '全局索引',
                                       `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                       `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                       `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                       PRIMARY KEY (`id`),
                                       KEY `idx_work_id` (`work_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1716047657112289282 DEFAULT CHARSET=utf8 COMMENT='作品小节关系表';

CREATE TABLE `__pin_image_old` (
                                   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                   `pin_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '资源 id',
                                   `img_hash` varchar(255) NOT NULL DEFAULT '' COMMENT '图片hash',
                                   `img_meta` varchar(1024) NOT NULL DEFAULT '' COMMENT '图片元数据，如宽高',
                                   `sort` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '图片排序',
                                   `resource_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '资源id',
                                   `relation_content_token` varchar(255) NOT NULL DEFAULT '' COMMENT '关联作品id，如长短篇、书单、活动页',
                                   `relation_content_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '关联作品类型 short_story/long_story/book_list/h5_activity',
                                   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                   `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `uniq_resource_order` (`pin_id`,`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=25572 DEFAULT CHARSET=utf8mb4 COMMENT='笔记相册表';

CREATE TABLE `channel_content_relation` (
                                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                            `content_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '内容 id',
                                            `content_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '1 表示内容 id 为作品 id',
                                            `channel_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '渠道 id',
                                            `channel_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '1 表示业务 id 为渠道节点 id, 2 表示业务 id 为渠道组 id',
                                            `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
                                            `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '最新一次流通时间',
                                            `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                            `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                            PRIMARY KEY (`id`),
                                            UNIQUE KEY `uniq_content_channel` (`content_id`,`channel_id`,`channel_type`,`content_type`)
) ENGINE=InnoDB AUTO_INCREMENT=236443 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='渠道内容关系表';

CREATE TABLE `___vip_pin_old` (
                                  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                  `member_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '发布人 id',
                                  `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'pin标题',
                                  `public_status` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '上下架状态 0 未上架 1上架 2已下架',
                                  `is_delete` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '删除状态 0 未删除 1已删除',
                                  `owner_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '角色类型',
                                  `extra` varchar(5000) NOT NULL DEFAULT '' COMMENT '额外数据，如笔记首图以及对应宽高',
                                  `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人 id',
                                  `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人 id',
                                  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                  `pin_type` int(8) NOT NULL DEFAULT '0' COMMENT '笔记类型，0默认类型，1系统笔记',
                                  PRIMARY KEY (`id`),
                                  KEY `idx_poster_id_is_delete` (`member_id`,`is_delete`)
) ENGINE=InnoDB AUTO_INCREMENT=1730619139318616065 DEFAULT CHARSET=utf8mb4 COMMENT='想法笔记表';

CREATE TABLE `_operate_log_old` (
                                    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                    `bussiness_id` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '业务ID',
                                    `bussiness_name` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '业务方名称',
                                    `old_value` varchar(2048) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '数据快照',
                                    `new_value` varchar(2048) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '更新后的数据快照',
                                    `operator` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
                                    `operate_mark` varchar(1024) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作日志信息(核心)',
                                    `operate_type` int(11) NOT NULL DEFAULT '0' COMMENT '操作类型',
                                    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
                                    PRIMARY KEY (`id`),
                                    KEY `idx_operate_type_bussiness_id` (`operate_type`,`bussiness_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2718071 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='日志表';

CREATE TABLE `__vip_pin_old` (
                                 `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                 `member_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '发布人 id',
                                 `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'pin标题',
                                 `public_status` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '上下架状态 0 未上架 1上架 2已下架',
                                 `is_delete` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '删除状态 0 未删除 1已删除',
                                 `extra` varchar(5000) NOT NULL DEFAULT '' COMMENT '额外数据，如笔记首图以及对应宽高',
                                 `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人 id',
                                 `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人 id',
                                 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                 `pin_type` int(8) NOT NULL DEFAULT '0' COMMENT '笔记类型，0默认类型，1系统笔记',
                                 PRIMARY KEY (`id`),
                                 KEY `idx_poster_id_is_delete` (`member_id`,`is_delete`)
) ENGINE=InnoDB AUTO_INCREMENT=1730619139318616065 DEFAULT CHARSET=utf8mb4 COMMENT='想法笔记表';

CREATE TABLE `_pin_image_old` (
                                  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                  `pin_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '资源 id',
                                  `img_hash` varchar(255) NOT NULL DEFAULT '' COMMENT '图片hash',
                                  `img_meta` varchar(1024) NOT NULL DEFAULT '' COMMENT '图片元数据，如宽高',
                                  `sort` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '图片排序',
                                  `relation_content_token` varchar(255) NOT NULL DEFAULT '' COMMENT '关联作品id，如长短篇、书单、活动页',
                                  `relation_content_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '关联作品类型 short_story/long_story/book_list/h5_activity',
                                  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `uniq_resource_order` (`pin_id`,`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=25572 DEFAULT CHARSET=utf8mb4 COMMENT='笔记相册表';

CREATE TABLE `__content_work_old` (
                                      `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                      `title` varchar(256) NOT NULL DEFAULT '' COMMENT '作品名称',
                                      `artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '横版封面',
                                      `tab_artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '竖版封面',
                                      `description_keypoint` varchar(256) NOT NULL DEFAULT '' COMMENT '一句话简介',
                                      `description_text` varchar(16383) NOT NULL DEFAULT '' COMMENT '内容详情',
                                      `leading_role` varchar(512) NOT NULL DEFAULT '' COMMENT '主角名称',
                                      `categories_type` int(10) NOT NULL DEFAULT '0' COMMENT '内容分类1故事2知识',
                                      `space_type` int(10) NOT NULL DEFAULT '0' COMMENT '是否为长篇1：长篇，2：短篇',
                                      `expected_word_count` int(32) NOT NULL DEFAULT '0' COMMENT '预计作品字数（万）',
                                      `product_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '商品类型',
                                      `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                      `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                      `updater` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
                                      `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                      `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                      `cell_type` int(10) NOT NULL DEFAULT '0' COMMENT 'Cell 类型 对应视频、音频、图文',
                                      `source_from` int(11) NOT NULL DEFAULT '0' COMMENT '来源0稿件后台1买断',
                                      `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                      `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                      `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                      `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                      PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1715711518966960129 DEFAULT CHARSET=utf8 COMMENT='内容库作品表';

CREATE TABLE `___content_work_old` (
                                       `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                       `title` varchar(256) NOT NULL DEFAULT '' COMMENT '作品名称',
                                       `artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '横版封面',
                                       `tab_artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '竖版封面',
                                       `description_keypoint` varchar(256) NOT NULL DEFAULT '' COMMENT '一句话简介',
                                       `description_text` varchar(16383) NOT NULL DEFAULT '' COMMENT '内容详情',
                                       `leading_role` varchar(512) NOT NULL DEFAULT '' COMMENT '主角名称',
                                       `categories_type` int(10) NOT NULL DEFAULT '0' COMMENT '内容分类1故事2知识',
                                       `space_type` int(10) NOT NULL DEFAULT '0' COMMENT '是否为长篇1：长篇，2：短篇',
                                       `expected_word_count` int(32) NOT NULL DEFAULT '0' COMMENT '预计作品字数（万）',
                                       `product_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '商品类型',
                                       `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                       `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                       `updater` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
                                       `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                       `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                       `cell_type` int(10) NOT NULL DEFAULT '0' COMMENT 'Cell 类型 对应视频、音频、图文',
                                       `clue_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '线索id',
                                       `source_from` int(11) NOT NULL DEFAULT '0' COMMENT '来源0稿件后台1买断',
                                       `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                       `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                       `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                       `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                       PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1715711518966960129 DEFAULT CHARSET=utf8 COMMENT='内容库作品表';

CREATE TABLE `____content_work_old` (
                                        `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                        `title` varchar(256) NOT NULL DEFAULT '' COMMENT '作品名称',
                                        `artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '横版封面',
                                        `tab_artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '竖版封面',
                                        `description_keypoint` varchar(256) NOT NULL DEFAULT '' COMMENT '一句话简介',
                                        `description_text` varchar(16383) NOT NULL DEFAULT '' COMMENT '内容详情',
                                        `leading_role` varchar(512) NOT NULL DEFAULT '' COMMENT '主角名称',
                                        `categories_type` int(10) NOT NULL DEFAULT '0' COMMENT '内容分类1故事2知识',
                                        `space_type` int(10) NOT NULL DEFAULT '0' COMMENT '是否为长篇1：长篇，2：短篇',
                                        `expected_word_count` int(32) NOT NULL DEFAULT '0' COMMENT '预计作品字数（万）',
                                        `product_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '商品类型',
                                        `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                        `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                        `updater` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
                                        `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                        `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                        `cell_type` int(10) NOT NULL DEFAULT '0' COMMENT 'Cell 类型 对应视频、音频、图文',
                                        `clue_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '线索id',
                                        `source_from` varchar(16) NOT NULL DEFAULT '' COMMENT '来源',
                                        `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                        `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                        `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1715711518966960129 DEFAULT CHARSET=utf8 COMMENT='内容库作品表';

CREATE TABLE `_vip_pin_old` (
                                `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                `member_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '发布人 id',
                                `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'pin标题',
                                `public_status` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '上下架状态 0 未上架 1上架 2已下架',
                                `is_delete` int(16) unsigned NOT NULL DEFAULT '0' COMMENT '删除状态 0 未删除 1已删除',
                                `extra` varchar(5000) NOT NULL DEFAULT '' COMMENT '额外数据，如笔记首图以及对应宽高',
                                `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人 id',
                                `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人 id',
                                `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                PRIMARY KEY (`id`),
                                KEY `idx_poster_id_is_delete` (`member_id`,`is_delete`)
) ENGINE=InnoDB AUTO_INCREMENT=1721244734931079169 DEFAULT CHARSET=utf8mb4 COMMENT='想法笔记表';

CREATE TABLE `_content_work_old` (
                                     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                     `title` varchar(256) NOT NULL DEFAULT '' COMMENT '作品名称',
                                     `artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '横版封面',
                                     `tab_artwork` varchar(256) NOT NULL DEFAULT '' COMMENT '竖版封面',
                                     `description_keypoint` varchar(256) NOT NULL DEFAULT '' COMMENT '一句话简介',
                                     `description_text` varchar(16383) NOT NULL DEFAULT '' COMMENT '内容详情',
                                     `leading_role` varchar(512) NOT NULL DEFAULT '' COMMENT '主角名称',
                                     `categories_type` int(10) NOT NULL DEFAULT '0' COMMENT '内容分类1故事2知识',
                                     `space_type` int(10) NOT NULL DEFAULT '0' COMMENT '是否为长篇1：长篇，2：短篇',
                                     `expected_word_count` int(32) NOT NULL DEFAULT '0' COMMENT '预计作品字数（万）',
                                     `product_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '商品类型',
                                     `producer` varchar(128) NOT NULL DEFAULT '' COMMENT '负责的制作人',
                                     `creator` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                     `updater` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
                                     `published_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '首次上架时间',
                                     `public_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上架状态',
                                     `cell_type` int(10) NOT NULL DEFAULT '0' COMMENT 'Cell 类型 对应视频、音频、图文',
                                     `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试',
                                     `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态',
                                     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1715711518966960129 DEFAULT CHARSET=utf8 COMMENT='内容库作品表';

CREATE TABLE `sex_transform_section` (
                                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                         `section_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '小节id',
                                         `transform_section_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '转换后小节id',
                                         `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '转向 1:男 2:女',
                                         `operator` varchar(256) NOT NULL DEFAULT '' COMMENT '上传人',
                                         `remark` varchar(1024) NOT NULL DEFAULT '' COMMENT '备注',
                                         `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                         `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                         PRIMARY KEY (`id`),
                                         UNIQUE KEY `uniq_section_id` (`section_id`),
                                         KEY `idx_transform_section_id` (`transform_section_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3297 DEFAULT CHARSET=utf8mb4 COMMENT='性转小节表';

CREATE TABLE `group_key_relation` (
                                      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                      `group_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '渠道组 id',
                                      `node_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '渠道节点 id',
                                      `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                      `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                      PRIMARY KEY (`id`),
                                      KEY `idx_group_id` (`group_id`),
                                      KEY `idx_node_id` (`node_id`)
) ENGINE=InnoDB AUTO_INCREMENT=210174 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='渠道组和节点关系表';

CREATE TABLE `content_channel_group` (
                                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                         `group_key` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '渠道组 key',
                                         `group_name` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '渠道组名字',
                                         `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                         `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                         PRIMARY KEY (`id`),
                                         UNIQUE KEY `uniq_group_key` (`group_key`)
) ENGINE=InnoDB AUTO_INCREMENT=1766505919615864833 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='渠道组表';

CREATE TABLE `content_channel_node` (
                                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                        `node_key` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '节点 key',
                                        `node_name` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '节点名字',
                                        `node_type` int(10) NOT NULL DEFAULT '0' COMMENT '1 app, 2 web, 3 h5, 4 抖音小程序, 5 快手小程序',
                                        `rights` varchar(512) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '权益支持, json 字符串',
                                        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                        `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                        PRIMARY KEY (`id`) USING BTREE,
                                        UNIQUE KEY `uniq_node_key` (`node_key`)
) ENGINE=InnoDB AUTO_INCREMENT=1767936457631338497 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='渠道节点表';

CREATE TABLE `content_relational` (
                                      `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                      `source_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '源 id',
                                      `source_content_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'source_id 的类型, 1 为 work, 2 为 section',
                                      `relation_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '与 源 id 相关联的 id',
                                      `relation_content_type` bigint(20) NOT NULL DEFAULT '0' COMMENT 'relation_id 的类型, 1 为 work, 2 为 section',
                                      `relation_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '关系类型, 1 relation_id 为 source_id 的原著 id',
                                      `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                      `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                      PRIMARY KEY (`id`),
                                      KEY `idx_source_id` (`source_id`)
) ENGINE=InnoDB AUTO_INCREMENT=623148 DEFAULT CHARSET=utf8 COMMENT='内容关系表';