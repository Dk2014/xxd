
$(document).ready(function(){
    $('.json_bool_editor, .json_number_editor, .json_content_editor'
            +' .json_rawjson_editor, .json_text_editor')
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
