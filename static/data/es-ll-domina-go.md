# Domina Go

https://www.linkedin.com/learning/domina-go

- duration: 00:50:00
- language: es
- topics: go
- rank: 4.96
- description: short videos but seems to cover basic topics, and from 2023, the first section have extremely simple concepts but after that, these are very useful if you are new to Go
- status: finished
- year: 2023

## intro, 0:43, 2024-11-25

https://www.linkedin.com/learning/domina-go/aplicaciones-con-go?resume=false

- basic intro

## Concatenar cadenas en Go, 1:02, 2024-11-25

- just showed how to use +, extremely simple

## Crear cadenas multilínea en Go, 1:10, 2024-11-25

- use backticks for multlines

## Conversión de tipos de datos en Go, 2:26, 2024-11-25

- using int(), float64(), strconv.Itoa(), strconv.Atoi()
- int doesn't round the number

## Crear una función con parámetros opcionales en Go, 1:42, 2024-11-26

https://www.linkedin.com/learning/domina-go/crear-una-funcion-con-parametros-opcionales-en-go?autoSkip=true&resume=false

- variadic parameters
- `func sum(numbers ...int) int {`

## Obtener el último elemento de un slice en Go, 1:10, 2024-11-26

https://www.linkedin.com/learning/domina-go/obtener-el-ultimo-elemento-de-un-slice-en-go?autoSkip=true&resume=false

- just gets the last elements from a slice

## Verificar si una cadena contiene una subcadena en Go, 1:31, 2024-11-27

https://www.linkedin.com/learning/domina-go/verificar-si-una-cadena-contiene-una-subcadena-en-go?resume=false

- uses Contains

## Llamado de una función desde otro archivo en Go, 1:24, 2024-11-27

https://www.linkedin.com/learning/domina-go/llamado-de-una-funcion-desde-otro-archivo-en-go?autoSkip=true&resume=false

- she imports from another package
- doing an example with tools
- doing this in new project soflash
  - `go mod init github.com/edwardtanguay/soflash`
  - then made "cmd/test/main.go"
  - worked

## Verificar si un map contiene una clave en Go, 1:40, 2024-11-28

https://www.linkedin.com/learning/domina-go/verificar-si-un-map-contiene-una-clave-en-go?autoSkip=true&resume=false

- maps are key/value storage containers
- keys are unique
- can use exist on end:

```
userID := "u456"
if user, exists := users[userID]; exists {
	fmt.Printf("User %s: Name = %s, Email = %s\n", userID, user["name"], user["email"])
} else {
	fmt.Printf("User with ID %s not found.\n", userID)
}
```

## Concatenar dos slices en Go, 1:16, 2024-11-28

https://www.linkedin.com/learning/domina-go/concatenar-dos-slices-en-go?autoSkip=true&resume=false

- uses append() to combine two slices append(first, second...)

## Leer un archivo en Go, 1:30, 2024-11-29

https://www.linkedin.com/learning/domina-go/leer-un-archivo-en-go?autoSkip=true&resume=false

- she converts the data she gets back to a string
- and indeed, the data that comes back from os.ReadFile is an array of bytes ([]bytes)
  - the reason it works in my example was that I used fmt.Printf with %s, which converts it to a string
  - if I use %v, then it shows the bytes

## Enviar un JSON en una solicitud POST en Go, 2:13, 2024-11-29

https://www.linkedin.com/learning/domina-go/enviar-un-json-en-una-solicitud-post-en-go?autoSkip=true&resume=false

- checks if method is post with:

```
if r.Method == http.MethodPost
```

## Eliminar clave en el map en Go, 0:48, 2024-11-29

https://www.linkedin.com/learning/domina-go/eliminar-clave-en-el-map-en-go?autoSkip=true&resume=false

- simply uses delete()

## Eliminar un elemento de un slice en Go, 2:15, 2024-11-29

https://www.linkedin.com/learning/domina-go/eliminar-un-elemento-de-un-slice-en-go?autoSkip=true&resume=false

- uses append to connect up to point and after point to be deleted

## Iterar con range sobre un arreglo de números enteros en Go, 0:56, 2024-11-29

https://www.linkedin.com/learning/domina-go/iterar-con-range-sobre-un-arreglo-de-numeros-enteros-en-go?autoSkip=true&resume=false

- very simple explanation of range

