<script lang="ts">
	import Post from '$lib/components/Post.svelte';
	import { formatDate } from '$lib/utils';

	let { data } = $props();
</script>

<h2 class="centered">My Thoughts</h2>
<p class="centered">Some thoughts I want to share</p>

<div class="mx-10 mt-12">
	{#if data.posts && data.posts.length > 0}
		{#each data.posts as post, index}
			<Post title={post.title} date={formatDate(post.created_at)} slug={post.slug}/>
			{#if index < data.posts.length - 1}
				<hr class="my-1 border-dashed" />
			{/if}
		{/each}
	{:else}
		<p class="centered secondary">no posts found</p>
	{/if}
</div>
