<script>
	let { children } = $props();
</script>

<a href="/blog" class="secondary underline">(back)</a>

<article class="mt-2 min-h-svh bg-white p-6 shadow">{@render children()}</article>
