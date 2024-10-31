# Continuous Integration and Continuous Delivery with GitLab

- https://www.linkedin.com/learning/continuous-integration-and-continuous-delivery-with-gitlab
- 01:27:00
- english-spanish
- GitLab
- 4.7

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

## Creating a pipeline, 3:34, nnn

- 

## vocab - spanish

```
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
