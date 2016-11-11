
$(document).ready(function(){
    $('.json_IsRelative_editor')
            .each(function(ind, ele){
                $(ele).focusout(function(event){
                    var data=$(this)
                    .add($(this).siblings('.'+$(this).attr('class')))
                    .add($('<input>', {
                        'type': 'hidden',
                        'name': 'path',
                        'value': $(this).getPath2Editor()
                    }))
                    .add($('<input>', {
                        'type': 'hidden',
                        'name': 'update',
                        'value': 'Update'
                    })).serialize();
                    $.post($(location).attr('pathname'), data, function(rtn){

                    }); 
                });
            });
});
