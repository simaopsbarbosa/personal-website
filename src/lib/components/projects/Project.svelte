<script>
	let { link, logo, title, description } = $props();
</script>

<div class="flex flex-col md:flex-row items-center">
	<div class="flex w-38 items-center justify-center my-6 md:my-0">
		<a href={link}><img src={logo} alt="{title} logo" class="max-h-10 w-auto max-w-20" /></a>
	</div>
	<div class="w-9/10 md:border-l md:border-dashed md:pl-4 text-center md:text-left">
		<p>{description}</p>
		<a href={link} target="_blank" rel="noopener noreferrer" class="secondary hover:underline"
			>({link.slice(8)})</a
		>
	</div>
</div>
