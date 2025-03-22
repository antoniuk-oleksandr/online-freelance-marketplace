<script lang="ts">
  import SearchPageLayout from '@/pages/search/SearchPageLayout.svelte'
  import SearchLeftSide from '@/pages/search/components/SearchLeftSide/SearchLeftSide.svelte'
  import SearchRightSide from '@/pages/search/components/SearchRightSide/SearchRightSide.svelte'
  import { useRouter } from 'svelte-routing'
  import type { GetFilterParamsRequestResponse } from '@/types/GetFilterParamsRequestResponse.ts'
  import { searchFilterDrawerStore } from '@/pages/search/stores/search-filter-drawer-store'
  import { searchCursorStore } from '@/pages/search/stores/search-cursor-store'
  import type { SearchCursorStore } from '@/types/SearchCursorStore.ts'
  import { handleSearchRefresh } from '@/pages/search/handlers'
  import { request } from '@/api/request'
  import { errorStore } from '@/common-stores/error-store'
  import { onDestroy } from 'svelte'

  let searchCursorData = $state<SearchCursorStore | undefined>()
  searchCursorStore.subscribe((value) => (searchCursorData = value))

  const unsubscribe = useRouter().routerBase.subscribe((value) => {
    if (value.uri !== '/search' || !searchCursorData) return
    handleSearchRefresh()
  })

  onDestroy(() => unsubscribe())

  let defaultFilterParamsRequestResponse = $state<GetFilterParamsRequestResponse | undefined>()

  request<GetFilterParamsRequestResponse>('GET', '/filter-params/all').then((response) => {
    if (response.status !== 200) {
      errorStore.set({ shown: true, error: response.data.error })
    }
    
    defaultFilterParamsRequestResponse = response
  })

  let isFiltersModalOpen = $state(false)
  searchFilterDrawerStore.subscribe((value) => (isFiltersModalOpen = value))
</script>

{#if defaultFilterParamsRequestResponse && defaultFilterParamsRequestResponse.status === 200}
  <SearchPageLayout>
    <SearchLeftSide />
    <SearchRightSide
      defaultFilterParams={defaultFilterParamsRequestResponse.data}
      {isFiltersModalOpen}
    />
  </SearchPageLayout>
{/if}
