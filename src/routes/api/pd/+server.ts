import { exec } from 'child_process';
import type { RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async () => {
	return new Promise((resolve, reject) => {
		exec('npm run pd', (error, stdout, stderr) => {
			if (error) {
				console.error(stderr);
				reject(new Response(stderr, { status: 500 }));
			} else {
				resolve(new Response(JSON.stringify({message: "data parsed"}), { status: 200 }));
			}
		});
	});
};
