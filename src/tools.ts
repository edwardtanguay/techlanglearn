/**
 * REPLACE ALL OCCURANCES IN A STRING:
 *
 * qstr.replaceAll("This is a tost.", "o", "e");
 *
 * "This is a test."
 */
export const replaceAll = (text: string, search: string, replace: string) => {
	return text.split(search).join(replace);
};

/**
 * Get a certain number of random items from an array
 *
 * nnn
 */
export const getRandomItemsFromArray = <T>(arr:T[], amount: number) : T[] => {
	const shuffled = [...arr].sort(() => .5 - Math.random());
	return shuffled.slice(0, amount);
};