var content = document.getElementById("lmContent")

function getLMs() {
    var GetRequest = new XMLHttpRequest()

    GetRequest.open("GET", "http://localhost:4088/lessonmaterial/all")

    GetRequest.onload = function() {
        
        var data = JSON.parse(this.parse)
        var lmList = Object.keys(data.LearningMaterials)

        content.innerHTML = ""

        lmList.forEach((lmId, index) => {

            var lm = Object.keys(data.LearningMaterials[lmId])

            var card = document.createElement('div');
            card.className = 'col-lg-6 mt-4';

            var cardDetail = document.createElement('div');
            cardDetail.className = 'material d-flex align-items-start';

            var cardImg = document.createElement('div');
            cardImg.className = 'materialimg';
            
            var cardImgSrc = document.createElement('img');
            cardImgSrc.src = '<source edit>';

            cardImg.appendChild(cardImgSrc)

            var cardLMInfo = document.createElement('div');
            cardLMInfo.className = 'material-info';

            lm.forEach((details, index2) => {
                //Learning Material Title
                var cardLMTitle = document.createElement('h4')
                //Learning Material Summary
                var cardLMSum = document.createElement('p')

                cardLMTitle.innerHTML += ((attribute == "Topic") ? data.LearningMaterials[lmId][details] : '')
                cardLMSum.innerHTML += ((attribute == "Summary") ? data.LearningMaterials[lmId][details] : '')
            })

            cardLMInfo.appendChild(cardLMTitle)
            cardLMInfo.appendChild(cardLMSum)

            var cardArrow = createElement('div')
            cardArrow.innerHTML += "<button onclick=\"getLMDetails(\'"+lmId+"\')\"> Learn More</button>"

            cardDetail.appendChild(cardImg)
            cardDetail.appendChild(cardLMInfo)
            cardDetail.appendChild(cardArrow)

            card.appendChild(cardDetail)

        });
    }
    GetRequest.send()
}

function addLM() {

    var addRequest = new XMLHttpRequest()

    const newID = document.getElementById("id").ariaValueMax
    addRequest.open("POST", "http://localhost:4088//lessonmaterial/" + newID)

    const newLMJSON = {
        "Topic": document.getElementById("topic").value,
        "Summary": document.getElementById("summary").value,
        "Ceated On": getTime(),
    }


    addRequest.onload = function () {
        if(addRequest.status == 202) {
            content.innerHTML = "Course" + newID + "created successfully"
        } else if (requestAnimationFrame.status == 409) {
            windows.location.href = 'http://localhost:4088//'
        }
    }
}