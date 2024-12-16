import type { Category } from "./types";

export const categories = (): Category[] => {
	return [
		{ idCode: 'current', topics: ['react', 'go'], description: '' },
		{
			idCode: 'retro',
			topics: ['vba', 'php', 'perl', 'java', 'googleSheets', 'regex'],
			description: ''
		},
		{ idCode: 'fringe', topics: [], description: '' },
		{ idCode: 'extend', topics: [], description: '' },
		{ idCode: 'linux', topics: ['sed', 'awk'], description: '' }
	];
}
