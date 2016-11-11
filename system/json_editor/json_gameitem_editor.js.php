<?php

require_once '../lib/global.php';


$all_items=get_all_item();

?>

$(document).ready(function(){
    <?=get_items_json($all_items, 'items');?>
    
    var json_gameitem_editor_initist=function(ele){
        if($(ele).attr('name')=='dummy_dummy_display'){
            $(ele).autocomplete({
                    lookup: items,
                    minChars: 0,
                    onSelect: function(s){
                        $(this).siblings('.json_gameitem_editor[name=value]').val(s.id);
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
                    }
            });
        }
    };
    
    $['json_gameitem_editor_initist']=json_gameitem_editor_initist;

    $('.json_gameitem_editor[name=dummy_dummy_display]:not(.table_addition_empty_row .json_gameitem_editor[name=dummy_dummy_display])')
    .each(function(ind, ele){
        json_gameitem_editor_initist(ele);
    });
});

<?php
