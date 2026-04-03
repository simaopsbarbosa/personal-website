<script>
  import { resolve } from '$app/paths';

  let { title, date, slug } = $props();
</script>

<div class="px-4 py-2 bg-white shadow flex justify-between items-center">
  <div>
    <p class="secondary">{date.toUpperCase()}</p>
    <h3><a href={resolve('/blog/[slug]', { slug })}>{title}</a></h3>
  </div>
  
</div>
