CREATE USER 'student'@'localhost' IDENTIFIED BY 'etistudentpwd';
GRANT ALL ON *.* TO 'student'@'localhost';

CREATE DATABASE IF NOT EXISTS student_db DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE student_db;

CREATE TABLE IF NOT EXISTS Student (
StudentID int NOT NULL AUTO_INCREMENT,
FirstName varchar (50) NOT NULL,
LastName varchar (50) NOT NULL,
Email varchar (50) NOT NULL,
Password varchar (50) NOT NULL,
PRIMARY KEY (`StudentID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

INSERT INTO Student (StudentID, FirstName, LastName, Email, Password)
VALUES(1, 'Yi Ting', 'Koay', 'yiting@gmail.com', '123456');

SELECT * FROM Student;