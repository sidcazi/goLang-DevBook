$('#nova-publicacao').on('submit',CriarPublicacao);

function CriarPublicacao(evento){
    evento.preventDefault();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data:{
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        } 
    }).done(function(){
      window.location = "/home";
    }).fail(function(){
        alert("Erro ao criar Publicacao");
    })
}
