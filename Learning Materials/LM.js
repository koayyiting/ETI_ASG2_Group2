var content = document.querySelector("ul.box-info#lmContent");

function getAllLM() {

    var GetRequest = new XMLHttpRequest()
    GetRequest.open("GET", "http://localhost:4088/lessonmaterial/all")

    GetRequest.onload = function () {
        var data = JSON.parse(this.response)
        var lmList = Object.keys(data.Materials)

        content.innerHTML = "";

        lmList.forEach((lmId, index) => {

            var lm = Object.keys(data.Materials[lmId])

            //HTML
            var lmItem = document.createElement("li");

            var lmLink = document.createElement("a");
            lmLink.addEventListener("click", () => {
                lmSummary(lmId)
            });

            console.log(lmId)

            var lmIcon = document.createElement("i");
            lmIcon.className = "bx bxs-chevron-right";

            var lmBody = document.createElement("span");
            lmBody.className = "text";

            var lmHeader = document.createElement("h3");
            // var lmSummary = document.createElement("p");
            var lmCreated = document.createElement("footer");
            lmCreated.className = "blockquote-footer";

            lm.forEach((material, index2) => {
                lmHeader.innerHTML +=  ((material == "Topic") ? data.Materials[lmId][material] : '')
                lmCreated.innerHTML += ((material == "Created on") ? formatDate(data.Materials[lmId][material]) : '')
                // lmSummary.innerHTML +=((material == "Summary") ? data.Materials[lmId][material] : '')
            })

            console.log(lmItem)

            lmLink.appendChild(lmIcon)
            lmBody.appendChild(lmHeader);
            // lmBody.appendChild(lmSummary);

            var brline = document.createElement("br")
            lmBody.appendChild(brline)
            lmBody.appendChild(lmCreated);
            
            lmItem.appendChild(lmLink);
            lmItem.appendChild(lmBody);

            content.appendChild(lmItem);

        });
    }

    console.log(getId())
    GetRequest.send()
}

function formatDate(datetimeStr) {
    var date = new Date(datetimeStr)
    var formattedDate = date.toLocaleString('en-US', {
        day: '2-digit',
        month: 'long',
        year: 'numeric',
        hour12: true
    }).replace(/ at /g, ' ');

    return formattedDate

}

function getId() {

    var id;
    fetch("http://localhost:4088/lessonmaterial/all")
        .then(response => response.json())
        .then(data => {
            var lmList = Object.keys(data.Materials)
            id = lmList.length
            console.log(id)
        })

    return id;

}

function getTutorId() {

    var tutorId;
    fetch("http://localhost:5211/api/v1/tutor/tutorobject")
        .then(response => response.json())
        .then(data => {
            var tutor = Object.keys(data.Tutor)
            tutorId = tutor.ID
        })

    return tutorId
}

function lmSummary(id) {
    
    var summary = document.querySelector("ul.box-info#lmSummary");
    var lmTitle = document.getElementById("materialTitle");

    var GetRequest = new XMLHttpRequest()
    GetRequest.open("GET", "http://localhost:4088/lessonmaterial/material/" + id)

    GetRequest.onload = function () {
        var data = JSON.parse(this.response)
        var lmNest = Object.keys(data.Material)

        console.log(lmNest)

        summary.innerHTML = "";

        lmNest.forEach((lmId, index) => {

            var lm = Object.keys(data.Material[lmId])

            //HTML
            var lmSummary = document.createElement("p");
            var lmCreated = document.createElement("span");

            lm.forEach((material, index2) => {

                lmTitle.innerHTML +=  ((material == "Topic") ? data.Materials[lmId][material] : '')
                lmSummary.innerHTML +=((material == "Summary") ? data.Materials[lmId][material] : '')
                lmCreated.innerHTML += ((material == "Created on") ? formatDate(data.Materials[lmId][material]) : '')

            })

            console.log(summary)

            summary.appendChild(lmSummary);
            summary.appendChild(lmCreated)

        });
    }

    window.location.href="../Learning Materials/LMSummary.html"
    
}

function addLM() {

    var addRequest = new XMLHttpRequest()
    const newID = getId()
    addRequest.open("POST", "http://localhost:5000//lessonmaterial/" + newID)

    console.log(addRequest)

    const newLMJSON = {
        "TutorID" : document.getElementById("tutorId").value,
        "Topic": document.getElementById("topic").value,
        "Summary": document.getElementById("summary").value,
        "Created On": $now(),
    }

    addRequest.onload = function () {
        if(addRequest.status == 202) {
            alert('Learning Material is successfully created')
            windows.location.href="../Learning Materials/LM.html"
        } else if (addRequest.status == 409) {
            alert('Learning Material is not created')
            windows.location.href="../Learning Materials/LM.html"
        }
    }

    addRequest.send(JSON.stringify(newLMJSON))


}