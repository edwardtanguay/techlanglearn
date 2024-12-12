# Domina Python: FastAPI

https://www.linkedin.com/learning/domina-python-fastapi

- duration: 01:24:00
- language: sp
- topics: python, fastapi
- rank: 4.7
- description: looks interesting, in Spanish, good way to learn practical Python

## FastAPI y Python, 0:36, 2024-11-29

https://www.linkedin.com/learning/domina-python-fastapi/fastapi-y-python

- simple introduction

## Para qué sirve FastAPI, 1:10, 2024-11-29

https://www.linkedin.com/learning/domina-python-fastapi/para-que-sirve-fastapi?autoSkip=true&resume=false

- was created in 2018 by Sebastián Ramírez
- used by Uber and Netflix

## Cómo instalar FastAPI, 2:05, 2024-12-10

https://www.linkedin.com/learning/domina-python-fastapi/como-instalar-fastapi?autoSkip=true&resume=false

- she says FastAPI uses the latest version of Python, 3.8 (2024-02)
  - I have Python 3.12.2
- the version of FastAPI seems to be 0.115.5 at the moment
- she recommends using a virtual environment
- install with this command: `pip install "fastapi[all]"
  - this uses Uvicorn, the server that FastAPI uses
- you can also use `pip install fastapi "unicorn[standard]"
- create **requirements.txt** manually, not with pip freeze

```
fastapi
uvicorn[standard]
```

- then `pip install -r requirements.txt`

## Configuración inicial de un API Rest con FastAPI, 3:08, 2024-12-10

https://www.linkedin.com/learning/domina-python-fastapi/configuracion-inicial-de-un-api-rest-con-fastapi?autoSkip=true&resume=false

- **main.py**

```
from fastapi import FastAPI

app = FastAPI()

@app.get("/")
async def hello_world():
	return {"message": "hello world"}
```

- `uvicorn main:app --reload` (reload makes it restart when code changes)
- runs it at port 8000
- `uvicorn main:app --reload --port 3344` for a port

## Documentación en FastAPI - Swagger y ReDoc, 4:04, 2024-12-11

https://www.linkedin.com/learning/domina-python-fastapi/documentacion-en-fastapi-swagger-y-redoc?autoSkip=true&resume=false

- https://www.openapis.org/
- http://localhost:3344/docs shows all the routes with Swagger##thefityroute
- http://localhost:3344/redoc shows all routes with ReDoc##wthredoc
- adds `app.title = "Hello World API"`

## Estructura de archivos de FastAPI, 2:09, 2024-12-12

https://www.linkedin.com/learning/domina-python-fastapi/estructura-de-archivos-de-fastapi?autoSkip=true&resume=false

- the structure recommended:##thestruct
- create **init**.py in every folder
  - this means it is a package
  - package-level initialization code
  - include metadata such as version, author
  - expose to public API
  - can control what happens when a package is imported
  - another structure example##anotheridffast

## Cómo crear un endpoint con FastAPI, 3:59, 2024-12-12

https://www.linkedin.com/learning/domina-python-fastapi/como-crear-un-endpoint-con-fastapi?autoSkip=true&resume=false

- endpoints can also return dictionaries or lists
- made an endpoint with a list
- `source env/Scripts/activate`
- `uvicorn main:app --port 3344 --reload`

```
@app.get("/flashcards")
async def get_flashcards():
	return [
		{
			"id": 1,
			"front": "house",
			"back": "das Haus"
		},
		{
			"id": 2,
			"front": "tree",
			"back": "der Baum"
		}
	]
```

## Parámetros del Path, 4:53, nnn

https://www.linkedin.com/learning/domina-python-fastapi/parametros-del-path?autoSkip=true&resume=false

- nnn

## VOCAB - SPANISH

```
later in this course
más adelante en este curso

neither of them meets the objective
ninguno de los dos cumple con el objetivo;cumplir = to fulfill, carry out, reach (an age); conj=cumplir

since this endpoint
ya que este endpoint;pr=JAW-kay-ES-tah

the endpoint allows the sending and receiving of data
el endpoint permite el envío y el recepción de datos

it is recommended to have a defined structure
se recomienda tener una estructura definida; conj=recomendar; recomienda = simple present

since this documentation is interactive
ya que esta documentación es interactiva; pr=jaw-KAY-stah

if there are any
en caso de haberlos

the information will be displayed
se desplegará la information

by clicking on this box
al darle clic a este recuadro

next we see the endpoint route
seguido vemos la ruta del endpoint

if we click on it, it will open the file
si le damos clic, abrirá el archivo

which points to the file
que apunta al archivo

just below the title we find a url
justo debajo del titulo encontramos una url

first we come across the API name
primero nos encontramos con el nombre de la API

pr. --reload
guión, guión, reload; pr=gee-UN

we run it by executing the command
lo corremos ejecutando el comando

the solution we have created
la solución que tenemos creada

as it grows
a medida que crece

high performance
alto rendimiento

don't miss the opportunity to take your skills to the next level
no te pierdas la oportunidad de llevar tus habilidades al siguiente nivel

```
