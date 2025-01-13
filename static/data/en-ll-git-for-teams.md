# Git for Teams

https://www.linkedin.com/learning/git-for-teams/using-git-for-team-collaboration

- duration: 02:37:00
- language: en
- topics: git
- rank: 4.93
- description: Kevin Bowersox, lots of good tips for teams, Gitlab installation, secrets, Git flow, GitLab Runner, CI pipeline, cherry-picking, squashing, reverting
- year: 2023

## Using Git for team collaboration, 0:59, 2025-01-04

- he is a Java developer

## What you need to know, 1:30, 2025-01-06

https://www.linkedin.com/learning/git-for-teams/what-you-need-to-know?autoSkip=true&resume=false

- is not a basic Git course
- will use vim to save time
- uses Java projects built with Maven

## Fundamentals of Git collaboration overview, 1:43, 2025-01-07

https://www.linkedin.com/learning/git-for-teams/fundamentals-of-git-collaboration-overview?autoSkip=true&resume=false

- talks about using Git in a team

## Common pitfalls: Untracked pulls, 4:27, 2025-01-08

https://www.linkedin.com/learning/git-for-teams/common-pitfalls-untracked-pulls?autoSkip=true&resume=false

- you could use git reset --hard
- he resolved the merge, but in my example, I had to do a rebase

## Common pitfalls: Force push, 3:16, 2025-01-09

https://www.linkedin.com/learning/git-for-teams/common-pitfalls-force-push?autoSkip=true&resume=false

- one of the most dangerous commands: push --force
- `git ls-tree --full-tree -r HEAD`

## Best practices: Committing and syncing, 3:29, 2025-01-09

https://www.linkedin.com/learning/git-for-teams/best-practices-committing-and-syncing?autoSkip=true&resume=false

- commit early and often
- focus on a single feature per commit

## Best practices: gitignore, 4:20, 2025-01-09

https://www.linkedin.com/learning/git-for-teams/best-practices-gitignore?autoSkip=true&resume=false

- uses a Java project in Eclipse
- ignore directory with forward slash at end: bin/ and \*.class
- `**bin` in any subdirectory

## Standardize line endings with autocrlf, 5:20, 2025-01-10

https://www.linkedin.com/learning/git-for-teams/standardize-line-endings-with-autocrlf-22334239?autoSkip=true&resume=false

- `git config --list`
- look for: `core.autocrlf=true` (for Windows)
- look for: `core.autocrlf=false` (for macOS/Linux)
- he writes `git config core.autocrlf input` ensures that the repository always stores files with consistent LF line endings
- he was in a different branch and did `git checkout master`

## Branch naming, 1:46, 2025-01-13

https://www.linkedin.com/learning/git-for-teams/branch-naming?autoSkip=true&resume=false

- keep branch names short, consistent and informative
- bug/123/menu-issue
- feat/125/new-email-feature

## Write descriptive commit messages, nnn

https://www.linkedin.com/learning/git-for-teams/write-descriptive-commit-messages?autoSkip=true&resume=false

- @1:00 setting up editor for a different kind of commit

## VOCAB - ITALIAN

