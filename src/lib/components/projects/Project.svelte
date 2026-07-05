<script>
	let { link, logo, title, description } = $props();
</script>

<div class="flex">
	<div class="flex w-38 items-center justify-center">
		<a href={link}><img src={logo} alt="{title} logo" class="max-h-10 w-auto max-w-20" /></a>
	</div>
	<div class="w-9/10 border-l border-dashed pl-4">
		<p>{description}</p>
		<a href={link} target="_blank" rel="noopener noreferrer" class="secondary hover:underline"
			>({link.slice(8)})</a
		>
	</div>
</div>
