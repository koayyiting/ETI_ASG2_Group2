//display all the schedules
function listSchedules() {
    // Make a GET request to the server endpoint
    const tutor_id = parseInt(localStorage.getItem('tutorId'));
    const url = `http://localhost:1000/api/v1/getSchedules/` + tutor_id;
    fetch(url)
      .then(response => {
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }
        return response.json();
      })
      .then(data => {
        console.log("Data from server:", data);
  
        // Get the table body element
        var tableBody = document.getElementById('schedules_table').getElementsByTagName('tbody')[0];
  
        // Clear existing rows
        tableBody.innerHTML = '';
  
        // Iterate through the received data and append rows to the table
        if (data != null){
            data.forEach(user => {
            var row = tableBody.insertRow();
            const formattedStartTime = new Date(user.start_time).toLocaleString('en-US', { 
                year: 'numeric', 
                month: '2-digit', 
                day: '2-digit',
                hour: 'numeric', 
                minute: '2-digit', 
                hour12: true 
            });
            
            const formattedEndTime = new Date(user.end_time).toLocaleString('en-US', { 
                year: 'numeric', 
                month: '2-digit', 
                day: '2-digit',
                hour: 'numeric', 
                minute: '2-digit', 
                hour12: true 
            });
            
            row.innerHTML = `<td>${user.lesson_name}</td>
                                <td>${formattedStartTime}</td>
                                <td>${formattedEndTime}</td>
                                <td>${user.location}</td><td>
                                    <button style="border:none; padding: 5px 15px; border-radius:10px; margin: 5px 5px;" onclick="return modifySchedule(${user.schedule_id}, '${user.lesson_id}', '${user.lesson_name}', '${user.location}', '${user.start_time}', '${user.end_time}')">Modify</button>
                                    <button style="border:none; padding: 5px 15px; border-radius:10px; margin: 5px 5px;" onclick="return deleteSchedule(${user.schedule_id})">Delete</button>
                                </td>`;
                            });
        } else{
            var emptyMessageDiv = document.getElementById('empty_message');
            emptyMessageDiv.innerHTML = `<p style="padding: 20px 0px">You have 0 Schedule</p>`;
        }
      })
      .catch(error => console.error('Error fetching user details:', error));
  }

//switch to create schedule form page
function switchToSchFormpage(){
    window.location.href = 'schedule_form.html';
}

//switch to index page
function indexPage(){
    window.location.href = 'index.html';
}

function create_schedule() {
    // Get tutor ID from localStorage
    const tutor_id = parseInt(localStorage.getItem('tutorId'));

    // Get form element
    const form = document.getElementById('createScheduleForm');

    // Validate form fields
    const lesson_id = parseInt(form.elements['lesson_id_create'].value, 10);
    const startTimeInput = document.getElementById('starttime_schedule_form').value;
    const endTimeInput = document.getElementById('endtime_schedule_form').value;
    const locationInput = document.getElementById('location_schedule_form').value;
    const topicName = form.elements['topic_name_create'].value;

    // Check if any of the required fields are empty
    if (!lesson_id || !startTimeInput || !endTimeInput || !locationInput) {
        alert('Please fill in all required fields.');
        return false; // Prevent form submission
    }

    // Format start and end times
    const startTime = startTimeInput.replace('T', ' ') + ':00';
    const endTime = endTimeInput.replace('T', ' ') + ':00';

    // Create XMLHttpRequest object
    var request = new XMLHttpRequest();
    const curl = 'http://localhost:1000/api/v1/schedule';

    // Set up request
    request.open("POST", curl);
    request.setRequestHeader('Content-Type', 'application/json');

    // Define onload and onerror handlers
    request.onload = function () {
        if (request.status >= 200 && request.status < 400) {
            console.log('Schedule created successfully');
            window.location.href = "user_schedules.html";
        } else {
            console.error('Error creating schedule:', request.statusText);
        }
    };

    request.onerror = function () {
        console.error('Network error');
    };

    // Send request with form data
    request.send(JSON.stringify({
        "tutor_id": tutor_id,
        "lesson_id": lesson_id,
        "start_time_str": startTime,
        "end_time_str": endTime, 
        "location": locationInput, 
        "lesson_name": topicName
    }));

    return false; // Prevent default form submission
}


