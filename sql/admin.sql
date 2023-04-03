CREATE TABLE `user` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name` varchar(64) NOT NULL COMMENT '用户名',
    `passwdMd5` varchar(64) NOT NULL COMMENT '密码md5',
    `status` varchar(16) NOT NULL COMMENT '状态',
    `isDelete` tinyint(1) NOT NULL COMMENT '是否删除',
    `cTime` datetime NOT NULL COMMENT '创建时间',
    `mTime` datetime DEFAULT NOW() COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '用户';


CREATE TABLE `message` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `userId` int(10)  NOT NULL COMMENT '用户id',
    `roomId` int(10)  NOT NULL COMMENT '房间id',
    `isSendAll` tinyint(1) NOT NULL COMMENT '是否发布给全部用户',
    `toUserIds` json NOT NULL COMMENT '发送用户列表',
    `content` varchar(512) NOT NULL COMMENT '消息内容',
    `cTime` datetime NOT NULL COMMENT '创建时间',
    `mTime` datetime DEFAULT NOW() COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idxUserId`(`userId`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '消息';