## Generar un UUID con Go, 1:11, 2024-11-29

https://www.linkedin.com/learning/domina-go/generar-un-uuid-con-go?autoSkip=true&resume=false

- just shows how to use uuid.New()

## Convertir interfaz{} en una cadena en Go, 1:23, 2024-11-29

https://www.linkedin.com/learning/domina-go/convertir-interfaz-en-una-cadena-en-go?autoSkip=true&resume=false

- a simple explanation of Sprintf and %v

## Obtener los parámetros de consulta en Go, 1:50, 2024-12-01

https://www.linkedin.com/learning/domina-go/obtener-los-parametros-de-consulta-en-go?autoSkip=true&resume=false

- idCode := r.URL.Query().Get("idCode")
- w.Write([]byte(idCode))

## Devolver en respuesta XML en un método GET, 2:01, 2024-12-03

https://www.linkedin.com/learning/domina-go/devolver-en-respuesta-xml-en-un-metodo-get?autoSkip=true&resume=false

- uses xml.MarshalIndent()

## Crear un token con JWT en Go, 2:11, 2024-12-03

https://www.linkedin.com/learning/domina-go/crear-un-token-con-jwt-en-go?autoSkip=true&resume=false

- tied example, works

## Escribir en archivos de texto en Go, 1:56, 2024-12-03

https://www.linkedin.com/learning/domina-go/escribir-en-archivos-de-texto-en-go?autoSkip=true&resume=false

- I had os.Create and file.WriteString
- she uses the same

## Cree un servidor web en Go usando Gin, 2:01, 2024-12-04

https://www.linkedin.com/learning/domina-go/cree-un-servidor-web-en-go-usando-gin?autoSkip=true&resume=false

- she runs it on the default port 8080
- the example worked, I specified the port with r.Run(":3355")

## Pasar una función como parámetro en Go, 2:25, 2024-12-05

https://www.linkedin.com/learning/domina-go/pasar-una-funcion-como-parametro-en-go?autoSkip=true&resume=false

- sends a function via parameter

## Crear un middleware en Go, 3:28, 2024-12-07

https://www.linkedin.com/learning/domina-go/crear-un-middleware-en-go?autoSkip=true&resume=false

- she shows how to use http.NewServerMux

## Retornar una imagen en un método GET, 2:10, 2024-12-08

https://www.linkedin.com/learning/domina-go/retornar-una-imagen-en-un-metodo-get?resume=false

- works nicely, serves an image from route

## Retornar un archivo xlsx en un método GET, 3:15, 2024-12-09

https://www.linkedin.com/learning/domina-go/retornar-un-archivo-xlsx-en-un-metodo-get?autoSkip=true&resume=false

- she creates an xslx file on the fly and returns it with Content-Disposition
- you can download it in Postman and then open it, worked first time

## WebSocket en Go, 5:20, 2024-12-10

https://www.linkedin.com/learning/domina-go/websocket-en-go?autoSkip=true&resume=false

- uses "interruptions"
- she runs two terminals, one server and one client

## VOCAB - SPANISH

