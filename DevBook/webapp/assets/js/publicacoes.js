$('#nova-publicacao').on('submit',CriarPublicacao);

function CriarPublicacao(evento){
    evento.preventDefault();

    $.ajax({
        url: "/publicacoes",
        mehtod: "POST",
        data:{
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        } 
    }).done(function(){
        alert("Publicação Criada com Sucesso");
    }).fail(function(){
        alert("Erro ao criar Publicacao");
    })
}
