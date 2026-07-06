<script>
	import { resolve } from '$app/paths';

	let { title, date, slug } = $props();
</script>

<div class="items-center justify-between py-2">
	<h3><a class="" href={resolve('/blog/[slug]', { slug })}>{title}</a></h3>

		<p class="secondary">{date}</p>
	
</div>
