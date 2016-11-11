
$(document).ready(function(){

    $('.json_timestamp_editor[name=value]').each(function(ind, ele){
        var date=new Date(parseInt($(ele).val())*1000);
        var date_str=date.getFullYear()+'/'+(date.getMonth()+1)+'/'+date.getDate()
            +' '+date.getHours()+':'+date.getMinutes();
        $(ele).val(date_str);
    });

    var json_timestamp_editor_initist=function(ele){
        if($(ele).attr('name')=='value'){
            $(ele).datetimepicker({
                format: 'Y/m/d H:i'
            });
        }
    };
    
    $['json_timestamp_editor_initist']=json_timestamp_editor_initist;
    
    $('.json_timestamp_editor[name=value]:not(.table_addition_empty_row .json_timestamp_editor[name=value])')
    .each(function(ind, ele){
        json_timestamp_editor_initist(ele);
    });
    
    $('.json_timestamp_editor[name=value]').focusout(function(event){
        var timestamp=Date.parse($(this).val())/1000;
        var data=$(this).siblings('.'+$(this).attr('class'))
        .add($('<input>', {
            'type': 'hidden',
            'name': 'path',
            'value': $(this).getPath2Editor()
        }))
        .add($('<input>', {
            'type': 'hidden',
            'name': 'value',
            'value': ''+timestamp
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
