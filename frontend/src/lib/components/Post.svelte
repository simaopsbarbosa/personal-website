<script>
	import { resolve } from '$app/paths';

	let { title, date, slug, tags = [] } = $props();
</script>

<div class="items-center justify-between py-2">
	<h3><a class="" href={resolve('/blog/[slug]', { slug })}>{title}</a></h3>
	<div class="my-0.5 flex gap-2">
		<p class="secondary">{date}</p>
		<div class="flex gap-1">
			{#each Object.values(tags) as tag}
				<span class="secondary bg-black px-1 pt-0.5 text-xs font-medium text-white!"
					>{tag.toLowerCase()}</span
				>
			{/each}
		</div>
	</div>
</div>
