# Skeleton UI Library for Svelte

- https://www.youtube.com/watch?v=P_A0qQ7AuK8
- 16:12
- english
- Skeleton, Svelte
- 4.6
- intro to SkeletonUI for Svelte

## watchlog

- 2024-11-02 - 01:20
- 2024-11-02 - 16:12

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

- now creating site with Skeleton itself: `npm create skeleton-app@latest my-skeleton-app` but it installs 4 not 5 
- but it has the whole background styled

## create theme

- theme, create them
- show theme source, copy
- create file with that content in root called `theme002.ts`
- change these two lines

```
export const theme002: CustomThemeConfig = {
	name: 'theme002',
```

- **tailwind.config.ts**

```
import { theme002 } from './theme002';

plugins: [
	skeleton({
		themes: {
			custom: [theme002]
		}
	})
]
```

- **app.html**

```
<html class="dark">
<body data-theme="theme001">
```

## more features

- variant-ghost-primary
- talking about app shell that is deprecated
- svelte:fragment
- drawers
- skeleton has a local storage store, which is like the svelte store
- input chips


## VOCAB - SPANISH

```

```
