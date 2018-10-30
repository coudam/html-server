var socket = new WebSocket("ws://lab.posevin.com:3341/ws");
alert("The socket is upped");

document.getElementById("input").focus();

socket.onopen = function() {
	alert("...");
};

function sendMess(){
	let mess = document.getElementById("input").value;
	// ...
	alert("The message \"" + mess + "\" have been sent");
}

function getMess(){
	let ans= "lol";
	// ... 
	document.getElementById("output").value=ans;
	alert("The anwser is hear, look)");
}