console.log("hello world");

function ajaxRequestFunction(method, url) {
    console.log("Requesting localhost:8080");
    let xhr = new XMLHttpRequest();

    // let's make the request
    xhr.open(method, url);
    xhr.send();

    let data = "";
    xhr.onload = () => {
        if(xhr.status == 200) {
            data = xhr.responseText;
        }
    }
    return data;
}

function fetchRequestExample(method, url) {
    fetch(url)
        .then(data => data.text())
        .then(data => console.log(data));

}

let url = "http://localhost:8080";
let method = "GET";
let requestButton = document.getElementById("request-button");
// requestButton.addEventListener("click", () => ajaxRequestFunction(method, url));
requestButton.addEventListener("click", () => fetchRequestExample(method, url));
