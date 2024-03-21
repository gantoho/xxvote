/*
 Navicat Premium Data Transfer

 Source Server         : Windows11
 Source Server Type    : MySQL
 Source Server Version : 80300 (8.3.0)
 Source Host           : localhost:3306
 Source Schema         : vote

 Target Server Type    : MySQL
 Target Server Version : 80300 (8.3.0)
 File Encoding         : 65001

 Date: 21/03/2024 23:14:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `created_time` datetime NULL DEFAULT NULL,
  `updated_time` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'admin', 'admin', '2024-03-21 14:37:06', '2024-03-21 14:37:09');

-- ----------------------------
-- Table structure for vote
-- ----------------------------
DROP TABLE IF EXISTS `vote`;
CREATE TABLE `vote`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `type` int NULL DEFAULT NULL COMMENT '0单选 1多选',
  `status` int NULL DEFAULT NULL COMMENT '0正常 1超时',
  `time` bigint NULL DEFAULT NULL COMMENT '有效时长',
  `user_id` bigint NULL DEFAULT NULL COMMENT '创建人',
  `created_time` datetime NULL DEFAULT NULL,
  `updated_time` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of vote
-- ----------------------------
INSERT INTO `vote` VALUES (1, '今天晚上吃什么？', 0, 0, 86400, 1, '2024-03-21 22:46:56', '2024-03-21 22:46:58');
INSERT INTO `vote` VALUES (2, '去哪里玩？', 0, 0, 86400, 1, '2024-03-21 22:51:20', '2024-03-21 22:51:22');

-- ----------------------------
-- Table structure for vote_opt
-- ----------------------------
DROP TABLE IF EXISTS `vote_opt`;
CREATE TABLE `vote_opt`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `vote_id` bigint NULL DEFAULT NULL,
  `count` int NULL DEFAULT NULL,
  `created_time` datetime NULL DEFAULT NULL,
  `updated_time` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of vote_opt
-- ----------------------------

-- ----------------------------
-- Table structure for vote_opt_user
-- ----------------------------
DROP TABLE IF EXISTS `vote_opt_user`;
CREATE TABLE `vote_opt_user`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NULL DEFAULT NULL,
  `vote_id` bigint NULL DEFAULT NULL,
  `vote_opt_id` bigint NULL DEFAULT NULL,
  `created_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of vote_opt_user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
