import { defineConfig } from 'vitest/config';
import { sveltekit } from '@sveltejs/kit/vite';

export default defineConfig({
	plugins: [
		sveltekit(),
	],

	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	},
	logLevel: 'error',
	css: {
		preprocessorOptions: {
			scss: {
				api: 'modern-compiler' // or "modern"
			}
		}
	}
});
