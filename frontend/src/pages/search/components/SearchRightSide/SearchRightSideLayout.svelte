<script lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import type { LayoutProps } from '@/types/LayoutProps.ts'
  import SearchLeftSideBackdrop from '@/pages/search/components/SearchLeftSideBackdrop/SearchLeftSideBackdrop.svelte'
  import { useIsClosedAfterAnimation } from '@/pages/search/helpers'
  import { searchFilterArrayStore } from '@/pages/search/stores/search-filter-array-store'

  type SearchRightSideLayoutProps = LayoutProps & {
    isFiltersModalOpen: boolean
  }

  let isMobile = $state(window.innerWidth < 1024)
  const setIsMobile = (value: boolean) => (isMobile = value)
  useIsClosedAfterAnimation(1024, setIsMobile)

  let isClosedAfterAnimation = $state(false)

  let isSearchFilterArrOpened = $state(false)
  searchFilterArrayStore.subscribe((value) => (isSearchFilterArrOpened = value.isOpened))

  const { isFiltersModalOpen, children }: SearchRightSideLayoutProps = $props()

  $effect(() => {
    let timeout: number

    if (isFiltersModalOpen) isClosedAfterAnimation = true
    else timeout = setTimeout(() => (isClosedAfterAnimation = false), 300)

    return () => clearTimeout(timeout)
  })

  $effect(() => {
    if (isMobile && isFiltersModalOpen) {
      document.body.style.overflow = 'hidden'
      document.body.style.height = '100svh'
    } else {
      document.body.style.overflow = 'auto'
      document.body.style.height = 'auto'
    }
  })
</script>

{#if isClosedAfterAnimation || !isMobile}
  <div
    class="fixed lg:sticky lg:top-22 lg:max-h-search-sidebar {!isMobile
      ? ''
      : isFiltersModalOpen
        ? 'animate-[fade-in_300ms]'
        : 'animate-[fade-out_300ms]'} z-30 lg:z-10 left-0 top-0 h-svh lg:h-full w-full text-base"
  >
    <SearchLeftSideBackdrop />
    <PaperElement
      styles="absolute overflow-hidden !p-0 flex {!isMobile
        ? ''
        : isFiltersModalOpen
          ? 'animate-drawer-open'
          : 'animate-drawer-close'}  shadow-lg lg:dark:shadow-none right-0 top-0 w-full sm:w-112 !rounded-none h-svh lg:h-full ring-0 w-full lg:dark:ring-1 lg:!rounded-xl lg:static lg:w-full"
    >
      <div
        class="flex h-full duration-200 ease-out {isSearchFilterArrOpened
          ? '-translate-x-[100svw] sm:-translate-x-112 lg:-translate-x-search-sidebar'
          : ''}"
      >
        {@render children()}
      </div>
    </PaperElement>
  </div>
{/if}
