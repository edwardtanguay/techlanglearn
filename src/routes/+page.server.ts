import { marked } from 'marked';
import { PUBLIC_BASE_URL } from '$env/static/public'; // Import your base URL
// import * as qstr from '../tools';

export async function load() {
	// KEEP: CODE THAT LOADS FROM FILE SYSTEM INSTEAD OF FETCHING FROM STATIC FILE
	// const filePath = path.join(process.cwd(), 'src', 'data', 'svelte5GuideForBeginners.md');
	// const markdown = fs.readFileSync(filePath, 'utf-8');

	const tutorialIdCodes = [
		// 'skeletonUiForSvelte',
		// 'masterGoIdiomsTestingSpanish',
		// 'svelte5GuideForBeginners',
		// 'cicdGitLab',
		// 'dominaGoAplicaciones',
		'svelte5IsATriumph'
	];

	let rawHtmlContent = '';
	for (const tutorialIdCode of tutorialIdCodes) {
		const fileUrl = `${PUBLIC_BASE_URL}/data/${tutorialIdCode}.md`;
		const res = await fetch(fileUrl);
		const markdown = await res.text();
		rawHtmlContent += marked(markdown);
	}

	return { htmlContent: rawHtmlContent };
}
