<script>
    import Post from '$lib/components/Post.svelte';

    const posts = [
        { title: "How I built this website and why I am proud of it", date: "1 jan 2026", slug: "how-i-built-this-website", tags: ['web', 'development'] },
        { title: "Why the Womier SK75 is a perfect keyboard", date: "2 jan 2026", slug: "why-the-womier-sk75-is-a-perfect-keyboard", tags: ['hardware', 'keyboard'] },
        { title: "What the **** is Sioyek?", date: "3 jan 2026", slug: "what-the-****-is-sioyek", tags: ['technology', 'productivity'] },
    ];
</script>

<h2 class="centered">My Thoughts</h2>
<p class="centered">Some thoughts and lessons I want to share</p>

<div class="mt-12 mx-10">
    {#each posts as post, index}
        <Post {...post} />
        {#if index < posts.length - 1}
            <hr class="my-1 border-dotted" />
        {/if}
    {/each}
</div>