<script lang="ts">
	import ProjectLinks from './ProjectLinks.svelte';
	import type { ProjectLink } from './types';

	let {
		links,
		banner,
		title,
		description
	}: {
		links: [ProjectLink, ...ProjectLink[]];
		banner: string;
		title: string;
		description: string;
	} = $props();
</script>

<div class="flex flex-col items-center">
	<a href={links[0].href} target="_blank" rel="noopener noreferrer"
		><img
			src={banner}
			alt="{title} banner"
			class="max-h-60 w-auto max-w-[95vw] p-6 md:max-h-75 md:max-w-140"
		/></a
	>
	<div class="w-2/3 pb-6 text-center">
		<p>{description}</p>
		<ProjectLinks {links} centered />
	</div>
</div>
