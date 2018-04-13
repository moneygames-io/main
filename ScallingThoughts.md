Some thoughts on scalling:

Frontend (static assets js/html/css/future native binary apps) will be hosted on content delivery networks. They're free, super fast, and super easy to update.

The frontends will talk to a Matchmaker which will be a traditional rest server. Scaling this is pretty easy and well understood.  Many options, low cost, etc. We will scale this by region.

The matchmaker will keep track of all our game servers. These game servers are responsible for running our game logic. The game servers will tell matchmakers when they come online. Based on region and load the matchmaker will handle the player queues, payment processing, and game state. 
