DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `first_name` varchar(45) CHARACTER SET utf8 COLLATE utf8_spanish2_ci DEFAULT NULL,
  `last_name` varchar(45) CHARACTER SET utf8 COLLATE utf8_spanish2_ci DEFAULT NULL,
  `email` varchar(45) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL,
  `date_created` datetime DEFAULT NULL,
  `status` varchar(45) CHARACTER SET utf8 COLLATE utf8_spanish2_ci DEFAULT NULL,
  `password` varchar(45) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;