<script lang="ts">
	import ProjectLinks from './ProjectLinks.svelte';
	import type { ProjectLink } from './types';

	let {
		links,
		logo,
		title,
		description
	}: {
		links: [ProjectLink, ...ProjectLink[]];
		logo: string;
		title: string;
		description: string;
	} = $props();
</script>

<div class="flex flex-col items-center md:flex-row">
	<div class="my-6 flex w-38 items-center justify-center md:my-0">
		<a href={links[0].href} target="_blank" rel="noopener noreferrer"
			><img src={logo} alt="{title} logo" class="max-h-10 w-auto max-w-20" /></a
		>
	</div>
	<div class="w-[60vw] text-center md:w-9/10 md:border-l md:border-dashed md:pl-4 md:text-left">
		<p>{description}</p>
		<ProjectLinks {links} />
	</div>
</div>
