export default class Gameserver {
	constructor(ctx, gs) {
      this.ctx = canvasContext;
      console.log(gs);
	}

    render() {
      this.ctx.clearRect(0, 0, this.ctx.canvas.width, this.ctx.canvas.height);
    }
}
