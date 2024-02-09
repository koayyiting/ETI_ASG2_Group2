CREATE TABLE user (
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) PRIMARY KEY NOT NULL 
   
);

CREATE TABLE post (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    content VARCHAR(200) NOT NULL,
    email VARCHAR(100)  NOT NULL ,
    FOREIGN KEY (email) REFERENCES user(email)
);

CREATE TABLE comment (
    id INT AUTO_INCREMENT PRIMARY KEY,
    content VARCHAR(200) NOT NULL,
    postid INT  ,
    email VARCHAR(100)  NOT NULL ,
    FOREIGN KEY (postid) REFERENCES post(id),
    FOREIGN KEY (email) REFERENCES user(email)
);
