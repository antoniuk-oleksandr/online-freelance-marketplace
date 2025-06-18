<script lang="ts">
  import { formatOrderDate } from '@/pages/order-submit-requirements/helpers'
  import Icon from '@iconify/svelte'
  import OrderByIdDeliveryTabCompleteContentHeaderLayout from './OrderByIdDeliveryTabCompleteContentHeaderLayout.svelte'
  import type { DeliveryDataInfo } from '@/types/DeliveryDataInfo'

  type OrderByIdDeliveryTabCompleteContentHeaderProps = {
    data: DeliveryDataInfo
  }

  const { data }: OrderByIdDeliveryTabCompleteContentHeaderProps = $props()
</script>

<OrderByIdDeliveryTabCompleteContentHeaderLayout>
  <div class="flex items-center justify-between w-full">
    <div class="flex items-center gap-1 text-white">
      <Icon icon="hugeicons:checkmark-circle-01" width="28" height="28" />
      <h2 class="text-xl font-bold">Work Delivered</h2>
    </div>
    <div class="flex items-center gap-1 text-green-100 text-base">
      <Icon icon="hugeicons:calendar-03" width="24" height="24" />
      <span>{formatOrderDate(data.date)}</span>
    </div>
  </div>
</OrderByIdDeliveryTabCompleteContentHeaderLayout>
