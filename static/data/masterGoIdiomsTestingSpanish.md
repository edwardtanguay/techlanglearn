# Master Go: Idioms and Code Testing

- https://www.linkedin.com/learning/domina-go-idioms-y-pruebas-de-codigo
- 01:10:00
- spanish-spanish
- Go
- 4.8
- advanced Go, coding practices, testing (unit and integration)

## intro - 0:49, 2024-11-01

- says course is about advanced Go, coding practices, testing (unit and integration)

## Declaración de paquetes en proyectos Go, 2:41, 2024-11-01

- break code up into packets
- the name of the package is the name of the directory
- lowercase is hidden##lowercasehidd

## Dependencias cíclicas en proyectos Go, 2:17, 2024-11-01

- Go detects circular dependenciees
- it helps you detect if package 1 uses package 2 but also package 2 uses package 1
- solves this by creating an abstraction

## El paquete Internal de Go, 1:44, 2024-11-01

- it's important to have libraries on the same level as main.go
- talks about loading in imports that you need
- e.g. import "myapp/internal/helper"

## Implementación del Standard Layout en Go, 3:13, 2024-11-09

- a collection of patterns
- https://github.com/golang-standards/project-layout
- cmd
  - executables
  - each subfolder contains a main.go for that application
- internal
  - can only be imported by packages in your project
  - protects internal details
- pkg
  - reusable code that can be used by other projects
- api
  - everything that pertains to an API
  - HTTP handlers, GraphQL, etc.
- web
  - for websites, contains the HTML, JavaScript and CSS files
  - also assets like images, audios that are served from the server
- configs
  - yaml and json
- scripts
  - all scripts for building, testing and deploying
- test
  - test helpers and data
  - \*\_test.go can go in other directories where the files reside that they test

## Ejemplo de layout: servicio de ArdanLabs en Go, 3:21, 2024-11-09

- he looks at another project with a slightly different version of the starter kit
  - app
    - the logic of the applicatin
  - business
    - business logic
  - foundation
    - the infrastructure of the application
  - vendor
    - dependencies of the project
  - zarf
- business can use code of foundation but not of app
- he types `go run ./...`
  - tells Go to find all .go files in the current directory and all subdirectories
- he imports the packages with github urls##thehandlsi
- you can only call functions that are lower in alphabetical order##onlydownkjd
  - this avoids circular references

## Utilizando la biblioteca estándar de Go, 3:13, 2024-11-10

https://www.linkedin.com/learning/domina-go-idioms-y-pruebas-de-codigo/utilizando-la-biblioteca-estandar-de-go?resume=false

- he talks about the standard library in general

## Utilizando opciones funcionales en Go, 3:59, 2024-11-10

https://www.linkedin.com/learning/domina-go-idioms-y-pruebas-de-codigo/utilizando-opciones-funcionales-en-go?autoSkip=true&resume=false

- creating a struct with a constructor
- withName##thewithtyp
- this is the difference##thisistheddddif
- nnn

## Manipulación de punteros en Go, 4:15, 2024-11-10

https://www.linkedin.com/learning/domina-go-idioms-y-pruebas-de-codigo/manipulacion-de-punteros-en-go?autoSkip=true&resume=false

- talks about & and \*
- when you send a struct to a function, a copy is made

## Gestión de contexto en aplicaciones Go, 4:39, 2024-11-12

https://www.linkedin.com/learning/domina-go-idioms-y-pruebas-de-codigo/gestion-de-contexto-en-aplicaciones-go?autoSkip=true&resume=false

- contexts have a life time which begins when they are created
- if a parent context is canceled, the child context is canceled as well
- WithValue
- use contexts to prevent resource hogging, e.g. in web servers

## Explorando el paquete Sync en Go, 2:59, 2024-11-13

https://www.linkedin.com/learning/domina-go-idioms-y-pruebas-de-codigo/explorando-el-paquete-sync-en-go?autoSkip=true&resume=false

- goroutine
  - a lightweight, concurrent function execution
  - a smaller, more efficient way to run tasks in parallel compared to traditional threads in other programming languages
- mutex
  - synchronization mechanism used to safely manage access to shared resources when multiple goroutines might access and modify them concurrently
  - goroutines do not execute in a predictable order
- sync.Once
  - se utiliza para garantizar que una operación se ejecute solo una vez

## Creación de Goroutines en Go, 1:13, 2024-11-14

- goroutines are used for
  - reading files
  - downloading data from the Internet

## Creación de canales en Go, 2:36, 2024-11-16

- goroutine
  - lightweight, concurrent function or method that runs independently
  - prefixed by "go"
  - great for tasks that need to run simultaneously, such as handling requests in a web server
- channels
  - provide a way for goroutines to communicate and synchronize by passing data between them
- goroutines provide the ability to run tasks concurrently
  - while channels manage the communication between these tasks

## Domina Go: Idioms y pruebas de código, 2:59, 2024-11-17

