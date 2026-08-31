<script lang="ts">
	import { formatDate, calculateAge } from '$lib/utils';
	import { MetaTags } from 'svelte-meta-tags';
	import { page } from '$app/state';

	let { data } = $props();
</script>

<MetaTags
	title={`${data.post.title} | Simão Barbosa`}
	description={data.post.content.substring(0, 160).replace(/<[^>]+>/g, '')}
	canonical={page.url.href}
	openGraph={{
		url: page.url.href,
		title: `${data.post.title} | Simão Barbosa`,
		description: data.post.content.substring(0, 160).replace(/<[^>]+>/g, '')
	}}
/>

<h2>{data.post.title}</h2>
<p class="secondary pb-6">
	{formatDate(data.post.created_at)}
</p>

<div class="prose max-w-none">
	{@html data.post.content}
</div>

<div class="secondary centered pt-4">
	At the time of writing, I was {calculateAge(data.post.created_at).formatted} old
</div>
