import type { Flashcard, Tutorial, Stats } from '../types';
import rawStats from './data/stats.json';
import rawTutorials from './data/tutorials.json';
import * as appTools from '../appTools';
import * as importModel from './importModel';

export let numberOfErrors = 0;

const [sourceFlashcards, errorCount] = importModel.getSourceFlashcards();
numberOfErrors += errorCount;

export const getFlashcards = (): Flashcard[] => {
	const flashcards: Flashcard[] = [];
	for (const sourceFlashcard of sourceFlashcards) {
		const flashcard: Flashcard = {
			language: sourceFlashcard.language,
			front: sourceFlashcard.front,
			back: sourceFlashcard.back,
			whenCreated: sourceFlashcard.whenCreated,
			extras: sourceFlashcard.extras,
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
			fileIdCode: rawTutorial.fileIdCode,
			status: rawTutorial.status === '' ? 'not yet started' : rawTutorial.status,
			category: appTools.getCategoryItem(rawTutorial.topics)
		};
		tutorials.push(tutorial);
	}
	return tutorials;
};
