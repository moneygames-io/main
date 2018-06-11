export default class Gameserver {
	constructor(ctx, gs, fixDPI) {
      this.ctx = ctx;
      this.gs = gs;
      this.fixDPI = fixDPI;
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
       this.ctx.clearRect(0, 0, this.ctx.canvas.width, this.ctx.canvas.height);

	   let canvasWidth = this.ctx.canvas.width;
	   let canvasHeight = this.ctx.canvas.height;

	   let gameAreaSize = Math.min(canvasWidth, canvasHeight);

	   let gameAreaOffsetW = canvasWidth - gameAreaSize;
	   let gameAreaOffsetH = canvasHeight - gameAreaSize;
	   
      for (let r = 0; r < this.colors.length; r++) {
        for (let c = 0; c < this.colors[r].length; c++) {
          //this.ctx.fillStyle = "#" + this.colors[r][c].toString(16);
          this.ctx.fillStyle = "#C0C0C0";
          this.ctx.fillRect(
	    				(gameAreaSize / this.colors.length) * r + this.offset + (gameAreaOffsetW / 2),
	    				(gameAreaSize / this.colors[r].length) * c + this.offset + (gameAreaOffsetH / 2),
	    				(gameAreaSize / this.colors.length) - (2*this.offset),
	    				(gameAreaSize / this.colors[r].length) - (2*this.offset)
          );
        }
      }
    }

    render() {
      this.fixDPI();
      this.ctx.clearRect(0, 0, this.ctx.canvas.width, this.ctx.canvas.height);
      this.drawColors();
    }
}
