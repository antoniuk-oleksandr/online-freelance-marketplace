<script lang="ts">
  import { formatOrderDate } from '@/pages/order-submit-requirements/helpers'
  import Icon from '@iconify/svelte'
  import OrderByIdOverviewTabReceiptHeaderLayout from './OrderByIdOverviewTabReceiptHeaderLayout.svelte'
  import type { OrderByIdOverviewTabData } from '@/types/OrderByIdOverviewTabData'
  import OrderByIdOverviewTabHeaderTopLine from '../OrderByIdOverviewTabHeaderTopLine/OrderByIdOverviewTabHeaderTopLine.svelte'

  type OrderByIdOverviewTabReceiptHeaderProps = {
    order: OrderByIdOverviewTabData
  }

  const { order }: OrderByIdOverviewTabReceiptHeaderProps = $props()
</script>

<OrderByIdOverviewTabReceiptHeaderLayout>
  <OrderByIdOverviewTabHeaderTopLine createdAt={order.createdAt} />
  <div class="uppercase grid font-semibold gap-6 grid-cols-order-overview-receipt">
    <span>Package name</span>
    <span>Description</span>
    <span>Duration</span>
  </div>
</OrderByIdOverviewTabReceiptHeaderLayout>
