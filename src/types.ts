import { z } from 'zod';

export const SourceFlashcardSchema = z.object({
	language: z.enum(['it', 'es']),
	front: z.string().min(1, { message: 'cannot be empty' }),
	back: z.string().min(1, { message: 'cannot be empty' }),
	whenCreated: z.string().refine(
		(value) => /^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$/.test(value), // Regex to match the format
		{
			message: 'Invalid date/time format. Expected YYYY-MM-DD HH:mm:ss'
		}
	),
	extras: z.string()
});

export type SourceFlashcard = z.infer<typeof SourceFlashcardSchema>;

export type Flashcard = {
	language: string;
	front: string;
	back: string;
	whenCreated: string;
	extras: string;
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
