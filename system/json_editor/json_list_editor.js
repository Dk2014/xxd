
json_list_id_onclick=function(e){
    $(this).attr('href',
                 '?path='
                 +$(this).parent().parent().parent().parent().getPath2Editor()
                 +'.'+$(this).text());
};

json_list_row_rm_onclick=function(e){
    var list_row=$(this).parent().parent();
    var deleting_id=list_row.find('.json_list_item_id').text();
    var data=$('<input>', {
        'type': 'hidden',
        'name': 'path',
        'value': list_row.parent().parent().getPath2Editor()+'.'+deleting_id
    }).add($('<input>', {
        'type': 'hidden',
        'name': 'delete',
        'value': 'Delete'
    })).serialize();
    $.post($(location).attr('pathname'), data, function(rtn){
        list_row.remove();
    });
};

json_list_new_row_onclick=function(e){
    var list_row=$(this).parent().parent();
    var row_id=prompt('Row ID:', '');
    var data=$('<input>', {
        'type': 'hidden',
        'name': 'path',
        'value': list_row.parent().parent().getPath2Editor()+'.'+row_id
    }).add($('<input>', {
        'type': 'hidden',
        'name': 'append',
        'value': 'Append'
    })).serialize();
    $.post($(location).attr('pathname'), data, function(rtn){
        $('<tr>', {}).append($('<td>', {}).append($('<a>', {
            'class': 'json_list_item_id',
            'href': '#',
            'text': row_id
        }).click(json_list_id_onclick))).append($('<td>', {
            'html': '<textarea readonly></textarea>'
        })).append($('<td>', {}).append($('<button>', {
            'class': 'delete_list_row',
            'text': 'rm'
        }).click(json_list_row_rm_onclick)))
        .insertBefore(list_row);
    });
    e.preventDefault();
};

$(document).ready(function(){
    $('.json_list_editor .json_list_item_id')
    .click(json_list_id_onclick);

    $('.json_list_editor .delete_list_row')
    .click(json_list_row_rm_onclick);

    $('.json_list_editor .new_list_row')
    .click(json_list_new_row_onclick);
});
