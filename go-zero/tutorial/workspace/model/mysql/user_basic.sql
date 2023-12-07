
-- 实体表
DROP TABLE class_info, school_info, work_info, teacher_info,  homework_info, student_info

CREATE TABLE `school_info`  (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
    `img`   varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT 'img',
    `desc`  varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
    `place` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '校区位置',
    `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态 0正常 1冻结',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '学校表';

CREATE TABLE `class_info`  (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '班级名',
    `img`  varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT ' 班级标图',
    `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT ' 班级描述',
    `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态 0正常 1冻结',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '班级表';

CREATE TABLE `teacher_info`  (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `gender` tinyint NOT NULL DEFAULT 1 COMMENT '用户性别 1女 2男',
    `role` tinyint NOT NULL DEFAULT 1 COMMENT '用户角色 1学生 2老师',
    `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录名',
    `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录密码',
    `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '系统用户' COMMENT '用户昵称',
    `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
    `authority_id` bigint UNSIGNED NOT NULL DEFAULT 888 COMMENT '用户角色ID',
    `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 0 COMMENT '用户手机号',
    `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户邮箱',
   `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态 0正常 1冻结',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '老师表';

CREATE TABLE `student_info`  (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `gender` tinyint NOT NULL DEFAULT 1 COMMENT '用户性别 1女 2男',
    `role` tinyint NOT NULL DEFAULT 1 COMMENT '用户角色 1学生 2老师',
    `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录名',
    `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录密码',
    `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '系统用户' COMMENT '用户昵称',
    `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
    `authority_id` bigint UNSIGNED NOT NULL DEFAULT 888 COMMENT '用户角色ID',
    `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 0 COMMENT '用户手机号',
    `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户邮箱',
   `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态 0正常 1冻结',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '学生表';

CREATE TABLE `work_info`  (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `images` varchar(256) NOT NULL DEFAULT '' COMMENT '作品',
    `place` tinyint NOT NULL DEFAULT 1 COMMENT '位置信息',
   `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态 0正常 1冻结',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '学生表';

CREATE TABLE `homework_info`  (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `images` varchar(256) NOT NULL DEFAULT '' COMMENT '作业图片',
    `desc` varchar(256) NOT NULL DEFAULT '' COMMENT '作业描述',
   `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态 0正常 1冻结',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '学生表';



-- 关系表
CREATE TABLE `school_class`  (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `school_id` bigint UNSIGNED NOT NULL DEFAULT 0,
    `class_id` bigint UNSIGNED NOT NULL DEFAULT 0,
    `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态 0正常 1冻结',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '学校 - 班级表';

CREATE TABLE `class_teacher`  (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `class_id` bigint UNSIGNED NOT NULL DEFAULT 0,
    `teacher_id` bigint UNSIGNED NOT NULL DEFAULT 0,
    `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态 0正常 1冻结',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '学校 - 班级表';

CREATE TABLE `class_student`  (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `class_id` bigint UNSIGNED NOT NULL DEFAULT 0,
    `student_id` bigint UNSIGNED NOT NULL DEFAULT 0,
    `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态 0正常 1冻结',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '班级 - 学生表';

CREATE TABLE `class_homework`  (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `class_id` bigint UNSIGNED NOT NULL DEFAULT 0,
    `homework_id` bigint UNSIGNED NOT NULL DEFAULT 0,
    `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态 0正常 1冻结',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '班级 - 作业表';
