class Matchmaker {
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
	this.ctx.beginPath();
	this.ctx.arc(95,50,40,0,Math.PI);
	this.ctx.stroke();
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
}

module.exports = Matchmaker;
