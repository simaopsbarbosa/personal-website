import type { PageLoad } from './$types';
import { error } from '@sveltejs/kit';
import { supabase } from '$lib/supabase';

export const prerender = false;

export const load: PageLoad = async ({ params }) => {
	try {
		const { data: post, error: dbError } = await supabase
			.from('posts')
			.select('id, slug, title, content, created_at, updated_at')
			.eq('slug', params.slug)
			.single();

		if (dbError || !post) {
			console.error('failed to fetch post from Supabase:', dbError?.message);
			throw error(404, 'post not found');
		}

		return { post };
	} catch (e) {
		console.error('failed to fetch post from Supabase:', e);
		throw error(404, 'post not found');
	}
};

