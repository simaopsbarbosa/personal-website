import type { PageLoad } from './$types';
import { supabase } from '$lib/supabase';

export const prerender = false;

export const load: PageLoad = async () => {
	try {
		const { data: posts, error } = await supabase
			.from('posts')
			.select('id, slug, title, content, created_at, updated_at')
			.eq('draft', false)
			.order('created_at', { ascending: false });

		if (error) {
			console.error('failed to fetch posts from Supabase:', error.message);
			return { posts: [] };
		}

		return { posts: posts || [] };
	} catch (e) {
		console.error('failed to fetch posts from Supabase:', e);
	}

	return { posts: [] };
};
