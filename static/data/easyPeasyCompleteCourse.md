# React Redux Easy Peasy Complete Course

- https://www.youtube.com/watch?v=sPFaDYaTUeI&list=PLNTXksYYFsn9jnS_7Vpk9Qi2vb5N-MdOM
- 00:47:00
- english-spanish
- easy peasy
- 4.82
- 2020
- old course but checking it out, uses TypeScript, only doing parts 2-4, part 1 was too simple

## Actions - 8:18, 2024-11-22

https://www.youtube.com/watch?v=8NDIQyXEQok&list=PLNTXksYYFsn9jnS_7Vpk9Qi2vb5N-MdOM&index=2

- he uses "state" not "store" for the action variable
- note he also uses <this>
- he uses a state variable redundantly

## Thunks - 10:18, 2024-11-22

https://www.youtube.com/watch?v=1L6i1ZMPrvc&list=PLNTXksYYFsn9jnS_7Vpk9Qi2vb5N-MdOM&index=3

- mostly used for asynchronous operations
- instead of separatinb based on symantic categories, he separates into state, action and thunks, an interface for each of them
- thunks are loaded with the actions
- he calls his thunks "...Thunk" at the end
- strange that he makes state variables to shadow all his easy-peasy variables

## Advanced - 19:04, 2024-11-22

https://www.youtube.com/watch?v=2HGwD0qWVuc&list=PLNTXksYYFsn9jnS_7Vpk9Qi2vb5N-MdOM&index=4

- thunks can use other store
- he has a two stores: books and users
- he has index.ts instead of store.ts
- and he calls it Model and not StoreModel
- I changed dataModel to dataManager to distinguish it from the store models
- he's extending the stores but it is not clear why##hesjisext
- he create interfaces, instead of types
- has Book | any
  - then changed it to question marks on the properties
- makes computer field
- he only has Computed<this> but that doesn't work for me
- thunk can receive injections
- he does: borrowBookThunk: Thunk<this, number, undefined, Model>
- he uses getStoreState
  - good use is authenticataion
- he installs axios
- he says see you in the next one, but apparently he never made it

## VOCAB - SPANISH

```
let's change it to no
combiémoslo a no

so we need to provide that name
por lo que debemos proporcionar ese nombre

what this function is about and what it does
qué se trata esta función y qué hace

I like that standard because
me gusta esa estandár porque

así es como lo hace
this is how you do it

let me save it
déjeme guardarlo

I'll explain it to you a little later
explicártelo un poco más adelante

```
