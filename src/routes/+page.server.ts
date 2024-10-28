// src/routes/+page.js
import fs from 'fs';
import path from 'path';
import { marked } from 'marked';

export async function load() {
	const filePath = path.join(process.cwd(), 'src', 'data', 'svelte5GuideForBeginners.md');
	const markdown = fs.readFileSync(filePath, 'utf-8');
	const content = marked(markdown);

	return { content }

	// return {
	// 	props: { content }
	// };
}
