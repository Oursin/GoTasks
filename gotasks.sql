DROP DATABASE IF EXISTS `gotasks`;

CREATE DATABASE `gotasks`;

USE `gotasks`;

CREATE TABLE `user` (
  `user_id` INT NOT NULL AUTO_INCREMENT,
  `username` TEXT NOT NULL,
  `password`  TEXT NOT NULL,
  `email` TEXT NOT NULL,
  `first name` TEXT NOT NULL,
  `name` TEXT NOT NULL,
  PRIMARY KEY (user_id),
);

CREATE TABLE `project` (
  `project_id` INT NOT NULL AUTO_INCREMENT,
  `title` TEXT NOT NULL,
  `desc` TEXT NOT NULL,
  PRIMARY KEY (project_id),
);

CREATE TABLE `task` (
  `task_id` INT NOT NULL AUTO_INCREMENT,
  `title`   TEXT NOT NULL,
  `begin`   DATETIME default CURRENT_TIMESTAMP,
  `end`     DATETIME,
  `status`  VARCHAR(12) DEFAULT 'not started',
  `project_id` INT NOT NULL,
  `creator_id` INT NOT NULL,
  PRIMARY KEY (`task_id`),
  FOREIGN KEY (project_id) REFERENCES project(project_id) ON DELETE CASCADE,
  FOREIGN KEY (creator_id) REFERENCES user(user_id),
);

CREATE TABLE `user_has_task` (
  `fk_user_id` INT NOT NULL,
  `fk_task_id` INT NOT NULL,
  PRIMARY KEY (fk_user_id, fk_task_id),
  FOREIGN KEY (fk_user_id) REFERENCES user(user_id),
  FOREIGN KEY (fk_task_id) REFERENCES task(task_id),
)

CREATE TABLE `user_has_project` (
  `fk_user_id` INT NOT NULL,
  `fk_project_id` INT NOT NULL,
  PRIMARY KEY (fk_project_id, fk_user_id),
  FOREIGN KEY (fk_user_id) REFERENCES user(user_id),
  FOREIGN KEY (fk_project_id) REFERENCES project(project_id),
)

CREATE TABLE `project_has_task` (
  `fk_project_id` INT NOT NULL,
  `fk_task_id` INT NOT NULL,
  PRIMARY KEY (fk_project_id, fk_task_id),
  FOREIGN KEY (fk_project_id) REFERENCES project(project_id),
  FOREIGN KEY (fk_task_id) REFERENCES task(task_id),
)