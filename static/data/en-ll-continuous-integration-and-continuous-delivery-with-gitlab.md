# Continuous Integration and Continuous Delivery with GitLab

https://www.linkedin.com/learning/continuous-integration-and-continuous-delivery-with-gitlab

- duration: 01:27:00
- language: en
- topics: gitlab
- rank: 4.7
- description: good basics of GitLab, from 2022 so not that old

## intro - 0:57, 2024-10-28

- esta es solo una introducción básica

## what is GitLab - 1:39, 2024-10-28

- GitLab es más una solución de extremo a extremo que GitHub
- te permite reemplazar Jenkins
- el CD rara vez se automatiza

## what problem needs solving - 2:54, 2024-10-29

- Software Development Lifecycle (SDLC)
  - analysis
  - design
  - development
  - deployment and delivery
  - maintenance

## set up a project - 3:33, 2024-10-31

- me estoy registrando para una cuenta de GitLab.com
- "create blank project": gitlabcicd-showcase
- project deployment target: vercel
- Web IDE

## Lean manufacturing models - 3:18, 2024-10-31

- just-in-time manufacturing
  - smaller batches
  - los materiales se compran justo antes de que se necesiten
  - los inventarios de productos terminados se mantienen bajos
- theory of constraints
  - what is holding it back
- **increasing cadence**
  - more frequent and smaller changes
  - get features to users more quickly
  - improved quality and stability
  - reduce the delta, i.e. the amount of change in each release
  - easier to fix bugs when there are smaller changes
  - hotfixes just go out with the next scheduled deployment

## Continuous integration - 5:15, 2024-10-31

- an automated system
- merge right after the changes are complete
- alternative is where some poor developer is tasked with combining all the changes into one deployment
- you need to QA every change as soon as it's ready to merge
- require that a test is passed before code can be merged
- driven by when the work is ready
- los cambios principales se han fusionado a lo largo del sprint.
- las pruebas más caras deberían ser las últimas
- testing
  - syntax testing and linting - making sure the text of your code is valid and conform
  - unit testing - individual functions, components, APIs
  - integration testing - whether specific functions, components, APIs work correctly together
  - end-to-end testing - test the application as the user experiences it, all the real-world scenarios that your code might encounter
- your tests should happen in a pipeline
- find out if there is something broken within minutes of submitting your merge request
- fail early and often
- quit as soon as something goes wrong
- quickly fix issue and retest
- put time and effort into comprehensive tests

## Creating a pipeline, 3:34, 2024-10-31

- I don't have a template like he does
  - ok, I have to choose, on the left: Build > Pipelines
  - use **Bash template**
  - will be using built-in containers in this course
  - you don't need containers, you could deploy to a cloud service
  - before_script is a global section
  - test1 and test2 run in parallel
  - deploy at end
  - doesn't make changes at the moment
  - main branch and commit changes

## Running your pipeline, 2:43, 2024-10-31

- click on the hash to go to pipeline
- can click on rerun icon
- regardless of which you click on, they all go to same logs
- rerunning my second test, it failed

## Going deeper with pipelines, 5:05, 2024-10-31

- there is no interaction between the stages
- I don't have lint button
- edit pipeline
- "file.txt" assumes root of repository

```
- stages:
  - "build"
  - "test"
  - "deploy"
```

- and then this:

```
build1:
  stage: build
  script:
    - echo "Do your build here"
    - echo "this is a file"|tee file.txt
  artifacts:
    paths:
      - "file.txt"

test1:
  stage: test
  script:
    - echo "Do a test here"
    - echo "For example run a test suite"
    - cat file.txt
  dependencies:
    - "build1"
```

- he advises to start with basic pipelines

## Adding a test, 6:32, 2024-10-31

- testing this:

```
test1:
  stage: test
  script:
    - echo "testing for string 'gitlab'"
    - grep "gitlab" file.txt
  dependencies:
    - "build1"
```

- worked, test1 failed, but it didn't show exactly where
- now adding variable
- it's not clear if there is one pipeline or many
- can't see where he added the variable to override
- my variable in settings/CICD didn't work
- he has "Run pipeline"
- I have "New pipeline" same place, same button
- okay, but "New pipeline" goes to "Run pipeline" page
- great, it worked then

```
image: busybox:latest

variables:
  MY_WORD2: "gitlab"

stages:
  - build
  - test
  - deploy

before_script:
  - echo "note_beforescript"

after_script:
  - echo "note_afterscript"

build1:
  stage: build
  script:
    - echo "note_build1"
    - echo "$APP_TITLE"
    - echo "this is a file $MY_WORD2"|tee file.txt
  artifacts:
    paths:
      - "file.txt"

test1:
  stage: test
  script:
    - echo "note_test"
    - echo "testing for string 'gitlab'"
    - grep "gitlab" file.txt
  dependencies:
    - "build1"

test2:
  stage: test
  script:
    - echo "note_test2"
    - echo "Do another parallel test here"
    - echo "For example run a lint test"

deploy1:
  stage: deploy
  script:
    - echo "note_deploy"
  environment: production
```

## Generate a website, 3:48, 2024-10-31

- now a more realistic example
- create markdown file: website.md
- `image: alpine:latest`, it can do more than busybox
- this works:

