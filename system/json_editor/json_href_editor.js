
$(document).ready(function(){
    $('.json_href_editor, .json_href_editor_empty').click(function(){
        var is_empty=($(this).attr('class')=='json_href_editor_empty');
        var full_path=$(this).getPath2Editor();
        var split_ind=full_path.lastIndexOf('.');
        var path=full_path.substring(0, split_ind);
        var key=full_path.substring(split_ind+1);
        if(is_empty){
            var data=$('<input>', {
                'type': 'hidden',
                'name': 'path',
                'value': path
            }).add($('<input>', {
                'type': 'hidden',
                'name': 'key',
                'value': key
            })).add($('<input>', {
                'type': 'hidden',
                'name': 'append',
                'value': 'Append'
            })).serialize();
            var a_tag=$(this);
            $.post($(location).attr('pathname'), data, function(rtn){
                a_tag.attr('class', 'json_href_editor');
                a_tag.text(a_tag.text()+'...');
            });
        }else{
            $(this).attr('href', '?path='+$(this).getPath2Editor());
        }
    });
});
