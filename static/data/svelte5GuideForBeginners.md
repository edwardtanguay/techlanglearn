# Svelte 5 Guide For Beginners

- https://www.youtube.com/watch?v=tErKyuUTzsM
- 21:40
- english
- svelte5
- 4.8

## watchlog

- 2024-10-28: 00:00 - 04:18

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
