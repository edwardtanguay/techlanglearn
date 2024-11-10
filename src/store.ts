const siteLocation = import.meta.env.VITE_SITE_LOCATION === 'dev' ? 'dev' : 'online';

export const getStore = () => {
	return {
		get siteLocation() {
			return siteLocation;
		}
	};
};