```
code changes
modifiche al codice;moh-DEE-fee-keh
2025-01-13 08:20:07

let's take a look at how you can do it
diamo un'occhiata a come puoi farlo; pr=oon-oh-key-AH-tah;conj=dare;diamo = let's give / we give;rule: drop final "e", so fare + lo becomes farlo
2025-01-13 08:20:01

here are the options you have
ecco le opzioni che hai
2025-01-10 20:19:28

we have already seen
abbiamo già visto
2025-01-10 20:17:32

this can cause problems
ciò puo causare problemi;pr=choh-POO-oh
2025-01-10 19:29:05

Git provides a setting
Git fornisce un'impostazione;pr=for-NEE-shay
2025-01-10 19:16:27

they use both a carriage return and a line feed
loro utilizzano sia un ritorno a capo che un line feed; rank=4.95
2025-01-10 19:14:01

but it's just nice to know
ma è solo bello saperlo
2025-01-09 17:00:09

you can do things like use two asterisks
puoi fare cose come usare due asterischi;pr=puh-OY;pr=ah-stair-REE-skee;conj=potere
2025-01-09 16:57:18

this means
ciò significa;pr=CHOH
2025-01-09 16:52:41

so anywhere inside our working directory
quindi ovunque all'interno della nostra directory lavoro;pr=oh-VUN-kway
2025-01-09 16:41:42

these lines
queste righe
2025-01-09 16:37:47

at the same time
allo stesso tempo
2025-01-09 14:59:36

that you might use
che potresti usare;conj=potere; potresti = conditional
2025-01-09 14:45:23

easy for your team
facile per il tuo team;pr=FACH-ee-lay
2025-01-09 14:41:52

as we will see in a minute
come vedremo tra un minuto
2025-01-09 08:25:15

those files need to be saved
quei file devono essere salvati
2025-01-09 08:20:17

they will commit the file
eseguiranno il commit del file;pr=eh-zeh-gwee-RAH-noh;conj=eseguire;eseguire = to execute
2025-01-09 08:15:22

so in this scenario
quindi in questo scenario;pr=sheh-NAR-ee-oh
2025-01-09 08:12:39

this may cause commits to be overwritten
ciò può causare la sovrascrittura dei commit;pr=chow-POO-oh
2025-01-09 08:10:40

one of the mistakes you need to watch out for
uno degli errori a cui devi fare attenzione
2025-01-09 08:06:36

we have made our commit
abbiamo fatto il nostro commit
2025-01-09 01:28:41

their changes
le loro modifiche;pr=moh-DEE-fee-keh
2025-01-09 00:59:16

they open the file with vim
aprono il file con vim;conj=aprire
2025-01-09 00:57:19

user two would also like to edit the file
l'utente due vorrebbe anche modificare il file;conj=volere; vorrebbe = conditional
2025-01-09 00:54:05

and it's pretty simple
e piuttosto semplice;pr=SEM-plee-chay
2025-01-08 21:16:55

in this scenario
in questo scenario; pr=sheh-NAH-ree-oh
2025-01-08 21:06:39

with their team
con il loro team
2025-01-08 21:05:31

keep this in mind
tienilo a mente;conj=tenere; tieni = imperative/present
2025-01-07 22:30:28

one more thing i should point out
un'altra cosa che dovrei sottolineare;conj=dovere; dovrei = conditional
2025-01-07 22:26:50

conj-potere-pres
posso, puoi, può, possiamo, potete, possono
2025-01-07 22:20:04

your team can adopt
il tuo team può adottare; conj=potere
2025-01-07 20:08:25

from there, we'll talk about conventions
da lì, parleremo di convenzioni; parleremo = futuro semplice;conj=parlare
2025-01-07 20:04:28

some practices
alcune pratiche;pr=PRAH-tee-kay
2025-01-07 20:03:07

at this point
a questo punto
2025-01-07 19:50:09

this chapter will focus on the basics
questo capitolo concentrerà sulla basi;conj=concentrare
2025-01-07 19:45:55

that this course provides you
che questo corso ti fornisca;pr=for-NEE-skah
2025-01-07 00:17:52

I will be there to guide you through every step
sarò lì per guidarti in ogni fase
2025-01-07 00:12:05

if you understand these concepts
se capisci questi concetti;pr=kah-PEE-shee
2025-01-07 00:08:41

life cycle
ciclo di vita;pr=CHEEK-loh
2025-01-07 00:01:03

I will guide you, step by step
ti guiderò passo dopo passo;pr=gwee-dair-OH
2025-01-05 03:13:52

we will guide you, step by step
guideremo, passo dopo passo;pr=gwee-dair-RAY-moh
2025-01-05 03:13:33

source code
codice sorgente
2025-01-05 03:09:35
```
