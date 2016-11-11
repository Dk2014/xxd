
(function($){
    $.fn.getPath2Editor=function(){
        var ed=this;
        while(!$.isFunction(ed[0].getPath2Editor)){
            ed=ed.parent();
        }
        return(ed[0].getPath2Editor());
    };
    
    $.fn.initJsonEditor=function(){
        this.each(function(ind, ele){
            if($.isFunction($[$(ele).attr('class')+'_initist'])){
                $[$(ele).attr('class')+'_initist'](ele);
            };
        });
    };
    
    $.fn.serializeURLParams=function(){
        var result={};
    
        if(this.attr("href").indexOf("?")==-1) 
            return(result);
    
        var pairs=this.attr("href").split("?")[1].split('&');
        pairs.forEach(function(pair){
            pair=pair.split('=');
            var name=decodeURI(pair[0]);
            var value=decodeURI(pair[1]);
            if(name.length)
                if(result[name]!==undefined){
                    if(!result[name].push){
                        result[name]=[result[name]];
                    }
                    result[name].push(value||'');
                }else{
                    result[name]=value||'';
                }
        });
        return(result);
    };
}(jQuery));

$(document).ready(function(){
    $('body')[0].getPath2Editor=function(){
        var query_obj=$(location).serializeURLParams();
        var id;
        if(query_obj.path===undefined)
            id='';
        else
            id=query_obj.path;
        return(id);
    };
});
