<script lang="ts">
  import OrderByIdDeliveryTabCanceledContentHeaderLayout from './OrderByIdDeliveryTabCanceledContentHeaderLayout.svelte'
  import { formatOrderDate } from '@/pages/order-submit-requirements/helpers'
  import Icon from '@iconify/svelte'
  import type { DeliveryCancellation } from '@/types/DeliveryCancellation'

  type OrderByIdDeliveryTabCanceledContentHeaderProps = {
    data: DeliveryCancellation
  }

  const { data }: OrderByIdDeliveryTabCanceledContentHeaderProps = $props()
</script>

<OrderByIdDeliveryTabCanceledContentHeaderLayout>
  <div class="flex items-center justify-between w-full">
    <div class="flex items-center gap-3 text-white">
      <Icon icon="heroicons:x-circle" width="28" height="28" />
      <h2 class="text-xl font-bold">Order Canceled</h2>
    </div>
    <div class="flex items-center gap-1 text-red-100 text-base">
      <Icon icon="hugeicons:calendar-03" width="24" height="24" />
      <span>{formatOrderDate(data.cancelledAt)}</span>
    </div>
  </div>
</OrderByIdDeliveryTabCanceledContentHeaderLayout>
