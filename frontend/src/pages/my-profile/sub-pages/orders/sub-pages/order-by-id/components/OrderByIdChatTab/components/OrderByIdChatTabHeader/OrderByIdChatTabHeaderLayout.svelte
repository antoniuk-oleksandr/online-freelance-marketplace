<script lang="ts">
  const { children } = $props()
</script>

<div
  class="h-22 bg-light-palette-action-hover items-center flex gap-3 py-3 rounded-t-xl px-6 dark:bg-dark-palette-action-hover border-b border-light-palette-divider dark:border-dark-palette-divider"
>
  {@render children()}
</div>