```
timed out
tiempo de espera agotado

but if the connection was successful
pero si la conexión fuera exitosa; fuera = subjunctive imperfect; conj=ser

we have the following server
tenemos el siguiente servidor

it looks kind of weird
se ve medio raro

we have been asked to return a client's information through a service
nos han solicitado devolver por medio de un servicio la information de una clienta

resource not found
recurso no encontrado

I'm going to go to my code
voy a irme a mi código

if there was an error
si hubiera un error

each request that arrives on the route
cada solicitud que llegue a la ruta; pr=YAY-gay

after this
luego de esto

the first thing we are going to do is
lo primero que vamos hacer es

add, subtract and multiply
sumar, restar y multiplicar

do you know anything about it
sabes algo al respecto

let's do the test
hagamos la prueba

we look at how we can create a web server
vemos cómo podemos crear un servidor web

a high performance web framework
un framework web de alto rendimiento

if there was an error
si hubiera un error; conj=haber; hubiera = imperfect subjunctive of haber

it allows you to format the output with indents
él permite formatear la salida con sangrías

we send the query
enviamos la consulta

let's go to the code
vamos a irnos al código

the information that comes in this parameter
la información que venga en esta parámetro; venga = subjunctive of venir; conj=venir

after the question mark
después del signo de pregunta

pr. alphabet, letters with standard pronunciations
a, b, c, d, e, _, _, _, i, _, k, _, _, _, _, o, p, q, _, _, t, u, _, _, _, _, _

pr. alphabet, letters with different pronuncations
f = EFF-ay, g = HEH, h = AH-chah, j = HOH-ta, l = ELL-ay, m = EM-ay, n = EN-ay, n = EN-yeh, r = ERR-ay, s = ESS-ay, v = OO-veh, w = oo-veh-DUB-lay, x = EK-ees, y = ee-grey-EGG-ah, z = SET-ah

each %v will be replaced by a value from our map
cada %v va a ser sustituido por un valor de nuestro mapa

and our wish is to turn it into a string
y nuestro deseo es lo convertirlo en un string

divided into five groups separated by hyphens
divididos en cinco grupos separados por guiones; pr=gee-OH-nays

we are going to go through the slice using a for loop
vamos a recorrer el slice por medio de un for

so we want to remove it from the slice
por lo que deseamos lo eliminar del slice; conj=desear; deseamos is present desear, to wish, regular

conj. dar present
doy das da damos dais dan

conj.dar preterite
di diste dio dimos disteis dieron

however, we realized that Thursday is a holiday
sin embargo, nos dimos cuenta que el día jueves es feriado; conj=dar; dimos is preterite of dar

and I click on Send
y le doy clic en Enviar; question=what does the "le" do here

I make sure it is a POST method
me aseguro que sea un método POST

let's set the header
vamos a setear el header

where we indicate that there was an error
donde indicamos que hubo un error; hubo = preterite of haber; conj=haber

by means of an error
por medio de un error

but if everything is successful
pero si todo es exitoso

in case there was an error
en caso de que existiera un error; conj=existir, existiera = imperfect subjunctive of existir

this helps us unpack the elements
esto nos ayuda a desempaquetar los elementos

to check if we saved the file
para comprobar si guardamos el archivo

otherwise we will indicate that it does not exist
caso contrario vamos a indicar que no existe

and in square brackets
y entre paréntesis cuadrados

let's see in the case of React
veamos en el caso de React; veamos = imperitive and subjunctive present

conj. ver present
veo, ves, ve, vemos, veis, ven

conj. ver preterite
vi viste vio vimos visteis vieron

it has a name associated with it
tiene asociada un nombre

two integers
dos numeros enteros

what we are going to do is place an if-statement here
lo que vamos a hacer es colocar una declaración if aquí

the last name
el apellido

we may not remember the name
tal vez no recordemos el nombre; conj=recordar; recordemos is present subjunctive of recordar, regular but o>ue

we subtract 1
le restamos 1

let's do that
hagamos eso; hagamos = present subjunctive of hacer; conj=hacer

we can do it through the index
lo podemos hacer por medio del índice

since this will be the one to whom the prize will be awarded
dado que esta será al que se le entregue el premio; entregue = subjunctive present of entregar, almost regular (yo preterite: entregué); conj=entregar

we have realized that there was a mistake
nos hemos dado cuenta que hubo un error; dar = dado, dando; hubo = preterite of haber; rank=4.92

the winning numbers of a raffle
los números ganadores de una rifa

when you don't know the amount of arguments
cuando desconozcas la cantidad de argumentos

to loop through this string array
para recorrer este arreglo de string; rank=4.89

in case it happens
en caso de que pase

this is because it may be possible that the string is not a valid number
esto se debe a que puede ser posible que la cadena no sea un número válido.

this can be done to make the data compatible
eso se puede hacer para que los datos sean compatibles; sean = ser present subjuctive

let's hit Enter
vamos a darle Enter; "le" is just a filler pronoun

double quotes
las comillas dobles

we place a backtick in front and at the end
colocamos un backtick adelante y al final; colocar = to place, almost regular, preterite: yo coloqué

we have achieved our goal
hemos logrado nuestro objetivo

as part of the task
como parte de la tarea

in the matter we have
en el asunto tenemos

we have been asked to send an email
nos han solicitado enviar un correo electrónico

the solution to common challenges
la solución desafíos comunes

although you may have encountered challenges
aunque quizás te hayas encontrado con retos; hayas = present perfect subjunctive

you have doubts
te surgen dudas; literally "doubts emerge for you", surgir = arise, emerge, appear

has it happened to you that while you are programming
te ha pasado que mientras estás programando

```
