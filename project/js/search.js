var R_current = document.querySelector('div.R_current');
var hot_search_header = document.querySelector('div.hot_search_header');
var hot_search = document.querySelector('div.hot_search');
var hot_expend = R_current.querySelector('.expend');
var hot_text = R_current.querySelector('.text');
R_current.onclick = function(){
    if(hot_search.style.display != 'none'){
        hot_search.style.display = 'none';  
        hot_search_header.className = 'hot_search_header hot_search_headerClick';
        hot_expend.className = 'expend expendclick';
        hot_text.innerHTML="展开热门";
    }else{
        hot_search.style.display = 'block';
        hot_search_header.className = 'hot_search_header';
        hot_expend.className = 'expend';
        hot_text.innerHTML="收起热门";
    }   
    
}