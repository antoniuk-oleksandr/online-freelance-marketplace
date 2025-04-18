<script lang="ts">
  import Stepper from '@/common-components/Stepper/Stepper.svelte'
  import type { Package } from '@/types/Package.ts'
  import OrderRequestPageLayout from '@/pages/order-request/OrderRequestPageLayout.svelte'
  import OrderRequestPageContent from '@/pages/order-request/components/OrderRequestPageContent/OrderRequestPageContent.svelte'
  import {
    fetchServiceDetailsAndPackage,
    parseOrderRequestQuery,
  } from '@/pages/order-request/helpers'
  import type { StepperItem } from '@/types/StepperItem.ts'
  import type { RestrictedService } from '@/types/RestrictedService.ts'

  let serviceData = $state<RestrictedService | undefined>()
  const setServiceData = (newServiceData: RestrictedService) => (serviceData = newServiceData)

  let serviceId = $state<string | undefined>()
  const setServiceId = (value: string) => (serviceId = value)
  let packageId = $state<string | undefined>()
  const setPackageId = (value: string) => (packageId = value)
  parseOrderRequestQuery(setServiceId, setPackageId)

  let selectedPackage = $state<Package | undefined>()
  const setSelectedPackage = (newPackage: Package) => (selectedPackage = newPackage)

  $effect(() => {
    if (!serviceId || !packageId) return
    fetchServiceDetailsAndPackage(serviceId, packageId, setServiceData, setSelectedPackage)
  })

  const steps: StepperItem[] = [
    { text: 'Order details' },
    { text: 'Confirm & pay' },
    { text: 'Submit requirements' },
  ]
</script>

{#if serviceData && selectedPackage !== undefined}
  <OrderRequestPageLayout>
    <Stepper styles="col-span-2 capitalize flex h-fit justify-center" activeStepIndex={0} {steps} />
    <OrderRequestPageContent {serviceData} {selectedPackage} {setSelectedPackage} />
  </OrderRequestPageLayout>
{/if}
