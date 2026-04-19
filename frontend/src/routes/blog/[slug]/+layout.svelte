<script>
    let { children } = $props();
</script>

<a href="/blog" class="underline secondary">&larr; Back to Blog</a>

<article class="bg-white shadow min-h-svh p-6 mt-2">{@render children()}</article>