import { marked } from 'marked';
import { PUBLIC_BASE_URL } from '$env/static/public'; // Import your base URL

console.log(111, PUBLIC_BASE_URL);

export async function load() {
	// KEEP: CODE THAT LOADS FROM FILE SYSTEM INSTEAD OF FETCHING FROM STATIC FILE
	// const filePath = path.join(process.cwd(), 'src', 'data', 'svelte5GuideForBeginners.md');
	// const markdown = fs.readFileSync(filePath, 'utf-8');

	const tutorialIdCodes = ['svelte5GuideForBeginners', 'cicdGitLab'];

	let content = '';
	for (const tutorialIdCode of tutorialIdCodes) {
		const fileUrl = `${PUBLIC_BASE_URL}/data/${tutorialIdCode}.md`;
		const res = await fetch(fileUrl);
		const markdown = await res.text();
		content += marked(markdown);
	}

	return { content };
}
