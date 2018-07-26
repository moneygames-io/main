# sneks (this document is currently out of date)

Sneks is going to be a really simple game moddeled after [slither.io](slither.io), except we're going to conduct bouts [battle royale style](https://en.wikipedia.org/wiki/Battle_royale_game), and allow people to use [crypto currencies](https://en.wikipedia.org/wiki/Cryptocurrency) to gamble on the outcomes of games.

Additionally other game modes may be added in the future for different playstyles. A game mode that encourages the collection of map resources and "escaping" to cash out might be fun as well. The playstyle differences will largely be a fast paced "winner take all" vs. a slower and more rewarding "grind" to try to grind more effectively than other players on the map.

## Game Mechanics Overview

### Matchmaking and Buy In
	
Players deposit a small amount of money into a specified wallet (different for each player). Upon receiving funds for the wallet they will enter the matchmaking queue. At any point prior to the game starting they may cancel for a penalty. After a fixed amount of time (based on # of active players), the game will begin. If more money is entered into their designated wallet, only the minimum buy in will be debitted, the remaining will be refunded to sender's address.

### What are we doing with the buy in?

Once a fixed number of players are found, each player's buy in will be utilized in the following manner: 80% towards the *length* of their snake, 10% converted to *food* distributed randomly throughout the map, 10% site fee. 

### Game mechanics

The *map* will simply be a large 2D array of tuples (which player, and height). Players will play on this *map*. Players will manifest themselves as *snakes*. *Snakes* have a *head* and a *tail*. *Snakes* will be represented as a _circular_, _doubly linked list_. Some sort of data structure will keep track of all the *heads*.

If a *head* of a *snake* colides with a *wall*, or with the *tail* of another *snake* the *snake* *dies*. Upon *death* that *snake*'s current *length* will be converted to food and scattered in the general vicinity of the body of that *snake*. Once a *snake* is *dead* they are removed from the game, and brought back to the main screen where they can buy in again. 

As the number of snakes, and the food available decreases, the map will shrink to keep game play lively.

*Snakes* can *sprint* in order to try to cut other *snakes* off. They do this at the expense of *length* which is converted to *food* behind their *tail*.

The last *snake* remaining after all *snakes* have been eliminated earns the entire prize pool (90% of everyone's buyin).

### Misc rules for interesting gameplay

+ If the final 2 *snakes* consume all the food on the map, they split the pot relative to their size. They can excrete more food simply by sprinting. 
+ Players outside map area will continue to lose length (food distributed randomly in actual play area). If length falls under initial buy in amount, they will have a fixed amount of time where they can go through some redemption. Can work through details of that later.

### Game logic

Game opperations that need to be conducted "every frame":

For all heads:

0. Is head intersecting a tail? Dead
1. Delete the tail 
2. Add a head in next location 
3. Is head intersecting a tail? Dead

Should be O(n) where n is number of heads.

## Architecture

All game computation happens server-side. 

Clients are responsbile for:
1. Asking for a map centered on head location at specified zoom level
2. Receiving a map, and showing it

Will likely be a very low amount of information being sent back and forth at a very high rate.

# Build and run instructions

The game is comprised of 3 microservices: a *frontend* that talks to our servers. A *matchmaker* that pools clients and assigns them to gameservers. And *gameserver*s, where actual gameplay is processed. 

Here are platform agnostic instructions to run these services.

Dependencies:
* Go
* npm

## frontend

```
cd frontend
npm i
npm run serve
```

## matchmaker

```
cd matchmaker
# TODO fetch deps
go build && ./matchmaker
```

## gameserver

```
cd gameserver
# TODO fetch deps
go build && ./gameserver
```
