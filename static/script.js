// ポストの投稿チェック
function checkSubmit(bodyText){
    if(window.confirm("この内容で送信してよろしいですか？\n"+bodyText)){
        return true;
    }else{
        window.alert("送信がキャンセルされました。");
        return false;
    }
}

// ポストを削除できるかチェック
// (一旦削除ボタンの挙動確認。後にDB連携)

function isPostOwner(userId=0, postUserId=0){
    if(userId === postUserId){
        return true;
    }else{
        return false;
    }
}