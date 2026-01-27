// 内容の表示を後で実装する
function check(bodyText){
    if(window.confirm("この内容で送信してよろしいですか？\n"+bodyText)){
        return true;
    }else{
        window.alert("送信がキャンセルされました。");
        return false;
    }
}