function deleteSchedule(scheduleID) {
    console.log('Deleting Schedule with ID:', scheduleID);
    const url = `http://localhost:1000/api/v1/oneSchedule/${scheduleID}`;
  
    // Confirm deletion with the user (you can customize this)
    if (confirm("Are you sure you want to delete this Schedule?")) {
        // Make a DELETE request to the server endpoint
        fetch(url, {
          method: 'DELETE',
        })
        .then(response => {
          if (!response.ok) {
              throw new Error(`HTTP error! Status: ${response.status}`);
          }
          return response.text(); // assuming the server returns a text response
        })
        .then(data => {
          console.log('Server response:', data);
          
          // Optionally, you can call listUsers() again to refresh the user list
          listSchedules();
        })
        .catch(error => console.error('Error deleting Schedule:', error));
    }
}

function modifySchedule(schedule_id, lesson_id, lesson_name, location, start_time, end_time) {
    // Build the URL with parameters
    const params = new URLSearchParams({
        schedule_id_update: encodeURIComponent(schedule_id),
        lesson_id_update: encodeURIComponent(lesson_id),
        topic_name_update: encodeURIComponent(lesson_name),
        location_schedule_update: encodeURIComponent(location), // Since location is not provided
        starttime_update: encodeURIComponent(start_time),
        endtime_update: encodeURIComponent(end_time)
    });

    const url = `update_schedule.html?${params.toString()}`;

    // Redirect to the new page with parameters
    window.location.href = url;
}

async function update_schedule(sid) {
    console.log("in update_schedule")
    // Get form element
    const form = document.getElementById('updateScheduleForm');

    // Validate form fields
    const lesson_id = parseInt(form.elements['lesson_id_update'].value, 10);
    const startTimeInput = document.getElementById('starttime_schedule_update').value;
    const endTimeInput = document.getElementById('endtime_schedule_update').value;
    const locationInput = document.getElementById('location_schedule_update').value;
    const topicName = form.elements['topic_name_update'].value;

    // Check if any of the required fields are empty
    if (!lesson_id || !startTimeInput || !endTimeInput || !locationInput) {
        alert('Please fill in all required fields.');
        return false; // Prevent form submission
    }

    // Format start and end times
    const startTime = startTimeInput.replace('T', ' ') + ':00';
    const endTime = endTimeInput.replace('T', ' ') + ':00';

    // Create XMLHttpRequest object
    const url = `http://localhost:1000/api/v1/oneSchedule/`+sid;
    console.log(url);

    try {
        const response = await fetch(url, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "schedule_id_str": sid,
            "lesson_id_str": lesson_id,
            "start_time_str": startTime,
            "end_time_str": endTime, 
            "location": locationInput, 
            "lesson_name": topicName
        }),
        });
        console.log("test");
        if (response.ok) {
        console.log("Update successful");
        window.location.href = "user_schedules.html";
        } else {
        const errorText = await response.text();
        alert("Error updating the Schedule Details. Status: " + response.status + "\n" + errorText);
        }
    } catch (error) {
        console.error("Error updating the Schedule Details:", error);
        alert("An unexpected error occurred. Please try again later.");
    }
}

// async function update_schedule(sid) {
//     event.preventDefault();
//     console.log("in update_schedule")
//     console.log(sid)
//     const form = document.getElementById('updateScheduleForm');
    
//     const url = `http://localhost:1000/api/v1/oneSchedule/${sid}`;
    
//     const starttime_update = form.elements['starttime_update'].value;
//     const endtime_update = form.elements['endtime_update'].value;
    
//     try {
//         const response = await fetch(url, {
//         method: 'PUT',
//         headers: {
//             'Content-Type': 'application/json',
//         },
//         body: JSON.stringify({
//             "schedule_id_str": sid,
//             "start_time_str": starttime_update,
//             "end_time_str": endtime_update,
//         }),
//         });
    
