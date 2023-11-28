const socket = new WebSocket("wss://servidor-ws.onrender.com/chat");

const nickname = localStorage.getItem('nickname');

function carregarMensagens() {
    fetch("/listar")
        .then(response => response.json())
        .then(data => {
            const chat = document.getElementById("chat");
            chat.innerHTML = "";

            for (let i = 0; i < data.mensagens.length; i++) {
                const mensagem = data.mensagens[i];
                const mensagemTexto = document.createElement("p");
                mensagemTexto.textContent = mensagem.usuario + ": " + mensagem.texto;
                if (mensagem.usuario !== nickname) {
                    mensagemTexto.style.textAlign = "right";
                }
                mensagemTexto.style.color = "white";
                chat.appendChild(mensagemTexto);
            }
        })
        .catch(error => {
            console.error("Erro ao carregar as mensagens:", error);
        });
}

document.addEventListener("DOMContentLoaded", carregarMensagens);

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

    if (message.usuario !== nickname) {

        mensagem.style.textAlign = "right";

    }

    mensagem.textContent = message.usuario + ":";

    mensagem.appendChild(document.createTextNode(" " + message.texto));

    mensagem.style.color = "white";

    chat.appendChild(mensagem);

}

function sendMessage() {
    const messageInput = document.getElementById("message");
    const message = messageInput.value;
    const msg = { usuario: nickname, texto: message };

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
        // Após enviar a mensagem e receber a resposta, rola a área do chat para baixo
        const chatArea = document.getElementById("chat");
        chatArea.scrollTop = chatArea.scrollHeight;
    });

    messageInput.value = "";
}
