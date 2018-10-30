var socket = new WebSocket("ws://lab.posevin.com:4011/ws");
alert("The socket is upped");

document.getElementById("input").focus();

socket.onopen = function() {
	alert("...");
};

function sendMess(){
	let mess = document.getElementById("input").value;
	socket.sent(mess);
	alert("The message \"" + mess + "\" have been sent");
}

socket.onmessage = function(event){
	document.getElementById("output").value=event.data;
	alert("The anwser is hear, look)");
}

