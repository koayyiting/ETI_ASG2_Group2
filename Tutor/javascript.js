const sign_in_btn = document.querySelector("#sign-in-btn");
const sign_up_btn = document.querySelector("#sign-up-btn");
const container = document.querySelector(".container");

sign_up_btn.addEventListener("click", () => {
  container.classList.add("sign-up-mode");
});

sign_in_btn.addEventListener("click", () => {
  container.classList.remove("sign-up-mode");
});

//tutor features
function tutorSignup(){
    var request = new XMLHttpRequest();
    const form = document.getElementById('tutorSignupForm');

    const curl = 'http://localhost:5211/api/v1/tutor';

    const tutorEmail = form.elements['tutor_signup_email'].value;
    const tutorPassword = form.elements['tutor_signup_password'].value;
    const tutorFirstName = form.elements['tutor_signup_firstname'].value;
    const tutorLastName = form.elements['tutor_signup_lastname'].value;

    console.log(tutorEmail);
    console.log(tutorPassword);

    request.open("POST", curl);
    request.send(JSON.stringify({
        "tutorEmail": tutorEmail,
        "tutorPassword": tutorPassword, 
        "tutorFirstName": tutorFirstName,
        "tutorLastName": tutorLastName
        
    }));
    form.reset();
    alert("Tutor account registered successfuly.");
    return false
}

function tutorLogin(){
  var request = new XMLHttpRequest();
  const form = document.getElementById('tutorLoginForm');
  const tutorEmail = form.elements['tutor_login_email'].value;
  const tutorPassword = form.elements['tutor_login_password'].value;
  console.log(tutorEmail);
  console.log(tutorPassword);

  const curl = 'http://localhost:5211/api/v1/tutor?tutorEmail=' + encodeURIComponent(tutorEmail) + '&tutorPassword=' + encodeURIComponent(tutorPassword);
  console.log(curl);

  request.open("GET", curl);
  request.onreadystatechange = function() {
    if (request.readyState === 4) {
      if (request.status === 200) {
        const responseData = JSON.parse(request.responseText);
        console.log(responseData);

        // Save user information to localStorage
        localStorage.setItem('tutorId', responseData.tutorId);
        localStorage.setItem('tutorFirstName', responseData.tutorFirstName);
        localStorage.setItem('tutorLastName', responseData.tutorLastName);
        localStorage.setItem('tutorEmail', responseData.tutorEmail);

        console.log('localStorage:', localStorage);

        // Successful login, redirect to main page
        location.href = "tutor_main.html";
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