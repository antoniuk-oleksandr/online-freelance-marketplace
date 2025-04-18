<script lang="ts">
  import Title from '@/common-components/Title/Title.svelte'
  import OrderByIdSidebarLayout from './OrderByIdSidebarLayout.svelte'
  import type { MyProfileOverviewRequestResponse } from '@/types/MyProfileOverviewRequestResponse'
  import DividerElement from '@/common-components/DividerElement/DividerElement.svelte'
  import OrderRequestSummaryItem from '@/pages/order-request/components/OrderRequestSummary/components/OrderRequestSummaryItem/OrderRequestSummaryItem.svelte'
  import { formatOrderDate } from '@/pages/order-submit-requirements/helpers'
  import OrderByIdSidebarTop from './components/OrderByIdSidebarTop/OrderByIdSidebarTop.svelte'

  type OrderByIdSidebarProps = {
    orderData: MyProfileOverviewRequestResponse | undefined
  }

  const { orderData }: OrderByIdSidebarProps = $props()
</script>

<OrderByIdSidebarLayout>
  {#if orderData && orderData.status === 200}
    <Title text="Order details" />
    <OrderByIdSidebarTop {...orderData} />
    <DividerElement styles="my-3" />
    <OrderRequestSummaryItem
      valueStyles="!font-semibold text-cyan-500 !normal-case hover:underline"
      title="ordered from"
      valueLink="/users/{orderData.data.freelancer.id}"
      value={orderData.data.freelancer.username}
    />
    <OrderRequestSummaryItem
      valueStyles="!font-semibold"
      title="delivery date"
      value={formatOrderDate(orderData.data.deliveryDate)}
    />
    <OrderRequestSummaryItem
      valueStyles="!font-semibold"
      title="total price"
      value="{orderData.data.totalPrice}$"
    />
    <OrderRequestSummaryItem
      valueStyles="!font-semibold"
      title="order number"
      value={orderData.data.id}
    />
  {/if}
</OrderByIdSidebarLayout>
