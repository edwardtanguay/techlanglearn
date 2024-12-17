# Learn Ink in 15 minutes

https://www.youtube.com/watch?v=KSRpcftVyKg

- duration: 00:14:46
- language: en
- topics: ink
- rank: 4.2
- description: a quick video to learn how to make dynamic conversations using Ink, a tool used in gaming
- year: 2021
- status: finished

## watchlog

- 2024-11-12 - 14:46

## intro

- download Inky the editor
- CTRL + and = to make smaller
- CTRL / to comment
- or block comments
- knot = a named section of the story script
- divert - directs the story flow to a different part of the script
  - -> END and -> DONE
- asterisk is a choice
- [Yes] will not print ou the choice
- loops
- basic

```
Where would you like to go?
    * [to the store] store
        Then you need some money.
    * [to the beach] beach
        Then you need some sun tan lotion.
```

- default choice

```
-> main

=== main ===
Which flavor?
    * [cherry]
        That is red.
        -> main
    * [blueberry]
        That is blue.
        -> main
    * [lemon]
        That is yellow.
        -> main
    * ->
        There are no more flavors to choose from.
-> END
```

- sticky choices
  - plus instead of star

```
-> main

=== main ===
Which flavor?
    + [cherry]
        That is red.
        -> main
    + [blueberry]
        That is blue.
        -> main
    + [lemon]
        That is yellow.
        -> main
-> END
```

- variables
  - VAR
  - temp
  - CONST
- values can be text, boolean, numbers and diverts
- show choice

```
VAR flavor = ""

-> main

=== main ===
Which flavor?
    + [cherry]
        ~ flavor = "cherry"
        -> chosen
    + [blueberry]
        ~ flavor = "blueberry"
        -> chosen
    + [lemon]
        ~ flavor = "lemon"
        -> chosen

=== chosen ===
you choose {flavor}

-> END
```

- pass parameter

```
-> main

=== main ===
Which flavor?
    + [cherry]
        -> chosen("cherry")
    + [blueberry]
        -> chosen("blueberry")
    + [lemon]
        -> chosen("lemon")

=== chosen(flavor) ===
You choose {flavor}.

-> END
        -> chosen("cherry")
```

- syntax: weave
- gathers

```
-> main

=== main ===
Which flavor?
    * cherry
    * blueberry
    * lemon

- Then enjoy!

-> END
```

- nested choices

```
-> main

=== main ===
Which flavor?
    * cherry
        dark or light red?
            * * dark red
            * * light red
        - - ok
    * blueberry
        dark or light blue?
            * * dark blue
            * * light blue
        - - ok then
    * lemon
        dark or light yellow?
            * * dark yellow
            * * light yellow
        - - very well then

- enjoy your ice cream

-> END
```

- conditional logic

```
VAR hasFinished = false
VAR kilometers = 20

{hasFinished: congratulations|keep on going}
{kilometers < 10: you have just started|you are quite a ways}
{kilometers:
- 1: one
- 2: two
- 3: three
- else: (unknown)
}

Your race status: {getRaceStatus(kilometers)}

Your choice
    * {hasFinished} claim your medal
    * {hasFinished} get a drink
    * {!hasFinished} keep running, {getRaceStatus(kilometers)}
    * {!hasFinished} quit

=== function getRaceStatus(k) ===
~ return "you have {42 - k}K to go"
```

- see all built-in functions in the [Writer's Manual](https://github.com/inkle/ink/blob/master/Documentation/WritingWithInk.md)
- stitches
  - a knot is like a scene, while a stitch is like a sub-scene or a specific part within that scene
- includes
  - you can do this with the editor as well
- variable texts
- tunnels
- threads
- advanced state tracking

## VOCAB - SPANISH

```
but above all I hope this has been useful
pero sobre todo espero que esto haya sido ùtil

anyway, please give it a like so more people can see it
de todos modos, dale me gusta para que más personas lo vean

this is something you should keep in mind
esto es algo que debes tener en cuenta

this means that it can handle any character
esto significa que puede manejar cualquier carácter

the good news is that ink has no limitations
la buena noticia es que ink no tiene limitaciones

as the story grows, you may want to do things
a medida que la historia crece, es posible que desees hacer cosas

the variable cannot be changed after it has been declared for the first time
la variable no se puede cambiar después de haber sido declarada por primera vez

next, let's add a variable
a continuación, agreguemos una variable

there is nothing more to show
no hay nada más que mostrar

as you may have guessed
como habrás adivinado

this is because we need to activate it
esto se debe a que necesitamos activarlo

the text does not yet appear in the story
el texto todavía no aparece en la historia

if you go to the drop-down menu on the top left
si vas al menu desplegable en la parte superior izquierda

we will go to a new line
iremos a una nueva línea

to take a step back
para retroceder un paso

once downloaded, you can start the program
una vez descargado, puedes iniciar el programa
```
