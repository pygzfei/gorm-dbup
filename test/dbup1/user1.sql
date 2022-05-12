/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50733
 Source Host           : localhost:3306
 Source Schema         : prolbem_reflection

 Target Server Type    : MySQL
 Target Server Version : 50733
 File Encoding         : 65001

 Date: 14/03/2022 23:19:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `id` bigint(20) unsigned NOT NULL,
  `name` varchar(256) NOT NULL,
  `age` tinyint(1) NOT NULL DEFAULT '0',
  `sex` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES ('2022-02-15 16:30:06', '2022-02-15 16:30:06', 257372176924741, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:11', '2022-02-15 16:30:11', 257372195954757, 'anson', 27, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:11', '2022-02-15 16:30:11', 257372196646981, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:11', '2022-02-15 16:30:11', 257372197384261, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:11', '2022-02-15 16:30:11', 257372198072389, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:11', '2022-02-15 16:30:11', 257372198715461, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:12', '2022-02-15 16:30:12', 257372199411781, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:12', '2022-02-15 16:30:12', 257372200050757, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:12', '2022-02-15 16:30:12', 257372200755269, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:12', '2022-02-15 16:30:12', 257372201447493, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:12', '2022-02-15 16:30:12', 257372202086469, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:12', '2022-02-15 16:30:12', 257372202770501, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:13', '2022-02-15 16:30:13', 257372203421765, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:13', '2022-02-15 16:30:13', 257372204109893, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:13', '2022-02-15 16:30:13', 257372204793925, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:13', '2022-02-15 16:30:13', 257372205494341, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:13', '2022-02-15 16:30:13', 257372206186565, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:13', '2022-02-15 16:30:13', 257372206850117, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:14', '2022-02-15 16:30:14', 257372207525957, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:30:14', '2022-02-15 16:30:14', 257372208140357, 'anson', 28, 0);
INSERT INTO `user` VALUES ('2022-02-15 16:37:26', '2022-02-15 16:37:26', 257373979717701, 'anson', 28, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
