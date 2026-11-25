<script lang="ts">
	import { formatDate, calculateAge } from '$lib/utils';

	let { data } = $props();
</script>

<h1>{data.post.title}</h1>
<p class="secondary pb-6">
	{formatDate(data.post.created_at)}
</p>

<div class="prose max-w-none">
	{@html data.post.content}
</div>

<div class="secondary centered pt-4">
	At the time of writing, I was {calculateAge(data.post.created_at).formatted} old
</div>
