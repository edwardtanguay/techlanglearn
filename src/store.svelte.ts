import type { Flashcard, PageStatus } from './types';
import * as dataModel from './db/dataModel';
import { getRandomItemsFromArray } from './tools';

// values
const siteLocation = import.meta.env.VITE_SITE_LOCATION === 'dev' ? 'dev' : 'online';
let pageStatus: PageStatus = $state('ready');
const flashcards = $state(dataModel.getFlashcards());
const tutorials = $state(dataModel.getTutorials());

// errorMessage
let errorMessage = $state(
	dataModel.numberOfErrors !== 0
		? `Number of import errors: ${dataModel.numberOfErrors} (see console)`
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
		toggleSortColumn: (fieldIdCode: string) => {
			if (fieldIdCode === 'rank') {
				tutorials.sort((a, b) => (a.rank < b.rank ? 1 : -1));
			}
			if (fieldIdCode === 'year') {
				tutorials.sort((a, b) =>
					String(a.year) + String(a.rank) < String(b.year) + String(b.rank) ? 1 : -1
				);
			}
			if (fieldIdCode === 'language') {
				tutorials.sort((a, b) =>
					a.language + String(a.rank) < b.language + String(b.rank) ? 1 : -1
				);
			}
		}
	};
};
