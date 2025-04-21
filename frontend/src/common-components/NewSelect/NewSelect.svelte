<script lang="ts">
  import type { NewSelectItem } from '@/types/NewSelectItem'
  import Dropdown from '../Dropdown/Dropdown.svelte'
  import SelectTrigger from './components/SelectTrigger/SelectTrigger.svelte'
  import { makeSelectValues, performOnSelectChange, setInitialSelectIndex } from './helpers'

  type SelectProps = {
    modalHeaderTitle: string
    initialSelectValues: NewSelectItem[]
    onSelectChange: (value: string | string[]) => void
    noneSelectedText?: string
    error?: string | null
    triggerIcon?: string
    triggerStyles?: string
    initialSelectIndex?: number
    multipleInitialSelectIndexes?: number[]
    selectWidth?: string
    layoutStyles?: string
    positionX?: 'left' | 'right'
    positionY?: 'top' | 'bottom'
    menuStyles?: string
  }

  let shownValue = $state(false)

  let {
    initialSelectIndex,
    multipleInitialSelectIndexes,
    initialSelectValues,
    selectWidth,
    onSelectChange,
    ...rest
  }: SelectProps = $props()

  const initialIndexes = setInitialSelectIndex(initialSelectIndex, multipleInitialSelectIndexes)
  let selectedIndexes = $state<number[]>(initialIndexes)
  const setSelectedIndexes = (value: number[]) => {
    selectedIndexes = value
    performOnSelectChange(items, selectedIndexes, onSelectChange)
  }

  let items = $derived(makeSelectValues(initialSelectValues, selectedIndexes, setSelectedIndexes))
</script>

<Dropdown bind:shownValue menuWidth={selectWidth} {items} {...rest}>
  <SelectTrigger {shownValue} {selectWidth} {items} selectedIndex={selectedIndexes[0]} {...rest} />
</Dropdown>
