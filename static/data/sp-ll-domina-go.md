# Domina Go

https://www.linkedin.com/learning/domina-go

- duration: 00:50:00
- language: sp
- topics: go
- rank: 4.82
- description: short videos but seems to cover basic topics, and from 2023, but some videos are very simple, e.g. how to get the last item out of a slice

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

## Enviar un JSON en una solicitud POST en Go, 2:13, nnn

https://www.linkedin.com/learning/domina-go/enviar-un-json-en-una-solicitud-post-en-go?autoSkip=true&resume=false

- checks if method is post with:

```
if r.Method == http.MethodPost
```

## VOCAB - SPANISH

```
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
