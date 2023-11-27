const socket = new WebSocket("wss://servidor-ws.onrender.com/chat");

const nickname = localStorage.getItem('nickname');

socket.onmessage = function(event) {

    const chat = document.getElementById("chat");

    const message = JSON.parse(event.data);

    const mensagem = document.createElement("p");

    if (message.username !== nickname) {

        mensagem.style.textAlign = "right";

    }

    mensagem.textContent = message.username + ":";

    mensagem.appendChild(document.createTextNode(" " + message.content));

    mensagem.style.color = "white";

    chat.appendChild(mensagem);

}

function sendMessage() {

    const messageInput = document.getElementById("message");

    const message = messageInput.value;

    const msg = {username: nickname, content: message };

    socket.send(JSON.stringify(msg));

    messageInput.value = "";

}