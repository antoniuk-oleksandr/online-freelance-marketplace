<script lang="ts">
  import type { OrderByIdOverviewTabData } from '@/types/OrderByIdOverviewTabData'
  import OrderByIdSidebarServiceImage from '../OrderByIdSidebarServiceImage/OrderByIdSidebarServiceImage.svelte'
  import OrderByIdSidebarTopLayout from './OrderByIdSidebarTopLayout.svelte'
    import StatusElement from '@/common-components/StatusElement/StatusElement.svelte'

  type OrderByIdSidebarTopProps = {
    data: OrderByIdOverviewTabData
  }

  const { data }: OrderByIdSidebarTopProps = $props()
</script>

<OrderByIdSidebarTopLayout>
  <OrderByIdSidebarServiceImage alt="service-image" src={data.service.image} />
  <div class="flex flex-col justify-between h-20">
    <p class="line-clamp-2">{data.service.title}</p>
    <StatusElement styles="size-fit" status={data.status} />
  </div>
</OrderByIdSidebarTopLayout>
