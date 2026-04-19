<script>
  import { resolve } from '$app/paths';

  let { title, date, slug, tags = [] } = $props();
</script>

<div class="py-2 justify-between items-center">
    <h3><a class="" href={resolve('/blog/[slug]', { slug })}>{title}</a></h3>
    <div class="flex gap-2 my-0.5">
      <p class="secondary">{date}</p>
      <div class="flex gap-1">
        {#each Object.values(tags) as tag}
          <span class="bg-black secondary text-white! text-xs font-medium px-1 pt-0.5 ">{tag.toLowerCase()}</span>
        {/each}
      </div>
    </div>
</div>
