import type { Flashcard, PageStatus } from './types';
import * as dataModel from './dataModel';
import { getRandomItemsFromArray } from './tools';

// values
const siteLocation = import.meta.env.VITE_SITE_LOCATION === 'dev' ? 'dev' : 'online';
let pageStatus: PageStatus = $state('ready');
let errorMessage = $state('');
const flashcards = $state(dataModel.getFlashcards());
const tutorials = $state(dataModel.getTutorials());

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
		sortTutorials: () => {
			// error: tutorials = any
			tutorials.sort((a, b) => (a.rank > b.rank ? 1 : -1));
		}
	};
};