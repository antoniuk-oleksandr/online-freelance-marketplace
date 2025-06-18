<script lang="ts">
  import Tabs from '@/common-components/Tabs/Tabs.svelte'
  import type { TabType } from '@/types/TabType'
  import MyProfileOrderByIdPageLayout from './MyProfileOrderByIdPageLayout.svelte'
  import { useRouter } from 'svelte-routing'
  import {
    fetchOrderTabData,
    getInitialTabComponentsData,
    getOrderByIdTabData,
    handleOrderByIdTabChange,
  } from './helpers'
  import { onDestroy, setContext } from 'svelte'
  import OrderByIdSidebar from './components/OrderByIdSidebar/OrderByIdSidebar.svelte'
  import type { MyProfileOrderByIdData } from '@/types/MyProfileOrderByIdData'
  import { writable } from 'svelte/store'

  type MyProfileOrderByIdPageProps = {
    orderId: string
  }

  const { orderId }: MyProfileOrderByIdPageProps = $props()

  let tabComponentsData = $state<MyProfileOrderByIdData>(getInitialTabComponentsData())
  const setTabComponentsData = (value: MyProfileOrderByIdData) => (tabComponentsData = value)

  let tabIndex = writable<number>(-1)
  setContext('tabIndex', tabIndex)

  const unsubscribe = useRouter().routerBase.subscribe((route) => {
    fetchOrderTabData(route, orderId, tabIndex, tabComponentsData, setTabComponentsData)
  })

  onDestroy(() => unsubscribe())

  const tabsData: TabType[] = $derived(getOrderByIdTabData(tabComponentsData))
</script>

{#if $tabIndex !== -1}
  <MyProfileOrderByIdPageLayout>
    <Tabs
      elementStyles="!px-1 xs:!px-2.5"
      orderId={parseInt(orderId)}
      {tabComponentsData}
      setTabIndex={(tabIndex) => handleOrderByIdTabChange(orderId, tabIndex)}
      tabIndex={$tabIndex}
      tabs={tabsData}
    />
    <OrderByIdSidebar {orderId} orderData={tabComponentsData[0]} />
  </MyProfileOrderByIdPageLayout>
{/if}
