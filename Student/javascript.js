const sign_in_btn = document.querySelector("#sign-in-btn");
const sign_up_btn = document.querySelector("#sign-up-btn");
const container = document.querySelector(".container");

sign_up_btn.addEventListener("click", () => {
  container.classList.add("sign-up-mode");
});

sign_in_btn.addEventListener("click", () => {
  container.classList.remove("sign-up-mode");
});

//student features
function studentSignup(){
    var request = new XMLHttpRequest();
    const form = document.getElementById('studentSignupForm');

    const curl = 'http://localhost:5212/api/v1/student';

    const studentEmail = form.elements['student_signup_email'].value;
    const studentPassword = form.elements['student_signup_password'].value;
    const studentFirstName = form.elements['student_signup_firstname'].value;
    const studentLastName = form.elements['student_signup_lastname'].value;

    console.log(studentEmail);
    console.log(studentPassword);

    request.open("POST", curl);
    request.send(JSON.stringify({
        "studentEmail": studentEmail,
        "studentPassword": studentPassword, 
        "studentFirstName": studentFirstName,
        "studentLastName": studentLastName
        
    }));
    form.reset();
    alert("Student account registered successfuly.");
    return false
}

function studentLogin(){
  var request = new XMLHttpRequest();
  const form = document.getElementById('studentLoginForm');
  const studentEmail = form.elements['student_login_email'].value;
  const studentPassword = form.elements['student_login_password'].value;
  console.log(studentEmail);
  console.log(studentPassword);

  const curl = 'http://localhost:5212/api/v1/student?studentEmail=' + encodeURIComponent(studentEmail) + '&studentPassword=' + encodeURIComponent(studentPassword);
  console.log(curl);

  request.open("GET", curl);
  request.onreadystatechange = function() {
    if (request.readyState === 4) {
      if (request.status === 200) {
        const responseData = JSON.parse(request.responseText);
        console.log(responseData);

        // Save user information to localStorage
        localStorage.setItem('studentId', responseData.studentId);
        localStorage.setItem('studentFirstName', responseData.studentFirstName);
        localStorage.setItem('studentLastName', responseData.studentLastName);
        localStorage.setItem('studentEmail', responseData.studentEmail);

        console.log('localStorage:', localStorage);

        // Successful login, redirect to main page
        location.href = "../Book/user_bookings.html";
      } else if (request.status === 401) {
        // Login failed, handle error
        form.reset();
        document.getElementById('error-message').innerHTML = 'Incorrect Email or Password.';
      } else {
        // Handle other status codes or network errors
        document.getElementById('error-message').innerHTML = 'An error occurred. Please try again later.';
      }
    }
  };
  request.send();
  return false
}