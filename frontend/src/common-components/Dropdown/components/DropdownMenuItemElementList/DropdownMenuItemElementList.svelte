<script lang="ts">
  import DropdownMenuItemElement from '../DropdownMenuItemElement/DropdownMenuItemElement.svelte'
  import type { DropdownItem } from '@/types/DropdownItem'

  type DropdownMenuItemElementListProps = {
    items: DropdownItem[]
    setShown: (value: boolean) => void
  }

  const { items, setShown }: DropdownMenuItemElementListProps = $props()
</script>

<div class="flex flex-col max-h-81.5  overflow-auto">
  {#each items as item}
    <DropdownMenuItemElement {setShown} {...item} />
  {/each}
</div>
