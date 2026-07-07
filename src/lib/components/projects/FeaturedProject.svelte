<script>
	let { link, banner, title, description } = $props();
</script>

<div class="flex flex-col items-center">
	<a href={link}
		><img
			src={banner}
			alt="{title} banner"
			class="max-h-60 w-auto max-w-[95vw] p-6 md:max-h-75 md:max-w-140"
		/></a
	>
	<div class="w-2/3 pb-6 text-center">
		<p>{description}</p>
		<a
			href={link}
			target="_blank"
			rel="noopener noreferrer"
			class="secondary dotted-underline break-all">({link.slice(8)})</a
		>
	</div>
</div>
