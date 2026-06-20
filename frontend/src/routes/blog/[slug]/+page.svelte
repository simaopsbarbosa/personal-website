<script lang="ts">
	import { formatDate } from '$lib/utils';

	let { data } = $props();
</script>

<h2>{data.post.title}</h2>
<p class="secondary pb-6">{formatDate(data.post.created_at)}</p>

<div class="prose max-w-none">
	{@html data.post.content}
</div>

<style>
	:global(.prose h1) {
		font-size: 1.8rem;
		font-weight: bold;
		margin-top: 1.5rem;
		margin-bottom: 0.5rem;
	}
	:global(.prose h2) {
		font-size: 1.5rem;
		font-weight: bold;
		margin-top: 1.2rem;
		margin-bottom: 0.4rem;
	}
	:global(.prose p) {
		margin-bottom: 1rem;
		line-height: 1.6;
	}
	:global(.prose code) {
		background-color: #f3f4f6;
		padding: 0.2rem 0.4rem;
		border-radius: 0.25rem;
		font-family: monospace;
	}
	:global(.prose pre) {
		background-color: #1f2937;
		color: #f9fafb;
		padding: 1rem;
		border-radius: 0.375rem;
		overflow-x: auto;
		margin-bottom: 1rem;
	}
</style>
