var socket = new WebSocket("ws://localhost:8080/echo");

document.getElementById("input").focus();

socket.onopen = function() {
	alert("...");
};

function sendMess(){
	let mess = document.getElementById("input").value;
	socket.send(mess);
	alert("The message \"" + mess + "\" have been sent");
};

socket.onmessage = function(event){
	alert();
	document.getElementById("output").value=event.data;
	alert("The anwser is hear, look)");
};

socket.onerror = function(error) {
  alert("Ошибка " + error.message);
};