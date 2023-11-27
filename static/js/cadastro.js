document.getElementById('btCadastrar').addEventListener('click', async function () {

    const nickname = document.querySelector("#nickname").value;

    const senha = document.querySelector("#senha").value;

    const confSenha = document.querySelector("#confSenha").value;

    if (senha != confSenha) {

        alert("Senhas não conferem!");  

        return;

    } else {

        const respostaLogin = await fetch("/cadastrar", {

            method: "POST",

            body: JSON.stringify({ nickname, senha })

        });

        const informacoesLogin = await respostaLogin.json();

        if (informacoesLogin.message === "Registro bem-sucedido") {

            alert('Registro efetuado com sucesso!');

            window.location.href = "/index";

        } else {

            alert('Ops! Registro não efetuado.');

        }

    }

});

