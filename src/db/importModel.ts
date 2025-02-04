/* eslint-disable @typescript-eslint/no-explicit-any */
import { SourceFlashcardSchema, type SourceFlashcard } from '../types';
import _rawFlashcards from './data/flashcards.json';

// raw = untrustworthy data
// source = validated, cleaned and structured data

const rawFlashcards = _rawFlashcards as any[];

export const getSourceFlashcards = (): [SourceFlashcard[], number] => {
	let numberOfErrors = 0;
	const sourceFlashcards: SourceFlashcard[] = [];
	for (const rawFlashcard of rawFlashcards) {
		const parseResult = SourceFlashcardSchema.safeParse(rawFlashcard);
		if (parseResult.success) {
			sourceFlashcards.push(rawFlashcard);
		} else {
			let r = '';
			r += `INVALID FLASHCARD IN IMPORT: ${JSON.stringify(rawFlashcard, null, 2)}\n`;
			parseResult.error.errors.forEach((err) => {
				r += `Error in field "${err.path.join('.')}" - ${err.message}\n`;
			});
			console.error(r);
			numberOfErrors++;
		}
	}
	return [sourceFlashcards, numberOfErrors];
};
