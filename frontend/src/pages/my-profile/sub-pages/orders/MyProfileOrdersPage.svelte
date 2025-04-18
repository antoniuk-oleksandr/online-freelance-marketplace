<script lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import Title from '@/common-components/Title/Title.svelte'
  import OrdersTable from './components/OrdersTable/OrdersTable.svelte'
  import Pagination from '@/common-components/Pagination/Pagination.svelte'
  import { onDestroy } from 'svelte'
  import { navigate, useRouter } from 'svelte-routing'
  import Loader from '@/common-components/Loader/Loader.svelte'
  import { fetchMyOrdersData } from './helpers'
  import type { MyProfileOrdersData } from '@/types/MyProfileOrdersData'
  import NoContentMessage from '@/common-components/NoContentMessage/NoContentMessage.svelte'
  import { getPageParam } from '@/utils/utils'

  let currentPage = $state<undefined | number>()
  let ordersData = $state<undefined | MyProfileOrdersData>()
  const setOrdersData = (data: MyProfileOrdersData) => (ordersData = data)

  const unsubscribe = useRouter().routerBase.subscribe((value) => {
    if (value.uri !== '/my-profile/orders') return

    currentPage = getPageParam()
    fetchMyOrdersData(currentPage, setOrdersData)
  })
  onDestroy(() => unsubscribe())
</script>

<PaperElement
  styles="min-h-app !ring-0 lg:dark:!ring-1 text-base !bg-transparent dark:!bg-transparent !shadow-none !p-0 flex flex-col gap-6 md:!p-6 md:!shadow-md dark:!shadow-none md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block"
>
  {#if currentPage && ordersData}
    <Title text="My orders" />
    {#if ordersData.orders.length === 0}
      <NoContentMessage text="You have no orders yet" />
    {:else}
      <OrdersTable orders={ordersData.orders} />
    {/if}

    {#if ordersData.totalPages > 1}
      <Pagination
        setPage={(value) => navigate(`/my-profile/orders?page=${value}`)}
        {currentPage}
        totalPages={ordersData.totalPages}
      />
    {/if}
  {:else}
    <Loader />
  {/if}
</PaperElement>
