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

## Implementación del Standard Layout en Go, 3:13, nnn

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
  - *_test.go can go in other directories where the files reside that they test

## VOCAB - SPANISH

```
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
