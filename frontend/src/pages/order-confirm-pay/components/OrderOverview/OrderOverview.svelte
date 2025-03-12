<script lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import Button from '@/common-components/Button/Button.svelte'
  import Spinner from '@/common-components/Spinner/Spinner.svelte'
  import type { Package } from '@/types/Package.ts'
  import DividerElement from '@/common-components/DividerElement/DividerElement.svelte'
  import OrderOverviewTopBlock from '@/pages/order-confirm-pay/components/OrderOverview/components/OrderOverviewTopBlock/OrderOverviewTopBlock.svelte'
  import OrderOverviewBottomBlock from '@/pages/order-confirm-pay/components/OrderOverview/components/OrderOverviewBottomBlock/OrderOverviewBottomBlock.svelte'
  import type { RestrictedService } from '@/types/RestrictedService.ts'
  import { getContext } from 'svelte'

  type OrderOverviewProps = {
    serviceData?: RestrictedService
    pkg?: Package
  }

  let isSubmitting = $state<boolean>(false)
  getContext<any>('orderComfirmPayFormIsSubmitting').subscribe(
    (value: boolean) => (isSubmitting = value),
  )

  const props: OrderOverviewProps = $props()
</script>

<PaperElement
  styles="flex flex-col {props.serviceData
    ? '!h-fit'
    : 'justify-center items-center !h-order-overview'} gap-6 !shadow-none md:!shadow-md !p-0 md:!p-6 !bg-transparent !ring-0 md:dark:!ring-1 md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block"
>
  {#if props.serviceData && props.pkg}
    <OrderOverviewTopBlock serviceData={props.serviceData} pkg={props.pkg} />
    <DividerElement />
    <OrderOverviewBottomBlock pkg={props.pkg} />
    <Button loading={isSubmitting} type="submit">Confirm & Pay</Button>
  {:else}
    <Spinner />
  {/if}
</PaperElement>
