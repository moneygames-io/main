export default class Matchmaker {
  constructor(matchmakingURL, canvasContext) {
    this.url = matchmakingURL;
    this.ctx = canvasContext;
	this.progress = 0;
  }

  joinQueue() {
    this.socket = new WebSocket(this.url);
    this.socket.onopen = this.socketOpened.bind(this);
    this.socket.onmessage = this.matchmakingMessage.bind(this);
  }

  socketOpened() {
    window.currentDraw = this.render.bind(this);
    window.requestAnimationFrame(window.currentDraw);
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
  	this.progress = s.CurrentClients / s.TargetClients;
	this.render();
  }

  joinGameServer(gs) {
    console.log(gs);
  }
  
  render() {
  	console.log("rendered");
    this.ctx.clearRect(0, 0, this.ctx.canvas.width, this.ctx.canvas.height);
	this.ctx.beginPath();
	this.ctx.lineWidth = 5;
	this.ctx.strokeStyle = 'rgb(108, 116, 128)';
	this.ctx.arc(
		this.ctx.canvas.width / 2, 
		this.ctx.canvas.height / 2, 
		50, 
		1.5 * Math.PI, 
		this.progress * 2 * Math.PI + 1.5 * Math.PI
	);
	this.ctx.stroke();
  }
}
