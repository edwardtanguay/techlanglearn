import type { Flashcard, PageStatus } from './types';
import * as dataModel from './dataModel';
import { getRandomItemsFromArray } from './tools';

// values
const siteLocation = import.meta.env.VITE_SITE_LOCATION === 'dev' ? 'dev' : 'online';
let pageStatus: PageStatus = $state('ready');
let errorMessage = $state('');
const flashcards = $state(dataModel.getFlashcards());

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
			console.log(11111, 'switching');
			f.isOpen = !f.isOpen;
		}
	};
};
