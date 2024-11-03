import rawFlashcards from './data/flashcards.json';
import type { Flashcard } from './types';

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
