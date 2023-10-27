function login() {

    const email = document.getElementById('email').value;

    const password = document.getElementById('password').value;

    const data = { email, password };

    fetch('/login', {

        method: 'POST',

        headers: {

            'Content-Type': 'application/json'

        },

        body: JSON.stringify(data)

    })

    .then(response => {

        if (response.ok) {

            window.location.href = '/index';

        } else {

            alert('Email ou senha incorretos');

        }

    })

    .catch(error => {

        alert('Ocorreu um erro ao fazer login');

    });

}

document.getElementById('btEntrar').addEventListener('click', login);
