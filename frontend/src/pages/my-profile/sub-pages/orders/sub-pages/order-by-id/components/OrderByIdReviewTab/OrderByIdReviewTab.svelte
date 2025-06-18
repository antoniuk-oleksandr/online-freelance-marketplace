<script lang="ts" module>
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import type { MyProfileReviewRequestResponse } from '@/types/MyProfileReviewRequestResponse'
  import OrderByIdReviewTabForm from './components/OrderByIdReviewTabForm/OrderByIdReviewTabForm.svelte'
  import { ResponseErrorEnum } from '@/types/ResponseErrorEnum'
  import RequestByIdReviewTabSuccessContent from '@/pages/my-profile/sub-pages/requests/sub-pages/request-by-id/components/RequestByIdReviewTab/components/RequestByIdReviewTabSuccessContent.svelte'
  import OrderByIdReviewTabContent from './components/OrderByIdReviewTabContent/OrderByIdReviewTabContent.svelte'

  export { orderByIdReviewTabComponent }
</script>

{#snippet orderByIdReviewTabComponent(
  response: MyProfileReviewRequestResponse | undefined,
  orderId: string,
)}
  <PaperElement
    styles=" !text-base !ring-0 lg:dark:!ring-1 !p-0 lg:!p-6 !shadow-none !bg-transparent lg:!bg-light-palette-background-block lg:dark:!bg-dark-palette-background-block  lg:!shadow-md dark:!shadow-none gap-6 flex flex-col flex-grow text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
  >
    {#if response}
      <OrderByIdReviewTabContent {response} {orderId} />
    {/if}
  </PaperElement>
{/snippet}
