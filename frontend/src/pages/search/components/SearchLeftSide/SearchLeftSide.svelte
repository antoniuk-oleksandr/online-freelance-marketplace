<script lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import SearchTopBar from '@/pages/search/components/SearchTopBar/SearchTopBar.svelte'
  import SearchServices from '@/pages/search/components/SearchServices/SearchServices.svelte'
  import SearchLeftSideLayout from '@/pages/search/components/SearchLeftSide/SearchLeftSideLayout.svelte'
</script>

<SearchLeftSideLayout>
  <PaperElement
    styles="min-h-app !ring-0 lg:dark:!ring-1 text-base !bg-transparent dark:!bg-transparent !shadow-none !p-0 flex flex-col gap-6 md:!p-6 md:!shadow-md dark:!shadow-none md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block"
  >
    <SearchTopBar />
    <SearchServices />
  </PaperElement>
</SearchLeftSideLayout>
