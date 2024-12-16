import type { Category } from './types';

export const categories = (): Category[] => {
	return [
		{
			idCode: 'current',
			topics: ['react', 'go'],
			description: 'technology needed for current job'
		},
		{
			idCode: 'retro',
			topics: ['vba', 'php', 'perl', 'java', 'googleSheets', 'regex'],
			description: 'technology that reminds me of developing back in the day'
		},
		{
			idCode: 'fringe',
			topics: [],
			description: "interesting technology that probably won't be needed for a job"
		},
		{
			idCode: 'design',
			topics: ['grid'],
			description: 'technology needed for web design'
		},
		{
			idCode: 'extend',
			topics: ['nextjs'],
			description: 'technology that might be needed for future jobs'
		},
		{
			idCode: 'linux',
			topics: ['sed', 'awk'],
			description: 'technology that makes you a Linux expert'
		}
	];
};
