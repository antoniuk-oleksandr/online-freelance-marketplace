<script lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import SearchTopBarLayout from '@/pages/search/components/SearchTopBar/SearchTopBarLayout.svelte'
  import { searchFilterDrawerStore } from '@/pages/search/stores/search-filter-drawer-store'
  import SearchInputForm from '@/pages/search/components/SearchInputForm/SearchInputForm.svelte'
  import SearchLeftSideTitle from '@/pages/search/components/SearchLeftSideTitle/SearchLeftSideTitle.svelte'
  import SearchSelect from '../SearchSelect/SearchSelect.svelte'
</script>

<SearchTopBarLayout>
  <SearchLeftSideTitle />
  <SearchSelect />
  <SearchInputForm />
  <Button
    clickAction={() => searchFilterDrawerStore.set(true)}
    styles="lg:!hidden w-fit row-start-1 col-start-2"
    >Filters
  </Button>
</SearchTopBarLayout>
