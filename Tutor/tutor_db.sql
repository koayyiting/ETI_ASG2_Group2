CREATE USER 'tutor'@'localhost' IDENTIFIED BY 'etitutorpwd';
GRANT ALL ON *.* TO 'tutor'@'localhost';

CREATE DATABASE IF NOT EXISTS tutor_db DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE tutor_db;

CREATE TABLE IF NOT EXISTS Tutor (
TutorID int NOT NULL AUTO_INCREMENT,
FirstName varchar (50) NOT NULL,
LastName varchar (50) NOT NULL,
Email varchar (50) NOT NULL,
Password varchar (50) NOT NULL,
PRIMARY KEY (`TutorID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

INSERT INTO Tutor (TutorID, FirstName, LastName, Email, Password)
VALUES(1, 'Zi Yi', 'Ng', 'ziyi@gmail.com', 'tutor1');

SELECT * FROM Tutor;