<script>
	let { title, link, isClicked } = $props();
</script>

<div
	class="min-w-20 px-4 py-1 text-center outline outline-black"
	class:bg-black={isClicked}
	class:text-[#fff]={isClicked}
>
	<a href={link} class=" hover:underline">{title}</a>
</div>
