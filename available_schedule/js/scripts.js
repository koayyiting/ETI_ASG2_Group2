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
    
            // Get the container where you want to append cards
            const cardContainer = document.getElementById('cardContainer');
    
            // Clear existing content
            cardContainer.innerHTML = '';
    
            // Loop through the data and create cards
            data.forEach(schedule => {
                // Create a new card element
                const card = document.createElement('div');
                card.className = 'col-md-4 mb-5';
    
                // Set the card HTML content
                card.innerHTML = `
                    <div class="card h-100">
                        <div class="card-body">
                            <h2 class="card-title" id="scheduled_lesson_id">Lesson ID: ${schedule.lesson_id}</h2>
                            <p class="card-text" id="scheduled_start_time">Start Time: ${schedule.start_time}</p>
                            <p class="card-text" id="scheduled_end_time">End Time: ${schedule.end_time}</p>
                            <p class="card-text" id="scheduled_schedule_id">Schedule ID: ${schedule.schedule_id}</p>
                        </div>
                        <button type="submit" class="btn btn-outline-secondary" onclick="return create_booking(${schedule.schedule_id})">Book</button>
                    </div>
                `;
    
                // Append the card to the container
                cardContainer.appendChild(card);
            });
        })
        .catch(error => console.error("Error fetching data:", error));
}

function create_booking(sid){
    var request = new XMLHttpRequest();
    const cardContainer = document.getElementById('cardContainer');

    const curl = 'http://localhost:1765/api/v1/book';
    console.log(sid)

    request.open("POST", curl);
    request.setRequestHeader('Content-Type', 'application/json');

    request.onload = function () {
        if (request.status >= 200 && request.status < 400) {
            console.log('Booking created successfully');
            window.location.href = "/Book/index_booking.html";
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
        "student_id": 1,
        "schedule_id": sid
    }));
    return false //prevent default submission
}
