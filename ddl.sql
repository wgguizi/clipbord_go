CREATE TABLE `c_data` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `code` char(32) NOT NULL DEFAULT '',
  `content` mediumtext,
  `ip` char(15) NOT NULL DEFAULT '',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`),
  KEY `created_at` (`created_at`)
) ENGINE=MyISAM AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;

