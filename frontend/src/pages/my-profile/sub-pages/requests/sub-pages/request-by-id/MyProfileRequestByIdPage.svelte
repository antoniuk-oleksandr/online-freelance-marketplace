<script lang="ts">
  import Tabs from '@/common-components/Tabs/Tabs.svelte'
  import type { TabType } from '@/types/TabType'
  import { useRouter } from 'svelte-routing'
  import {
    fetchOrderTabData,
    getInitialTabComponentsData,
    getOrderByIdTabData,
    handleOrderByIdTabChange,
    processOrderTabComponentsData,
  } from './helpers'
  import { onDestroy, onMount } from 'svelte'
  import type { MyProfileOrderByIdData } from '@/types/MyProfileOrderByIdData'
  import MyProfileOrderByIdPageLayout from '../../../orders/sub-pages/order-by-id/MyProfileOrderByIdPageLayout.svelte'
  import OrderByIdSidebar from '../../../orders/sub-pages/order-by-id/components/OrderByIdSidebar/OrderByIdSidebar.svelte'
  import { orderByIdSidebarStore } from '../../../orders/sub-pages/order-by-id/stores/order-by-id-sidebar-store'

  type MyProfileOrderByIdPageProps = {
    orderId: string
  }

  const { orderId }: MyProfileOrderByIdPageProps = $props()

  let tabComponentsData = $state<MyProfileOrderByIdData>(getInitialTabComponentsData())
  const setTabComponentsData = (value: MyProfileOrderByIdData) => (tabComponentsData = value)

  const unsubOrderStore = orderByIdSidebarStore.subscribe((value) => {
    tabComponentsData = processOrderTabComponentsData(tabComponentsData, value)
  })

  let tabIndex = $state(-1)
  const setTabIndex = (value: number) => (tabIndex = value)

  const unsubRouterBase = useRouter().routerBase.subscribe((route) => {
    fetchOrderTabData(route, orderId, setTabIndex, tabComponentsData, setTabComponentsData)
  })

  onMount(() => orderByIdSidebarStore.set(undefined))
  onDestroy(() => {
    unsubRouterBase()
    unsubOrderStore()
  })

  const tabsData: TabType[] = $derived(getOrderByIdTabData(tabComponentsData))
</script>

{#if tabIndex !== -1}
  <MyProfileOrderByIdPageLayout>
    <Tabs
      orderId={parseInt(orderId)}
      {tabComponentsData}
      setTabIndex={(tabIndex) => handleOrderByIdTabChange(orderId, tabIndex)}
      {tabIndex}
      tabs={tabsData}
    />
    <OrderByIdSidebar showOrderByIdSidebarSelect {orderId} orderData={tabComponentsData[0]} />
  </MyProfileOrderByIdPageLayout>
{/if}
