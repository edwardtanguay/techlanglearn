import type { Category } from './types';

export const categories = (): Category[] => {
	return [
		{
			idCode: 'current',
			name: 'current job',
			topics: ['react', 'go', 'gitlab', 'ink', 'easyPeasy', 'docker', 'git'],
			description: 'technology needed for current job'
		},
		{
			idCode: 'retro',
			name: 'retro tech',
			topics: [
				'vba',
				'php',
				'perl',
				'java',
				'googleSheets',
				'regex',
				'canvas',
				'npm',
				'webpack',
				'scala'
			],
			description: 'technology that reminds me of developing back in the day'
		},
		{
			idCode: 'fringe',
			name: 'fringe tech',
			topics: ['webflow'],
			description: "interesting technology that probably won't be needed for a job"
		},
		{
			idCode: 'personalPublishing',
			name: 'personal publishing',
			topics: ['obsstudio'],
			description: 'tech that helps produce YouTube tech videos, etc.'
		},
		{
			idCode: 'design',
			name: 'design skills',
			topics: ['grid', 'css', 'flexbox'],
			description: 'technology needed for web design'
		},
		{
			idCode: 'extend',
			name: 'future jobs',
			topics: [
				'nextjs',
				'monorepos',
				'flask',
				'python',
				'kinde',
				'svelte',
				'svelte5',
				'sveltekit',
				'mux',
				'githubActions',
				'laravel',
				'cypress',
				'playwright',
				'trpc',
				'kotlin'
			],
			description: 'technology that might be needed for future jobs'
		},
		{
			idCode: 'linux',
			name: 'linux expert',
			topics: ['sed', 'awk'],
			description: 'technology that makes you a Linux expert'
		},
		{
			idCode: 'general',
			name: 'general skills',
			topics: ['androidUser'],
			description: 'general skills and technology'
		}
	];
};
