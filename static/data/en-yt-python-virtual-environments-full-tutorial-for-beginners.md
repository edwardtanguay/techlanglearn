# Python Virtual Environments - Full Tutorial for Beginners

https://www.youtube.com/watch?v=Y21OR1OPC9A

- duration: 00:09:04
- language: en
- topics: python
- rank: 4.2
- description: shows how to install a Python virtual environment

## watchlog

09:04 - 2024-12-10

## notes

- there are tools that can do this
  - Conda: https://docs.conda.io/projects/conda/en/latest/user-guide/getting-started.html
  - virtenv: https://virtualenv.pypa.io/en/latest/
  - pipenv: https://pipenv.pypa.io/en/latest/
  - poetry: https://python-poetry.org/
- but venv is built into Python and is the most common way to create virtual environments
- create an environment on Windows
  - `python -m venv env` ("env" is convention)
  - `env\Scripts\activate.bat`
    - it then puts an (env) in front of the directory
  - in Git Bash: `source env/Scripts/activate`
- to deactivate it: `deactivate`
- installing packages
  - `pip list`
  - `pip install requests`
- packaging your project to be run elsewhere
  - `pip freeze > requirements.txt`
- setting up virtual environment with these requirements
  - create directory
  - copy requirements.txt to it (and other Python files)
  - `python -m venv env`
  - `pip install -r requirements.txt`
  - create **main.py** that uses requests

```
import requests

response = requests.get("https://edwardtanguay.vercel.app/share/skills.json")

if response.status_code == 200:
	data = response.json()
	print(data)
else:
	print("API request failed.")
```

## VOCAB - ITALIAN

```
for your Python projects
per i tuoi progetti Python

```
