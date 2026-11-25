import { definePageMetaTags } from 'svelte-meta-tags';
import type { PageLoad } from './$types';

export const load: PageLoad = () =>
	definePageMetaTags({
		title: 'Projects | Simão Barbosa',
		description:
			'Explore software projects by Simão Barbosa, including uni for University of Porto students and Inclui for finding accessible places.'
	});
