-- MySQL dump 10.16  Distrib 10.2.14-MariaDB, for Win64 (AMD64)
--
-- Host: localhost    Database: mlm
-- ------------------------------------------------------
-- Server version	10.2.14-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `members`
--

DROP TABLE IF EXISTS `members`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `members` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf32 COLLATE utf32_unicode_ci DEFAULT NULL,
  `leader_id` int(11) DEFAULT NULL,
  `level` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `leader_id_fk_idx` (`leader_id`),
  CONSTRAINT `member_id_fk_leader_id` FOREIGN KEY (`leader_id`) REFERENCES `members` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `members`
--

LOCK TABLES `members` WRITE;
/*!40000 ALTER TABLE `members` DISABLE KEYS */;
INSERT INTO `members` VALUES (10001,'อะธีน่า',20001,1),(10002,'โพไซดอน',20002,1),(10003,'อริสตา',20003,1),(10004,'อาบิเกล',20004,1),(10005,'เอเรีย',20005,2),(10006,'เอมมี่',20006,2),(10007,'อลิสสา',20007,2),(10008,'แองเจล่า',20008,2),(10009,'เอมเบอร์',20009,3),(10010,'อมีเรีย',20010,3),(10011,'เอมี่',20011,3),(10012,'เอวา',20012,3),(10013,'แอนน์',20013,3),(10014,'อลิซ',20014,3),(10015,'เอนิซซา',20015,4),(10016,'แอนเดรีย',20016,4),(10017,'จตุรภัทร',20017,4),(10018,'จรีภรณ์',20018,4),(10019,'กชนันท์',20019,4),(10020,'อณิกา',20020,4),(10021,'อลินดา',20021,4),(10022,'ไอริน',20022,5),(10023,'อานนท์',20023,5),(10024,'อาชา',20024,5),(10025,'บอน',20025,5),(10026,'บุญ',20026,5),(10027,'โบ',20027,5),(10028,'เบญจมินทร์',20028,5),(10029,'ชนา',20029,6),(10030,'จันทรา',20030,6),(10031,'ชายา',20031,6),(10032,'แดน',20032,6),(10033,'ดานิกา',20033,6),(10034,'ดาวิน',20034,6),(10035,'ดอน',20035,6),(10036,'เอมมาริน',20036,6),(10037,'ฟิน',20037,7),(10038,'จีน',20038,7),(10039,'มะนาว',20039,7),(10040,'กิ๊บ',20040,7),(10041,'กาย',20041,7),(10042,'จีน่า',20042,7),(10043,'แอน',20043,7),(10044,'ไอรา',20044,8),(10045,'เจนนี่',20045,8),(10046,'จีซู',20046,8),(10047,'โรเซ่',20047,8),(10048,'ลิซ่า',20048,8),(10049,'จอน',20049,8),(10050,'ฮยอนอา',20050,8),(10051,'กะทิ',20051,9),(10052,'กิต',20052,9),(10053,'เคน',20053,9),(10054,'บอม',20054,9),(10055,'เฟิร์ส',20055,9),(10056,'ไลลา',20056,9),(10057,'หลิน',20057,9),(10058,'มลิกา',20058,9),(10059,'มาริสา',20059,10),(10060,'มานี',20060,10),(10061,'มะลิ',20061,10),(10062,'มิรา',20062,10),(10063,'นานา',20063,10),(10064,'เนตร',20064,10),(10065,'พลอย',20065,10),(10066,'แพรว',20066,11),(10067,'โอม',20067,11),(10068,'พัด',20068,11),(10069,'กาน',20069,11),(10070,'เล็ก',20070,11),(10071,'กอล์ฟ',20071,11),(10072,'รุธ',20072,11),(20001,'ชื่อของหัวหน้าทีมอะธีน่า',NULL,1),(20002,'ชื่อของหัวหน้าทีมโพไซดอน',NULL,1),(20003,'ชื่อของหัวหน้าทีมอริสตา',NULL,1),(20004,'ชื่อของหัวหน้าทีมอาบิเกล',NULL,1),(20005,'ชื่อของหัวหน้าทีมเอเรีย',NULL,1),(20006,'ชื่อของหัวหน้าทีมเอมมี่',NULL,1),(20007,'ชื่อของหัวหน้าทีมอลิสสา',NULL,1),(20008,'ชื่อของหัวหน้าทีมแองเจล่า',NULL,1),(20009,'ชื่อของหัวหน้าทีมเอมเบอร์',NULL,1),(20010,'ชื่อของหัวหน้าทีมอมีเรีย',NULL,1),(20011,'ชื่อของหัวหน้าทีมเอมี่',NULL,1),(20012,'ชื่อของหัวหน้าทีมเอวา',NULL,1),(20013,'ชื่อของหัวหน้าทีมแอนน์',NULL,1),(20014,'ชื่อของหัวหน้าทีมอลิซ',NULL,1),(20015,'ชื่อของหัวหน้าทีมเอนิซซา',NULL,1),(20016,'ชื่อของหัวหน้าทีมแอนเดรีย',NULL,1),(20017,'ชื่อของหัวหน้าทีมจตุรภัทร',NULL,1),(20018,'ชื่อของหัวหน้าทีมจรีภรณ์',NULL,1),(20019,'ชื่อของหัวหน้าทีมกชนันท์',NULL,1),(20020,'ชื่อของหัวหน้าทีมอณิกา',NULL,1),(20021,'ชื่อของหัวหน้าทีมอลินดา',NULL,1),(20022,'ชื่อของหัวหน้าทีมไอริน',NULL,1),(20023,'ชื่อของหัวหน้าทีมอานนท์',NULL,1),(20024,'ชื่อของหัวหน้าทีมอาชา',NULL,1),(20025,'ชื่อของหัวหน้าทีมบอน',NULL,1),(20026,'ชื่อของหัวหน้าทีมบุญ',NULL,1),(20027,'ชื่อของหัวหน้าทีมโบ',NULL,1),(20028,'ชื่อของหัวหน้าทีมเบญจมินทร์',NULL,1),(20029,'ชื่อของหัวหน้าทีมชนา',NULL,1),(20030,'ชื่อของหัวหน้าทีมจันทรา',NULL,1),(20031,'ชื่อของหัวหน้าทีมชายา',NULL,1),(20032,'ชื่อของหัวหน้าทีมแดน',NULL,1),(20033,'ชื่อของหัวหน้าทีมดานิกา',NULL,1),(20034,'ชื่อของหัวหน้าทีมดาวิน',NULL,1),(20035,'ชื่อของหัวหน้าทีมดอน',NULL,1),(20036,'ชื่อของหัวหน้าทีมเอมมาริน',NULL,1),(20037,'ชื่อของหัวหน้าทีมฟิน',NULL,1),(20038,'ชื่อของหัวหน้าทีมจีน',NULL,1),(20039,'ชื่อของหัวหน้าทีมมะนาว',NULL,1),(20040,'ชื่อของหัวหน้าทีมกิ๊บ',NULL,1),(20041,'ชื่อของหัวหน้าทีมกาย',NULL,1),(20042,'ชื่อของหัวหน้าทีมจีน่า',NULL,1),(20043,'ชื่อของหัวหน้าทีมแอน',NULL,1),(20044,'ชื่อของหัวหน้าทีมไอรา',NULL,1),(20045,'ชื่อของหัวหน้าทีมเจนนี่',NULL,1),(20046,'ชื่อของหัวหน้าทีมจีซู',NULL,1),(20047,'ชื่อของหัวหน้าทีมโรเซ่',NULL,1),(20048,'ชื่อของหัวหน้าทีมลิซ่า',NULL,1),(20049,'ชื่อของหัวหน้าทีมจอน',NULL,1),(20050,'ชื่อของหัวหน้าทีมฮยอนอา',NULL,1),(20051,'ชื่อของหัวหน้าทีมกะทิ',NULL,1),(20052,'ชื่อของหัวหน้าทีมกิต',NULL,1),(20053,'ชื่อของหัวหน้าทีมเคน',NULL,1),(20054,'ชื่อของหัวหน้าทีมบอม',NULL,1),(20055,'ชื่อของหัวหน้าทีมเฟิร์ส',NULL,1),(20056,'ชื่อของหัวหน้าทีมไลลา',NULL,1),(20057,'ชื่อของหัวหน้าทีมหลิน',NULL,1),(20058,'ชื่อของหัวหน้าทีมมลิกา',NULL,1),(20059,'ชื่อของหัวหน้าทีมมาริสา',NULL,1),(20060,'ชื่อของหัวหน้าทีมมานี',NULL,1),(20061,'ชื่อของหัวหน้าทีมมะลิ',NULL,1),(20062,'ชื่อของหัวหน้าทีมมิรา',NULL,1),(20063,'ชื่อของหัวหน้าทีมนานา',NULL,1),(20064,'ชื่อของหัวหน้าทีมเนตร',NULL,1),(20065,'ชื่อของหัวหน้าทีมพลอย',NULL,1),(20066,'ชื่อของหัวหน้าทีมแพรว',NULL,1),(20067,'ชื่อของหัวหน้าทีมโอม',NULL,1),(20068,'ชื่อของหัวหน้าทีมพัด',NULL,1),(20069,'ชื่อของหัวหน้าทีมกาน',NULL,1),(20070,'ชื่อของหัวหน้าทีมเล็ก',NULL,1),(20071,'ชื่อของหัวหน้าทีมกอล์ฟ',NULL,1),(20072,'ชื่อของหัวหน้าทีมรุธ',NULL,1),(29001,'ลูกทีมของคุณชนา คนที่หนึ่ง',10029,6),(29002,'ลูกทีมของคุณชนา คนที่สอง',10029,6),(29003,'ลูกทีมของคุณจันทรา คนที่หนึ่ง',10030,6),(29004,'ลูกทีมของคุณจันทรา คนที่สอง',10030,6),(29005,'ลูกทีมของคุณชายา คนที่หนึ่ง',10031,6),(29006,'ลูกทีมของคุณชายา คนที่สอง',10031,6),(29007,'ลูกทีมของคุณแดน คนที่หนึ่ง',10032,6),(29008,'ลูกทีมของคุณแดน คนที่สอง',10032,6),(29009,'ลูกทีมของคุณดานิกา คนที่หนึ่ง',10033,6),(29010,'ลูกทีมของคุณดานิกา คนที่สอง',10033,6),(29011,'ลูกทีมของคุณดาวิน คนที่หนึ่ง',10034,6),(29012,'ลูกทีมของคุณดาวิน คนที่สอง',10034,6),(29013,'ลูกทีมของคุณดอน คนที่หนึ่ง',10035,6),(29014,'ลูกทีมของคุณดอน คนที่สอง',10035,6),(29015,'ลูกทีมของคุณเอมมาริน คนที่หนึ่ง',10036,6),(29016,'ลูกทีมของคุณเอมมาริน คนที่สอง',10036,6),(30001,'กานกนก',20072,6),(30002,'ขนมโมจิ',30001,4),(30003,'แพรวนภา',30001,4),(99999,'TEST_ID',NULL,1);
/*!40000 ALTER TABLE `members` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `point_records`
--

DROP TABLE IF EXISTS `point_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `point_records` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `member_id` int(11) DEFAULT NULL,
  `point` int(11) DEFAULT NULL,
  `create_date` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `member_id_fk_idx` (`member_id`),
  CONSTRAINT `member_id_fk_point_record` FOREIGN KEY (`member_id`) REFERENCES `members` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=210 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `point_records`
--

LOCK TABLES `point_records` WRITE;
/*!40000 ALTER TABLE `point_records` DISABLE KEYS */;
INSERT INTO `point_records` VALUES (1,10029,650,'2019-05-01 04:27:19'),(2,10029,350,'2019-07-25 04:27:49'),(3,29001,9500,'2019-07-25 04:31:43'),(4,29002,9500,'2019-07-25 04:31:50'),(5,10009,100,'2019-07-25 06:47:18'),(6,30001,4000,'2019-07-25 09:10:42'),(7,30002,10000,'2019-07-25 09:10:42'),(8,30003,10000,'2019-07-25 09:10:42'),(9,10030,650,'2019-05-01 04:27:19'),(10,10030,350,'2019-07-25 04:27:49'),(11,29003,9500,'2019-07-25 04:27:49'),(12,29004,9500,'2019-07-25 04:27:49'),(13,10031,650,'2019-05-01 04:27:19'),(14,10031,350,'2019-07-25 04:27:49'),(15,29005,12000,'2019-07-25 04:27:49'),(16,29006,12000,'2019-07-25 04:27:49'),(17,10032,650,'2019-05-01 04:27:19'),(18,10032,350,'2019-07-25 04:27:49'),(19,29007,12000,'2019-07-25 04:27:49'),(20,29008,12000,'2019-07-25 04:27:49'),(21,10033,650,'2019-05-01 04:27:19'),(22,10033,350,'2019-07-25 04:27:49'),(23,29009,7000,'2019-07-25 04:27:49'),(24,29010,7000,'2019-07-25 04:27:49'),(25,10034,650,'2019-05-01 04:27:19'),(26,10034,350,'2019-07-25 04:27:49'),(27,29011,9000,'2019-07-25 04:27:49'),(28,29012,9950,'2019-07-25 04:27:49'),(29,10035,650,'2019-05-01 04:27:19'),(30,10035,350,'2019-07-25 04:27:49'),(31,29013,7000,'2019-07-25 04:27:49'),(32,29014,7000,'2019-07-25 04:27:49'),(33,10036,650,'2019-05-01 04:27:19'),(34,10036,350,'2019-07-25 04:27:49'),(35,29015,9000,'2019-07-25 04:27:49'),(36,29016,10000,'2019-07-25 04:27:49'),(205,10001,550,'2019-08-01 03:28:22'),(206,10002,550,'2019-08-01 03:28:22');
/*!40000 ALTER TABLE `point_records` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-08-01 13:31:32
