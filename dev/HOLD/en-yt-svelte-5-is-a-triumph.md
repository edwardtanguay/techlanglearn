# Svelte 5 Is A Triumph

https://www.youtube.com/watch?v=tC9KSwnczxQ

- duration: 00:22:03
- language: en
- topics: svelte5
- rank: 4.8
- description: another good Svelte 5 video, recorded all concepts in SvelteKit 5 showcase
- year: 2024
- status: finished

## watchlog

- 2024-11-09 - 5:48
- 2024-11-11 - 10:38
- 2024-11-14 - 12:47
- 2024-11-15 - 22:03

## reactivity

- $state() is used inside and outside svelte components
- in Svelte 4 you had to spread values for state arrays
  - no longer the case in Svelte 5
- $inspect(nnn) only works in development
- made store
- in 5, use onclick instead of on:click
- ref works with $effect##thearotoref
- you don't have to input $effect()
- onMount() is replaced with $effect()

## props

- props have simple syntax##propstypeds
- couldn't get children to work instead of slot, although it says slot is deprecated##couldchild

## slots

- will be deprecated
- replaced by {@render children()}

## form binding

- simply `bind:checked={isAvailable}`
- some changes in 5

## webapp

- everything is a function in 5
- mount(App, ...)

## fetching data

- if you use SvelteKit, you probably won't be using this
