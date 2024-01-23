function book(){
    var request = new XMLHttpRequest();
    const form = document.getElementById('booking_details_form');

    const curl = 'http://localhost:1765/api/v1/book';

    const name = form.elements['name'].value;
    const email = form.elements['email'].value;
    const lessonId = form.elements['lesson_id'].value;
    console.log(name);
    console.log(email);
    console.log(lessonId);

    request.open("POST", curl);
    request.send(JSON.stringify({
        "Student Name": name,
        "Student Email": email,
        "Lesson ID": lessonId
    }));
    form.reset();
    return false //prevent default submission
}

function deleteCourse(email){
    var request = new XMLHttpRequest();
    console.log(screen)
    request.open('DELETE', 'http://localhost:1765/api/v1/book/'+email);
    request.send();
}