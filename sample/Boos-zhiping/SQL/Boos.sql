/*
 Navicat MySQL Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 50720
 Source Host           : localhost:3306
 Source Schema         : demo

 Target Server Type    : MySQL
 Target Server Version : 50720
 File Encoding         : 65001

 Date: 16/11/2018 19:15:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for Boos
-- ----------------------------
DROP TABLE IF EXISTS `Boos`;
CREATE TABLE `Boos` (
  `Id` int(8) NOT NULL AUTO_INCREMENT,
  `Position` text COMMENT '职位',
  `Url` varchar(255) DEFAULT NULL COMMENT 'Url',
  `des` text COMMENT '职位描述',
  `CompanyInfo` text COMMENT '公司介绍',
  `Location` varchar(255) DEFAULT NULL COMMENT '公司位置',
  `Company` varchar(255) DEFAULT NULL COMMENT '公司名称',
  `Slary` varchar(255) DEFAULT NULL COMMENT '工资',
  `InsertTime` datetime DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of Boos
-- ----------------------------
BEGIN;
INSERT INTO `Boos` VALUES (1, NULL, 'www.baidu.com', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `Boos` VALUES (2, 'Golang开发工程师(业务中间件)', 'https://www.zhipin.com/job_detail/b94c0fa8d0c981f11XZy3t-9Flo~.html', '岗位职责：1. 负责滴滴核心业务的稳定性和效率研发工作；2. 推动项目落地，持续跟进和完善项目；3. 研发和完善基础组件、中间件等，更好的支持业务迭代；4. 善于学习新技术，掌握技术发展方向，应用合适的技术到项目中。任职要求：1. 至少两年Linux下开发经验，Golang至少一年开发经验；2. 熟悉Web开发，对操作系统、TCP/IP有深入理解；3. 基础扎实，熟悉常用数据结构和算法，对分布式系统有所了解；4. 熟悉常用DB、缓存应用和优化；5. 熟悉C/C++、PHP、java等其他语言优先；6. 善于发现问题，解决问题，并能落地解决方案；7. 有较强的抽象、逻辑分析能力，善于总结；8. 有团队精神，善于沟通。', '滴滴出行（www.didiglobal.com）是全球领先的一站式移动出行平台；为5.5亿用户提供出租车、快车、专车、豪华车、顺风车、公交、小巴、代驾、企业级、共享单车、共享电单车、共享汽车、外卖等全面的出行和运输服务，日订单已达3000万。在滴滴平台，超过3000万车主及司机获得灵活赚取收入的机会', '北京海淀区未名视通研发楼', '北京嘀嘀无限科技发展有限公司', '25K-40K', '2018-11-16 19:05:42');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
