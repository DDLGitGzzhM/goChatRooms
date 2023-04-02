CREATE TABLE `user` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `nameMd5` varchar(64) NOT NULL COMMENT '用户名md5',
    `passwdMd5` varchar(64) NOT NULL COMMENT '密码md5',
    `avatar` varchar(64) NOT NULL COMMENT '头像',
    `status` varchar(16) NOT NULL COMMENT '状态', //0下线 1上线 拉黑定义在这里是否合适？
    `isDelete` tinyint(1) NOT NULL COMMENT '是否删除', // 0否 1是
    `cTime` datetime NOT NULL COMMENT '创建时间',
    `mTime` datetime DEFAULT NOW() COMMENT '更新时间',
    PRIMARY KEY (`id`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '用户吧';


CREATE TABLE `message` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `userId` int(10)  NOT NULL COMMENT '用户id',
    `roomId` int(10)  NOT NULL COMMENT '房间id',
    `isSendAll` tinyint(1) NOT NULL COMMENT '是否发布给全部用户', // 0否 1是
    `toUserIds` json NOT NULL COMMENT '发送用户列表', //列表
    `content` varchar(512) NOT NULL COMMENT '消息内容', //注意什么
    `cTime` datetime NOT NULL COMMENT '创建时间',
    `mTime` datetime DEFAULT NOW() COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idxUserId`(`userId`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '消息';