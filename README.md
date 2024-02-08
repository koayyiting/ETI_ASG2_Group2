# ETI-ASG2
## Website Description
FrEd (Free Education) is your one-stop solution for accessible, high-quality education for all. 
Our platform empowers volunteers to become tutors, where they can create lesson materials and schedule sessions effortlessly. 
Students can easily sign up, browse available lessons, and book sessions at their convenience. With a shared discussion forum, 
both tutors and students can engage, ask questions, and build a supportive learning community. Join us in advancing 
Sustainable Development Goal 4 by making education accessible to everyone, regardless of background or location.

## Design Consideration of our microservices
1. Student Microservice
The Student Microservice is designed to handle a student's information to verify their identity.
A registered student will be able to access various utilities of the site,
such as communicating with different users in the discussion forum or booking a lesson.

2. Tutor Microservice
The Tutor Microservice handles the creation of new tutor accounts and is designed to retrieve
tutor information based on user input to verify their identity. It utilizes the tutor table in the
'tutor_db' MySQL database to store tutor-related data. A registered tutor will be able to access various
features of the site, such as creating lesson materials, scheduling lessons, and communicating with different users in the discussion forum.

3. Schedule Microservice

4. Booking Microservice

5. Lesson Material Microservice

6. Discussion Microservice
As for the discussion form you are able to log in as a user and can post question edit the question and
delete the question and once the user logs out another user can log in and comment on the post made by other previous owner.


## Architecture Diagram
![ETI ASG2 Architecture Diagram](https://github.com/koayyiting/ETI_ASG2_Group2/assets/93900155/bc41ca39-34f8-4bd5-8807-3452f053174b)
   - Each microservice has its own MySQL database

## Instructions for setting up and running our solution
1. Database Setup
   - Execute the SQL script 'studentDB.sql' to create the required database and table (Student)
   - Execute the SQL script 'tutor_db.sql' to create the required database and table (Tutor)
   - Execute the SQL script 'schedule_db_setup.sql' to create the required database and table (Schedule)
   - Execute the SQL script 'booking_db_setup.sql' to create the required database and table (Booking)
   - Execute the SQL script 'LessonMaterial.sql' to create the required database and table (LessonMaterials)
   - Execute the SQL script 'database.sql' to create the required database and table (Post, Comment)

2. Microservices Setup
   - After cloning the repository, navigate to the root directory of each microservice and run 'go run Student.go', 'go run tutor.go', 'go run schedule.go', 'go run book.go', 'go run LM.go', 'go run main.go'.
   - Student microservice runs on port 5212, Tutor microservice runs on port 5211, Schedule microservice runs on port 1000, Booking microservice runs on port 1765, Lesson Material microservice runs on port 4088, Discussion microservice runs on port 4090.

3. User Interface Setup
   - From Visual Studio Code, select "Go Live" from any HTML page to start using the respective microservice
   - Or use these links to directly access each page:
     - Tutor Login/Signup page (https://koayyiting.github.io/ETI_ASG2_Group2/)
     - Student Login/Signup page (https://koayyiting.github.io/ETI_ASG2_Group2/Student/student_signup_login.html)
     - Schedule (https://koayyiting.github.io/ETI_ASG2_Group2/Schedule/user_schedules.html)
     - Booking (https://koayyiting.github.io/ETI_ASG2_Group2/Book/user_bookings.html)
     - Learning Materials (https://koayyiting.github.io/ETI_ASG2_Group2/Learning%20Materials/LM.html)
     - Discussion Page (can only access through localhost)
