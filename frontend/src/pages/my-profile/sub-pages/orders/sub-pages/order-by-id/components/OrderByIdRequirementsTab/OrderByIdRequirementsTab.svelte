<script lang="ts" module>
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import OrderByIdRequirementsTabQaList from './components/OrderByIdRequirementsTabQAList/OrderByIdRequirementsTabQAList.svelte'
  import type { MyProfileRequirementsRequestResponse } from '@/types/MyProfileRequirementsRequestResponse'

  export { orderByIdRequirementsTabComponent }
</script>

{#snippet orderByIdRequirementsTabComponent(response: MyProfileRequirementsRequestResponse)}
  <PaperElement
    styles="!ring-0 lg:dark:!ring-1 !p-0 lg:!p-6 !shadow-none !bg-transparent lg:!bg-light-palette-background-block lg:dark:!bg-dark-palette-background-block  lg:!shadow-md dark:!shadow-none gap-6 flex flex-col flex-grow text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
  >
    {#if response && response.status === 200}
      <OrderByIdRequirementsTabQaList serviceQAs={response.data.questionsAnswers} />
    {/if}
  </PaperElement>
{/snippet}
