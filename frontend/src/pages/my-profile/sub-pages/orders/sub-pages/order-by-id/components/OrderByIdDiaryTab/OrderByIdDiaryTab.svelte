<script lang="ts" module>
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import OrderByIdDiaryTabTable from './components/OrderByIdDiaryTabTable/OrderByIdDiaryTabTable.svelte'
  import OrderByIdDiaryTabTopLine from './components/OrderByIdDiaryTabTopLine/OrderByIdDiaryTabTopLine.svelte'
  import type { MyProfileDiaryRequestResponse } from '@/types/MyProfileDiaryRequestResponse'

  export { orderByIdDiaryTabComponent }
</script>

{#snippet orderByIdDiaryTabComponent(response: MyProfileDiaryRequestResponse)}
  <PaperElement
    styles="!ring-0 lg:dark:!ring-1 !p-0 lg:!p-6 !shadow-none !bg-transparent lg:!bg-light-palette-background-block lg:dark:!bg-dark-palette-background-block  lg:!shadow-md dark:!shadow-none gap-6 flex flex-col flex-grow text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
  >
    {#if response && response.status === 200}
      <OrderByIdDiaryTabTopLine />
      <OrderByIdDiaryTabTable diaryItems={response.data.dates} />
    {/if}
  </PaperElement>
{/snippet}
