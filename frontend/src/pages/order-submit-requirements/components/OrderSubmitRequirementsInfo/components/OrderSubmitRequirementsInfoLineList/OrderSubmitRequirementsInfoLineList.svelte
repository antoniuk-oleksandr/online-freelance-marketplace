<script lang="ts">
  import OrderSubmitRequirementsInfoLineListLayout from '@/pages/order-submit-requirements/components/OrderSubmitRequirementsInfo/components/OrderSubmitRequirementsInfoLineList/OrderSubmitRequirementsInfoLineListLayout.svelte'
  import type { OrderSubmitRequirementsData } from '@/types/OrderSubmitRequirementsData.ts'
  import OrderSubmitRequirementsInfoLine from '@/pages/order-submit-requirements/components/OrderSubmitRequirementsInfo/components/OrderSubmitRequirementsInfoLine/OrderSubmitRequirementsInfoLine.svelte'
  import { calcPriceWithServiceFees } from '@/utils/utils'
  import OrderSubmitRequirementsInfoStatusLine from '@/pages/order-submit-requirements/components/OrderSubmitRequirementsInfo/components/OrderSubmitRequirementsInfoStatusLine/OrderSubmitRequirementsInfoStatusLine.svelte'
  import { formatOrderDate } from '@/pages/order-submit-requirements/helpers'

  type OrderSubmitRequirementsInfoLineListProps = {
    orderData: OrderSubmitRequirementsData
  }

  const { orderData }: OrderSubmitRequirementsInfoLineListProps = $props()
</script>

<OrderSubmitRequirementsInfoLineListLayout>
  <OrderSubmitRequirementsInfoStatusLine {orderData} />
  <OrderSubmitRequirementsInfoLine key="Order" value={'#' + orderData.order.id} />
  <OrderSubmitRequirementsInfoLine
    key="Order Date"
    value={formatOrderDate(orderData.order.createdAt)}
  />
  <OrderSubmitRequirementsInfoLine
    key="Price"
    value={calcPriceWithServiceFees(orderData.service.package.price) + '$'}
  />
</OrderSubmitRequirementsInfoLineListLayout>