```
image: alpine:latest

stages:
  - build
  - test
  - deploy

render_website:
  stage: build
  script:
    - apk add markdown
    - markdown website.md | tee index.html
  artifacts:
    paths:
      - "index.html"

test_website:
  stage: test
  script:
    - apk add libxml2-utils
    - xmllint --html index.html
  dependencies:
    - "render_website"

deploy_website:
  stage: deploy
  script:
    - echo "note_deploy"
  environment: production
```

## CD concepts, 5:26, 2024-10-31

- continuous delivery happens in an environment
- environment
  - standalone hardware and software stack
  - can run entire application
  - configured for a specific purpose
    - development
      - developer's laptop
      - one development environment per developer
      - easily repared or replaced if broken
      - if they delete data in the database by mistake, they should be able to quickly restore it
    - QA
      - intended for testing
      - similar to production
      - usually has only dummy data
    - staging
      - a clone of production
      - matches as closely as possible to production
      - including data in database as long as it doesn't violate privacy laws
      - point is to test something in production without releasing it in production
    - production
      - customer-facing servers
      - also automated testing in prod
      - should be easy to roll back if there is a problem
- code promotion
  - your environments correlate to your branches
  - GitFlow
    - dev and QA is where feature branches can be developed and tested
    - long-running development branch
    - dev branch is deployed to QA
    - push to release branch and tagged
    - pushed to prod

## Environments, 6:43, 2024-11-01

- creates new user with IAM server
- you get an access key ID and a secret access key##accessandsec
- sets up a bucket
- in GitLab he creates a new environment and calls it qa##thenwissinew

## Environment variables, 4:28, 2024-11-01

- settings > ci/cd > variables
- adds variables for AWS
- build > pipeline editor
- creates variable: S3_BUCKET
- in render_website, adds `mkdir -p public`
- deploy_website
  - `apk add aws-cli`
  - `aws s3 cp ./public/ s3://$S3_BUCKET/ --recursive`
  - `dependencies:`
    - `render_website`

## First deployment pipeline, 5:22, 2024-11-01

- view pipeline
- it was deployed
- changes name from `deploy_website` to `deploy_qa`
- every commit to our repo will trigger a deployment
  - so we will add a rules section to the deploy job

```
rules:
  - when: manual
```

- adds `deploy_staging`
- adds job-level variable
- now he has three deploy jobs##nowhas3deploy

## Automating deployments, 5:40, nnn

https://www.linkedin.com/learning/continuous-integration-and-continuous-delivery-with-gitlab/automating-deployments?resume=false

- replaces 'when manual' with: if: $CI_COMMIT_BRANCH == "main"
- @@1:11

## VOCAB - SPANISH

```
GitLab provides us with some variables
GitLab nos proporciona algunas variables

we will return to our project
volveremos a nuestro proyecto

for everything else
para todo lo demás

I saved that earlier
lo guardé antes

it's right here
está justo aquí

we need to allow the user to upload files
necesitamos permitir al usuario cargar archivos

we will need to add some permissions
necesitaremos agregar algunos permisos

and we will have to check this box
y tendremos que marcar esta casilla

and then scroll down and then back up
y luego desplácese hacia abajo y luego hacia arriba

you can repeat the steps
puedes repetir los pasos

you'll want to copy this password
querrás copiar esta contraseña

you will have to regenerate it
tendrás de regenerarlo

make sure you save them because otherwise you will lose them
asegúrate de guardarlos porque de lo contrario los perderás

I am going to use it to paste some text into the input box
voy a usarlo para pegar algo de texto en el cuadro de entrada

just follow exactly what I do and you should be fine
solo sigue exactamente lo que hago y deberías estar bien

we didn't do anything with it
no hicimos nada con él

and I will delete some of these comments
y borraré algunos de estos comentarios

it may be tempting to create
puede ser tentador crear

let's go back
regresemos

I got to this page by browsing
llegué a esta página navegando

so we can move on
así que podemos seguir adelante

something starts to happen
algo comienza a suceder

below the image
debajo de la imagen

so let's take a look
así que echemos un vistazo

that investment will pay for itself many times over
esa inversión amortizará muchos veces

it would be nice to find out
sería bueno averiguarlo

the next level
el siguiente nivel

each layer of testing
cada capa de prueba

if you have ever seen a documentary
si alguna vez has visto un documental

it's pretty easy to find out
es bastante fácil averiguar

due to the delta
debido al delta

both approaches
ambos enfoques

pr. 1930
meel NOH-vay see-EN-troh TRAYN-tah

in the last century
en el siglo pasado

click on the tab
haga clic en la pestaña

we will not cover in this course
no cubriremos en este curso

take a couple of minutes and
tómate un par de minutos y

a branch called main
una rama llamada main

a couple of things I want to point out
un par de cosas que quiero señalar

we will leave these unset
dejaremos estos sin establecer

work back and forth
trabajar de un lado de otro

the other thing that happens
la otra cosa que sucede

I'm testing a hunch
pruebo una corazonada

testing a hunch
probando una corazonada

app developer
desarrollador de aplicaciones

this is because it is very difficult to distinguish
esto se debe a que es muy difícil distinguir

it is assumed that
se supone que

this concept helps to emphasize
Este concepto ayuda a enfatizar

let's drink to software development
brindemos por el desarrollo de software; [day-sah-ROY-yoh]

it consists of the following parts
es consta de las siguientes partes; [see-g|ee-EN-tays]

so let's take a look at the general process
así que echemos un vistazo el proceso general

continuous delivery
entrega continua

to a large extent
en gran medida
```
