## Pen.go [WIP]
 Inspired by the classic arcade game [Pengo](https://en.wikipedia.org/wiki/Pengo_(video_game)). I have an original table top arcade version of this game and thought that since it has Go in the name..

The game takes place in an overhead maze made of ice blocks, where Pengo (red block) fights the Sno-Bees (yellow blocks). The objective of the game is for Pengo to survive a series of rounds by eliminating all Sno-Bees, whil amassing bonuses by bringing together the three diamonds dispersed in the maze.

### Installation

~~~
git clone https://github.com/fergstar/go_game_jam
go get -u github.com/JoelOtter/termloop
cd go_game_jam
go run *.go
~~~

### Controls 
Use the arrow keys to move Pengo around the maze.
Press the space bar to push an Iceblock.

### Game Jam Post Mortem

Although I didn't get as much done in the two days of the game jam as I would have liked, it doesn't help having a 5 month old baby to look after. 

I'm going to finish this until I have a version of Pengo that matches, as close as possible, to the old arcade version.

### TODO
- [DONE] Pengo
- [DONE] Iceblocks
- [DONE] Push blocks
- Diamond Blocks
- Pengo can crush blocks
- SnoBee's with Movement/AI
- Crushing SnoBee's
- Scoring
- Pengo Lives
- Additional levels
- Maze generation like original
- Wobbly walls
- Intermission screens
- Time bonuses
