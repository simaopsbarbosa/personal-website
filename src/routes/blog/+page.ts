import type { PageLoad } from './$types';
import { supabase } from '$lib/supabase';
import { definePageMetaTags } from 'svelte-meta-tags';

export const prerender = false;

const pageMeta = definePageMetaTags({
	title: 'Blog | Simão Barbosa',
	description: "Some thoughts I want to share - Simão Barbosa's blog."
});

export const load: PageLoad = async () => {
	try {
		const { data: posts, error } = await supabase
			.from('posts')
			.select('id, slug, title, content, created_at, updated_at')
			.eq('draft', false)
			.order('created_at', { ascending: false });

		if (error) {
			console.error('failed to fetch posts from Supabase:', error.message);
			return { posts: [], ...pageMeta };
		}

		return { posts: posts || [], ...pageMeta };
	} catch (e) {
		console.error('failed to fetch posts from Supabase:', e);
	}

	return { posts: [], ...pageMeta };
};
