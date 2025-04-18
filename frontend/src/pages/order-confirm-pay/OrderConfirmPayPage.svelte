<script lang="ts">
  import Stepper from '@/common-components/Stepper/Stepper.svelte'
  import OrderConfirmPayPageLayout from '@/pages/order-confirm-pay/OrderConfirmPayPageLayout.svelte'
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import PaymentForm from '@/pages/order-confirm-pay/components/PaymentForm/PaymentForm.svelte'
  import OrderOverview from '@/pages/order-confirm-pay/components/OrderOverview/OrderOverview.svelte'
  import {
    fetchServiceDetailsAndPackage,
    parseOrderRequestQuery,
  } from '@/pages/order-request/helpers'
  import type { Package } from '@/types/Package.ts'
  import type { StepperItem } from '@/types/StepperItem.ts'
  import type { RestrictedService } from '@/types/RestrictedService.ts'

  let serviceId = $state<string | undefined>()
  const setServiceId = (value: string) => (serviceId = value)
  let packageId = $state<string | undefined>()
  const setPackageId = (value: string) => (packageId = value)
  parseOrderRequestQuery(setServiceId, setPackageId)

  let selectedPackage = $state<Package | undefined>()
  const setSelectedPackage = (newPackage: Package) => (selectedPackage = newPackage)

  let serviceData = $state<RestrictedService | undefined>()
  const setServiceData = (newServiceData: RestrictedService) => (serviceData = newServiceData)

  $effect(() => {
    if (!serviceId || !packageId) return
    fetchServiceDetailsAndPackage(serviceId, packageId, setServiceData, setSelectedPackage)
  })

  let steps: StepperItem[] = $derived([
    {
      text: 'Order details',
      link: `/orders/request?serviceId=${serviceId}&packageId=${selectedPackage?.id}`,
    },
    { text: 'Confirm & pay' },
    { text: 'Submit requirements' },
  ])
</script>

{#if serviceId && packageId}
  <OrderConfirmPayPageLayout {serviceId} {packageId}>
    <Stepper styles="col-span-2 capitalize flex h-fit justify-center" activeStepIndex={1} {steps} />
    <div class="flex flex-col lg:grid grid-cols-search-page gap-6 h-full">
      <PaperElement
        styles="flex flex-col gap-6  !shadow-none md:!shadow-md !p-0 md:!p-6 !bg-transparent !ring-0 md:dark:!ring-1 md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block"
      >
        <PaymentForm />
      </PaperElement>
      <OrderOverview pkg={selectedPackage} {serviceData} />
    </div>
  </OrderConfirmPayPageLayout>
{/if}
