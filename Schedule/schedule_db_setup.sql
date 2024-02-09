-- CREATE USER 'schedule_system'@'localhost' IDENTIFIED BY 'ETI_Group2_Schedule';
-- GRANT ALL ON *.* TO 'schedule_system'@'localhost';

-- SELECT User, Host FROM mysql.user WHERE User = 'schedule_system' AND Host = 'localhost';
CREATE DATABASE IF NOT EXISTS ETI_Schedule;
USE ETI_Schedule;

-- drop table Schedule;
CREATE TABLE IF NOT EXISTS Schedule (
    ScheduleID INT AUTO_INCREMENT PRIMARY KEY,
    TutorID INT NOT NULL,
    LessonID INT NOT NULL,
    LessonName varchar(100),
    Location varchar(1000),
    StartTime datetime NOT NULL,
    EndTime datetime NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- , 
);

-- insert into Booking (StudentName, StudentEmail, LessonID) values("Zi Yi","test","1");
-- DELETE FROM Booking WHERE ID = 1;

-- INSERT INTO Schedule (TutorID, LessonID, StartTime, EndTime)
-- VALUES
--     (1, 1, '2022-12-01 08:00:00', '2022-12-01 10:00:00'),
--     (2, 2, '2022-12-02 13:30:00', '2022-12-02 15:30:00'),
--     (3, 3, '2022-12-03 10:45:00', '2022-12-03 12:45:00'),
--     (3, 4, '2022-12-04 11:45:00', '2022-12-04 13:45:00');

-- SELECT TutorID, LessonID, StartTime, EndTime FROM Schedule
-- DELETE FROM Schedule WHERE ScheduleID = 16;
-- DELETE FROM Schedule WHERE ScheduleID = 17;
-- DELETE FROM Schedule WHERE ScheduleID = 18;
-- UPDATE Schedule SET StartTime="2022-12-04 11:40:00", EndTime="2022-12-04 13:45:00" WHERE ScheduleID=4;
SELECT * FROM Schedule