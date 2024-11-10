import type { PageStatus } from './types';

const siteLocation = import.meta.env.VITE_SITE_LOCATION === 'dev' ? 'dev' : 'online';
let pageStatus: PageStatus = $state('ready');
let errorMessage = $state('');

export const getStore = () => {
	return {
		// values
		get siteLocation() {
			return siteLocation;
		},
		get pageStatus() {
			return pageStatus;
		},
		get errorMessage() {
			return errorMessage;
		},
		// actions
		setPageStatus: (_pageStatus: PageStatus) => {
			pageStatus = _pageStatus;
		},
		setErrorMessage: (_errorMessage: string) => {
			errorMessage = _errorMessage;
		}
	};
};
