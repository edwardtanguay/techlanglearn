/* eslint-disable @typescript-eslint/no-explicit-any */
import { SourceFlashcardSchema, type SourceFlashcard } from '../types';
import _rawFlashcards from './data/flashcards.json';

// raw = untrustworthy data
// source = validated, cleaned and structured data

const rawFlashcards = _rawFlashcards as any[];

export const getSourceFlashcards = (): SourceFlashcard[] => {
	const sourceFlashcards: SourceFlashcard[] = [];
	for (const rawFlashcard of rawFlashcards) {
		const parseResult = SourceFlashcardSchema.safeParse(rawFlashcard);
		if (parseResult.success) {
			sourceFlashcards.push(rawFlashcard);
		} else {
			console.error(`INVALID FLASHCARD IN IMPORT: ${JSON.stringify(rawFlashcard, null, 2)}`);
			parseResult.error.errors.forEach((err) => {
				console.error(`Error in field "${err.path.join('.')}" - ${err.message}`);
			});
		}
	}
	return sourceFlashcards;
};
