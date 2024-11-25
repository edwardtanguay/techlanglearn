import { marked } from 'marked';
import { PUBLIC_BASE_URL } from '$env/static/public'; // Import your base URL
// import * as qstr from '../tools';

export async function load() {
	// KEEP: CODE THAT LOADS FROM FILE SYSTEM INSTEAD OF FETCHING FROM STATIC FILE
	// const filePath = path.join(process.cwd(), 'src', 'data', 'svelte5GuideForBeginners.md');
	// const markdown = fs.readFileSync(filePath, 'utf-8');

	const tutorialIdCodes = [
		// 'inkInFifteenMinutes',
		// 'skeletonUiForSvelte',
		// 'masterGoIdiomsTestingSpanish',
		// 'svelte5GuideForBeginners',
		// 'cicdGitLab',
		// 'dominaGoAplicaciones',
		// 'svelte5IsATriumph'
	] as string[];

	let rawHtmlContent = '';
	let rawMenuHtmlContent = '';
	let rawHtmlBodyContent = '';

	rawMenuHtmlContent += '<div id="menu">&nbsp;</div><div class="bg-slate-300 p-3 mt-3 flex flex-wrap gap-3">';
	rawMenuHtmlContent += tutorialIdCodes
		.map((m) => `<a href="#${m}" class="underline">${m}</a>`)
		.join('');
	rawMenuHtmlContent += '</div>';

	rawHtmlContent += rawMenuHtmlContent;

	let count = 1;
	for (const tutorialIdCode of tutorialIdCodes) {
		const fileUrl = `${PUBLIC_BASE_URL}/data/${tutorialIdCode}.md`;
		const res = await fetch(fileUrl);
		const markdown = await res.text();
		const topLink = count === 1 ? `<div class="mt-3">^ <a href="#vocab" class="underline">vocab</a></div>` : `<div class="mt-3">^ <a href="#menu" class="underline">menu</a></div>`;
		rawHtmlBodyContent += `<div id="${tutorialIdCode}">${topLink}</div>` + marked(markdown);
		count++;
	}

	rawHtmlContent += rawHtmlBodyContent;

	return { htmlContent: rawHtmlContent };
}
