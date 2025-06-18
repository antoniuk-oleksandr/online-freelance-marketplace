<script lang="ts">
  import type { DeliveryData } from '@/types/DeliveryData'
  import { StatusEnum } from '@/types/StatusEnum'
  import OrderByIdDeliveryTabCanceledContentHeader from '../OrderByIdDeliveryTabCanceledContentHeader/OrderByIdDeliveryTabCanceledContentHeader.svelte'
  import OrderByIdDeliveryTabContentMessage from '../OrderByIdDeliveryTabContentMessage/OrderByIdDeliveryTabContentMessage.svelte'
  import OrderByIdDeliveryTabCanceledContentRefund from '../OrderByIdDeliveryTabCanceledContentRefund/OrderByIdDeliveryTabCanceledContentRefund.svelte'
  import OrderByIdDeliveryTabCanceledContentNextSteps from '../OrderByIdDeliveryTabCanceledContentNextSteps/OrderByIdDeliveryTabCanceledContentNextSteps.svelte'
  import OrderByIdDeliveryTabCanceledContentLayout from './OrderByIdDeliveryTabCanceledContentLayout.svelte'

  type OrderByIdDeliveryTabCanceledContentProps = {
    data: DeliveryData
    showRefundAndNextSteps?: boolean
  }

  const { data, showRefundAndNextSteps = true }: OrderByIdDeliveryTabCanceledContentProps = $props()
</script>

{#if data.status === StatusEnum.Cancelled && data.cancellation && data.payment}
  <OrderByIdDeliveryTabCanceledContentLayout>
    <OrderByIdDeliveryTabCanceledContentHeader data={data.cancellation} />
    <div class="flex-1 p-6 flex flex-col gap-6 overflow-auto">
      <OrderByIdDeliveryTabContentMessage
        title="Cancellation Reason"
        message={data.cancellation.cancellationReason}
      />
      {#if showRefundAndNextSteps}
        <OrderByIdDeliveryTabCanceledContentRefund data={data.payment} />
        <OrderByIdDeliveryTabCanceledContentNextSteps />
      {/if}
    </div>
  </OrderByIdDeliveryTabCanceledContentLayout>
{/if}
