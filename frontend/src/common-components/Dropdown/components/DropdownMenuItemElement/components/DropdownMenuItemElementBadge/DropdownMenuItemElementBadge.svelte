<script lang="ts">
  import type { DropdownItem } from '@/types/DropdownItem'

  const { badge }: DropdownItem = $props()
</script>

{#if badge}
  <span class="bg-light-p">{badge}</span>
{/if}
