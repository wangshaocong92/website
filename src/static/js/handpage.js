
function changeBackgroude() {
    var img = document.getElementById("icon_img");
    var href = document.getElementById("icon_href");
    var sel = document.getElementById("select_engine")
    var str =sel.options[sel.selectedIndex].value;
    if (str == "baidu"){
        img.src="../image/baidu_logo.png";
        href.href="https://www.baidu.com/";
    }
    else if (str == "google"){
        href.href="https://www.google.com/";
        img.src="../image/google_logo.png"

    }else{
        href.href="https://cn.bing.com/";
        img.src="https://upload.wikimedia.org/wikipedia/commons/thumb/3/35/Bing_Logo_2016.svg/2000px-Bing_Logo_2016.svg.png";
    }
}
