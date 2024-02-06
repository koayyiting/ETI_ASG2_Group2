-- CREATE USER 'etiLessonMaterial'@'localhost' IDENTIFIED BY 'eti2024';
-- GRANT ALL ON *.* TO 'etiLessonMaterial'@'localhost';

-- [!!] Uncomment next statement for initial execution of database
-- CREATE database lessonmaterial_db;

USE lessonmaterial_db;

DROP TABLE IF EXISTS LessonMaterials;

CREATE TABLE LessonMaterials (
LMID VARCHAR(5) NOT NULL PRIMARY KEY,
TutorID	INT NOT NULL,
Topic VARCHAR(100) NULL,
Summary VARCHAR(400) NULL,
Created DATETIME NULL 
);

INSERT INTO LessonMaterials (LMID, TutorID, Topic, Summary, Created)
VALUES("1","1", "Basic Go Programming Language", "Go, often referred to as Golang, is a statically-typed programming language developed by Google. It was created to address issues like slow compilation times and complex dependencies in other languages.", "2023-01-02 15:04:05"),
	("2","1", "Basic Python Programming Language", "Python's easy-to-read syntax, resembling plain English, makes it perfect for beginners and experienced programmers alike. It also boasts concise code, often requiring fewer lines compared to other languages, saving you time and effort.", "2023-11-21 15:32:47"),
	("3","1", "Basic C# Programming Language", "C# is a modern language developed by Microsoft for building various applications, known for its strong typing, automatic garbage collection, and object-oriented programming support. It's often used with the .NET framework and is highly integrated with the Windows operating system.", "2024-06-08 09:15:02"),
    ("4","2","Basic C++ Programming Language","The C programming language, created by Dennis Ritchie in 1972, is a powerful and efficient language used for system programming, embedded systems, and application development. It provides low-level access to memory and hardware, making it suitable for developing operating systems and other performance-critical software.","2024-06-05 09:15:02"),
    ("5","2","Basic C Programming Language","The C programming language, created by Dennis Ritchie in 1972, is a powerful and widely used language for system programming, embedded systems, and application development. Known for its efficiency and flexibility, C has influenced many other programming languages and remains an important tool for software development.","2024-04-27 16:44:50"),
    ("6","2","DevOps Fundamentals","DevOps is a software development approach that combines development (Dev) and operations (Ops) to improve collaboration and automation. It emphasizes communication, integration, automation, and measurement to streamline the software delivery process and ensure high-quality, rapid deployment of applications.","2023-11-21 15:32:47")
