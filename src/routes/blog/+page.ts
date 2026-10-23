import type { PageLoad } from './$types';
import { supabase } from '$lib/supabase';

export const prerender = false;

export const load: PageLoad = async () => {
	try {
		const { data: posts, error } = await supabase
			.from('posts')
			// ⚡ Bolt Optimization: Only select needed fields to prevent over-fetching
			// Previously selected 'content' and 'updated_at' which aren't used in the list view
			.select('id, slug, title, created_at')
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
