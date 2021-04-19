

CREATE TABLE `edgex_service_item` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
	`edgex_name` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'edgex服务名',
	`prefix` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'edgex网关前缀',
	`status` tinyint NOT NULL DEFAULT '0' COMMENT '状态: 0-inactive 1-active',
	`deleted` tinyint NOT NULL DEFAULT '0' COMMENT '0-未删除 1-已删除',
	`created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`description` text COLLATE utf8mb4_general_ci COMMENT '描述',
    `location` text COLLATE utf8mb4_general_ci COMMENT 'edgex服务位置信息',
    `extra` text COLLATE utf8mb4_general_ci COMMENT '额外信息',
    PRIMARY KEY (`id`),
    KEY `idx_created_time` (`created_time`),
    KEY `idx_modify_time`  (`modified_time`),
	KEY `idx_user_id` (`user_id`),
    KEY `idx_edgex_name` (`edgex_name`)
) ENGINE=InnoDB AUTO_INCREMENT=1000000000000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='edgex服务表';