marked.setOptions({
    renderer: new marked.Renderer(),
    gfm: true,
    tables: true,
    breaks: false,
    pedantic: false,
    sanitize: true,
    smartLists: true,
    smartypants: false,
});

var conf = {
    EDIT_MODEL: 1,
    MODIFY_TIME: 0
};

$(function () {
    $("#editor-area").on("keydown",function (e) { // metaKey->command
        conf.MODIFY_TIME = (new Date()).valueOf();
        switch (e.keyCode){
            case 9: { // tab
                e.preventDefault();
                var indent = '    ';
                var start = this.selectionStart;
                var end = this.selectionEnd;
                var selected = window.getSelection().toString();
                selected = indent + selected.replace(/\n/g,'\n'+indent);
                this.value = this.value.substring(0,start) + selected + this.value.substring(end);
                this.setSelectionRange(start+indent.length,start+selected.length);
                break;
            }
            case 90: { // ctrl+z
                if (e.metaKey){
                    keyToCtrlZ(e);
                    break;
                }
            }
            case 83: { // ctrl+s
                if (e.metaKey){

                    break;
                }
                if (e.metaKey && e.shiftKey){

                }
                break;
            }
        }
    });
    // 定时保存
    saveLog();
    window.setInterval(function () {
        translate();
    }, 1000);
    // test
    $("#btn-marked").click(function () {
        if (conf.EDIT_MODEL == 1) {
            conf.EDIT_MODEL = 0;
            $("#btn-marked").removeClass("fa-eye");
            $("#btn-marked").addClass("fa-eye-slash")
        }else{
            conf.EDIT_MODEL = 1;
            $("#btn-marked").removeClass("fa-eye-slash")
            $("#btn-marked").addClass("fa-eye");
        }
        changeview();
    });
});
// key ctrl+z
var record = [];
function keyToCtrlZ(e){
    e.preventDefault();
    record.pop();
    $("#editor-area").val(record[record.length - 1]).blur();
    $("#editor-area").focus();
}
function saveLog(){
    window.setInterval(function () {
        if (record[record.length - 1] != $("#editor-area").val()) {
            record[record.length] = $("#editor-area").val();
        }
    }, 1500);
}
// get content
jQuery.editor ={
    getcontent: function(){
        return new Array($("#editor-title").val(),$("#editor-area").val())
    }   
}
// scheduler
function translate(){
    var now = (new Date()).valueOf();
    if (now - conf.MODIFY_TIME > 600 && (now - conf.MODIFY_TIME) < 1601){
        $("#editor-view").html(marked($("#editor-area").val()));
    }
}
// 转换编辑模式
function changeview(){
    switch (conf.EDIT_MODEL){
        case 1:{ // 即时浏览
            $("#editor-view").css("display","inline-block");
            $("#editor-area").css({"width":"50%"});
            break;
        }
        default :{ // 普通
            $("#editor-view").css("display","none");
            $("#editor-area").css({"width":"100%"});
        }
    }
}
