<script>
  let { title, date, description, link } = $props();
</script>

<div class="px-4 py-2">
  <div class="flex items-center gap-4 my-2">
    <h3>{title}</h3>
    <p class="secondary">@ {date.toUpperCase()}</p>
  </div>
  <p class="mb-2">{description}</p>
  <a href={link} target="_blank" rel="noopener noreferrer" class="underline secondary block mb-1">GitHub</a>
</div>
