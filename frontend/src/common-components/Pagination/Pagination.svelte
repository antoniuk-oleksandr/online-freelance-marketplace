<script lang="ts">
  import PaginationButton from './components/PaginationButton/PaginationButton.svelte'
  import PaginationLayout from './components/PaginationLayout.svelte'
  import PaginationMovePageButton from './components/PaginationMovePageButton/PaginationMovePageButton.svelte'
  import { getPaginationRange } from './helpers'

  type PaginationProps = {
    currentPage: number
    totalPages: number
    setPage: (page: number) => void
  }

  let { currentPage, totalPages, setPage }: PaginationProps = $props()
</script>

<PaginationLayout>
  <PaginationMovePageButton
    clickAction={() => setPage(--currentPage)}
    type="prev"
    disabled={currentPage === 1}
  />
  {#each getPaginationRange(currentPage, totalPages) as page}
    {#if page === '...'}
      <PaginationButton disabled={true} currentPage={false}>...</PaginationButton>
    {:else}
      <PaginationButton
        clickAction={() => setPage(page as number)}
        disabled={false}
        currentPage={currentPage === page}
      >
        {page}
      </PaginationButton>
    {/if}
  {/each}
  <PaginationMovePageButton
    clickAction={() => setPage(++currentPage)}
    type="next"
    disabled={currentPage === totalPages}
  />
</PaginationLayout>
