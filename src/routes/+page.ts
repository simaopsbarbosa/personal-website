import { definePageMetaTags } from 'svelte-meta-tags';
import type { PageLoad } from './$types';

export const load: PageLoad = () =>
	definePageMetaTags({
		title: 'Simão Barbosa | Software Engineer in Porto',
		description:
			'I’m Simão Barbosa, a Software Engineer studying at FEUP in Porto. Explore my projects, including uni and Inclui, and read my blog.'
	});
