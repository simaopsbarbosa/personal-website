<script>
    let { title, link, isClicked } = $props();
</script> 

<div class="px-4 py-1 min-w-20 text-center outline outline-black" class:bg-black={isClicked} class:text-white={isClicked}>
    <a href={link} class=" hover:underline">{title}</a>
</div>