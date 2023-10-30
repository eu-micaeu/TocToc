document.getElementById('btEntrar').addEventListener('click', async function () {

    const email = document.querySelector("#nickname").value;

    const senha = document.querySelector("#senha").value;

    const respostaLogin = await fetch("/login", {

        method: "POST",

        body: JSON.stringify({ nickname, senha })

    });

    const informacoesLogin = await respostaLogin.json();

    if (informacoesLogin.message === "Login efetuado com sucesso!") {

        window.location.href = "/index";

    } else {

        alert('Ops! Usuário inexistente');

    }

});

