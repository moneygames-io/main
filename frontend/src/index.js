import _ from 'lodash';

function canvas() {
  var canv = document.createElement('canvas');
  setupCanvas(canv);

  return canv;
}

function setupCanvas(c) {
	c.width = document.body.clientWidth;
	c.heigh = document.body.clientHeight;
	var link
}


document.body.appendChild(canvas());
