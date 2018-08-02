import Matchmaker from './matchmaker.js'
import Gameserver from './gameserver.js'

class Index {
	constructor() {
	  this.matchmaker = new Matchmaker("ws://" + window.location.hostname + ":8000/ws", this.gameserverReady.bind(this));
	  this.matchmaker.joinQueue();
	}

    gameserverReady(gs) {
        this.matchmaker.destroyCanvas();
        let gameserver = new Gameserver(gs);
        gameserver.connect();
    }
}

new Index();
