import type { Category } from './types';

export const getCategoryItem = (topicsList: string): Category | null => {
	const topics = topicsList.split(',').map(m => m.trim());
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

	for (const category of categories) {
		for (const categoryTopic of category.topics) {
			for (const topic of topics) {
				console.log('compare', categoryTopic, topic);
				if (categoryTopic === topic) {
					return category;
				}
			}
		}
	}
	return null;
};
