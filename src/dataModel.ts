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
	const tutorials: Tutorial[] = [];
	for (const rawTutorial of rawTutorials) {
		const tutorial: Tutorial = {
			topics: rawTutorial.topics,
			language: rawTutorial.language,
			duration: rawTutorial.duration,
			year: Number(rawTutorial.year),
			rank: Number(rawTutorial.rank),
			url: rawTutorial.url,
			title: rawTutorial.title,
			description: rawTutorial.description,
			platform: rawTutorial.platform,
			fileIdCode: rawTutorial.fileIdCode
		};
		tutorials.push(tutorial);
	}
	return tutorials;
};
