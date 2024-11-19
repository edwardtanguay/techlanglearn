# Build GraphQL Server with Go

- https://www.youtube.com/watch?v=wfBeEn8d-8A&list=PL3eAkoh7fypqxiqDq9OQGm3aJO7dZjSBI
- 02:30:00
- english
- go, graphql
- 2020
- 4.8234

## watchlog

- nnn

## What is GraphQL, 4:49, 2024-11-15

https://www.youtube.com/watch?v=wfBeEn8d-8A&list=PL3eAkoh7fypqxiqDq9OQGm3aJO7dZjSBI

- developed by Facebook in 2015
- you get what you want
- server declares resources it had
- get mulitple resources in a single request
- mutation/queries
- one endpoint
- users resolvers##useresolvere
- schemas
- mutations
  - run in series (first finishes before the second)

## Building a GraphQL Server in Golang, 14:08, nnn

https://www.youtube.com/watch?v=ocNw1GHovUI&list=PL3eAkoh7fypqxiqDq9OQGm3aJO7dZjSBI&index=2

- type safe
- he starts from the very beginning: a new directory: gographql-server
- repo: https://github.com/edwardtanguay/gographql-server
- `git init -b dev`
- `go mod init github.com/edwardtanguay/gographql-server`
- `go get github.com/99designs/gqlgen`
- generate project skeleton
  - `go run github.com/99designs/gqlgen init`
- will create GraphQL playground
- graph
  - schema.graphqls
    - it has schemas todo, user, etc.
  - generated.go
    - it has schemas todo, user, etc.
  - gqlgen.yml
    - mentions resolvers
  - schema.resolvers.go
    - our own implementation of resolvers
- now replace default schema
  - schema.graphqls

```
type Video {
  id: ID!
  title: String!
  url: String!
  author: User!
}

type User {
  id: ID!
  name: String!
}

type Query {
  videos: [Video!]!
}

input NewVideo {
  title: String!
  url: String!
  userId: String!
}

type Mutation {
  createVideo(input:NewVideo!): Video!
}
```

- 

## VOCAB - SPANISH

```
now let's move onto the terminal
ahora pasemos a la terminal

will return a result
va a devolver un resultado

just to give an example
solo para dar un ejemplo

```
