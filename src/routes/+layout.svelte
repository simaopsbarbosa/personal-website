<script lang="ts">
	import './layout.css';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import NavButton from '$lib/components/NavButton.svelte';
	import { preloadData } from '$app/navigation';
	import { onMount } from 'svelte';
	import { MetaTags } from 'svelte-meta-tags';

	let { children } = $props();

	type NavPath = '/' | '/blog' | '/projects';
	const pathname = $derived(page.url.pathname);
	const siteName = 'Simão Barbosa';
	const canonical = $derived(`https://simaob.com${pathname}`);
	const title = $derived(page.data.pageMetaTags?.title ?? siteName);
	const description = $derived(
		page.data.pageMetaTags?.description ||
			'Personal website of Simão Barbosa, a Software Engineer studying at FEUP.'
	);

	function isActive(path: NavPath): boolean {
		if (path === '/') return (pathname as string) === '/';
		return (pathname as string).startsWith(`${path}/`);
	}

	onMount(() => {
		preloadData('/blog').catch((err) => {
			console.error('Error preloading blog data:', err);
		});
	});

	const isException = $derived((pathname as string).startsWith('/admin/editor'));
</script>

<MetaTags
	{title}
	{description}
	{canonical}
	openGraph={{ type: 'website', url: canonical, title, description, siteName }}
	twitter={{ site: '@simaopsbarbosa', cardType: 'summary_large_image' }}
/>

<svelte:head>
	<link rel="icon" type="image/png" sizes="96x96" href="/favicon.png" />
</svelte:head>

{#if isException}
	<main class="relative z-10 w-full flex-1">
		{@render children()}
	</main>
{:else}
	<div class="relative flex min-h-screen flex-col">
		<!-- vertical borders -->
		<div
			class="pointer-events-none absolute inset-y-0 left-1/2 z-0 w-full max-w-3xl -translate-x-1/2 border-dotted md:border-x"
		></div>

		<!-- nav bar -->
		<nav
			class="relative z-10 mx-auto mt-6 mb-12 flex min-w-full justify-center border-y border-dotted md:mt-12"
		>
			<NavButton title="Home" link={resolve('/')} isClicked={isActive('/')} />
			<NavButton title="Projects" link={resolve('/projects')} isClicked={isActive('/projects')} />
			<NavButton title="Blog" link={resolve('/blog')} isClicked={isActive('/blog')} />
			<NavButton title="CV" link="/Simao_Barbosa_CV.pdf" isClicked={false} />
		</nav>

		<!-- content -->
		<main class="relative z-10 mx-auto w-full max-w-3xl flex-1">
			{@render children()}
		</main>
	</div>
{/if}
