-- 黑名单表
CREATE TABLE `black` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '拉黑记录ID',
  `type` int NOT NULL COMMENT '拉黑目标类型 1.ip 2.uid',
  `target` varchar(255) NOT NULL COMMENT '拉黑目标',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='拉黑记录表';

-- 用户IP信息表
CREATE TABLE `ip_detail` (
  `ip` varchar(255) NOT NULL COMMENT '注册时的ip',
  `isp` varchar(255) COMMENT '最新登录的ip',
  `isp_id` varchar(255) COMMENT 'ISP ID',
  `city` varchar(255) COMMENT '城市',
  `city_id` varchar(255) COMMENT '城市ID',
  `country` varchar(255) COMMENT '国家',
  `country_id` varchar(255) COMMENT '国家ID',
  `region` varchar(255) COMMENT '地区',
  `region_id` varchar(255) COMMENT '地区ID',
  PRIMARY KEY (`ip`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户IP信息表';

-- 用户IP信息表
CREATE TABLE `ip_info` (
  `create_ip` varchar(255) NOT NULL COMMENT '注册时的ip',
  `create_ip_detail` bigint(20) COMMENT '注册时的ip详情ID',
  `update_ip` varchar(255) COMMENT '最新登录的ip',
  `update_ip_detail` bigint(20) COMMENT '最新登录的ip详情ID',
  PRIMARY KEY (`create_ip`),
  FOREIGN KEY (`create_ip_detail`) REFERENCES `ip_detail`(`id`),
  FOREIGN KEY (`update_ip_detail`) REFERENCES `ip_detail`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户IP信息表';

-- 功能物品配置表
CREATE TABLE `item_config` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '物品ID',
  `type` int NOT NULL COMMENT '物品类型 1改名卡 2徽章',
  `img` varchar(255) COMMENT '物品图片',
  `describe` varchar(255) COMMENT '物品功能描述',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='功能物品配置表';

-- 角色表
CREATE TABLE `role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(255) NOT NULL COMMENT '角色名称',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- 用户表
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `name` varchar(255) NOT NULL COMMENT '用户昵称',
  `avatar` varchar(255) COMMENT '用户头像',
  `sex` int COMMENT '性别 1为男性，2为女性',
  `open_id` varchar(255) COMMENT '微信openid用户标识',
  `active_status` int COMMENT '上下线状态 1在线 2离线',
  `last_opt_time` datetime COMMENT '最后上下线时间',
  `ip_info` json COMMENT '用户IP信息',
  `item_id` bigint(20) COMMENT '佩戴的徽章ID',
  `status` int COMMENT '用户状态 0正常 1拉黑',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 用户申请表
CREATE TABLE `user_apply` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '申请ID',
  `uid` bigint(20) NOT NULL COMMENT '申请人UID',
  `type` int NOT NULL COMMENT '申请类型 1加好友',
  `target_id` bigint(20) NOT NULL COMMENT '接收人UID',
  `msg` varchar(255) COMMENT '申请信息',
  `status` int NOT NULL COMMENT '申请状态 1待审批 2同意',
  `read_status` int NOT NULL COMMENT '阅读状态 1未读 2已读',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户申请表';

-- 用户背包表
CREATE TABLE `user_backpack` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '背包ID',
  `uid` bigint(20) NOT NULL COMMENT '用户UID',
  `item_id` bigint(20) NOT NULL COMMENT '物品ID',
  `status` int NOT NULL COMMENT '使用状态 0.未失效 1失效',
  `idempotent` varchar(255) COMMENT '幂等号',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户背包表';

-- 用户表情包表
CREATE TABLE `user_emoji` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '表情包ID',
  `uid` bigint(20) NOT NULL COMMENT '用户UID',
  `expression_url` varchar(255) NOT NULL COMMENT '表情地址',
  `delete_status` int NOT NULL DEFAULT 0 COMMENT '逻辑删除(0-正常,1-删除)',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表情包表';

-- 用户联系人表
CREATE TABLE `user_friend` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '联系人ID',
  `uid` bigint(20) NOT NULL COMMENT '用户UID',
  `friend_uid` bigint(20) NOT NULL COMMENT '好友UID',
  `delete_status` int NOT NULL DEFAULT 0 COMMENT '逻辑删除(0-正常,1-删除)',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户联系人表';

-- 用户角色关系表
CREATE TABLE `user_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色关系ID',
  `uid` bigint(20) NOT NULL COMMENT '用户UID',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色关系表';

