<script lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import Title from '@/common-components/Title/Title.svelte'
  import ServicesTable from './components/ServicesTable.svelte'
  import type { MyProfileServicesData } from '@/types/MyProfileServicesData'
  import { onDestroy } from 'svelte'
  import { navigate, useRouter } from 'svelte-routing'
  import { getPageParam } from '@/utils/utils'
  import { fetchMyServicesData } from './helpers'
  import Loader from '@/common-components/Loader/Loader.svelte'
  import NoContentMessage from '@/common-components/NoContentMessage/NoContentMessage.svelte'
  import Pagination from '@/common-components/Pagination/Pagination.svelte'

  let currentPage = $state<undefined | number>()
  let servicesData = $state<undefined | MyProfileServicesData>()
  const setServicesData = (data: MyProfileServicesData) => (servicesData = data)

  const unsubscribe = useRouter().routerBase.subscribe((value) => {
    if (value.uri !== '/my-profile/services') return

    currentPage = getPageParam()
    fetchMyServicesData(currentPage, setServicesData)
  })
  onDestroy(() => unsubscribe())
</script>

<PaperElement
  styles="min-h-app !ring-0 lg:dark:!ring-1 text-base !bg-transparent dark:!bg-transparent !shadow-none !p-0 flex flex-col gap-6 md:!p-6 md:!shadow-md dark:!shadow-none md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block"
>
  {#if currentPage && servicesData}
    <Title text="My services" />
    {#if servicesData.services.length === 0}
      <NoContentMessage text="You have no services yet" />
    {:else}
      <ServicesTable services={servicesData.services} />
    {/if}

    {#if servicesData.totalPages > 1}
      <Pagination
        setPage={(value) => navigate(`/my-profile/services?page=${value}`)}
        {currentPage}
        totalPages={servicesData.totalPages}
      />
    {/if}
  {:else}
    <Loader />
  {/if}
</PaperElement>
