<script lang="ts">
  import Tabs from '@/common-components/Tabs/Tabs.svelte'
  import type { TabType } from '@/types/TabType'
  import { useRouter } from 'svelte-routing'
  import {
    fetchOrderTabData,
    getInitialTabComponentsData,
    getOrderByIdTabData,
    handleOrderByIdTabChange,
  } from './helpers'
  import { onDestroy } from 'svelte'
  import type { MyProfileOrderByIdData } from '@/types/MyProfileOrderByIdData'
  import MyProfileOrderByIdPageLayout from '../../../orders/sub-pages/order-by-id/MyProfileOrderByIdPageLayout.svelte'
  import OrderByIdSidebar from '../../../orders/sub-pages/order-by-id/components/OrderByIdSidebar/OrderByIdSidebar.svelte'

  type MyProfileOrderByIdPageProps = {
    orderId: string
  }

  const { orderId }: MyProfileOrderByIdPageProps = $props()

  let tabComponentsData = $state<MyProfileOrderByIdData>(getInitialTabComponentsData())
  const setTabComponentsData = (value: MyProfileOrderByIdData) => (tabComponentsData = value)

  let tabIndex = $state(-1)
  const setTabIndex = (value: number) => (tabIndex = value)

  const unsubscribe = useRouter().routerBase.subscribe((route) => {
    fetchOrderTabData(route, orderId, setTabIndex, tabComponentsData, setTabComponentsData)
  })

  onDestroy(() => unsubscribe())

  const tabsData: TabType[] = $derived(getOrderByIdTabData(tabComponentsData))
</script>

{#if tabIndex !== -1}
  <MyProfileOrderByIdPageLayout>
    <Tabs
      orderId={parseInt(orderId)}
      {tabComponentsData}
      tabChangeAction={(tabIndex) => handleOrderByIdTabChange(orderId, tabIndex)}
      initialActiveTabIndex={tabIndex}
      tabs={tabsData}
    />
    <OrderByIdSidebar orderData={tabComponentsData[0]} />
  </MyProfileOrderByIdPageLayout>
{/if}
