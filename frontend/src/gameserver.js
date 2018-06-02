export default class Gameserver {
	constructor(ctx, gs) {
      this.ctx = ctx;
      this.gs = gs;
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
      console.log(e);
    }

    render() {
      this.ctx.clearRect(0, 0, this.ctx.canvas.width, this.ctx.canvas.height);
    }
}
