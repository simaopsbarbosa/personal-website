import type { PageLoad } from './$types';

export const prerender = false;

export const load: PageLoad = async ({ fetch }) => {
	const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';
	
	try {
		const res = await fetch(`${API_URL}/posts`);
		if (res.ok) {
			const posts = await res.json();
			return { posts };
		}
	} catch (e) {
		console.error('failed to fetch posts from backend:', e);
	}
	
	return { posts: [] };
};