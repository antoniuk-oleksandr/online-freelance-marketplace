<script lang="ts">
  import type { SearchPageParams } from '@/types/SearchPageParams'
  import { searchStore } from '@/pages/search/stores/search-store'
  import { onDestroy } from 'svelte'
  import NewSelect from '@/common-components/NewSelect/NewSelect.svelte'
  import {
    makeSearchSelectTriggerIcon,
    makeInitialSearchSelectIndexes,
    getSearchSelectItems,
    handleSearchSelectChange,
  } from './helpers'

  let searchPageParams = $state<SearchPageParams | undefined>()
  const unsubscribe = searchStore.subscribe((value) => (searchPageParams = value))

  onDestroy(() => unsubscribe())
</script>

<NewSelect
  triggerIcon={makeSearchSelectTriggerIcon(searchPageParams)}
  triggerStyles="!justify-center flex-row-reverse h-10.5"
  selectWidth="min-w-48"
  multipleInitialSelectIndexes={makeInitialSearchSelectIndexes(searchPageParams)}
  modalHeaderTitle="Sort"
  initialSelectValues={getSearchSelectItems()}
  onSelectChange={(value) => handleSearchSelectChange(value, searchPageParams)}
/>
