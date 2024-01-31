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
1. create a .env file in the base directory with `API_KEY="your-openai-api-key-here"`
2. use `go run main.go` in the base directory
3. open `localhost:8000` in your browser


## A Long List of Future Features That May or May Not Get Done at Any Point 
- [ ] login
- [ ] persistent storage of a user's characters
- [ ] dice roller for saves/attacks/etc.
- [ ] upload notes for data retrieval from AI
- [ ] create multiple characters at the same time for encounters
- [ ] use DALL-E to create a picture of a character based on their character sheet or other information.
- [ ] incorporate subscription and actually make money?
- [ ] create a map from description of cities with information like population, images, government, etc.
