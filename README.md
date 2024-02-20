## What?
This is a webapp that acts as an interface to OpenAI API for Dungeon Masters. The first feature under construction takes a prompt like:
`create an enemy strong enough to fight 3 level 3 players that has spellcasting`
and displays a character sheet complete with ability scores, spell information, etc.

Much of this functionality is still a work in progress, but there is some basic functionality present currently.

## Why?
This is a small personal project that I am using as a way to both learn Go and brush up on my HTML/X.
Also, as someone who DM's once a week I think this could be a useful time-saving tool

## When?
This is a slow moving project as it is something I do in my spare time.

## How?
To run:

Requirements:
1. Go
2. Docker

Steps:
1. create a .env file in the base directory with 
```
DB_HOST="localhost"
DB_USER="postgres"
DB_PASSWORD="your-db-password-here"
DB_PORT="5432"
DB_NAME="dnd"
API_KEY="your-openai-api-key-here
```
2. run docker compose -f compose.yml up to start the postgres server
3. use `go run main.go` in the base directory to start the application
4. open `localhost:8000` in your browser to view the application
5. run docker compose -f compose.yml down to stop the postgres server

Note: 
There is currently no loading screen to know when the character creation is in progress.
Having openAI generate spell text will be the most time/token expensive part of the process. 
The solution I've implemented is storing this data in the SQL database so that is only generated once. 
That being said if you make a spell caster and you have not already generated the information for the spells that it gives
it will take quite a while, just be patient!

## A Long List of Future Features That May or May Not Get Done at Any Point 
- [ ] login
- [ ] persistent storage of a user's characters
- [ ] dice roller for saves/attacks/etc.
- [ ] upload notes for data retrieval from AI
- [ ] create multiple characters at the same time for encounters
- [ ] use DALL-E to create a picture of a character based on their character sheet or other information.
- [ ] create a map from description of cities with information like population, images, government, etc.
