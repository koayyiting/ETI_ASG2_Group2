//tutor features
function tutorSignup(){
    var request = new XMLHttpRequest();
    const form = document.getElementById('tutorSignupForm');

    const curl = 'http://localhost:5211/api/v1/tutor';

    const tutorUsername = form.elements['tutor_signup_username'].value;
    const tutorPassword = form.elements['tutor_signup_password'].value;
    const tutorTitle = form.elements['tutor_signup_title'].value;
    const tutorFirstName = form.elements['tutor_signup_firstname'].value;
    const tutorLastName = form.elements['tutor_signup_lastname'].value;

    console.log(tutorUsername);
    console.log(tutorPassword);

    request.open("POST", curl);
    request.send(JSON.stringify({
        "tutorUsername": tutorUsername,
        "tutorPassword": tutorPassword, 
        "tutorTitle": tutorTitle,
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
    const tutorUsername = form.elements['tutor_login_username'].value;
    const tutorPassword = form.elements['tutor_login_password'].value;
    console.log(tutorUsername);
    console.log(tutorPassword);

    const curl = 'http://localhost:5211/api/v1/tutor?tutorUsername=' + encodeURIComponent(tutorUsername) + '&tutorPassword=' + encodeURIComponent(tutorPassword);
    console.log(curl);

    request.open("GET", curl);
    request.onreadystatechange = function() {
      if (request.readyState === 4) {
        if (request.status === 200) {
          // Successful login, redirect to main page
          location.href = "tutor_main.html";
        } else if (request.status === 401) {
          // Login failed, handle error
          form.reset();
          document.getElementById('error-message').innerHTML = 'Incorrect Username or Password.';
        } else {
          // Handle other status codes or network errors
          document.getElementById('error-message').innerHTML = 'An error occurred. Please try again later.';
        }
      }
    };
    request.send();
    return false
}

