<script lang="ts">
	import type { ProjectLink } from './types';

	let { links, centered = false }: { links: ProjectLink[]; centered?: boolean } = $props();
</script>

<div
	class="mt-1 flex flex-wrap justify-center gap-x-4 gap-y-1 {centered ? '' : 'md:justify-start'}"
>
	{#each links as link (link.href)}
		<a href={link.href} target="_blank" rel="noopener noreferrer" class="secondary dotted-underline"
			>{link.label}</a
		>
	{/each}
</div>
