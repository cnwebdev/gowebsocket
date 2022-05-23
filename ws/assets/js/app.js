// "use strict";

let socket = null;

document.addEventListener("DOMContentLoaded", () => {
    socket = new WebSocket("ws://127.0.0.1:8080/ws");

    socket.onopen = () => {
        console.log("Client successfully connected");
    }

    socket.onclose = () => {
        console.log("connection closed");
    }

    socket.onerror = error => {
        console.log("there was an error");
    }

    socket.onmessage = msg => {
        console.log(msg);
    }
});