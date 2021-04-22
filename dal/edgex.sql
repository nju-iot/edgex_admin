

CREATE TABLE `edgex_service_item` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
	`edgex_name` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'edgex服务名',
	`prefix` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'edgex网关前缀',
	`status` tinyint NOT NULL DEFAULT '0' COMMENT '状态: 0-inactive 1-active',
	`deleted` tinyint NOT NULL DEFAULT '0' COMMENT '0-未删除 1-已删除',
	`address` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'edgex网关域名,e.g. 106.15.79.230:8080',
	`created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`description` text COLLATE utf8mb4_general_ci COMMENT '描述',
    `location` text COLLATE utf8mb4_general_ci COMMENT 'edgex服务位置信息',
    `extra` text COLLATE utf8mb4_general_ci COMMENT '额外信息',
    PRIMARY KEY (`id`),
	UNIQUE KEY `idx_prefix` (`prefix`, `deleted`),
    KEY `idx_created_time` (`created_time`),
    KEY `idx_modify_time`  (`modified_time`),
	KEY `idx_user_id` (`user_id`),
    KEY `idx_edgex_name` (`edgex_name`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='edgex服务表';


CREATE TABLE `edgex_related_user` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
	`username` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
	`edgex_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'edgex服务id',
	`edgex_name` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'edgex服务名',
	`status` tinyint NOT NULL DEFAULT '0' COMMENT '状态: 0-无关系 1-follow 2-unfollow',
	`created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
	UNIQUE KEY `idx_relation` (`user_id`, `edgex_id`),
    KEY `idx_created_time` (`created_time`),
    KEY `idx_modify_time`  (`modified_time`),
	KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户收藏edgex记录表';