import Matchmaker from './matchmaker.js'
import Gameserver from './gameserver.js'

function main() {
  var canvas = createCanvas();
  document.body.appendChild(canvas);
  fixDPI();

  var context = canvas.getContext('2d');
  var matchmaker = new Matchmaker("ws://127.0.0.1:8001/ws", context, gameserverReady.bind(this));
  var gameserver = new Gameserver(context, undefined);
  window.requestAnimationFrame(render);

  matchmaker.joinQueue();
  window.currentView = matchmaker;
}

function gameserverReady(gameserver) {
	console.log(gameserver);
}

function render() {
  fixDPI();
  window.currentView.render();
  window.requestAnimationFrame(render);
}

function createCanvas() {
  var canv = document.createElement('canvas');
  canv.id = "canv";

  window.addEventListener('resize', fixDPI, false);
  return canv;
}

function fixDPI() {
  var canv = document.getElementById('canv');
  var dpi = window.devicePixelRatio;
  let style = {
    height() {
      return +getComputedStyle(canv).getPropertyValue('height').slice(0,-2);
    },
    width() {
      return +getComputedStyle(canv).getPropertyValue('width').slice(0,-2);
    }
  }

  canv.setAttribute('width', style.width() * dpi);
  canv.setAttribute('height', style.height() * dpi);
}

main();
