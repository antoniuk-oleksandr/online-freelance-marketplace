<script lang="ts">
  import type { OrderByIdOverviewTabData } from '@/types/OrderByIdOverviewTabData'
  import OrderByIdOverviewTabReceiptBodyLayout from './OrderByIdOverviewTabReceiptBodyLayout.svelte'
  import { formatDeliveryTime } from '@/pages/my-profile/sub-pages/orders/helpers'

  type OrderByIdOverviewTabReceiptBodyProps = {
    order: OrderByIdOverviewTabData
  }

  const { order }: OrderByIdOverviewTabReceiptBodyProps = $props()
</script>

<OrderByIdOverviewTabReceiptBodyLayout>
  <span>{order.service.package.name}</span>
  <span class="truncate">{order.service.package.description}</span>
  <span>{formatDeliveryTime(order.service.package.deliveryTime)}</span>
</OrderByIdOverviewTabReceiptBodyLayout>
