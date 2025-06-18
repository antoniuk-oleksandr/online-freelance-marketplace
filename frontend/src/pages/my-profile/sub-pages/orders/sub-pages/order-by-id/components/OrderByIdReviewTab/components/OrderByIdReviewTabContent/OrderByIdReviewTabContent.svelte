<script lang="ts">
  import type { MyProfileReviewRequestResponse } from '@/types/MyProfileReviewRequestResponse'
  import { orderReviewStore } from '../../../../stores/order-review-store'
  import RequestByIdReviewTabSuccessContent from '@/pages/my-profile/sub-pages/requests/sub-pages/request-by-id/components/RequestByIdReviewTab/components/RequestByIdReviewTabSuccessContent.svelte'
  import { ResponseErrorEnum } from '@/types/ResponseErrorEnum'
  import { StatusEnum } from '@/types/StatusEnum'
  import OrderByIdAwaitingAcceptanceReviewTab from '../OrderByIdAwaitingAcceptanceReviewTab/OrderByIdAwaitingAcceptanceReviewTab.svelte'
  import OrderByIdReviewTabForm from '../OrderByIdReviewTabForm/OrderByIdReviewTabForm.svelte'

  type OrderByIdReviewTabContentProps = {
    response: MyProfileReviewRequestResponse
    orderId: string
  }

  const { response, orderId }: OrderByIdReviewTabContentProps = $props()
</script>

{#if $orderReviewStore !== undefined}
  <RequestByIdReviewTabSuccessContent data={$orderReviewStore} />
{:else if response.status === 404 && response.data.error === ResponseErrorEnum.OrderReviewNotFound && response.data.status === StatusEnum.AwaitingAcceptance}
  <OrderByIdAwaitingAcceptanceReviewTab />
{:else if response.status === 404 && response.data.error === ResponseErrorEnum.OrderReviewNotFound}
  <OrderByIdReviewTabForm {orderId} />
{:else if response.status === 200}
  <RequestByIdReviewTabSuccessContent data={response.data} />
{/if}
