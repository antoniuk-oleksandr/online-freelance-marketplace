<script lang="ts">
  import type { OrderByIdOverviewTabData } from '@/types/OrderByIdOverviewTabData'
  import OrderByIdOverviewTabReceiptFooterLayout from './OrderByIdOverviewTabReceiptFooterLayout.svelte'
    import OrderByIdOverviewTabReceiptRow from '../OrderByIdOverviewTabReceiptRow/OrderByIdOverviewTabReceiptRow.svelte'

  type OrderByIdOverviewTabReceiptFooterProps = {
    order: OrderByIdOverviewTabData
  }

  const { order }: OrderByIdOverviewTabReceiptFooterProps = $props()
</script>

<OrderByIdOverviewTabReceiptFooterLayout>
  <OrderByIdOverviewTabReceiptRow
    firstText="subtotal"
    secondText="{order.subtotal}$"
  />
  <OrderByIdOverviewTabReceiptRow
    firstText="service fee"
    secondText="{order.serviceFee}$"
  />
  <OrderByIdOverviewTabReceiptRow
    firstText="total"
    secondText="{order.totalPrice}$"
  />
</OrderByIdOverviewTabReceiptFooterLayout>
