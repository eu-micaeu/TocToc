var socket = new WebSocket('wss://servidor-ws.onrender.com/chat');

const nickname = localStorage.getItem('nickname');

document.addEventListener("DOMContentLoaded", function() {
    const messageInput = document.getElementById("message");
    messageInput.addEventListener("keydown", function(event) {
        if (event.key === "Enter") {
            sendMessage();
        }
    });
});

socket.onmessage = function (event) {

    const chat = document.getElementById("chat");

    const message = JSON.parse(event.data);

    const mensagem = document.createElement("p");

    mensagem.textContent = message.nickname + ":";

    mensagem.appendChild(document.createTextNode(" " + message.texto));

    mensagem.style.color = "white";

    chat.appendChild(mensagem);

}

function isOpen(ws) { 
    return ws.readyState === ws.OPEN; 
}

function sendMessage() {
    const messageInput = document.getElementById("message");
    const message = messageInput.value;

    const msg = { nickname: nickname, texto: message };

    if (!isOpen(socket)) return;

    socket.send(JSON.stringify(msg));

    fetch("/enviar", {

        method: "POST",

        body: JSON.stringify(msg),

        headers: {
            'Content-Type': 'application/json'
        }

    }).then(response => {

        if (response.ok) {

            console.log(msg);

            return response.json();

        } else {

            throw new Error('Erro na resposta da rede');

        }

    }).then(() => {

        const chatArea = document.getElementById("chat");

        chatArea.scrollTop = chatArea.scrollHeight;

    });

    messageInput.value = "";
}
