<script lang="ts">
  import OrderRequestPackagesBlock from '@/pages/order-request/components/OrderRequestPackagesBlock/OrderRequestPackagesBlock.svelte'
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import OrderRequestSummary from '@/pages/order-request/components/OrderRequestSummary/OrderRequestSummary.svelte'
  import OrderRequestServiceInfo from '@/pages/order-request/components/OrderRequestServiceInfo/OrderRequestServiceInfo.svelte'
  import type { Package } from '@/types/Package.ts'
  import OrderRequestPageContentLayout from '@/pages/order-request/components/OrderRequestPageContent/OrderRequestPageContentLayout.svelte'
  import type { RestrictedService } from '@/types/RestrictedService.ts'

  type OrderRequestPageContentProps = {
    serviceData: RestrictedService
    selectedPackage: Package
    setSelectedPackage: (newPackage: Package) => void
  }
  const props: OrderRequestPageContentProps = $props()
</script>

<OrderRequestPageContentLayout>
  <PaperElement
    styles="flex flex-col gap-6 !h-full !shadow-none md:!shadow-md !p-0 md:!p-6 !bg-transparent !ring-0 md:dark:!ring-1 md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block"
  >
    <OrderRequestServiceInfo {...props} />
    <OrderRequestPackagesBlock {...props} />
  </PaperElement>
  <OrderRequestSummary serviceId={props.serviceData.id} package={props.selectedPackage} />
</OrderRequestPageContentLayout>
