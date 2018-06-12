import Matchmaker from './matchmaker.js'
import Gameserver from './gameserver.js'

class Index {
	constructor() {
	  let matchmaker = new Matchmaker("ws://127.0.0.1:8001/ws", this.gameserverReady.bind(this));

	  matchmaker.joinQueue();
	}

	gameserverReady(gs) {
		let gameserver = new Gameserver(gs);
		gameserver.connect();
	}
}

new Index();
