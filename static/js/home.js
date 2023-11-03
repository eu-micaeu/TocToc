const socket = new WebSocket("wss://servidor-ws.onrender.com/chat");

        socket.onmessage = function(event) {
            const chat = document.getElementById("chat");
            const message = JSON.parse(event.data);
            const p = document.createElement("p");
            const strong = document.createElement("strong");
            strong.textContent = message.username + ":";
            p.appendChild(strong);
            p.appendChild(document.createTextNode(" " + message.content));
            chat.appendChild(p);
        };

        function sendMessage() {
            const usernameInput = document.getElementById("username");
            const messageInput = document.getElementById("message");
            const username = usernameInput.value;
            const message = messageInput.value;
            const msg = { username, content: message };
            socket.send(JSON.stringify(msg));
            messageInput.value = "";
        }