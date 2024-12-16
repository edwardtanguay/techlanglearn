import { getRandomIndex } from './tools';
import type { Category } from './types';

export const getCategoryItem = (topics: string): Category => {
	const categories: Category[] = [
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
	console.log(topics);
	return categories[getRandomIndex(categories.length)];
};
