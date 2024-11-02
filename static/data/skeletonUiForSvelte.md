# Skeleton UI Library for Svelte

https://www.youtube.com/watch?v=P_A0qQ7AuK8

- 00:16:12
- english
- Skeleton, Svelte
- 4.6
- intro to SkeletonUI for Svelte

## watchlog

01:20 - 2024-11-02

## intro

- a combination between Mantine and DaisyUI
- doesn't try to force a design system

## installation

- https://www.skeleton.dev/docs/get-started
- `npm i -D @skeletonlabs/skeleton @skeletonlabs/tw-plugin`
- `npm add -D @types/node`
- **tailwind.config.ts**

```
import type { Config } from 'tailwindcss';
import path from 'path';

import { skeleton } from '@skeletonlabs/tw-plugin';

const config = {
	darkMode: 'class',
	content: [
		'./src/**/*.{html,js,svelte,ts}',
		path.join(require.resolve('@skeletonlabs/skeleton'), '../**/*.{html,js,svelte,ts}')
	],
	theme: {
		extend: {}
	},
	plugins: [
		skeleton({
			themes: { preset: [ "skeleton" ] }
		})
	]
} satisfies Config;

export default config;
```

## VOCAB - SPANISH

```

```
