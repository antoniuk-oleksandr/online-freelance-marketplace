<script lang="ts">
  import type { SelectMenuProps } from '@/types/SelectMenuProps'
  import DividerElement from '@/common-components/DividerElement/DividerElement.svelte'
  import SelectItemElementList from '@/common-components/Select/components/SelectItemElementList/SelectItemElementList.svelte'
  import { handleSelectBackdropClick } from '@/common-components/Select/handlers'

  const props: SelectMenuProps = $props()

  const handleClick = (e: any) => handleSelectBackdropClick(e, props.setIsOpen)

  $effect(() => {
    document.addEventListener('click', handleClick)

    return () => document.removeEventListener('click', handleClick)
  })
</script>

<SelectItemElementList
  items={props.items}
  setIsOpen={props.setIsOpen}
  selectedItem={props.selectedItem}
  setSelectedItem={props.setSelectedItem}
/>
{#if props.additionalItems && props.selectedAdditionalItem}
  <DividerElement styles="Select-menu" />
  <SelectItemElementList
    selectedItem={props.selectedAdditionalItem}
    setSelectedItem={props.setSelectedAdditionalItem}
    items={props.additionalItems}
    setIsOpen={props.setIsOpen}
  />
{/if}
