# Domina Go: Creación y empaquetado de aplicaciones

- https://www.linkedin.com/learning/domina-go-creacion-y-empaquetado-de-aplicaciones
- 00:34:00
- spanish-spanish
- Go
- 4.6

## intro - 0:30, 2024-10-29

- eso fue solo una introducción

## Necesito una CLI en Go - 1:21, 2024-10-30

- habla sobre construir una CLI para consumir una API

## Implementación de comandos en CLI en Go, 2:41, 2024-10-30

- Cobra es un CLI framework para Go
- https://transform.tools/json-to-go
- tiene cmd/root.go
- usó cobra, ejecuta el comando y genera JSON

## Creación de subcomandos en CLI en Go, 3:05, 2024-10-30

- parece estar creando la aplicación del último vídeo

## Integración de prompts en CLI en Go, 4:08, 2024-10-30

- él va a usar Cobra y Survey
- go.mod con requires
- pokemon/pokemon.go
- cmd/root.go
- getCmd es un subcommando
- `go mod tidy`
- `go run main.go get`

## Configuración de un rúter y creación de Handlers en servicios web Go, 4:40, 2024-10-30

- él va a usar Gin
- él siempre crea primero un archivo go.mod
- /pokemon/pokemon.go
  - structs
- /pokemon/api.go
- el sitio usa json.Unmarshal
- /web/handlers.go
- r.Param, r.JSON
- /main.go
- las rutas
  - /api
  - /pokemon/:name
  - /pokemon/types
- `go mod tidy`
- `go run main.go`

## Uso de Middleware en servicios web Go, 6:35, 2024-11-10

https://www.linkedin.com/learning/domina-go-creacion-y-empaquetado-de-aplicaciones/uso-de-middleware-en-servicios-web-go?resume=false

- /web/middleware/counter.go
- un Mutex en Go es un primitivo de sincronización que se utiliza para controlar el acceso a recursos compartidos en la programación concurrente
- /web/middleware/logging.go
- /web/error/error.go
- /web/middleware/error.go
- /web/middleware/auth.go
- shows how middleware works in an API

## ¿Por qué es importante generar binarios en Go?, 1:39, 2024-11-10

https://www.linkedin.com/learning/domina-go-creacion-y-empaquetado-de-aplicaciones/por-que-es-importante-generar-binarios-en-go?autoSkip=true&resume=false

- `go build -o pkm`
- `./pkm get`
- he created it on Mac, so it will not work on Windows or Linux
  - and it won't even work on a Mac with ARM processors

## Generación de binarios en Go, 2:02, 2024-11-10

https://www.linkedin.com/learning/domina-go-creacion-y-empaquetado-de-aplicaciones/generacion-de-binarios-en-go?autoSkip=true&resume=false

- you can build for numerous operating systems
- variables GOOS and GOARCH
- creates a Makefile
- GOOS=darwin
- build-all##creatsoall

## Creación de imágenes Docker en Go, 3:06, 2024-11-11

https://www.linkedin.com/learning/domina-go-creacion-y-empaquetado-de-aplicaciones/creacion-de-imagenes-docker-en-go?autoSkip=true&resume=false

- have to create a Dockerfile##thedodockfil
  - FROM scratch
  - ENTRYPOINT ["./pkm"]
- now Makefile
  - docker-run
- executes
  - `make-all`
  - creates dist
  - `docker images`
  - `make docker-run`
  - it works

## Configuración de Goreleaser en proyectos Go: ejemplo práctico, 4:44, 2024-11-11

https://www.linkedin.com/learning/domina-go-creacion-y-empaquetado-de-aplicaciones/configuracion-de-goreleaser-en-proyectos-go-ejemplo-practico?autoSkip=true&resume=false

- GoReleaser
  - a release automation tool for Go projects
  - simplifies and streamlines the process of building, packaging, and distributing Go applications
  - widely used in continuous integration (CI) and continuous deployment (CD) pipelines to create consistent, ready-to-release binaries, Docker images, and release assets
  - install
    - `brew install goreleaser/tap/goreleaser
    - add **.goreleaser.yml**
    - it is a long file, 200+ lines
    - announce: (in social media)
  - now we are in **root.go**
  - mpw om **main.go**
    - add variables for GoReleaser
  - SBOM
    - Software Bill of Materials

## VOCAB - SPANISH

```
the file will be overwritten
el archivo se sobrescribirá

to a large extent
en gran medida

which does everything that has been done in the previous videos
que hace todo lo que hace hecho en los vídeos anteriores

this process will take some time
este proceso tardará un tiempo

so that it is capable
para que sea capaz

as we see, the binaries will be generated
como vemos, se va a generar los binarios

having built it on our machine
al haberlo construido en nuestra máquina

Go stands out for the small size of its binaries
Go destaca por el pequeño tamaño de sus binarios

since we would be forcing him to install Go
puesto que le estaríamos obligando a instalar Go [gr. present conditional]

but we can't deliver it to the user like this
pero no podemos entregarla así al usuario

let's start with a small application
partamos de una pequeña aplicación

in which they are going to be executed
en el que se vayan a ejecutar [gr. subjunctive present]

we usually run our programs directly from the source code
solemos ejecutar nuestros programas directamente desde el código fuente

common tasks such as counting requests
tareas comunes como el conteo de peticiones

we should see the message
deberíamos ver el mensaje [gr. present conditional]

this time, we should get the correct answer
esta vez, deberíamos recibir la respuesta correcta [gr. present conditional]

as expected, we have an error
como era de esperar, tenemos un error

whenever we returned an error previously
siempre que devolviéramos un error anteriormente [gr. subjunctive imperfect 1]

we will leave these set as they are
dejaremos estos configurados como están

in order to save it
para guardarlo

I will save that later
lo guardaré más tarde

it's right there
está justo ahí

as dark as it can be
tan oscuro como puede ser

click on the button
haga clic en el botón

click on the left side
haga clic en el lado izquierdo

on the one hand
por un lado

since we consider it an error
puesto que lo consideramos un error

we will return an error
retornaremos un error

by our web layer
por nuestra capa web

including the value
incluyendo el valor

the time that has elapsed
el tiempo que ha transcurrido

the first thing it does is get the current time
lo primero que hace es obtener la hora actual

that is in charge of counting the requests
que se encargue de contar las peticiones

that is in charge of counting the requests
que se encargue de contar las peticiones [en-KAR-gay]

although we could have used any other
aunque podríamos haber utilizado cualquier otro

we access the URL
accedemos a la URL

start the server on port 8080
arranca el servidor en el puerto 8080

in isolation
de manera aislada

will make a request
hará una petition

it will return the list
devolverá la lista

to test it
para probarlo

finally
por último

let's ask
vamos a pedir

I will store in the variable
almacenaré en la variable

he copies text in all at once, he doesn't type it
copia el texto de una sola vez, no lo escribe

by prompt we understand
entendemos por prompt

you can see on the last line
puedes verlo en la última línea

we will read the answer
leeremos la respuesta

the root directory
el directorio raíz

we will return to the subject
volveremos al tema

which is responsible for making the request to the API
que se engarga de hacer la petición a la API

in the following videos
en siguientes vídeos [see-g|ee-EN-tays]

wouldn't it be better to be able to do something like this
no sería mejor ser capaces de hacer algo así

tens, hundreds
decenas, cientos

text box
caja de texto

the mouse
el ratón

via our browser
mediante nuestro navegador
```
