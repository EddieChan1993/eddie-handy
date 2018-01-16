$(function(){
    var name = getUrlParam("name");
    var websocket_domain = "ws://115.28.215.160:8686/chat";
    // var websocket_domain = "ws://127.0.0.1:8686/chat";
    var ws =new WebSocket(websocket_domain)
    ws.onopen = function (res) {
        ws.send(sendMessAll(name,"进入房间",'connect'))
    };

    ws.onmessage = function (res) {
        var data = jQuery.parseJSON(res.data);
        var mess="<strong>"+data.name+":</strong>"+data.content
        switch (data.type){
            case "tel_self":
                $("#typed").typed({
                    strings:[mess],
                    callback:function (res) {
                    }
                });
                break
            case "join_room":
                $("#typed").typed({
                    strings:[mess],
                    callback:function (res) {
                    }
                });
                break
            case "send_all":
                $("#typed").typed({
                    strings:["<strong>"+data.name+":</strong>"+data.content]
                });
                break;
        }
    };
    ws.onclose = function (res) {
        console.log(res)
    };


    $('.send-btn').click(function () {
        layer.prompt({title: '输入你传达的', formType: 2}, function(text, index){
            ws.send(sendMessAll(name,text, "all"));
            layer.close(index);
            var mess="<strong>自己:</strong>"+text
            $("#typed").typed({
                strings:[mess],
                callback:function () {
                }
            });
        });
    });

    $(".reset").click(function(){
        $("#typed").typed({stringsElement:$("#typed-strings")});
    });

});

function newTyped(){ console.log("reset")/* A new typed object */ }

function foo(){}

function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
    var r = window.location.search.substr(1).match(reg);  //匹配目标参数
    if (r != null) return unescape(r[2]); return null; //返回参数值
}

function sendMessAll(name, data, type) {
    return JSON.stringify(
        {
            data: JSON.stringify({
                name:name,
                age:123
            }),
            type: type
        }
    )
}