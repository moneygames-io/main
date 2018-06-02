import Matchmaker from './matchmaker.js'
import Gameserver from './gameserver.js'

class Index {
	constructor() {
	  this.canvas = this.createCanvas();
	  document.body.appendChild(this.canvas);
	  this.fixDPI();

	  this.context = this.canvas.getContext('2d');
	  let matchmaker = new Matchmaker("ws://127.0.0.1:8001/ws", this.context, this.gameserverReady.bind(this));
	  window.requestAnimationFrame(this.render.bind(this));

	  matchmaker.joinQueue();
	  window.currentView = matchmaker;
	}

	gameserverReady(gs) {
		let gameserver = new Gameserver(this.ctx, gs);
		gameserver.connect();
	}

	/**
	 * TODO should anyone else call render? 
	 * TODO I think everyone should handle their own rendering and we should be renamed to something like "clear" that gets called at the start of every animation frame. (can remove request then) and not keep track of window.currentView
	 */
	render() {
	  this.fixDPI();
	  window.currentView.render();
	  window.requestAnimationFrame(this.render.bind(this));
	}

	createCanvas() {
	  let canv = document.createElement('canvas');
	  canv.id = "canv";

	  window.addEventListener('resize', this.fixDPI, false);
	  return canv;
	}

	height() {
	  return +getComputedStyle(this.canvas).getPropertyValue('height').slice(0,-2);
	}

	width() {
	  return +getComputedStyle(this.canvas).getPropertyValue('width').slice(0,-2);
	}

	fixDPI() {
	  let dpi = window.devicePixelRatio;

	  this.canvas.setAttribute('width', this.width() * dpi);
	  this.canvas.setAttribute('height', this.height() * dpi);
	}

}

new Index();
