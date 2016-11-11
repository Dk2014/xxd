
json_table_row_rm_onclick=function(e){
    var table_row=$(this).parent().parent();
    var deleting_id=table_row.find('.json_table_row_id').text();
    var data=$('<input>', {
        'type': 'hidden',
        'name': 'path',
        'value': table_row.parent().parent().getPath2Editor()+'.'+deleting_id
    }).add($('<input>', {
        'type': 'hidden',
        'name': 'delete',
        'value': 'Delete'
    })).serialize();
    $.post($(location).attr('pathname'), data, function(rtn){
        table_row.remove();
    });
};

json_table_new_row_onclick=function(e){
    var table_row=$(this).parent().parent();
    var dummy_row=table_row.prev('tr');
    var new_row=dummy_row.clone(true);
    new_row.attr('class', '');
    var row_id;
    if(!$.isNumeric(dummy_row.find('.json_table_row_id').text())){
        row_id=prompt('Row ID:', '');
        new_row.find('.json_table_row_id').text(row_id);
    }else{
        row_id=new_row.find('.json_table_row_id').text();
        dummy_row.find('.json_table_row_id').text(
                parseInt(row_id)+1
        );
    }
    new_row[0].getPath2Editor=table_tr_getPath2Editor;
    new_row.find('td').each(function(ind, ele){
        ele.getPath2Editor=table_tr_td_getPath2Editor;
        $(ele).children().initJsonEditor();
    });
    var data=$('<input>', {
        'type': 'hidden',
        'name': 'path',
        'value': table_row.parent().parent().getPath2Editor()
    }).add($('<input>', {
        'type': 'hidden',
        'name': 'key',
        'value': row_id
    })).add($('<input>', {
        'type': 'hidden',
        'name': 'append',
        'value': 'Append'
    })).serialize();
    $.post($(location).attr('pathname'), data, function(rtn){
        new_row.show();
        new_row.insertBefore(dummy_row);
    });
    e.preventDefault();
};

table_tr_getPath2Editor=function(){
    return($(this).parent().getPath2Editor()
            +'.'+$(this).find('.json_table_row_id').text());
};

table_tr_td_getPath2Editor=function(){
    var table=$(this).parent().parent().parent();
    var ind=$(this).prevAll('td').length;
    var col_id=table.find('th:eq('+ind+') .json_table_column').val();
    return($(this).parent().getPath2Editor()+'.'+col_id);
};

$(document).ready(function(){
    $('.json_table_editor tbody tr').each(function(index, element){
        element.getPath2Editor=table_tr_getPath2Editor;
    });
    $('.json_table_editor tbody tr').each(function(index, element){
        $(element).find('td').each(function(ind, ele){
            ele.getPath2Editor=table_tr_td_getPath2Editor;
        });
    });
    
    $('.json_table_editor .table_addition_empty_row').hide();
    
    $('.json_table_editor .delete_table_row')
    .click(json_table_row_rm_onclick);
    
    $('.json_table_editor .new_table_row')
    .click(json_table_new_row_onclick);
});
