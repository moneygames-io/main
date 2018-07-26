import Matchmaker from './matchmaker.js'
import Gameserver from './gameserver.js'

class Index {
<<<<<<< HEAD
	constructor() {
	  this.matchmaker = new Matchmaker("ws://" + window.location.hostname + ":8000/ws", this.gameserverReady.bind(this));
	  this.matchmaker.joinQueue();
	}
=======
    constructor() {
        this.matchmaker = new Matchmaker("ws://" + window.location.hostname + ":8000/ws", this.gameserverReady.bind(this));

        this.matchmaker.joinQueue();
    }
>>>>>>> b3586b65ec65e049c9f5a66e7b0b7a924972b490

    gameserverReady(gs) {
        this.matchmaker.destroyCanvas();
        let gameserver = new Gameserver(gs);
        gameserver.connect();
    }
}

new Index();