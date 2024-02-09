// Display all the schedules based on bookings
function listBookings() {
    const studentId = localStorage.getItem('studentId');
    const bookingsUrl = `http://localhost:1765/api/v1/getBookings/` + studentId;

    fetch(bookingsUrl)
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then(bookingsData => {
            console.log("Bookings data from server:", bookingsData);

            // Get the table body element
            const tableBody = document.getElementById('bookings_table').getElementsByTagName('tbody')[0];

            // Clear existing rows
            tableBody.innerHTML = '';

            if (bookingsData != null) {
                bookingsData.forEach(booking => {
                    const scheduleId = booking.schedule_id;

                    // Fetch schedule information for the current schedule ID
                    fetchScheduleInfo(scheduleId)
                        .then(scheduleInfo => {
                            // Create a row for each schedule
                            const row = tableBody.insertRow();
                            const formattedStartTime = new Date(scheduleInfo.start_time).toLocaleString('en-US', { 
                                year: 'numeric', 
                                month: '2-digit', 
                                day: '2-digit',
                                hour: 'numeric', 
                                minute: '2-digit', 
                                hour12: true 
                            });
                            
                            const formattedEndTime = new Date(scheduleInfo.end_time).toLocaleString('en-US', { 
                                year: 'numeric', 
                                month: '2-digit', 
                                day: '2-digit',
                                hour: 'numeric', 
                                minute: '2-digit', 
                                hour12: true 
                            });

                            row.innerHTML = `<td>${scheduleInfo.lesson_name}</td>
                                             <td>${formattedStartTime}</td>
                                             <td>${formattedEndTime}</td>
                                             <td>${scheduleInfo.location}</td>
                                             <td><button style="border:none; padding: 5px 15px; border-radius:10px;" onclick="return deleteBooking(${booking.booking_id})">Cancel</button></td>`;
                        })
                        .catch(error => console.error('Error fetching schedule info:', error));
                });
            } else {
                var emptyMessageDiv = document.getElementById('empty_message');
                emptyMessageDiv.innerHTML = `<p style="padding: 20px 0px">You have 0 Bookings</p>`;
            }
        })
        .catch(error => console.error('Error fetching bookings:', error));
}

// Function to fetch schedule information by schedule ID
function fetchScheduleInfo(scheduleId) {
    const url = `http://localhost:1000/api/v1/oneSchedule/` + scheduleId;
    return fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .catch(error => console.error('Error fetching schedule info:', error));
}


// function create_booking(){
//     var request = new XMLHttpRequest();
//     const form = document.getElementById('createBookingForm');

//     const curl = 'http://localhost:1765/api/v1/book';

//     const student_id = form.elements['student_id_create'].value;
//     const schedule_id = form.elements['schedule_id_create'].value;

//     request.open("POST", curl);
//     request.setRequestHeader('Content-Type', 'application/json');

//     request.onload = function () {
//         if (request.status >= 200 && request.status < 400) {
//             console.log('Booking created successfully');
//             window.location.href = "index_booking.html";
//         } else {
//             // Error: handle the error response (if needed)
//             console.error('Error creating schedule:', request.statusText);
//         }
//     };

//     request.onerror = function () {
//         // Network error: handle the error (if needed)
//         console.error('Network error');
//     };

//     request.send(JSON.stringify({
//         "student_id": student_id,
//         "schedule_id": schedule_id
//     }));
//     form.reset();
//     return false //prevent default submission
// }

function deleteBooking(bookingID) {
    console.log('Deleting Booking with ID:', bookingID);
    const url = `http://localhost:1765/api/v1/oneBooking/${bookingID}`;
  
    // Confirm deletion with the user (you can customize this)
    if (confirm("Are you sure you want to cancel this Booking?")) {
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
            listBookings();
        })
        .catch(error => console.error('Error deleting Booking:', error));
    }
}

function available_booking_page(){
    window.location.href = 'available_bookings.html';
}

function user_booking_page(){
    window.location.href = 'user_bookings.html';
}


function listAvailableSchedule() {
    const url = "http://localhost:1000/api/v1/getAllSchedules";
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
        var tableBody = document.getElementById('booking_table_available').getElementsByTagName('tbody')[0];
  
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
                                <button style="border:none; padding: 5px 15px; border-radius:10px;" onclick="return create_booking(${user.schedule_id})">Book</button>
                                </td>`;
                            });
        } else{
            var emptyMessageDiv = document.getElementById('empty_message_available');
            emptyMessageDiv.innerHTML = `<p style="padding: 20px 0px">0 Lesson Available</p>`;
        }
      })
      .catch(error => console.error('Error fetching user details:', error));
}

function create_booking(sid){
    var request = new XMLHttpRequest();
    const studentid = localStorage.getItem('studentId');

    const curl = 'http://localhost:1765/api/v1/book/' + sid;
    console.log(sid)

    request.open("POST", curl);
    request.setRequestHeader('Content-Type', 'application/json');

    request.onload = function () {
        if (request.status >= 200 && request.status < 400) {
            console.log('Booking created successfully');
            window.location.href = "/Book/user_bookings.html";
        } else if (request.status == 409){
            // Error: handle the error response (if needed)
            console.error('Error creating schedule:', request.statusText);
        } else if (request.status == 422){
            alert("You have already Booked this Schedule")
        }
    };

    request.onerror = function () {
        // Network error: handle the error (if needed)
        console.error('Network error');
    };

    request.send(JSON.stringify({
        "student_id": parseInt(studentid),
        "schedule_id": sid
    }));
    return false //prevent default submission
}

function logout() {
    // Clear localStorage
    localStorage.clear();

    // Redirect to the logout page or any other desired destination
    window.location.href = "../Student/student_signup_login.html";
}