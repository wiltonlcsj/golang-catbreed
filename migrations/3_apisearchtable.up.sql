CREATE TABLE IF NOT EXISTS `hgbackend`.`api_search` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `query_param` VARCHAR(255) NOT NULL,
  `api_status` INT(3) NOT NULL,
  `response` TEXT NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
