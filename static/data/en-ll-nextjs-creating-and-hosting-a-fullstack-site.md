# Next.js: Creating and Hosting a Full-Stack Site

https://www.linkedin.com/learning/next-js-creating-and-hosting-a-full-stack-site/create-a-full-stack-site-with-next-js

- duration: 03:54:00
- language: en
- topics: nextjs, mongo
- rank: 4.99
- description: Shaun Wassell, looks like an excellent, long course using Next.js version 14.2.8, hosts on Vercel
- year: 2024
- status: watched intro videos

## Create a full-stack site with Next.js, 0:43, 2024-12-17

- simple intro

## What you should know, 1:34, 2024-12-17

https://www.linkedin.com/learning/next-js-creating-and-hosting-a-full-stack-site/what-you-should-know?autoSkip=true&resume=false

- know JavaScript, React, networking

## Basic setup and exercise files, 2:35, 2024-12-17

https://www.linkedin.com/learning/next-js-creating-and-hosting-a-full-stack-site/basic-setup-exercise-files?autoSkip=true&resume=false

- explains codespaces

## Project introduction, 00:47, 2024-12-17

https://www.linkedin.com/learning/next-js-creating-and-hosting-a-full-stack-site/project-introduction?autoSkip=true&resume=false

- host on Vercel

## What is Next.js?, 2:48, 2025-03-07

https://www.linkedin.com/learning/next-js-creating-and-hosting-a-full-stack-site/what-is-next-js?autoSkip=true&resume=false

- better for SEO
- uses TypeScript

## Setting up a Next.js project, 5:03, 2025-03-07

- using 14.2.8

## Creating pages, 5:41, 2025-03-07

- created 4 pages

## Creating a product list, 9:44, 2025-03-07

- shows how to use the Image command

## Using Next.js Links, 3:54, 2025-03-08

- he puts divs in Link, i.e. in <a> tags

## Customizing content with route parameters, 8:13, 2025-03-08

- he puts product-detail under products

## Creating a shopping cart list , 8:42, 2025-03-08

- very simple

## Creating a 404 page, 5:32, 2025-03-08

- I used not-found.tsx
- he called it the same
- it works no matter how many levels deep you are
- even takes care of this for the detail page

## Styling Next.js applications with Tailwind CSS, 9:09, 2025-03-09

- standard styling

## Challenge: Creating a navigation bar, 3:34, 2025-03-11

- shows basics of starting navbar

## Solution: Creating a navigation bar, 5:34, 2025-03-11

- completes navbar

## What are Next.js route handlers?, 2:12, 2025-03-11

- compared backend routs to frontend

## How do Next.js route handlers work?, 7:29, 2025-03-11

- calls api hello

## Testing route handlers with Postman, 7:33, 2025-03-11

- showed GET and POST in one file

## Creating a list endpoint for products, 5:42, 2025-03-11

- returned an array

## Using route parameters in route handlers , 8:10, 2025-03-12

- makes a route

## Creating a shopping cart endpoint, 9:05, 2025-03-12

- makes route: /api/users/[id]/cart

## Creating an add-to-cart endpoint, 9:07, 2025-03-12

- uses .concat

## Challenge: Creating a remove-from-cart endpoint, 1:23, 2025-03-12

- explains delete

## Solution: Creating a remove-from-cart endpoint, 5:53, 2025-03-12

- made delete

## What is MongoDB?, 1:42, 2025-03-12

- will be using MongoDB Atlas

## Setting up hosting for MongoDB, 4:54, 2025-03-12

- set up new cluster at Mongo DB Atlas

## Adding MongoDB to Next.js , 6:46, 2025-03-12

- set up db.ts

## Adding data to MongoDB , 7:32, 2025-03-12

- used the playground

## Rewriting the list endpoints , 7:29, 2025-03-12

- getting data from MongoDB

## Rewriting the load product endpoint, 4:02, 2025-03-13

- loads a product, two awaits

## Rewriting the shopping cart endpoint, 4:27, 2025-03-13

- gets all products for cart

## Rewriting the add-to-cart endpoint, 5:08, 2025-03-14

- saved cart

## Challenge: Rewriting the remove-from-cart endpoint, 0:57, 2025-03-14

- explained

## Solution: Rewriting the remove-from-cart endpoint, 4:02, 2025-03-14

- shows how and tests

## Data loading basics in Next.js , 1:13, 2025-03-14

- fetch does server and client side data fetching

## Loading data with Fetch, 5:57, 2025-03-14

- shows that is it server-side

## Loading all products , 3:39, 2025-03-14

- can use async in Next.js for components

## Loading individual products , 3:09, 2025-03-14

- adds loading

## Loading shopping cart items , 6:16, 2025-03-14

- child component that manages state

## Adding items to the cart , 10:05, 2025-03-14

- added no-cache

## Shopping cart improvements , 6:02

- same cache problem

## Challenge: Removing items from the cart , 1:28, 2025-03-14

- showed challenge

## Solution: Removing items from the cart , 5:34, 2025-03-14

- solved it

## Preparing an app for release , 6:56, 2025-03-14

- makes url dynamic
- in nextConfig, ignoreDuringBuilds: true, ignoreBuildErrors: true
- for each page: export const dynamic = 'force-dynamic'

