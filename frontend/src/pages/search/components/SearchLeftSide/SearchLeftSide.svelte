<script lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import SearchTopBar from '@/pages/search/components/SearchTopBar/SearchTopBar.svelte'
  import SearchServices from '@/pages/search/components/SearchServices/SearchServices.svelte'
  import SearchLeftSideLayout from '@/pages/search/components/SearchLeftSide/SearchLeftSideLayout.svelte'
</script>

<SearchLeftSideLayout>
  <PaperElement styles="p-3 flex flex-col w-full gap-y-3">
    <SearchTopBar />
    <SearchServices />
  </PaperElement>
</SearchLeftSideLayout>
