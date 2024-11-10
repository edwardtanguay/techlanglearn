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

## Uso de Middleware en servicios web Go, 6:35, 2024-10-30

- /web/middleware/counter.go
- un Mutex en Go es un primitivo de sincronización que se utiliza para controlar el acceso a recursos compartidos en la programación concurrente
- /web/middleware/logging.go
- /web/error/error.go
- /web/middleware/error.go
- /web/middleware/auth.go
- continue 3:45

## VOCAB - SPANISH

```
we will leave these set as they are
dejaremos estos configurados como están

in order to save it
para guardarlo

I will save that later
lo guardaré más tarde

it's right there
esta justo ahí

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

as soon as you enter
nada más entrar

that is in charge of counting the requests
que se encargue de contar las peticiones

that is in charge of counting the requests
que se encargue de contar las peticiones

although we could have used any other
aunque podríamos haber utilizado cualquier otro

we access the URL
accedemos a la URL

start the server on port 8080
arranca el servidor en el puerto 8080

in isolation
de manera aislada

as light as possible
lo más ligeros posibles

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
