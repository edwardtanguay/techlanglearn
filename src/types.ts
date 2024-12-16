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
	status: string;
	category: Category;
};

export type PageStatus = 'loading' | 'ready' | 'error';

export type Category = {
	idCode: string;
	topics: string[];
	description: string;
};
