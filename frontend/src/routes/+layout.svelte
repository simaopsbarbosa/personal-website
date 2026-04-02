
<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	
	let { children } = $props();

	type NavPath = '/' | '/blog' | '/projects';
	const pathname = $derived(page.url.pathname);

	function isActive(path: NavPath): boolean {
		if (path === '/') return pathname === '/';
		return pathname.startsWith(`${path}/`);
	}
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<div class="max-w-5xl mx-auto min-h-screen">	
	<nav class="flex gap-4 px-4 py-12 mx-auto justify-around max-w-md">
		<a href={resolve('/')} class:underline={isActive('/')}>Home</a>
		<a href={resolve('/blog')} class:underline={isActive('/blog')}>Blog</a>
		<a href={resolve('/projects')} class:underline={isActive('/projects')}>Projects</a>
	</nav>
	
	{@render children()}
</div>
