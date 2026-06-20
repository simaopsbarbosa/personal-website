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

<div class="relative flex min-h-screen flex-col">
	<!-- Vertical borders running full height from top to bottom -->
	<div class="pointer-events-none absolute inset-y-0 left-1/2 z-0 w-full max-w-3xl -translate-x-1/2 border-x border-dashed"></div>

	<!-- Navigation bar with horizontal lines spanning the full screen width -->
	<nav class="relative z-10 mx-auto my-12 flex min-w-full justify-center border-y border-dashed">
		<NavButton title="Home" link={resolve('/')} isClicked={isActive('/')} />
		<NavButton title="Projects" link={resolve('/projects')} isClicked={isActive('/projects')} />
		<NavButton title="Blog" link={resolve('/blog')} isClicked={isActive('/blog')} />
		<NavButton title="CV" link={cv} isClicked={false} />
	</nav>

	<!-- Main page content restricted to max-w-3xl and centered -->
	<main class="relative z-10 mx-auto w-full max-w-3xl flex-1">
		{@render children()}
	</main>
</div>
