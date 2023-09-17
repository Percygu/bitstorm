DROP TABLE IF EXISTS `t_prize`;
CREATE TABLE `t_prize`
(
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品名称',
    `prize_num` int(11) NOT NULL DEFAULT '-1' COMMENT '奖品数量，0 无限量，>0限量，<0无奖品',
    `left_num` int(11) NOT NULL DEFAULT '0' COMMENT '剩余数量',
    `prize_code` varchar(50) NOT NULL DEFAULT '' COMMENT '0-9999表示100%，0-0表示万分之一的中奖概率',
    `prize_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '发奖周期，多少天，以天为单位',
    `img` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品图片',
    `display_order` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '位置序号，小的排在前面',
    `prize_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '奖品类型，1-虚拟币，2-虚拟券，3-实物小奖，4-实物大奖',
    `prize_profile` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品扩展数据，如：虚拟币数量',
    `begin_time` datetime NOT NULL DEFAULT '1000-01-01 00:00:00' COMMENT '奖品有效周期：开始时间',
    `end_time` datetime NOT NULL DEFAULT '1000-01-01 00:00:00' COMMENT '奖品有效周期：结束时间',
    `prize_plan` mediumtext COMMENT '发奖计划，[[时间1,数量1],[时间2,数量2]]',
    `prize_begin` int(11) NOT NULL DEFAULT '0' COMMENT '发奖计划周期的开始',
    `prize_end` int(11) NOT NULL DEFAULT '0' COMMENT '发奖计划周期的结束',
    `sys_status` smallint(5) unsigned NOT NULL DEFAULT '1' COMMENT '状态，1-正常，2-删除',
    `sys_created` datetime NOT NULL DEFAULT '1000-01-01 00:00:00' COMMENT '创建时间',
    `sys_updated` datetime NOT NULL DEFAULT '1000-01-01 00:00:00' COMMENT'修改时间',
    `sys_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '操作人IP',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='奖品表';