## Releasing a Next.js app, 4:45, nnn

- adds env variables at Vercel

## Shutting down a Next.js app , 0:45, nnn

- how to shut down: deletes it

## VOCAB - ITALIAN

```
hopefully
si spera che
2025-03-14 14:12:11

and this reminds me that
e questo mi ricorda che
2025-03-14 10:37:55

after that
dopodichè
2025-03-14 09:30:00

just a warning here
solo un avvertimento qui
2025-03-14 09:09:45

besides that
oltre a ciò
2025-03-14 09:02:22

what's happening is that
quello che sta succedendo è che
2025-03-12 22:17:35

we will stop the execution of the project
fermeremo l'esecuzione del progetto
2025-03-12 21:45:48

would you like to do it
ti piacerebbe farlo
2025-03-12 19:25:25

the rest of the things
il resto delle cose
2025-03-12 16:01:56

among other things
tra le altre cose
2025-03-12 13:37:41

it is possible that an error happens during the compilation process
è possibile che accada un errore durante il processo di compilazione
2025-03-12 13:33:21

one string, two strings
la stringa, due stringhe; pr = both like "gear"
2025-03-12 13:03:35

let me copy, let me see, let me know
fammi copiare, fammi vedere, fammi sapere
2025-03-12 12:57:15

this is how it will be
ecco come sarà
2025-03-12 12:37:58

there was this little thing
c'era questa piccola cosa
2025-03-12 10:11:42

we were able to do
siamo stati in grado di fare
2025-03-12 10:05:41

you might remember that
potresti ricordare che
2025-03-12 10:04:18

I remembered to register and record the data.
Mi sono ricordato di registrarmi e registrare i dati.
2025-03-11 16:28:49

the test file, the text file
il file di prova, il file di testo
2025-03-11 15:32:51

drop down menu
menu a discesa
2025-03-11 15:23:35

here on the left side
qui sul lato sinistro
2025-03-11 15:16:47

we'll talk about that very shortly
ne parliamo molto a breve
2025-03-11 14:22:23

so let's do that
quindi facciamolo
2025-03-11 13:53:24

and we will notice that there are already many things
e noteremo che ci sono già molte cose
2025-03-10 12:22:01

in the right place
nel posto giusto
2025-03-10 12:17:30

just as a reminder
solo come promemoria
2025-03-10 12:09:48

everything that
tutto ciò che
2025-03-08 12:41:07

and that is that we have
e cioè che abbiamo
2025-03-08 11:30:26

what we want
quello che vogliamo
2025-03-08 11:18:38

that little spinner
quel piccolo spinner
2025-03-08 11:08:34

I will explain
spiegherò
2025-03-08 10:30:31

to make this happen
per far sì che ciò accada
2025-03-08 10:29:45

so what we're going to do next
quindi quello che faremo dopo
2025-03-07 22:30:38

alright, cool
bene, figo
2025-03-07 20:21:14

so here it will appear
quindi ecco che apparirà
2025-03-07 20:14:41

that is, the path of the folder
e cioè, il percorso della cartella; pr=choh-EH
2025-03-07 20:08:16

if you look here on the left side
se guardi qui sul lato sinistro
2025-03-07 20:06:25

in any case
in ogni caso
2025-03-07 18:29:29

and there you have it
ed ecco fatto
2025-03-07 18:27:25

each of these options
ciascuna di queste opzione
2025-03-07 17:57:33

what we will do for now
quello che faremo per ora
2025-03-07 17:56:08

and then we will press enter
e poi premeremo invio
2025-03-07 17:54:42

it doesn't matter what the name is
non importa quale sia il nome
2025-03-07 17:53:32

e poiché creeremo un frontend
and since we will create a frontend
2025-03-07 17:51:02

ma tutto quello che devi fare qui
but all you have to do here
2025-03-07 17:47:33

a special way
un modo speciale
2025-03-07 17:41:21

server side
lato server
2025-03-07 17:40:00

since we will be using Next.js
dato che utilizzeremo Next.js
2025-03-07 17:33:33

so, in any case
quindi, in ogni caso;pr=OHN-yee-KAH-zoh
2024-12-17 18:32:06

so for example
quindi, ad esempio;pr=ay-ZEM-pee-oh (not SEM)
2024-12-17 18:29:07

but it's also very easy
ma è anche molto facile;pr=FAH-chee-lay
2024-12-17 18:27:32

Therefore
quindi;pr=QUEEN-dee
2024-12-17 17:47:04

of fundamental networking concepts
dei concetti fondamentali di rete; pr=day-kon-CHET-tee; pr=RAY-teh
2024-12-17 17:43:36

and finally
e infine; pr=in-FEE-nah
2024-12-17 17:42:03

that you knew
che tu conoscessi; pr=kon-oh-SHEH-see
2024-12-17 17:37:47

and at the end of the course
e alla fine del corso; pr=KOR'-soh (not zoh)
2024-12-17 17:28:03

in recent years
negli ultimi anni;pr=nehleeoo-tee-mee
2024-12-17 17:23:50

for React developers
per gli sviluppatori di React
2024-12-17 17:25:27

and for this reason
e per questo motivo
2024-12-17 17:26:10




```
