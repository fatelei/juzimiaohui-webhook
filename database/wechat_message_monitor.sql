CREATE TABLE `wechat_message_monitor` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `wechat_id` varchar(1024) NOT NULL,
  `wechat_name` varchar(1024) DEFAULT NULL,
  `room_name` varchar(1024) DEFAULT NULL,
  `content` varchar(1024) NOT NULL,
  `msg_type` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `room_id` varchar(200) DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `idx_room` (`room_name`(191)),
  KEY `idx_wechat_name` (`wechat_name`(191))
) ENGINE=InnoDB AUTO_INCREMENT=1376 DEFAULT CHARSET=utf8mb4


CREATE TABLE `wechat_room` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `room_id` VARCHAR(200) NOT NULL,
    `room_name` VARCHAR(1024) DEFAULT "",
    `room_member_number` INT(11) DEFAULT 0,
    `open_monitor` TINYINT(1) DEFAULT 1,
    PRIMARY KEY (`id`),
    KEY `idx_room_id` (`room_id`),
    KEY `idx_room_name` (`room_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
