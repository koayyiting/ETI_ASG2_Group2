const login_btn = document.querySelector("#login-btn");
const register_btn = document.querySelector("#register-btn");
const container = document.querySelector(".container");

register_btn.addEventListener("click", () => {
  container.classList.add("register-mode");
});

login_btn.addEventListener("click", () => {
  container.classList.remove("register-mode");
});

//student features
function studentRegister(){
    var request = new XMLHttpRequest();
    const form = document.getElementById('studentRegisterForm');

    const curl = 'http://localhost:3306/api/v1/student';

    const studentEmail = form.elements['student_register_email'].value;
    const studentPassword = form.elements['student_register_password'].value;
    const studentFirstName = form.elements['student_register_firstname'].value;
    const studentLastName = form.elements['student_register_lastname'].value;

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

    const curl = 'http://localhost:3306/api/v1/student?StudentEmail=' + encodeURIComponent(StudentEmail) + '&StudentPassword=' + encodeURIComponent(StudentPassword);
    console.log(curl);

    request.open("GET", curl);
    request.onreadystatechange = function() {
      if (request.readyState === 4) {
        if (request.status === 200) {
          // Successful login, redirect to main page
          location.href = "studentmain.html";
        } else if (request.status === 401) {
          // Login failed, handle error
          form.reset();
          document.getElementById('error-message').innerHTML = 'Incorrect Email or Password.';
        } else {
          // Handle other status codes or network errors
          document.getElementById('error-message').innerHTML = 'Error encountered. Please try again later.';
        }
      }
    };
    request.send();
    return false
}

