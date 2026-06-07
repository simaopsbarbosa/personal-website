
<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.png';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import cv from '$lib/assets/cv.pdf';
	import NavButton from '$lib/components/NavButton.svelte';
	
	let { children } = $props();

	type NavPath = '/' | '/blog' | '/projects';
	const pathname = $derived(page.url.pathname);

	function isActive(path: NavPath): boolean {
		if (path === '/') return pathname === '/';
		return pathname.startsWith(`${path}/`);
	}
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<title>Simão Barbosa</title>
</svelte:head>

<nav class="flex my-12 mx-auto justify-center max-w-md min-w-full border border-dotted">
	<NavButton title="Home" link={resolve('/')} isClicked={isActive('/')} />
	<NavButton title="Projects" link={resolve('/projects')} isClicked={isActive('/projects')} />
	<NavButton title="Blog" link={resolve('/blog')} isClicked={isActive('/blog')} />
	<NavButton title="CV" link={cv} isClicked={false} />
</nav>

<div class="max-w-4xl mx-auto min-h-screen">
	{@render children()}
</div>
