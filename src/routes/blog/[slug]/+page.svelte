<script lang="ts">
	import { formatDate, calculateAge } from '$lib/utils';

	let { data } = $props();
</script>

<h2>{data.post.title}</h2>
<p class="secondary pb-6">{formatDate(data.post.created_at)} | {calculateAge(data.post.created_at).formatted} old</p>

<div class="prose max-w-none">
	{@html data.post.content}
</div>

<style>
	:global(.prose h1) {
		color: black;
		font-size: 1.8rem;
		font-weight: bold;
		margin-top: 1.5rem;
		margin-bottom: 0.5rem;
	}
	:global(.prose h2) {
		color: black;
		font-size: 1.5rem;
		font-weight: bold;
		margin-top: 1.2rem;
		margin-bottom: 0.4rem;
	}
	:global(.prose p) {
		color: black;
		margin-bottom: 1rem;
		line-height: 1.6;
	}
	:global(.prose code) {
		background-color: #000;
		color: white;
		padding: 0.2rem 0.4rem;
		font-family: monospace;
		font-weight: 400;
	}
	:global(.prose pre) {
		background-color: #000000;
		color: white;
		border-radius: 0 !important;
		padding: 1rem;
		overflow-x: auto;
		margin-bottom: 1rem;
	}
</style>
