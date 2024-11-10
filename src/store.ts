const siteLocation = 'dev3';

export const getStore = () => {
	return {
		get siteLocation() {
			return siteLocation;
		}
	};
};
