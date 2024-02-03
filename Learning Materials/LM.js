var content = document.getElementById("lmContent");

function getAllLM() {

    var GetRequest = new XMLHttpRequest()
    GetRequest.open("GET", "http://localhost:4088//lessonmaterial/all")

    GetRequest.onload = function () {
        var data = JSON.parse(this.response)
        var lmList = Object.keys(data.Materials)

        content.innerHTML = ""
        
        lmList.forEach((lmId, index) => {

            var lm = Object.keys(data.Materials[lmId])
            console.log(lm)

            //HTML
            var card = document.createElement('div')
            card.className = 'card'
            card.style = 'width: 18rem'

            var cardBody = document.createElement('div')
            cardBody.className = 'card-body'

            var cardTitle = document.createElement('h5')
            cardTitle.className = 'card-title'

            var cardSub = document.createElement('h6')
            cardSub.className = 'card-subtitle mb-2 text-muted'

            var cardText = document.createElement('p')
            cardText.className = 'card-text'

            lm.forEach((material, index2) => {
                console.log(material)
                cardTitle.innerHTML +=  ((material == "Topic") ? data.Materials[lmId][material] : '')

                cardSub.innerHTML += ((material == "Created on") ? formatDate(data.Materials[lmId][material]) : '')
                cardText.innerHTML +=((material == "Summary") ? data.Materials[lmId][material] : '')
            })

            cardBody.appendChild(cardTitle)
            cardBody.appendChild(cardSub)
            cardBody.appendChild(cardText)
        
            card.appendChild(cardBody)
            content.appendChild(card)

        });
    }

    GetRequest.send()
}

function formatDate(datetimeStr) {
    var date = new Date(datetimeStr)
    var formattedDate = date.toLocaleString('en-US', {
        day: '2-digit',
        month: 'long',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
        hour12: true
    }).replace(/ at /g, ' ');

    return formattedDate

}

function addLM() {

    var addRequest = new XMLHttpRequest()
    const newID = document.getElementById("id").value
    addRequest.open("POST", "http://localhost:5000//lessonmaterial/" + newID)

    const newLMJSON = {
        "TutorID" : document.getElementById("tutorId").value,
        "Topic": document.getElementById("topic").value,
        "Summary": document.getElementById("summary").value,
        "Created On": $now(),
    }

    addRequest.onload = function () {
        if(addRequest.status == 202) {
            alert('Learning Material is successfully created')
            windows.location.href="LM.html"
        } else if (addRequest.status == 409) {
            alert('Learning Material is not created')
            windows.location.href="LM.html"
        }
    }

    addRequest.send(JSON.stringify(newLMJSON))
}