<script lang="ts">
  import { StatusEnum } from '@/types/StatusEnum'
  import OrderByIdDeliveryTabCompleteContentHeader from '../OrderByIdDeliveryTabCompleteContentHeader/OrderByIdDeliveryTabCompleteContentHeader.svelte'
  import OrderByIdDeliveryTabCompleteContentReviewCTA from '../OrderByIdDeliveryTabCompleteContentReviewCTA/OrderByIdDeliveryTabCompleteContentReviewCTA.svelte'
  import OrderByIdDeliveryTabCompleteContentFileSection from '../OrderByIdDeliveryTabCompleteContentFileSection/OrderByIdDeliveryTabCompleteContentFileSection.svelte'
  import OrderByIdDeliveryTabCompleteContentLayout from './OrderByIdDeliveryTabCompleteContentLayout.svelte'
  import OrderByIdDeliveryTabContentMessage from '../OrderByIdDeliveryTabContentMessage/OrderByIdDeliveryTabContentMessage.svelte'
  import type { DeliveryData } from '@/types/DeliveryData'

  type OrderByIdDeliveryTabInCompleteContentProps = {
    data: DeliveryData
    orderId: string
    renderReviewCTA?: boolean
  }

  const {
    data,
    orderId,
    renderReviewCTA = true,
  }: OrderByIdDeliveryTabInCompleteContentProps = $props()
</script>

{#if data.status === StatusEnum.Completed && data.delivery}
  <OrderByIdDeliveryTabCompleteContentLayout>
    <OrderByIdDeliveryTabCompleteContentHeader data={data.delivery} />
    <div
      class="flex-1 p-6 flex flex-col gap-6 overflow-auto text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
    >
      <OrderByIdDeliveryTabContentMessage
        message={data.delivery.message}
        title="Delivery Message"
      />
      <OrderByIdDeliveryTabCompleteContentFileSection {data} />
      {#if renderReviewCTA}
        <OrderByIdDeliveryTabCompleteContentReviewCTA {orderId} />
      {/if}
    </div>
  </OrderByIdDeliveryTabCompleteContentLayout>
{/if}
