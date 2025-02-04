# Domina Python: FastAPI

https://www.linkedin.com/learning/domina-python-fastapi

- duration: 01:24:00
- language: es
- topics: python, fastapi
- rank: 4.9
- description: looks interesting, in Spanish, good way to learn practical Python
- year: 2024
- status: finished

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

## Cómo crear el API router en FastAPI, 5:42, 2024-12-17

https://www.linkedin.com/learning/domina-python-fastapi/como-crear-el-api-router-en-fastapi?autoSkip=true&resume=false

- API Router is a class from FastAPI
- allows you to separate routes into modules
- directory "routers"
- this has files:

```
__init__.py
support.py
todo.py
```

- the init file is a requirement for Python to treat the directory as a module that can be imported
- copies all todo routes to todo.py
- same for support routes
- then in each file, app become router
- then include_router()
- app.include_router(todo.router)
- goes to Swagger

## Tags para separar los endpoints y los routers en swagger, 4:59, 2024-12-18

https://www.linkedin.com/learning/domina-python-fastapi/tags-para-separar-los-endpoints-y-los-routers-en-swagger?autoSkip=true&resume=false

- Swagger uses tags
- created enums, works well
- everything worked except the openapi_tags

## SQLAlchemy para conexión a una base de datos SQL en FastAPI, 7:18, 2024-12-19

https://www.linkedin.com/learning/domina-python-fastapi/sqlalchemy-para-conexion-a-una-base-de-datos-sql-en-fastapi?autoSkip=true&resume=false

- she is going to use SQLite
- she has folders: models, routers, schemas
- models inherit Base
- lots of configuration that she did beforehand

## Cómo usar HTTPException en FastAPI, 3:00, 2024-12-20

https://www.linkedin.com/learning/domina-python-fastapi/como-usar-httpexception-en-fastapi?autoSkip=true&resume=false

- import status from FastAPI, then you automatically get the codes##statusok

## Cómo crear funciones para el manejo de errores, 5:09, 2024-12-21

https://www.linkedin.com/learning/domina-python-fastapi/como-crear-funciones-para-el-manejo-de-errores?autoSkip=true&resume=false

- creates custom exception handling
- also makes a custom exception handler with JSONResponse
- it's not clear how "app" works in her code
- part is defined in handler, part in exception

```
from app.custom_exceptions import CustomException
```

- from app.custom_exceptions import CustomException

```
app = FastAPI()

# Custom exception handler
@app.exception_handler(CustomException)
```

## Cómo crear un middleware en FastAPI, 5:43, 2024-12-21

https://www.linkedin.com/learning/domina-python-fastapi/como-crear-un-middleware-en-fastapi?autoSkip=true&resume=false

- middleware is executed before route are executed
- shows the time elapsed in header##sshowstehsitime
- it applies to all routes in all routers automatically
- then creates middleware.py and in main.py: app.add_middleware

## CORS con FastAPI, 4:35, 2024-12-21

https://www.linkedin.com/learning/domina-python-fastapi/cors-con-fastapi?autoSkip=true&resume=false

- three parts of the url that make it the same: protocal, domain, port

## VOCAB - SPANISH

```
we are indicating that
estamos indicando que
2024-12-21 21:45:15

with the time it has taken to be executed
con el tiempo que se ha demorado en ser ejecutado; conj=demorar; demorar = to delay, regular
2024-12-21 11:30:55

let's calculate how long the request takes
vamos a calcular cuánto se demora la petición; conj=demorar; demorar = to delay, regular
2024-12-21 11:24:06

he must always do three things
el siempre debe hace tres cosas
2024-12-21 11:20:25

this will be in charge of making the request
esta se va encargar de hacer la petición
2024-12-21 11:06:36

let's implement it
vamos a implementarlo
2024-12-21 11:02:27

since if the id is not found
ya que si no se encuentra el id; pr=JAH-kay-see
2024-12-21 02:08:11

where we are going to raise an exception
donde vamos a levantar una excepción
2024-12-21 02:06:30

the status codes that an API returns
los códigos de estado que una API retorna
2024-12-21 01:58:38

we can check it and it should bring us a list
lo podemos comprobar y este nos debería traer una lista
2024-12-19 11:21:16

when we ran the server, a file called todo.db was created
cuando corrimos el servidor, se ha creado un archivo llamada todo.db; conj=correr; corrimos = preterite; preterite for both -er and -ir verbs is -imos
2024-12-19 11:18:26

that we are labeling
que estamos etiquetando
2024-12-18 10:47:08

we are going to use the "id" key
vamos a usar la llave "id";pr=JAW-veh
2024-12-18 10:45:12

and no longer belongs to the default section
y que ya no pertenece a la sección default;pr=ee-kay-JAW-noh
2024-12-18 10:37:56

along with attribute we just defined
junto con atributo que acabamos de definir
2024-12-18 10:24:32

with the help of tags
con la ayuda de los tags
2024-12-18 10:14:06

the quotation marks and the commas
las comillas y las comas
2024-12-17 19:08:56

we leave the empty quotes
dejamos las comillas vacias
2024-12-17 19:08:04

that belong to them
que les pertenezcan; conj=pertenecer; pertenezcan = present subjunctive (instead of present indicative: pertenecen)
2024-12-17 19:03:15

let's remove the extra lines
vamos a quitar los líneas que sobran; rank=4.8
2024-12-17 18:55:25

and we change app for router
y cambiamos app por router
2024-12-17 18:52:55

indicates an error
señala un error
2024-12-17 18:49:17

This, in turn, allows the project to be more organized
esto, a su vez, permite que el proyecto sea más organizado
2024-12-17 18:42:41

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
