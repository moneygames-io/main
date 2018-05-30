export default class Matchmaker {
  constructor(matchmakingURL, canvasContext) {
    this.url = matchmakingURL;
    this.ctx = canvasContext;
  }

  joinQueue() {
    this.socket = new WebSocket(this.url);
    this.socket.onopen = this.socketOpened.bind(this);
    this.socket.onmessage = this.matchmakingMessage.bind(this);
  }

  socketOpened() {
    window.draw = this.render.bind(this);
    window.requestAnimationFrame(window.draw);
  }

  matchmakingMessage(e) {
    var data = JSON.parse(e.data);
    if (data['Url']) {
      this.joinGameServer(data['Url']);
    }

    if (data['Status']) {
      this.updateStatus(data['Status']);
    }
  }

  updateStatus(s) {
    console.log(s);
  }

  joinGameServer(gs) {
    console.log(gs);
  }
  
  render() {
    console.log('render called');
    this.ctx.clearRect(0, 0, this.ctx.canvas.width, this.ctx.canvas.height);
	this.ctx.beginPath();
	this.ctx.arc(this.ctx.canvas.width / 2, this.ctx.canvas.height / 2, 40,0,Math.PI);
	this.ctx.stroke();
    window.requestAnimationFrame(this.render.bind(this));
  }
}
