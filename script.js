//var socket = new WebSocket("ws://localhost:8080/echo");
var socket = new WebSocket("ws://lab.posevin.com:4111");
var pos =0;
document.getElementById("input").focus();


socket.onopen = function() {
};

var t = "encryption)";

function revers(){
	pos ++;
	pos%=2;
	t="encryption)";
	if (pos==1){ t= "decoding)"}
	document.getElementById("back").style.background= "url(f"+ pos+".png )";
	document.getElementById("status").innerHTML= "("+t;
}

function sendMess(){
	let mess = document.getElementById("input").value;
	mess=pos+mess;
	socket.send(mess);
	document.getElementById("input").placeholder="The message \"" + mess + "\" have been sent";
	//document.getElementById("input").value="";
};

socket.onmessage = function(event){
	document.getElementById("output").value=event.data;
	document.getElementById("status").innerHTML="The anwser is hear, look) ("+t ;
};

socket.onerror = function(error) {
  alert("Ошибка " + error.message);
};