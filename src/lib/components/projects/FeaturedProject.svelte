<script>
	let { link, banner, title, description } = $props();
</script>

<div class="flex flex-col items-center">
	<div class="flex w-38 items-center justify-center">
		<a href={link}
			><img src={banner} alt="{title} banner" class="max-h-75 w-auto max-w-140 p-6" /></a
		>
	</div>
	<div class="w-2/3 pb-6 text-center">
		<p>{description}</p>
		<a href={link} target="_blank" rel="noopener noreferrer" class="secondary hover:underline"
			>({link.slice(8)})</a
		>
	</div>
</div>
