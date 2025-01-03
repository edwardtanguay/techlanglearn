# Flask Essential Training

https://www.linkedin.com/learning/flask-essential-training-24681038

- duration: 02:02:00
- language: en
- topics: flask
- rank: 4.91
- description: seems like a solid course on Flask, Natasha has a nice clear accent, builds example app: uses Flask-WTF for forms, and SQLite and Postgres, builds app, chart.js
- year: 2024
- status: watched all intro videos, starting now with coding videos

## Build a full stack app in Python with Flask, 0:38, 2024-12-31

https://www.linkedin.com/learning/flask-essential-training-24681038/build-a-full-stack-app-in-python-with-flask?resume=false

- basic intro

## Prerequisites, 1:22, 2024-12-31

https://www.linkedin.com/learning/flask-essential-training-24681038/prerequisites?resume=false

- has everything installed

## What is Flask?, 2:30, 2024-12-31

https://www.linkedin.com/learning/flask-essential-training-24681038/what-is-flask?autoSkip=true&resume=false

- mentioned April fool's joke when it was created

## Comparing Flask and Django, 3:14, 2025-01-01

https://www.linkedin.com/learning/flask-essential-training-24681038/comparing-flask-and-django?autoSkip=true&resume=false

- Flask is very flexible, for smaller projects
- Django is large, structured, has less flexibility

## Project step 1: Intro to the Health Tracker Dashboard project, 2:08, 2025-01-02

https://www.linkedin.com/learning/flask-essential-training-24681038/project-step-1-introduction-to-the-health-tracker-dashboard-project?autoSkip=true&resume=false

- showed the app she will build, has quite a bit of interesting tech in it, e.g. sqlite, chart.js

## Setting up your Flask environment and first application, 5:52, 2025-01-03

https://www.linkedin.com/learning/flask-essential-training-24681038/setting-up-your-flask-environment-and-first-application?autoSkip=true&resume=false

- sets up virtual environment
- `mkdir flask-showcase-002`
- `python -m venv venv` (venv is more standard than env)
- `source venv/Scripts/activate`
- `pip install Flask`
- open VSCode
- create `app.py`

```
from flask import Flask

app = Flask(__name__)

@app.route('/')
def info():
	return 'Info Site'

if __name__ == '__main__':
	app.run(debug=True)
```

- `http://localhost:5000`

## Accessing the code repository on GitHub, 1:32, nnn

https://www.linkedin.com/learning/flask-essential-training-24681038/accessing-code-repository-on-github?autoSkip=true&resume=false

- nnn

## VOCAB - ITALIAN

```
so we are all ready
quindi siamo tutti pronti;pr=[pronti like "own"]
2025-01-03 07:47:34

so this will activate our virtual environment
quindi questo attiverà il nostro ambiente virtuale; conj=attivare; attiverà = futuro simplice
2025-01-03 07:41:32

let's fix this
correggiamo questo
2025-01-03 07:26:06

and to activate it
e per attivarlo
2025-01-03 07:22:45

each project
ciascun progetto;pr=chass-KOON
2025-01-02 16:46:37

this means you will be able to submit your data
ciò significa che sarai in grado di inviare i tuoi dati
2025-01-02 07:56:41

this form collects information
questo modulo raccoglie informazioni;pr=rah-KOH-lee-eh
2025-01-02 07:48:46

and I'll leave you enough space
e ti lascerò abbastanza spazio
2025-01-02 07:46:50

physical exercise
l'esercizio fisico
2025-01-02 07:43:43

during the course, we will build a project
durante il corso, costruiremo un progetto
2025-01-02 07:41:44

for science projects
per i progetti di scienza;pr=SHENT-sa
2025-01-01 09:36:19

where each framework excels
dove eccelle ogni framework
2025-01-01 09:25:56

it's light, keeps the core simple
è leggero, mantiene el nucleo semplice;pr=leh-JER-oh; pr=SEM-plih-cheh
2025-01-01 08:17:17

right for your project
giusto per il tuo progetto; pr=[JOOS-toh]
2025-01-01 08:14:54

sono entrambi framework popolari
they are both popular frameworks
2025-01-01 08:09:56

so why choose Flask
allora perché scegliere Flask;pr=sheh-lee-AIR-eh
2024-12-31 17:43:59

check out the documentation
dai un'occhiata alla documentazione;pr=oon-oh-key-AH-ta
2024-12-31 17:39:59

on a large scale
su larga scala
2024-12-31 17:38:49

with just a few lines
con solo poche righe;pr=POH-keh
2024-12-31 17:36:52

as your project grows
mentre il tuo progetto cresce;pr=CRAY-sha
2024-12-31 17:26:26

Flask can handle it all
Flask è in grado di gestire tutto; pr=jess-TEER-eh
2024-12-31 16:49:36

you can add
puoi aggiungere; pr=[poo-OH-ee ah-JOON-jer-eh]
2024-12-31 16:43:20

five lines of code
cinque righe di codice
2024-12-31 16:42:00

here are some key features of Flask
ecco alcune caratteristiche chiave di Flask
2024-12-31 16:37:55

this means you can decide
ciò significa che puoi decidere;pr=[chow, kay, POH-ee, day-CHEE-deh-reh]
2024-12-31 16:30:23

it's light, easy
è leggero, facile; pr=[leh-JEH-oh FAH-chee-lay	]
2024-12-31 16:16:06

if you can write
se sei in grado di scrivere;pr=[seh SAY]; pr=SKREE-veh-reh
2024-12-31 15:31:35

here's what I recommend
ecco cosa ti consiglio;pr=con-SEE-leeoh
2024-12-31 15:21:19

it's pretty simple
è piuttosto semplice;pr=[eh pee-oo-TOH-sto SEM-plih-cheh]
2024-12-31 15:13:42

in this course
in questo corso;pr=KWEH-stoh
2024-12-31 14:53:30

```
