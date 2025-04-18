<script lang="ts">
  const { children } = $props()
</script>

<div class="relative duration-200 focus-within:ring-1 ring-cyan-500 ease-out">
  {@render children()}
</div>
