<script lang="ts">
	import { MetaTags } from 'svelte-meta-tags';
	import { page } from '$app/state';
	import Post from '$lib/components/Post.svelte';
	import { formatDate } from '$lib/utils';

	let { data } = $props();
</script>

<MetaTags
	title="Blog | Simão Barbosa"
	description="Some thoughts I want to share - Simão Barbosa's blog."
	canonical={page.url.href}
	openGraph={{
		url: page.url.href,
		title: 'Blog | Simão Barbosa',
		description: "Some thoughts I want to share - Simão Barbosa's blog."
	}}
/>

<h2 class="centered">My Thoughts</h2>
<p class="centered">Some thoughts I want to share</p>

<div class="mx-10 mt-12">
	{#if data.posts && data.posts.length > 0}
		{#each data.posts as post, index}
			<Post title={post.title} date={formatDate(post.created_at)} slug={post.slug} />
			{#if index < data.posts.length - 1}
				<hr class="my-1 border-dotted" />
			{/if}
		{/each}
	{:else}
		<p class="centered secondary">no posts found</p>
	{/if}
</div>
