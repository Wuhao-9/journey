var topui = document.querySelector('.topui');
var lis = topui.querySelectorAll('li');
for (var i = 0; i < lis.length; i++){
    lis[i].onclick = function(){
        for (var i = 0; i<lis.length; i++){
            lis[i].className = 'middle_function';
        }
        this.className = 'middle_function current';
    }
}



