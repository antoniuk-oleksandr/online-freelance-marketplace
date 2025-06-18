<script lang="ts" module>
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import type { MyProfileReviewRequestResponse } from '@/types/MyProfileReviewRequestResponse'
  import { ResponseErrorEnum } from '@/types/ResponseErrorEnum'
  import RequestByIdReviewTabSuccessContent from './components/RequestByIdReviewTabSuccessContent.svelte'

  export { requestByIdReviewTabComponent }
</script>

{#snippet requestByIdReviewTabComponent(
  response: MyProfileReviewRequestResponse | undefined,
  orderId: string,
)}
  <PaperElement
    styles="!p-6 flex -mx-6 md:m-0 items-center justify-center !ring-0 lg:dark:!ring-1 !p-0 !shadow-none !bg-transparent lg:!bg-light-palette-background-block lg:dark:!bg-dark-palette-background-block  lg:!shadow-md dark:!shadow-none gap-3 flex flex-col flex-grow text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
  >
    {#if response}
      {#if response.status === 404 && response.data.error === ResponseErrorEnum.OrderReviewNotFound}
        <div>No review yet</div>
      {:else if response.status === 200}
        <RequestByIdReviewTabSuccessContent data={response.data} />
      {/if}
    {/if}
  </PaperElement>
{/snippet}
