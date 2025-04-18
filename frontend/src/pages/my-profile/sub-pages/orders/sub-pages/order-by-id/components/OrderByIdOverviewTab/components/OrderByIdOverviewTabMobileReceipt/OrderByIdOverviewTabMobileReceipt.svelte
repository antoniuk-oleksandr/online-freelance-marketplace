<script lang="ts">
  import type { OrderByIdOverviewTabData } from '@/types/OrderByIdOverviewTabData'
  import OrderByIdOverviewTabHeaderTopLine from '../OrderByIdOverviewTabHeaderTopLine/OrderByIdOverviewTabHeaderTopLine.svelte'
  import OrderByIdOverviewTabMobileReceiptLayout from './OrderByIdOverviewTabMobileReceiptLayout.svelte'
  import OrderByIdOverviewTabReceiptRow from '../OrderByIdOverviewTabReceiptRow/OrderByIdOverviewTabReceiptRow.svelte'
  import DividerElement from '@/common-components/DividerElement/DividerElement.svelte'
  import { formatDeliveryTime } from '@/pages/my-profile/sub-pages/orders/helpers'

  type OrderByIdOverviewTabMobileReceiptProps = {
    order: OrderByIdOverviewTabData
  }

  const { order }: OrderByIdOverviewTabMobileReceiptProps = $props()
</script>

<OrderByIdOverviewTabMobileReceiptLayout>
  <OrderByIdOverviewTabHeaderTopLine createdAt={order.createdAt} />
  <div class="mt-5"></div>
  <OrderByIdOverviewTabReceiptRow
    firstText="Package Name"
    secondText={order.service.package.name}
  />
  <OrderByIdOverviewTabReceiptRow
    firstText="Delivery time"
    secondText={formatDeliveryTime(order.service.package.deliveryTime)}
  />
  <OrderByIdOverviewTabReceiptRow
    firstText="Description"
    secondText={order.service.package.description}
    styles="flex-col !gap-0"
  />
  <DividerElement styles="my-4" />
  <OrderByIdOverviewTabReceiptRow firstText="subtotal" secondText="{order.subtotal}$" />
  <OrderByIdOverviewTabReceiptRow firstText="service fee" secondText="{order.serviceFee}$" />
  <OrderByIdOverviewTabReceiptRow firstText="total" secondText="{order.totalPrice}$" />
</OrderByIdOverviewTabMobileReceiptLayout>
