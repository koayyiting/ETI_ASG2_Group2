Create database student;

use student;
Create table student (studentID INT NOT NULL PRIMARY KEY, studentFirstName Varchar(50), studentLastName varchar(30), phoneNo varchar(20), studentEmail varchar (50), studentPassword varchar (20));


INSERT INTO student VALUES (1, 'Salman', 'Khan', '88237164', 'salmon@gmail.com', 'b0ll1vard'),
(2, 'Ting Wong', 'Sum', '88732192', 'sumtw@hotmail.com', 'wR0ngW4y'),
(3, 'Daniel', 'Middleton', '90293124', 'danmdt@hotmail.com', 'tr4yaURu5'),
(4, 'Jacky', 'Lee', '86234953', 'bruceno1@gmail.com', 'hw4CH4A');

