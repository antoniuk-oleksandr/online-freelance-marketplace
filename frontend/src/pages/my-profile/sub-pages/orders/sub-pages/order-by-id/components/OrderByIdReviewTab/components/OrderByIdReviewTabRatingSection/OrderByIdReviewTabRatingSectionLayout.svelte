<script lang="ts">
  const { children } = $props()
</script>

<div class="w-full flex gap-3 flex-col items-center">
  {@render children()}
</div>
