export type Flashcard = {
	front: string;
	back: string;
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
};

export type PageStatus = 'loading' | 'ready' | 'error';
