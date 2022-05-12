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

 Date: 15/03/2022 07:03:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint(20) unsigned NOT NULL,
  `name` varchar(255) NOT NULL COMMENT '''权限名称''',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `type` tinyint(1) unsigned NOT NULL COMMENT '''对应foundation type''',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (262183989186629, '組員', '2022-03-01 06:49:25', '2022-03-03 20:08:59', 0);
INSERT INTO `role` VALUES (262184008179781, '組長', '2022-03-01 06:49:30', '2022-03-03 08:41:36', 1);
INSERT INTO `role` VALUES (262184028958789, '部門負責人', '2022-03-01 06:49:35', '2022-03-01 06:49:35', 2);
INSERT INTO `role` VALUES (262184517984325, '公司負責人', '2022-03-01 06:51:34', '2022-03-01 06:51:34', 3);
INSERT INTO `role` VALUES (262184577556549, '集團負責人', '2022-03-01 06:51:49', '2022-03-01 06:51:49', 4);
INSERT INTO `role` VALUES (262184698970181, '超級管理員', '2022-03-01 06:52:18', '2022-03-01 06:52:18', 127);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
