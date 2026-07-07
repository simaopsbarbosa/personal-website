<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon/favicon.ico';
	import favicon16 from '$lib/assets/favicon/favicon-16x16.png';
	import favicon32 from '$lib/assets/favicon/favicon-32x32.png';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import NavButton from '$lib/components/NavButton.svelte';

	let { children } = $props();

	type NavPath = '/' | '/blog' | '/projects';
	const pathname = $derived(page.url.pathname);

	function isActive(path: NavPath): boolean {
		if (path === '/') return (pathname as string) === '/';
		return (pathname as string).startsWith(`${path}/`);
	}

	const isException = $derived((pathname as string).startsWith('/admin/editor'));
</script>

<svelte:head>
	<link rel="icon" href={favicon} sizes="any" />
	<link rel="icon" type="image/png" sizes="32x32" href={favicon32} />
	<link rel="icon" type="image/png" sizes="16x16" href={favicon16} />
	<title>Simão Barbosa</title>
	<meta
		name="description"
		content="Personal website of Simão Barbosa, a Software Engineer studying at FEUP."
	/>
	<meta property="og:title" content="Simão Barbosa" />
	<meta
		property="og:description"
		content="Personal website of Simão Barbosa, a Software Engineer studying at FEUP."
	/>
	<meta property="og:type" content="website" />
</svelte:head>

{#if isException}
	<main class="relative z-10 w-full flex-1">
		{@render children()}
	</main>
{:else}
	<div class="relative flex min-h-screen flex-col">
		<!-- vertical borders -->
		<div
			class="pointer-events-none absolute inset-y-0 left-1/2 z-0 w-full max-w-3xl -translate-x-1/2 border-dashed md:border-x"
		></div>

		<!-- nav bar -->
		<nav
			class="relative z-10 mx-auto mt-6 mb-12 flex min-w-full justify-center border-y border-dashed md:mt-12"
		>
			<NavButton title="Home" link={resolve('/')} isClicked={isActive('/')} />
			<NavButton title="Projects" link={resolve('/projects')} isClicked={isActive('/projects')} />
			<NavButton title="Blog" link={resolve('/blog')} isClicked={isActive('/blog')} />
			<NavButton title="CV" link="/cv_simao_barbosa.pdf" isClicked={false} />
		</nav>

		<!-- content -->
		<main class="relative z-10 mx-auto w-full max-w-3xl flex-1">
			{@render children()}
		</main>
	</div>
{/if}
