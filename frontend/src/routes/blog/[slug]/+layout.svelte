<script>
	let { children } = $props();
</script>

<div class="p-6">
	<a href="/blog" class="secondary">(back)</a>
	
	<article class="mt-2">{@render children()}</article>
</div>
