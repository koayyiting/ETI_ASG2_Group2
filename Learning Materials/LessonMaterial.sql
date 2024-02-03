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
Summary VARCHAR(9000) NULL,
Created DATETIME NULL 
);

INSERT INTO LessonMaterials (LMID, TutorID, Topic, Summary, Created)
VALUES("1","1", "Basic Go Programming Language", "Go, often referred to as Golang, is a statically-typed programming language developed by Google. It was created to address issues like slow compilation times and complex dependencies in other languages.", "2023-01-02 15:04:05")
