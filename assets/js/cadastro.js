$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento){
    evento.preventDefault();
    console.log("Dentro da funcao do usuario");

    if($('#senha').val() !== $('#confirmar-senha').val()){
        Swal.fire(
            "ops!",
            "senhas não coincidem",
            "error"
        );
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    }).done(function(){
        Swal.fire(
            "Sucesso",
            "Usuário cadastrado com sucesso",
            "success"
        ).then(function(){
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email: $('#email').val(),
                    senha: $('#senha').val()
                }
            }).done(function(){
                window.location = "/home"
            }).fail(function(){
                Swal.fire(
                    "OPS!",
                    "erro ao cadastrar o usuario",
                    "error"
                )
            })
        })
    }).fail(function(erro){
        Swal.fire(
            "Erro",
            "Erro ao cadastrar usuário",
            "error"
        );
    });
}