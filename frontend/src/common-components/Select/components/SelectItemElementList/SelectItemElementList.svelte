<script lang="ts">
  import SelectItemElement from '@/common-components/Select/components/SelectItemElement/SelectItemElement.svelte'
  import type { SelectItem } from '@/types/SelectItem'

  type SelectItemElementListProps = {
    items: SelectItem[]
    setIsOpen?: (value: boolean) => void
    selectedItem: string
    setSelectedItem?: (item: string) => void
  }

  const props: SelectItemElementListProps = $props()
</script>

<div class="flex flex-col gap-y-1 Select-menu py-1 lg:py-0">
  {#each props.items as item}
    <SelectItemElement {...props} {...item} />
  {/each}
</div>
