import type { PageLoad } from './$types';
import { error } from '@sveltejs/kit';

export const prerender = false;

export const load: PageLoad = async ({ fetch, params }) => {
	const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';
	
	try {
		const res = await fetch(`${API_URL}/posts/${params.slug}`);
		if (res.ok) {
			const post = await res.json();
			return { post };
		}
	} catch (e) {
		console.error('failed to fetch post from backend:', e);
	}
	
	throw error(404, 'post not found');
};
