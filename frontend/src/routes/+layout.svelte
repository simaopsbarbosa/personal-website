<script>
	import { resolve } from '$app/paths';
	import './layout.css';

	let { children } = $props();
</script>

<nav class="m-4 flex items-center justify-between gap-4 border border-neutral-700 px-8 py-4">
	<a class="text-xl font-bold" href={resolve('/')}>Simao Barbosa</a>
	<div class="flex gap-4">
		<a href={resolve('/about')}>about</a>
		<a href={resolve('/projects')}>projects</a>
		<a href={resolve('/blog')}>blog</a>
		<a href={resolve('/contact')}>contact</a>
	</div>
</nav>

<main class="mx-auto max-w-4xl">
	{@render children()}
</main>
