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

            var lmLink = document.createElement("button");
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

    console.log(getId())
    GetRequest.send()
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
    console.log(lmDate.textContent)

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

function addLM() {

    var addRequest = new XMLHttpRequest()
    const newID = getId()
    addRequest.open("POST", "http://localhost:4088//lessonmaterial/material/" + newID)

    const newLMJSON = {
        "TutorID" : tutor_id,
        "Topic": document.getElementById("topic").value,
        "Summary": document.getElementById("summary").value,
        "Created on": $now(),
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

function editLM(TutorID, Topic, Summary, CreatedOn) {

    updateContent.innerHTML = (
        "<form class=\"login active\" onsubmit=\"return updateLM()\">" +
        "<a href=\"../Learning Materials/LM.html\">"+
            "<i class='bx bxs-arrow-to-left' id=\"backbtn\" title=\"Back to Learning Material\" ></i>" +
        "</a>"+
        "<h2 class=\"title\">Update Learning Material</h2>" +
        "<div class=\"form-group\">" +
            "<label for=\"tutorid\">Tutor ID</label>" +
            "<div class=\"input-group\">" +
                "<input type=\"text\" id=\"tutorid\" placeholder=\"Tutor ID\">" + TutorID +
            "</div>" +
            "<span class=\"help-text\">Required</span>" +
        "</div>" +
        "<div class=\"form-group\">"+
            "<label for=\"topic\">Topic</label>"+
            "<div class=\"input-group\">"+
                "<input type=\"text\" id=\"topic\" placeholder=\"Topic\">" + Topic +
            "</div>" +
            "<span class=\"help-text\">Required</span>" +
        "</div>" +
        "<div class=\"form-group\">" +
            "<label for=\"topic\">Summary</label>" +
            "<div class=\"input-group\">" +
                "<textarea name=\"opinion\" rows=\"7\" type=\"text\" placeholder=\"Summary\" id=\"summary\">" + Summary +
                "</textarea>" +
            "</div>" +
            "<span class=\"help-text\">Required</span>" +
        "</div>" +
        
        "<button class=\"btn-submit\">Update</button>" +

    "</form>"
    )
}

function updateLM() {

    var updateRequest = new XMLHttpRequest()
    const id = document.getElementById("Id").value
    
    updateRequest.open("PUT", "http://localhost:4088/lessonmaterial/material/"+id)

    const updatedLMJSON = {
        "TutorID" : document.getElementById("tutorId").value,
        "Topic": document.getElementById("topic").value,
        "Summary": document.getElementById("summary").value,
        "Created on": $now(),
    }

    updateRequest.onload = function () {
        if(updateRequest.status == 202) {
            alert('Learning Material is successfully created')
            windows.location.href="../Learning Materials/LM.html"

        } else if (updateRequest.status == 404) {
            alert('Learning Material is not created')
            windows.location.href="../Learning Materials/LM.html"
        }
    }

    updateRequest.send(JSON.stringify(updatedLMJSON))
    return false

}