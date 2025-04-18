<script lang="ts">
  const { children } = $props()
</script>

<div class="grid grid-cols-7">
  {@render children()}
</div>