//         if (response.ok) {
//         console.log("Update successful");
//         window.location.href = "index.html";
//         } else {
//         const errorText = await response.text();
//         alert("Error updating the Schedule Details. Status: " + response.status + "\n" + errorText);
//         }
//     } catch (error) {
//         console.error("Error updating the Schedule Details:", error);
//         alert("An unexpected error occurred. Please try again later.");
//     }
// }

function logout() {
    // Clear localStorage
    localStorage.clear();

    // Redirect to the logout page or any other desired destination
    window.location.href = "../Tutor/tutor_signup_login.html";
}

// Function to populate dropdown with topics
function populateDropdown() {
    var request = new XMLHttpRequest();
    const tutorId = localStorage.getItem('tutorId');
    const url = `http://localhost:4088/lessonmaterial/tutor/` + tutorId;
    request.open("GET", url);
    request.onload = function () {
        var data = JSON.parse(this.response);
        if (request.status >= 200 && request.status < 400) {
            var lmList = Object.keys(data.Materials);
            // Get the dropdown content element
            const dropdownContent = document.getElementById('dropdown-content');
            // Clear previous dropdown content
            dropdownContent.innerHTML = '';
            // Iterate over the Material object and create dropdown list items
            for (const key in data.Materials) {
                if (data.Materials.hasOwnProperty(key)) {
                    const topic = data.Materials[key].Topic;
                    const lessonId = key; // Get the lesson ID
                    const listItem = document.createElement('a');
                    listItem.href = '#';
                    listItem.textContent = topic;
                    // Add event listener to capture lesson ID when clicked
                    listItem.addEventListener('click', function() {
                        // Set the lesson ID to a hidden input field in the form
                        document.getElementById('lesson_id_create').value = lessonId;
                        // Set the topic name to a hidden input field in the form
                        document.getElementById('topic_name_create').value = topic;
                        // Change the text of the dropdown button to the selected topic
                        document.querySelector('.dropbtn').textContent = topic;
                    });
                    dropdownContent.appendChild(listItem);
                }
            }
        } else {
            console.error('Error fetching data:', request.statusText);
        }
    };
    request.onerror = function () {
        console.error('Network error');
    };
    request.send();
}

function populateDropdown_update(urlParams) {
    var request = new XMLHttpRequest();
    const tutorId = localStorage.getItem('tutorId');
    const url = `http://localhost:4088/lessonmaterial/tutor/` + tutorId;
    request.open("GET", url);
    request.onload = function () {
        var data = JSON.parse(this.response);
        if (request.status >= 200 && request.status < 400) {
            var lmList = Object.keys(data.Materials);
            // Get the dropdown content element
            const dropdownContent = document.getElementById('dropdown-content');
            // Clear previous dropdown content
            dropdownContent.innerHTML = '';
            // Iterate over the Material object and create dropdown list items
            for (const key in data.Materials) {
                if (data.Materials.hasOwnProperty(key)) {
                    const topic = data.Materials[key].Topic;
                    const lessonId = key; // Get the lesson ID
                    const listItem = document.createElement('a');
                    listItem.href = '#';
                    listItem.textContent = topic;
                    // Add event listener to capture lesson ID when clicked
                    listItem.addEventListener('click', function() {
                        // Set the lesson ID to a hidden input field in the form
                        document.getElementById('lesson_id_update').value = lessonId;
                        console.log(lessonId);
                        // Set the topic name to a hidden input field in the form
                        document.getElementById('topic_name_update').value = topic;
                        // Set the lesson ID as the selected value in the dropdown
                        document.getElementById('schedule_update_drop').value = lessonId;
                        // Change the text of the dropdown button to the selected topic
                        document.querySelector('.dropbtn').textContent = topic;
                    });
                    // Check if lessonId and topicName are not null to preset the dropdown
                    if (lessonId === decodeURIComponent(urlParams.get('lesson_id_update')) && topic === decodeURIComponent(urlParams.get('topic_name_update'))) {
                        listItem.classList.add('selected');
                        // Change the text of the dropdown button to the selected topic
                        document.querySelector('.dropbtn').textContent = topic;
                        console.log("this");
                    }
                    dropdownContent.appendChild(listItem);
                }
            }
        } else {
            console.error('Error fetching data:', request.statusText);
        }
    };
    request.onerror = function () {
        console.error('Network error');
    };
    request.send();
}
