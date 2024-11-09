# Svelte 5 Guide For Beginners

- https://www.youtube.com/watch?v=tErKyuUTzsM
- 21:40
- english
- svelte5
- 4.8

## watchlog

- 04:18 - 2024-11-01
- 10:08 - 2024-11-01
- 10:58 - 2024-11-06
- 13:21 - 2024-11-07
- 13:38 - 2024-11-07
- 21:40 - 2024-11-09

## Svelte 3

- you have to use requestAnimationFrame()
- Svelte 5 improves on this
- Svelte 3 introduced stores
- writable(0)
- so Svelte introduced $
  - only works inside svelte components

## Svelte 5 and runes

- runes
  - reduce complexity
- create SvelteKit 5 site
  - `npx sv create`

## side effects

- effect, notice lack of dependency arrays

```ts
let count = $state(0);
let double = $derived(count * 2);

$effect(() => {
	console.log(111, count);
});
```

- don't do this:##dontdothis1
- instead do this: ##insteidjdo1
- runes:: function-like symbols that provide isntructions to the Svelte provider
- runes are easy to type because they look like functions
- signal:: a reactive variable that causes parts of the component to automatically update when the variable's value changes
- [example](https://github.com/edwardtanguay/sveltekit5-showcase/blob/dev/src/lib/state/createUser.svelte.ts)

## universal reactivity

- universal reactivity:: a concept that means you can use the same logic inside and outside of Svelte components
- problem:##countprobl
- do it like this, thanks to the magic of signals##thankmagisig

## nested reactivity

- these are the same:##thesesamems
- uses a "proxy under the hood" so you don't have to use a getter/setter
- update an array state variable simply with .push()##thepush23
- another example#gif#addex101

## shared state

- runes are just reactive primitives
- SolidJS is similar to Svelte because they both use signals
- runes are just function-like symbols that give instructions to the Svelte compiler which then get turned to signals under the hood

