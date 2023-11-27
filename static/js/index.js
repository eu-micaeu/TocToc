document.getElementById('btEntrar').addEventListener('click', async function () {

    const nickname = document.querySelector("#nickname").value;

    const senha = document.querySelector("#senha").value;

    const respostaLogin = await fetch("/login", {

        method: "POST",

        body: JSON.stringify({ nickname, senha })

    });

    const informacoesLogin = await respostaLogin.json();

    if (informacoesLogin.message === "Login efetuado com sucesso!") {

        localStorage.setItem('nickname', nickname);

        window.location.href = "/home";

    } else {

        alert('Ops! Usuário inexistente');

    }

});

