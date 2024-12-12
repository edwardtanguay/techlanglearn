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
- in addition, I use frameworks, programming languages, and tools that I want to get practice experience with and improve my skills in, including:
  - [Go](https://go.dev) to parse the .md files with a CLI script, using [Cobra](https://cobra.dev)
  - frontend is in [SvelteKit5](https://svelte.dev/blog/svelte-5-is-alive) released Oct 22, 2024
  - eventually the Go script will save the data in an [ArangoDB](https://arangodb.com) database and a Go backend will serve the data via [GraphQL](https://graphql.org)

## live site

- I practice **Minimal Viable Product** and **iterative development** in this project
- this means that, while the site is very new
- I am already watching videos and taking tech and language notes 
- these are available on this [live site](https://techlanglearn.vercel.app)
- many of the videos I am currently watching (Go, SvelteKit 5) help me build skills to develop the site

## setup

- `npm i`
- `npm run dev`
- **.env**

```
VITE_SITE_LOCATION=dev
```

## parse md files to json (under construction)

- `npm run pd` (pd = parse data)

## create new note files

- in **courses.tts.txt** add a line like this:

```
>>create>>python;en;00:09:04;2024;4.2;https://www.youtube.com/watch?v=Y21OR1OPC9A; Python Virtual Environments - Full Tutorial for Beginners ; shows how to install a Python virtual environment
```
