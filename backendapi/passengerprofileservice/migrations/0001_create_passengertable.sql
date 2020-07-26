-- -----------------------------------------------------
-- Schema cab
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `cab` DEFAULT CHARACTER SET utf8 ;
USE `cab` ;

-- -----------------------------------------------------
-- Table `cab`.`passenger`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cab`.`passenger` (
  `id` INT NOT NULL,
  `email` VARCHAR(45) NOT NULL,
  `name` VARCHAR(45) NULL,
  `mobileno` VARCHAR(12) NOT NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `mobileno_UNIQUE` (`mobileno` ASC),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC)
);
