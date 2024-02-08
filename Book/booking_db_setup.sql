-- CREATE USER 'booking_system'@'localhost' IDENTIFIED BY 'ETI_Group2_Booking';
-- GRANT ALL ON *.* TO 'booking_system'@'localhost';
USE ETI_Booking;
SELECT User, Host FROM mysql.user WHERE User = 'booking_system' AND Host = 'localhost';
-- CREATE database ETI_Booking;

drop table Booking
-- CREATE TABLE Booking (
-- BookingID INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
-- StudentID INT NOT NULL, 
-- ScheduleID INT NOT NULL,
-- LessonID varchar(100)
-- );

-- insert into Booking (StudentName, StudentEmail, LessonID) values("Zi Yi","test","1");
-- DELETE FROM Booking WHERE ID = 1;

-- ALTER TABLE Booking AUTO_INCREMENT = 1; 

select * from Booking;
