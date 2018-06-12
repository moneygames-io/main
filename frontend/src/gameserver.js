import Canvasobject from './canvasobject.js'

export default class Gameserver extends Canvasobject {
	constructor(gs, fixDPI) {
	  super();
      this.gs = gs;
      this.offset = 5;
	}

    connect() {
      this.socket = new WebSocket(this.gs["Url"]);
      this.socket.onopen = this.socketOpened.bind(this);
      this.socket.onmessage = this.mapReceived.bind(this);
    }

    socketOpened() {
      this.socket.send(JSON.stringify({
        'Name': 'Parth',
        'Token': 'token'
      }));
    }

    mapReceived(e) {
      this.colors = JSON.parse(e.data);
      window.requestAnimationFrame(this.render.bind(this));
    }

    drawColors() {
	   let canvasWidth = super.getContext().canvas.width;
	   let canvasHeight = super.getContext().canvas.height;

	   let gameAreaSize = Math.min(canvasWidth, canvasHeight);

	   let gameAreaOffsetW = canvasWidth - gameAreaSize;
	   let gameAreaOffsetH = canvasHeight - gameAreaSize;
	   
      for (let r = 0; r < this.colors.length; r++) {
        for (let c = 0; c < this.colors[r].length; c++) {
          super.getContext().fillStyle = "#" + this.colors[r][c].toString(16);
          super.getContext().fillRect(
	    				(gameAreaSize / this.colors.length) * r + this.offset + (gameAreaOffsetW / 2),
	    				(gameAreaSize / this.colors[r].length) * c + this.offset + (gameAreaOffsetH / 2),
	    				(gameAreaSize / this.colors.length) - (2*this.offset),
	    				(gameAreaSize / this.colors[r].length) - (2*this.offset)
          );
        }
      }
    }

    render() {
      super.getContext().clearRect(0, 0, super.getContext().canvas.width, super.getContext().canvas.height);
      this.drawColors();
    }
}
