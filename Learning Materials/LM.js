var content = document.querySelector("ul.box-info#lmContent");
const tutor_id = parseInt(localStorage.getItem('tutorId'));

function getAllLM() {

    var GetRequest = new XMLHttpRequest()
    GetRequest.open("GET", "http://localhost:4088/lessonmaterial/all")

    GetRequest.onload = function () {
        var data = JSON.parse(this.response)
        var lmList = Object.keys(data.Materials)

        content.innerHTML = "";

        console.log(data)

        lmList.forEach((lmId, index) => {

            var lm = Object.keys(data.Materials[lmId])

            //HTML
            var lmItem = document.createElement("li");

            var lmLink = document.createElement("a");
            lmLink.id = lmId
            lmLink.addEventListener("click", () => {
                loadSummary(lmId)
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
            })

            console.log(lmItem)

            lmLink.appendChild(lmIcon)
            lmBody.appendChild(lmHeader);

            var brline = document.createElement("br")
            lmBody.appendChild(brline)
            lmBody.appendChild(lmCreated);
            
            lmItem.appendChild(lmLink);
            lmItem.appendChild(lmBody);

            content.appendChild(lmItem);

        });
    }

    GetRequest.send()
}

function loadSummary(lmSummaryId) {

    fetch("http://localhost:4088/material/" + lmSummaryId)
        .then(response => response.json())
        .then(data => {
            if(data) {
                const queryString = new URLSearchParams({id : lmSummaryId});

                window.location.href = "../Learning Materials/LMSummary.html?" + queryString;

            } else {
                console.error("Failed to fetch material data for:", lmSummaryId)
            }
        })
        .catch(error => {
            console.error("Error fetching material data:", error);
        });
}

function lmSummary() {

    const urlParams = new URLSearchParams(window.location.search);
    const summaryId = urlParams.get("id");

    if (!summaryId){
        console.error("Missing ID in query string");
        return;
    }

    //HTML Element
    var lmMaterialId = document.getElementById("lmMaterialID");
    var lmTitle = document.getElementById("materialTitle");
    var lmParagraph = document.getElementById("lmSummary");
    var lmDate = document.getElementById("lmCreated");


    console.log(lmMaterialId.textContent);
    console.log(lmTitle.textContent);
    console.log(lmParagraph.textContent);
    console.log(lmDate.textContent);

    var GetRequest = new XMLHttpRequest();
    GetRequest.open("GET", "http://localhost:4088/material/" + summaryId)

    GetRequest.onload = function () {
        var data = JSON.parse(this.response)
        var lmObj = Object.keys(data.Material)

        console.log(lmObj)
        lmMaterialId.textContent = lmObj;

        lmTitle.textContent = (data.Material[lmObj]['Topic']);
        lmParagraph.textContent = (data.Material[lmObj]['Summary']);
        lmCreated.textContent = "Created On " + (formatDate(data.Material[lmObj]['Created on']));

    }

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

function formatDateSQL() {
    const date = new Date();
    const formattedDate = date.toISOString().replace(/T/, ' ').replace(/\.\d+Z$/, '');
    return formattedDate
}

function getId() {

    var id;
    fetch("http://localhost:4088/lessonmaterial/all")
        .then(response => response.json())
        .then(data => {
            var lmList = Object.keys(data.Materials)
            id = lmList.length
            id = id + 1
            console.log("new LMID output:", id)
                
            return id.toString();
        })
        .then(idString => {
            window.globalLmId = idString;
    });

}

function addLM() {

    //Get LMID, TutorId and DateTime
    const newID = window.globalLmId;
    const createdDate = formatDateSQL();
    // const tutorId = tutor_id;

    console.log(newID);
    console.log(createdDate);
    // console.log(tutorId);

    var addRequest = new XMLHttpRequest()
    addRequest.open("POST", "http://localhost:4088/lessonmaterial/material/" + newID)

     //Create Lesson Material JSON
     const newLMJSON = {
        "TutorID" : parseInt(document.getElementById("tutorid").value),
        "Topic": document.getElementById("topic").value,
        "Summary": document.getElementById("summary").value,
        "Created on": createdDate,
    }

    console.log(newLMJSON)

    addRequest.onload = function () {
        if(addRequest.status === 202) {
            alert('Learning Material is successfully created')
            windows.location.href="../Learning Materials/LM.html"
        } else if (addRequest.status === 409) {
            const errorMessage = addRequest.response || "Learning Material is not created. Please check inputs are within requirements!";
            alert(errorMessage)
        } else {
            console.error('Failed to add learning material. Status:', addRequest.status);
            alert('Failed to add learning material. Please try again later.');
        }
    };

    console.log(addRequest)
    console.log(addRequest.status)
    addRequest.send(JSON.stringify(newLMJSON))

}

function loadEditLM(lmUpdateId) {

    fetch("http://localhost:4088/material/" + lmUpdateId)
        .then(response => response.json())
        .then(data => {
            if(data) {
                const queryString = new URLSearchParams({id : lmUpdateId});
                window.location.href = "../Learning Materials/LMUpdate.html?" + queryString;

            } else {
                console.error("Failed to fetch material data for:", lmUpdateId)
            }
        })
        .catch(error => {
            console.error("Error fetching material data:", error);
        });
}

// function updateLM() {

//     const urlParams = new URLSearchParams(window.location.search);
//     const loadId = urlParams.get("id");

//     if (!summaryId){
//         console.error("Missing ID in query string");
//         return;
//     }

//     //HTML Element
//     var lmTutor = document.getElementById("tutorid");
//     var lmTopic = document.getElementById("topic");
//     var lmSummary = document.getElementById("summary");


//     console.log(lmTitle.textContent);
//     console.log(lmTopic.textContent);
//     console.log(lmSummary.textContent);

//     var loadRequest = new XMLHttpRequest();
//     loadRequest.open("GET", "http://localhost:4088/material/"+loadId)

//     loadRequest.onload = function() {
//         var data = JSON.parse(this.response)
//         var lmObj = Object.keys(data.Material)

//         console.log(lmObj)
//         lmMaterialId.textContent = lmObj;

//         lmTutor.textContent = (data.Material[lmObj]['TutorID']);
//         lmTopic.textContent = (data.Material[lmObj]['Topic']);
//         lmTopic.textContent = (data.Material[lmObj]['Summary']);
//     }

//     loadRequest.send()
// }

// function updateLM(updateId) {

//     var loadRequest = new XMLHttpRequest();
//     loadRequest.open("PUT", "http://localhost:4088/lessonmaterial/material/"+ updateId)

//     const updatedLMJSON = {
//         "TutorID" : document.getElementById("tutorId").value,
//         "Topic": document.getElementById("topic").value,
//         "Summary": document.getElementById("summary").value,
//     }

// }
    
//     updateRequest.onload = function () {
//         if(updateRequest.status == 202) {
//             alert('Learning Material is successfully created')
//             windows.location.href="../Learning Materials/LM.html"

//         } else if (updateRequest.status == 404) {
//             alert('Learning Material is not created')
//             windows.location.href="../Learning Materials/LM.html"
//         }
//     }

//     updateRequest.send(JSON.stringify(updatedLMJSON))
//     return false
