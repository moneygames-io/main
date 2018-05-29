import Game from './game.js';

function canvas() {
  var canv = document.createElement('canvas');
  setupCanvas(canv);

  return canv;
}

function testImport() {
	var test = document.createElement('p');
	var a = new Game();
	test.innerHTML = a.render();
	return test;
}

function setupCanvas(c) {
	c.width = document.body.clientWidth;
	c.heigh = document.body.clientHeight;
}


document.body.appendChild(testImport());
