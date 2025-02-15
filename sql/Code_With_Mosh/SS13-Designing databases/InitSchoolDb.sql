CREATE SCHEMA IF NOT EXISTS `school` DEFAULT CHARACTER SET utf8;
USE `school`;

CREATE TABLE IF NOT EXISTS `school`.`students` (
    student_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    date_registered DATETIME NOT NULL
)
ENGINE = InnoDB;

CREATE TABLE `school`.`instructors` (
    instructor_id SMALLINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL
)
ENGINE = InnoDB;


CREATE TABLE `school`.`courses` (
    course_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    price DECIMAL(5,2) NOT NULL,
    instructor VARCHAR(255) NOT NULL,
    tags VARCHAR(255) NOT NULL,
    instructor_id SMALLINT NOT NULL,
    FOREIGN KEY (instructor_id) REFERENCES `school`.`instructors`(instructor_id)
)
ENGINE = InnoDB;

CREATE TABLE `school`.`enrollments` (
    course_id INT NOT NULL,
    student_id INT NOT NULL,
    date DATETIME NOT NULL,
    price DECIMAL(5,2) NOT NULL,
    coupon VARCHAR(50),
    PRIMARY KEY (course_id, student_id),
    FOREIGN KEY (course_id) REFERENCES `school`.`courses`(course_id),
    FOREIGN KEY (student_id) REFERENCES `school`.`students`(student_id)
)
ENGINE = InnoDB;

ALTER TABLE `school`.`enrollments` ADD COLUMN coupon VARCHAR(50);

CREATE TABLE `school`.`tags` (
    tag_id TINYINT NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
)
ENGINE = InnoDB;

CREATE TABLE `school`.`course_tags` (
    course_id INT NOT NULL,
    tag_id TINYINT NOT NULL,
    PRIMARY KEY (course_id, tag_id),
    FOREIGN KEY (course_id) REFERENCES `school`.`courses`(course_id),
    FOREIGN KEY (tag_id) REFERENCES `school`.`tags`(tag_id)
)
ENGINE = InnoDB;