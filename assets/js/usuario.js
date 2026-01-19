$('#parar-de-seguir').on('click', pararDeSeguir);
$('#seguir').on('click', seguir)
$('#editar-usuarios').on('submit', editar)

function pararDeSeguir() {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/parar-de-seguir`,
        method: "POST"
    }).done(function () {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function () {
        Swal.fire("Ops", "Erro ao parar de seguir o usuario", "error");
        $('#parar-de-seguir').prop('disabled', false);
    })
}

function seguir() {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/seguir`,
        method: "POST"
    }).done(function () {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function () {
        Swal.fire("Ops", "Erro ao seguir o usuario", "error");
        $('#seguir').prop('disabled', false);
    })
}

function editar(evento){
    evento.preventDefault();

    $.ajax({
        url: "/editar-usuarios",
        method: "PUT",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val()
        }
    }).done(function(){
        SWal.fire("sucesso", "usuario atualizado com sucesso", "success")
        .then(function(){
            window.location = "/perfil";
        })
    }).fail(function(){
        Swal.fire("ops!", "erro ao atualizar usuarios", "error");
    })
}