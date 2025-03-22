<script lang="ts">
  import Loader from '@/common-components/Loader/Loader.svelte'
  import NewSelect from '@/common-components/NewSelect/NewSelect.svelte'
  import NoContentMessage from '@/common-components/NoContentMessage/NoContentMessage.svelte'
  import Pagination from '@/common-components/Pagination/Pagination.svelte'
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import Title from '@/common-components/Title/Title.svelte'
  import type { MyProfileRequestsData } from '@/types/MyProfileRequestsData'
  import { onDestroy } from 'svelte'
  import { navigate, useRouter } from 'svelte-routing'
  import RequestsTable from './components/RequestsTable.svelte'
  import {
    fetchMyRequestsData,
    getInitialRequestsStatusValues,
    getMyProfileRequestsInitialSelectIndex,
    getMyProfileRequestsPageParam,
    handleMyProfileRequestsSelectChange,
  } from './helpers'
  import type { StatusEnum } from '@/types/StatusEnum'

  let currentPage = $state<undefined | number>()
  let selectStatusValue = $state<StatusEnum | 0>(0)
  let requestsData = $state<undefined | MyProfileRequestsData>()
  const setRequestsData = (data: MyProfileRequestsData) => (requestsData = data)

  const unsubscribe = useRouter().routerBase.subscribe((value) => {
    if (value.uri !== '/my-profile/requests') return

    const { page, status } = getMyProfileRequestsPageParam()
    currentPage = page
    selectStatusValue = status
    fetchMyRequestsData(currentPage, setRequestsData, selectStatusValue)
  })
  onDestroy(() => unsubscribe())
</script>

<PaperElement
  styles="min-h-app !ring-0 lg:dark:!ring-1 text-base !bg-transparent dark:!bg-transparent !shadow-none !p-0 flex flex-col gap-6 md:!p-6 md:!shadow-md dark:!shadow-none md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block"
>
  {#if currentPage && requestsData}
    <Title text="My requests" />
    <NewSelect
      onSelectChange={(value) => handleMyProfileRequestsSelectChange(value)}
      initialSelectIndex={getMyProfileRequestsInitialSelectIndex(selectStatusValue)}
      selectWidth="min-w-60"
      initialSelectValues={getInitialRequestsStatusValues()}
      layoutStyles="self-end"
      modalHeaderTitle="Request status"
    />
    {#if requestsData.requests.length === 0}
      <NoContentMessage text="You have no requests yet" />
    {:else}
      <RequestsTable requests={requestsData.requests} />
    {/if}

    {#if requestsData.totalPages > 1}
      <Pagination
        setPage={(value) =>
          navigate(`/my-profile/requests?page=${value}&status=${selectStatusValue}`)}
        {currentPage}
        totalPages={requestsData.totalPages}
      />
    {/if}
  {:else}
    <Loader />
  {/if}
</PaperElement>

<!-- Add initial selected value(value after page reload)-->
<!-- Fix "http://localhost:3000/my-profile/requests?page=1&status=2" no select element -->
