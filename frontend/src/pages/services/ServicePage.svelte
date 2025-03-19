<script lang="ts">
  import ServicePageLayout from '@/pages/services/ServicePageLayout.svelte'
  import ServiceBreadcrumbs from '@/pages/services/components/ServiceBreadcrumbs/ServiceBreadcrumbs.svelte'
  import ServiceContent from '@/pages/services/components/ServiceContent/ServiceContent.svelte'
  import { request } from '@/api/request'
  import type { GetServiceByUserIdRequestResponse } from '@/types/GetServiceByIdRequestResponse.ts'
  import type { UpdateFunc } from '@/types/UpdateFunc.ts'
  import { errorStore } from '@/common-stores/error-store'

  type ServicePageProps = { id: string }

  const { id }: ServicePageProps = $props()

  let serviceResponse = $state<GetServiceByUserIdRequestResponse | undefined>()
  let setServiceResponse: UpdateFunc<GetServiceByUserIdRequestResponse | undefined> = (updater) => {
    ;(async () => {
      serviceResponse = await updater(serviceResponse)
    })()
  }

  request<GetServiceByUserIdRequestResponse>('GET', `/freelances/${id}`, 'GET').then((response) => {
    if (response.status === 200) {
      serviceResponse = response
    } else errorStore.set({ shown: true, error: response.data.error })
  })
</script>

{#if serviceResponse && serviceResponse.status === 200}
  <ServicePageLayout>
    <ServiceBreadcrumbs service={serviceResponse.data.service} />
    <ServiceContent {serviceResponse} {setServiceResponse} />
  </ServicePageLayout>
{/if}
