# Domina Python: FastAPI

https://www.linkedin.com/learning/domina-python-fastapi

- duration: 01:24:00
- language: sp
- topics: python, fastapi
- rank: 4.9
- description: looks interesting, in Spanish, good way to learn practical Python
- year: 2024
- status: 7 more videos, good course

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

## Parámetros del Path, 4:53, 2024-12-12

https://www.linkedin.com/learning/domina-python-fastapi/parametros-del-path?autoSkip=true&resume=false

- shows how in FastAPI and then in Swagger with CURL
- works

```
@app.get("/books")
async def get_books(completed: Union[bool, None] = None):
	if completed is not None:
		filtered_books = list(filter(lambda book: book["completed"] == completed, BOOKS))
		return filtered_books
	return BOOKS
```

## Parámetros Query, 4:16, 2024-12-12

https://www.linkedin.com/learning/domina-python-fastapi/parametros-query?autoSkip=true&resume=false

- worked well, have to import HTTPException

```
@app.get("/book/{id}")
async def get_book(id:int):
	try:
		book = next(book for book in BOOKS if book["id"] == id)
		return book
	except:
		raise HTTPException(status_code=404, detail="book not found")
```

## Parámetros del Body, 3:21, 2024-12-14

https://www.linkedin.com/learning/domina-python-fastapi/parametros-del-body?autoSkip=true&resume=false

- only post, patch, put may have a body
- imports body like this:

```
from fastapi import FastAPI, HTTPException, Body
```

## Uso de modelos en FastAPI, 4:16, 2024-12-14

https://www.linkedin.com/learning/domina-python-fastapi/uso-de-modelos-en-fastapi?autoSkip=true&resume=false

- uses Pydantic, which enables you to create Zod-like schemas
- e.g. you can define that a title must be between 5 and 40 characters, if not, it sends back 422 Unprocessable Entity
- and sends back a detailed JSON string that the front end can display or log##titeltitlebody

## Qué es un path operation en FastAPI, 3:24, 2024-12-15

https://www.linkedin.com/learning/domina-python-fastapi/que-es-un-path-operation-en-fastapi?autoSkip=true&resume=false

- path operation decorator
- response_model=Todo, name="Create a todo", summary="nnn", description="nnn"
- I got an error when I use reponse_model=Book

## Cuando usar async para definir una path operation, 2:43, 2024-12-15

https://www.linkedin.com/learning/domina-python-fastapi/cuando-usar-async-para-definir-una-path-operation?autoSkip=true&resume=false

- she accesses three urls at the same time
- import asyncio
- after implementing sleep, she runs the bash script three times again
- if you use SQLAlchemy, you will have to use synchronous calls

## Uso de formularios en FastAPI, 3:27, 2024-12-15

https://www.linkedin.com/learning/domina-python-fastapi/uso-de-formularios-en-fastapi?autoSkip=true&resume=false

- when receiving data from a form
- `pip install python-multipart`
- then import Form
- Form inherits Body
- imports Annotated from typing
- she uses the form in Swagger##usetheformsj

## Manipulación de archivos en un endpoint de FastAPI, 4:21, 2024-12-16

https://www.linkedin.com/learning/domina-python-fastapi/manipulacion-de-archivos-en-un-endpoint-de-fastapi?autoSkip=true&resume=false

- you have to use File and UploadFile, both from Python multipart
- worked, afterwards I also saved the file, works well

## Cómo crear el API router en FastAPI, 5:42, nnn

https://www.linkedin.com/learning/domina-python-fastapi/como-crear-el-api-router-en-fastapi?autoSkip=true&resume=false

- nnn

## VOCAB - SPANISH

```
since we can get metadata
puesto que podemos obtener metadatos

instead of defining the file
en vez de definir el archivo

delivers a 404 code
entrega una código 404

the element associated with the file you want to upload
el elemento asociado al archivo que se quiere subir

whose name will be stored
cuyo nombre será almacenado

storage service
servicio de almacenamiento

since what we want to see is how it works
puesto que lo que queremos ver es su funcionamiento

the function is going to be called
la función se va a llamar

we must wait for the instruction to be fulfilled
debemos esperar que la instrucción se cumple

since the function only returns a dictionary
puesto que la función solo retorna un diccionario; puesto = past participle of poner (irregular); conj=poner

another may be being processed at the same time
otra puede estar siendo procesada al mismo tiempo

that has been created
que a sido creada

and therefore
y por ende

in which it is placed
en el que se pone

which is responsible for creating the elements
que se encarga de crear los elementos

to test the model
para probar el modelo

that is part of the body
que hace parte del body

and add them to the list of items
y los agregue a la lista de elementos

they usually receive data through the body
suelen recibir datos a través del body

it is not always required that the request receives a body
no siempre es requerido que la petición recibe un body

who delivers the request to us
que nos entrega la petición

they should not be delivered by an endpoint
no deben ser entregadas por un endpoint

five hundred
quinientos

these parameters come as part of the route
estos parámetros llegan como parte de la ruta

it will be as if we had sent a None
será como hubiéramos enviado un None; conj=haber; hubiéramos = subjunctive imperfect

if we leave it empty
si lo dejamos vacío

but we must write it
sino que debemos escribirlo

we no longer have options
ya no nos salen opciones;conj=salen; salen = to leave, go out

that have already been completed
que ya hayan sido completados; pr=JAW-ha-jahn-SEE-doh

an assigned value
un valor asignado

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
