function listBookings() {
    // Make a GET request to the server endpoint
    const url = `http://localhost:1765/api/v1/getBookings`;
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
        var tableBody = document.getElementById('bookings_table').getElementsByTagName('tbody')[0];
  
        // Clear existing rows
        tableBody.innerHTML = '';

        if (data != null){
            data.forEach(user => {
                var row = tableBody.insertRow();
                row.innerHTML = `<td>${user.booking_id}</td>
                                    <td>${user.student_id}</td>
                                    <td>${user.schedule_id}</td>
                                        <button style="border:none; padding: 5px 15px; border-radius:10px;" onclick="return deleteBooking(${user.booking_id})">Delete</button>
                                    </td>`;
            });
        } else{
            var emptyMessageDiv = document.getElementById('empty_message');
            emptyMessageDiv.innerHTML = `<p style="padding: 20px 0px">You have 0 Bookings</p>
            <button style="border:none; padding: 5px 15px; border-radius:10px;" onclick="return available_booking_page()">Book Now</button>`;
        }
  
        // Iterate through the received data and append rows to the table
        
      })
      .catch(error => console.error('Error fetching user details:', error));
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
    if (confirm("Are you sure you want to delete this Booking?")) {
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
    const url = "http://localhost:1000/api/v1/getSchedules";
    
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
            var tableBody = document.getElementById('bookings_table').getElementsByTagName('tbody')[0];
      
            // Clear existing rows
            tableBody.innerHTML = '';
    
            if (data != null){
                data.forEach(user => {
                    var row = tableBody.insertRow();
                    row.innerHTML = `<td>${user.schedule_id}</td>
                                        <td>${user.start_time}</td>
                                        <td>${user.end_time}</td>
                                        <td>Temp Location</td>
                                        <td>
                                            <button style="border:none; padding: 5px 15px; border-radius:10px;" onclick="return create_booking(${user.schedule_id})">Book</button>
                                        </td>`;
                });
            } else{
                tableBody.innerHTML = '<p style="justify-content: center; align-items: center;">There are no available lessons</p>';
            }
      
            // Iterate through the received data and append rows to the table
            
          })
          .catch(error => console.error('Error fetching user details:', error));
}

function create_booking(sid){
    var request = new XMLHttpRequest();

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
        "student_id": 1,
        "schedule_id": sid
    }));
    return false //prevent default submission
}