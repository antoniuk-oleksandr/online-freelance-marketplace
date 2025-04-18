<script lang="ts">
  import type { DropdownItem } from '@/types/DropdownItem'
  import SelectTriggerLayout from './SelectTriggerLayout.svelte'
  import Spinner from '@/common-components/Spinner/Spinner.svelte'
  import SelectTriggerIcon from '../SelectTriggerIcon/SelectTriggerIcon.svelte'

  type SelectTriggerProps = {
    items: DropdownItem[]
    selectedIndex: number
    shownValue: boolean
    noneSelectedText?: string
    triggerStyles?: string
    selectWidth?: string
    triggerIcon?: string
    error?: string | null
  }

  let iconLoaded = $state(false)

  const { items, selectedIndex, triggerIcon, noneSelectedText, ...rest }: SelectTriggerProps = $props()
</script>

<SelectTriggerLayout {...rest}>
  {#if selectedIndex === -1}
    <span
      class="font-normal text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
      >{noneSelectedText ?? 'Select'}</span
    >
  {:else if triggerIcon == undefined || iconLoaded}
    <span>{items[selectedIndex].title}</span>
  {:else}
    <Spinner size="size-4" />
  {/if}
  <SelectTriggerIcon {triggerIcon} bind:iconLoaded {...rest} />
</SelectTriggerLayout>
