import type { Flashcard, PageStatus } from './types';
import * as dataModel from './db/dataModel';
import { getRandomItemsFromArray } from './tools';

// values
const siteLocation = import.meta.env.VITE_SITE_LOCATION === 'dev' ? 'dev' : 'online';
let pageStatus: PageStatus = $state('ready');
const flashcards = $state(
	dataModel.getFlashcards().sort((a, b) => (a.whenCreated < b.whenCreated ? 1 : -1))
);
const tutorials = $state(dataModel.getTutorials());
let filteredTutorials = $state(tutorials);

const topicList = (): string[] => {
	const topicCounts: Record<string, number> = {};

	for (const tutorial of tutorials) {
		const tutorialTopics = tutorial.topics.split(',').map((topic) => topic.trim());
		for (const topic of tutorialTopics) {
			topicCounts[topic] = (topicCounts[topic] || 0) + 1;
		}
	}

	const topics = Object.keys(topicCounts).sort((a, b) => {
		const countDiff = topicCounts[b] - topicCounts[a];
		if (countDiff !== 0) {
			return countDiff;
		}
		return a.localeCompare(b);
	});

	topics.unshift('ALL');
	return topics;
};

// errorMessage
// TODO: pass through full array of errors, not just number
let errorMessage = $state(
	dataModel.numberOfErrors !== 0
		? `Number of import errors: ${dataModel.numberOfErrors} (see browser console)`
		: ''
);

// computed values
const randomFlashcards = $state(getRandomItemsFromArray(flashcards, 3));

export const getStore = () => {
	return {
		// state
		get siteLocation() {
			return siteLocation;
		},
		get pageStatus() {
			return pageStatus;
		},
		get errorMessage() {
			return errorMessage;
		},
		get flashcards() {
			return flashcards;
		},
		get tutorials() {
			return tutorials;
		},
		get filteredTutorials() {
			return filteredTutorials;
		},

		// computed state
		get randomFlashcards() {
			return randomFlashcards;
		},

		// actions
		setPageStatus: (_pageStatus: PageStatus) => {
			pageStatus = _pageStatus;
		},
		setErrorMessage: (_errorMessage: string) => {
			errorMessage = _errorMessage;
		},
		handleRandomFlashcardToggle: (f: Flashcard) => {
			f.isOpen = !f.isOpen;
		},
		filterTutorials: (topic: string) => {
			if (topic === 'ALL') {
				filteredTutorials = tutorials;
			} else {
				filteredTutorials = tutorials.filter((m) => m.topics.includes(topic));
			}
		},
		getSelectableTopics: () => {
			return topicList();
		},
		toggleSortColumn: (fieldIdCode: string) => {
			if (fieldIdCode === 'rank') {
				filteredTutorials.sort((a, b) => (a.rank < b.rank ? 1 : -1));
			}
			if (fieldIdCode === 'year') {
				filteredTutorials.sort((a, b) =>
					String(a.year) + String(a.rank) < String(b.year) + String(b.rank) ? 1 : -1
				);
			}
			if (fieldIdCode === 'language') {
				filteredTutorials.sort((a, b) =>
					a.language + String(a.rank) < b.language + String(b.rank) ? 1 : -1
				);
			}
		}
	};
};
