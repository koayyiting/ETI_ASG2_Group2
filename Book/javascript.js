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
  
        // Iterate through the received data and append rows to the table
        data.forEach(user => {
            var row = tableBody.insertRow();
            row.innerHTML = `<td>${user.booking_id}</td>
                                <td>${user.student_id}</td>
                                <td>${user.schedule_id}</td>
                                    <button class="btn btn-outline-secondary" onclick="return deleteBooking(${user.booking_id})">Delete</button>
                                </td>`;
        });
      })
      .catch(error => console.error('Error fetching user details:', error));
}

function create_booking(){
    var request = new XMLHttpRequest();
    const form = document.getElementById('createBookingForm');

    const curl = 'http://localhost:1765/api/v1/book';

    const student_id = form.elements['student_id_create'].value;
    const schedule_id = form.elements['schedule_id_create'].value;

    request.open("POST", curl);
    request.setRequestHeader('Content-Type', 'application/json');

    request.onload = function () {
        if (request.status >= 200 && request.status < 400) {
            console.log('Booking created successfully');
            window.location.href = "index_booking.html";
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
        "student_id": student_id,
        "schedule_id": schedule_id
    }));
    form.reset();
    return false //prevent default submission
}

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

function switchToBookingFormpage(){
    window.location.href = 'create_schedule.html';
}

function indexPage(){
    window.location.href = 'index_booking.html';
}