https://www.linkedin.com/learning/domina-go-idioms-y-pruebas-de-codigo/explorando-el-paquete-sync-en-go?autoSkip=true&resume=false

- muestra por qué necesitamos un WaitGroup
- muestra cómo se ejecuta con y sin un WaitGroup
- once.Do

## Creación de Goroutines en Go, 1:13, nnn

https://www.linkedin.com/learning/domina-go-idioms-y-pruebas-de-codigo/creacion-de-goroutines-en-go?autoSkip=true&resume=false

- nnn

## VOCAB - SPANISH

```
just once, as we expected
una única vez, como esperábamos

we can ensure that
podemos asegurarnos que

this is because
esto se debe a que

that has the protections inside
que dentro tiene las protecciones

a counter and a container //comparison
un contador y un contenedor

he works like a teacher but he doesn't know how to solve this problem //comparison
él trabaja como profesor pero no sabe cómo solucionar este problema

we see how to use
vemos cómo utilizar

a shared resource
un recurso compartido

a set of primitives
un conjunto de primitivas

instead of writing this code
en lugar de escribir este código

expensive operations
operaciones costosas

expensive operations
operaciones costosas

a single thread
un solo hilo

this way we can make sure that all goroutines have finished
de esta manera, podemos asegurarnos que todas las gorutinas hayan finalizado

an integer value
un valor entero

the shared resource can be a variable
el recurso compartido puede ser una variable

race conditions
condiciones de carrera

as it can lead to over coupling
ya que puede llevar a un acoplamiento excesivo

the only one that fails
el único que falla

after which the context will be cancelled
tras el cual se cancelerá el contexto

the text string
la cadena de texto

let's make sure that
vamos a aseguarnos que

when this time is up
cuando se cumpla este tiempo; note=cumplir: to fulfill

thirty, forty, fifty, sixty, seventy, eighty, ninety
treinta, cuarenta, cincuenta, sesenta, setenta, ochenta, noventa

the breakfast, the lunch, the dinner
el desayuno, el almuerzo, la cena

let's look at this with different examples
veámoslo con diferentes ejemplos

which begins when it is created
que comienza cuando se crea

freeing up resources if necessary
liberando recursos si fuera necesario [gr. subjunctive imperfect 1]

being able to define a time to cancel the operation
pudiendo definir un tiempo para cancelar la operación

as we have seen
como hemos visto

the variable will have been modified
la variable habrá sido modificada

will be applied to the struct
se aplicarán a la struct

this pointer will point to the memory address of the value
este puntero apuntará a la dirección de memoria del valor

let's store the value 17
vamos a almacenar el valor 17

an integer value
un valor de tipo entero

leaving the code much cleaner
quedando el código mucho más limpio

let's see this in an example
veámoslo con un ejemplo

if there is any error
si hubiera algún error [gr. subjunctive imperfect 1]

each of which modifies a field
cada una de las cuales modifica un campo

a function that receives a pointer
una función que recibe un puntero

which value corresponds to which field
qué valor corresponde a qué campo

if we were to add new fields
si añadiésimos nuevos campos [gr. subjunctive imperfect 2]

its wide range of packages
su amplia gama de paquetes

Go places a strong emphasis on testing
Go pone un fuerte énfasis en las pruebas

digital signatures
firmas digitales

the library provides many useful tools
la librería proporciona muchas herramientas útiles

to a large extent
en gran medida

high performance
alto rendimiento

the library is highly optimized
la librería está altamente optimizada

file management
manejo de archivos

above the current package
por encima del paquete actual

in such a way that
de tal modo que

higher layers of the application
capas más altas de la aplicación

therefore
por tanto

below it
por debajo de él

such as
tales como

deployment code
código despliegue [des-SPLEE-gay]

let's look at it with a small example
veámoslo con un pequeño ejemplo

and finally, I want to highlight the test directory
y por último, quiero destacar el directorio de pruebas

from outside the project
desde fuera del proyecto

pr. cmd
[say ehm-MAY day]

we can find it in the repository
podemos encontrarlo en el repositorio

a set of organizational patterns
un conjunto de patrones de organización

our hierarchy
nuestra jerarquía

which helps us
lo cual nos ayuda

which is a mistake
lo cual es un error

although it seems to make sense
aunque parece tener sentido

for this reason
por ello

in such a way that
de tal forma de

internal scope of the package
ámbito interno del paquete

we will have to expose the functions
tendremos que exponer las funciones

since it is an implementation detail
ya que es un detalle de implementación

we are hiding it on purpose
lo estamos ocultando a propósito

place the files in that directory
colocar los archivos en ese directorio

code coverage
cobertura de código

it's just the beginning
es solo el principio

you will immerse yourself in the code
te sumergirás en el código

provide you with an understanding
brindarte una comprensión

you are going to start
vas a empezar

optimized performance
rendimiento optimizado

new heights
nuevas alturas

```
