
<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import { resolve } from '$app/paths';
	
	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<div class="max-w-5xl mx-auto min-h-screen">	
	<nav class="flex gap-4 px-4 py-12 mx-auto justify-around max-w-md">
		<a href={resolve('/')}>Home</a>
		<a href={resolve('/blog')}>Blog</a>
		<a href={resolve('/projects')}>Projects</a>
	</nav>
	
	{@render children()}
</div>
