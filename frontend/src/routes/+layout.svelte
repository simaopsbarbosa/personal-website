<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.ico';
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

<nav class="mx-auto my-12 flex max-w-md min-w-full justify-center border-y border-dotted">
	<NavButton title="Home" link={resolve('/')} isClicked={isActive('/')} />
	<NavButton title="Projects" link={resolve('/projects')} isClicked={isActive('/projects')} />
	<NavButton title="Blog" link={resolve('/blog')} isClicked={isActive('/blog')} />
	<NavButton title="CV" link={cv} isClicked={false} />
</nav>

<div class="mx-auto min-h-screen max-w-3xl">
	{@render children()}
</div>
