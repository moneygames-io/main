import Matchmaker from './matchmaker.js';

window.draw = function() {console.log("draw");}

function main() {
  var canvas = createCanvas();
  document.body.appendChild(canvas);
  window.onresize = function() {
	canvas.width = window.innerWidth;
	canvas.height = window.innerHeight;
	window.draw();
  }

  var context = canvas.getContext('2d');
  var matchmaker = new Matchmaker("ws://127.0.0.1:8001/ws", context);

  matchmaker.joinQueue();
}

function createCanvas() {
  var canv = document.createElement('canvas');
  canv.width = window.innerWidth;
  canv.height = window.innerHeight;
  return canv;
}

main();
