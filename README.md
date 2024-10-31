# techlanglearn

- on this site I record notes while watching a wide range of tech video tutorials in numerous foreign languages, mainly using the following sources for tutorials:
  - **LinkedInLearning** 
    - has tutorials in English, German, French and Spanish 
    - has accurate subtitles in dozens of languages 
    - so I can watch e.g. Spanish tutorials with Spanish subtitles 
    - I can also watch English tutorials with Dutch, Polish or Italian subtitles
  - **YouTube** 
    - has tutorials in English, German, French and Spanish
    - subtitles are automated and hence error prone, but help in understanding the language
    - you can also watch most videos with subtitles in another language, but of course also automated
- I not only record notes on the technical aspects of the videos
  - but also useful phrases that I want to learn in various languages
  - these are presented to me to take as flashcards along with the tech notes that I take
- hence this site helps me improve my web development skills as well as my foreign language skills, mostly in Spanish, French and Italian, hence the name **techlanglearn**
- in addition, the site uses tools, frameworks and programming languages that I want to learn including:
  - [Go](https://go.dev) to parse the .md files with a CLI script, using [Cobra](https://cobra.dev)
  - frontend is in [SvelteKit5](https://svelte.dev/blog/svelte-5-is-alive) released Oct 22, 2024
  - eventually the Go script will save the data in an [ArangoDB](https://arangodb.com) database and a Go backend will serve the data via [GraphQL](https://graphql.org)

## setup

- `npm i`
- `npm run dev`

## parse md files to json (under construction)

- `npm run pd` (pd = parse data)
