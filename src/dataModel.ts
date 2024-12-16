import rawFlashcards from './data/flashcards.json';
import type { Flashcard, Tutorial, Stats } from './types';
import rawStats from './data/stats.json';
import rawTutorials from './data/tutorials.json';

export const getFlashcards = (): Flashcard[] => {
	const flashcards: Flashcard[] = [];
	for (const rawFlashcard of rawFlashcards) {
		const flashcard: Flashcard = {
			front: rawFlashcard.front,
			back: rawFlashcard.back,
			isOpen: false
		};
		flashcards.push(flashcard);
	}
	return flashcards;
};

export const getStats = (): Stats => {
	return rawStats;
};

export const getTutorials = (): Tutorial[] => {
	return rawTutorials;
};
