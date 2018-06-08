export default class Gameserver {
	constructor(ctx, gs, fixDPI) {
      this.ctx = ctx;
      this.gs = gs;
      this.fixDPI = fixDPI;
      this.offset = 20;
	}

    connect() {
      this.socket = new WebSocket(this.gs["Url"]);
      this.socket.onopen = this.socketOpened.bind(this);
      this.socket.onmessage = this.mapReceived.bind(this);
    }

    socketOpened() {
      console.log("gameserver opened");
      this.socket.send(JSON.stringify({
        'Name': 'Parth',
        'Token': 'token'
      }));
    }

    mapReceived(e) {
      this.colors = e.data;
      window.requestAnimationFrame(this.render.bind(this));
    }

    drawColors() {
      this.ctx.clearRect(0, 0, this.ctx.canvas.width, this.ctx.canvas.height);
      for (let r = 0; r < this.colors.length; r++) {
        for (let c = 0; c < this.colors[r].length; c++) {
          //this.ctx.fillStyle = "#" + this.colors[r][c].toString(16);
          this.ctx.fillStyle = "#000000";
          // TODO the lengths might be wrong here
          this.ctx.fillRect(
            this.offset + (this.ctx.canvas.width * r) / this.colors.length,
            this.offset + (this.ctx.canvas.height * c) / this.colors[r].length,
            ((this.ctx.canvas.width * (c + 1)) / this.colors.length) - this.offset,
            ((this.ctx.canvas.height * (r + 1)) / this.colors.length) - this.offset
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
