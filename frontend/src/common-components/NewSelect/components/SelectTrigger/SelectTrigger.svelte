<script lang="ts">
  import type { DropdownItem } from '@/types/DropdownItem'
  import SelectTriggerLayout from './SelectTriggerLayout.svelte'
  import Spinner from '@/common-components/Spinner/Spinner.svelte'
  import SelectTriggerIcon from '../SelectTriggerIcon/SelectTriggerIcon.svelte'

  type SelectTriggerProps = {
    items: DropdownItem[]
    selectedIndex: number
    shownValue: boolean
    triggerStyles?: string
    selectWidth?: string
    triggerIcon?: string
  }

  let iconLoaded = $state(false)

  const { items, selectedIndex, triggerIcon, ...rest }: SelectTriggerProps = $props()
</script>

<SelectTriggerLayout {...rest}>
  {#if triggerIcon == undefined || iconLoaded}
    <span>{items[selectedIndex].title}</span>
  {:else}
    <Spinner size="size-4" />
  {/if}
  <SelectTriggerIcon {triggerIcon} bind:iconLoaded {...rest} />
</SelectTriggerLayout>
