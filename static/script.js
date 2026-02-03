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

// 未入力項目チェック
function checkInput(form){
    const username = form.username;
    const password = form.password;

    let hasError = false;

    // 前回の赤色を消す
    username.classList.remove("input-error");
    password.classList.remove("input-error");

    if(username.value==""){
        username.classList.add("input-error");
        hasError = true;
    }

    if (password.value ==""){
        password.classList.add("input-error");
        hasError = true;
    }

    if(hasError==true){
        window.alert("未入力の項目があります");
        return false;
    }
    return true;
}