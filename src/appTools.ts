import type { Category } from './types';
import * as config from './config';

export const getCategoryItem = (topicsList: string): Category | null => {
	const topics = topicsList.split(',').map(m => m.trim());

	for (const category of config.categories()) {
		for (const categoryTopic of category.topics) {
			for (const topic of topics) {
				if (categoryTopic === topic) {
					return category;
				}
			}
		}
	}
	return null;
};
