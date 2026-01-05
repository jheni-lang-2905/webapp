document.querySelector('#login').addEventListener('submit', fazerLogin);

async function fazerLogin(evento) {
    evento.preventDefault();

    const email = document.querySelector('#email').value;
    const senha = document.querySelector('#senha').value;

    try {
        const resposta = await fetch("/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ email, senha })
        });

        if (resposta.ok) {
            window.location.href = "/home";
        } else {
            Swal.fire("Ops...", "Usuário ou senha incorretos!", "error");
        }
    } catch (erro) {
        Swal.fire("Erro", "Não foi possível conectar ao servidor!", "error");
    }
}
