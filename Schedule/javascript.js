//display all the schedules
function listSchedules() {
    // Make a GET request to the server endpoint
    const url = `http://localhost:1000/api/v1/getSchedules`;
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
        data.forEach(user => {
          var row = tableBody.insertRow();
          row.innerHTML = `<td>${user.schedule_id}</td>
                            <td>${user.tutor_id}</td>
                            <td>${user.lesson_id}</td>
                            <td>${user.start_time}</td>
                            <td>${user.end_time}</td><td>
                                <button class="btn btn-outline-secondary" onclick="return modifySchedule(${user.schedule_id}, '${user.tutor_id}', '${user.lesson_id}', '${user.start_time}', '${user.end_time}')">Modify</button>
                                <button class="btn btn-outline-secondary" onclick="return deleteSchedule(${user.schedule_id})">Delete</button>
                            </td>`;
        });
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

function create_schdedule(){
    var request = new XMLHttpRequest();
    const form = document.getElementById('createScheduleForm');
    console.log(form)

    const curl = 'http://localhost:1000/api/v1/schedule';
    console.log(curl)

    const tutor_id = parseInt(form.elements['tutor_id_create'].value, 10);
    const lesson_id = parseInt(form.elements['lesson_id_create'].value, 10);
    const starttime = form.elements['starttime_create'].value.replace('T', ' ');
    const endtime = form.elements['endtime_create'].value.replace('T', ' ');
    console.log(tutor_id)

    request.open("POST", curl);
    request.setRequestHeader('Content-Type', 'application/json');

    request.onload = function () {
    if (request.status >= 200 && request.status < 400) {
        console.log('Schedule created successfully');
        window.location.href = "index.html";
    } else {
        // Error: handle the error response (if needed)
        console.error('Error creating schedule:', request.statusText);
    }
};

    request.onerror = function () {
        // Network error: handle the error (if needed)
        console.error('Network error');
    };

    request.send(JSON.stringify({
        "tutor_id": tutor_id,
        "lesson_id": lesson_id, 
        "start_time_str": starttime, 
        "end_time_str": endtime
        
    }));
    return false //prevent default submission
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

function modifySchedule(schedule_id, tutor_id, lesson_id, start_time, end_time) {
    // Build the URL with parameters
    const params = new URLSearchParams({
        schedule_id_update: encodeURIComponent(schedule_id),
        tutor_id_update: encodeURIComponent(tutor_id),
        lesson_id_update: encodeURIComponent(lesson_id),
        starttime_update: encodeURIComponent(start_time),
        endtime_update: encodeURIComponent(end_time)
    });

    const url = `update_schedule.html?${params.toString()}`;

    // Redirect to the new page with parameters
    window.location.href = url;
}
  

async function update_schedule(sid) {
    event.preventDefault();
    console.log("in update_schedule")
    console.log(sid)
    const form = document.getElementById('updateScheduleForm');
    
    const url = `http://localhost:1000/api/v1/oneSchedule/${sid}`;
    
    const starttime_update = form.elements['starttime_update'].value;
    const endtime_update = form.elements['endtime_update'].value;
    
    try {
        const response = await fetch(url, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "schedule_id_str": sid,
            "start_time_str": starttime_update,
            "end_time_str": endtime_update,
        }),
        });
    
        if (response.ok) {
        console.log("Update successful");
        window.location.href = "index.html";
        } else {
        const errorText = await response.text();
        alert("Error updating the Schedule Details. Status: " + response.status + "\n" + errorText);
        }
    } catch (error) {
        console.error("Error updating the Schedule Details:", error);
        alert("An unexpected error occurred. Please try again later.");
    }
}
