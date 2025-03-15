<script lang="ts">
  import Select from '@/common-components/Select/Select.svelte'
  import { getSelectedSearchSelectItem, makeASearchSelectItemList } from '@/pages/search/helpers'
  import SearchSelectTrigger from '@/pages/search/components/SearchSelectTrigger/SearchSelectTrigger.svelte'
  import type { SearchPageParams } from '@/types/SearchPageParams'
  import { searchStore } from '@/pages/search/stores/search-store'

  let searchPageParams = $state<SearchPageParams | undefined>()
  searchStore.subscribe((value) => (searchPageParams = value))

  const items = $derived(makeASearchSelectItemList(searchPageParams))
  const additionalItems = $derived(makeASearchSelectItemList(searchPageParams, true))
</script>

<Select
  title="Sorting"
  selectedItem={getSelectedSearchSelectItem(searchPageParams)}
  {items}
  {additionalItems}
  selectedAdditionalItem={getSelectedSearchSelectItem(searchPageParams, true)}
>
  <SearchSelectTrigger
    selectedItem={getSelectedSearchSelectItem(searchPageParams)}
    selectedAdditionalItem={getSelectedSearchSelectItem(searchPageParams, true)}
  />
</Select>
