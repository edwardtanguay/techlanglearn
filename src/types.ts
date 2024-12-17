import { z } from 'zod';

export const SourceFlashcardSchema = z.object({
	language: z.string(),
	front: z.string(),
	back: z.string(),
	whenCreated: z.string(),
	isOpen: z.boolean()
});

export type SourceFlashcard = z.infer<typeof SourceFlashcardSchema>;

export type Flashcard = {
	language: string;
	front: string;
	back: string;
	whenCreated: string;
	isOpen: boolean;
};

export type TimeUnit = {
	calendarDate: string;
	duration: string;
};

export type Stats = {
	totalDays: number;
	averageDurationPerDay: string;
	totalDuration: string;
	timeUnits: TimeUnit[];
};

export type Tutorial = {
	topics: string;
	language: string;
	duration: string;
	year: number;
	rank: number;
	url: string;
	title: string;
	description: string;
	platform: string;
	fileIdCode: string;
	status: string;
	category: Category | null;
};

export type PageStatus = 'loading' | 'ready' | 'error';

export type Category = {
	idCode: string;
	name: string;
	topics: string[];
	description: string;
};
