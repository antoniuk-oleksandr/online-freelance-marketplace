<script lang="ts">
  import { calcPriceWithServiceFees, getServiceFees, round } from '@/utils/utils'
  import OrderRequestSummaryItem from '@/pages/order-request/components/OrderRequestSummary/components/OrderRequestSummaryItem/OrderRequestSummaryItem.svelte'
  import OrderOverviewBottomBlockLayout from '@/pages/order-confirm-pay/components/OrderOverview/components/OrderOverviewBottomBlock/OrderOverviewBottomBlockLayout.svelte'
  import type { Package } from '@/types/Package.ts'

  type OrderOverviewBottomBlockProps = {
    pkg: Package
  }

  const serviceFees = getServiceFees()

  const { pkg }: OrderOverviewBottomBlockProps = $props()
</script>

<OrderOverviewBottomBlockLayout>
  <OrderRequestSummaryItem title="service fee" value="{round(serviceFees * pkg.price, 2)}$" />
  <OrderRequestSummaryItem
    styles="font-semibold"
    title="total"
    value="{calcPriceWithServiceFees(pkg.price)}$"
  />
  <OrderRequestSummaryItem
    title="delivery time"
    value="{pkg.deliveryDays} {pkg.deliveryDays > 1 ? 'days' : 'day'}"
  />
</OrderOverviewBottomBlockLayout>
