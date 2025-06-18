<script module lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import type { MyProfileDeliveryRequestResponse } from '@/types/MyProfileDeliveryRequestResponse'
  import RequestByIdDeliveryTabInProgresContent from './components/RequestByIdDeliveryTabInProgresContent/RequestByIdDeliveryTabInProgresContent.svelte'
  import OrderByIdDeliveryTabCompleteContent from '@/pages/my-profile/sub-pages/orders/sub-pages/order-by-id/components/OrderByIdDeliveryTab/components/OrderByIdDeliveryTabCompleteContent/OrderByIdDeliveryTabCompleteContent.svelte'
  import type { DeliveryData } from '@/types/DeliveryData'
  import { compareDeliveryData } from './helpers'
  import OrderByIdDeliveryTabCanceledContent from '@/pages/my-profile/sub-pages/orders/sub-pages/order-by-id/components/OrderByIdDeliveryTab/components/OrderByIdDeliveryTabCanceledContent/OrderByIdDeliveryTabCanceledContent.svelte'
    import OrderByIdDeliveryTabAwaitingAcceptanceContent from '@/pages/my-profile/sub-pages/orders/sub-pages/order-by-id/components/OrderByIdDeliveryTab/components/OrderByIdDeliveryTabAwaitingAcceptanceContent/OrderByIdDeliveryTabAwaitingAcceptanceContent.svelte'
    import RequestByIdDeliveryTabAwaitingAcceptanceContent from './components/RequestByIdDeliveryTabAwaitingAcceptanceContent/RequestByIdDeliveryTabAwaitingAcceptanceContent.svelte'

  let newDeliveryData = $state<DeliveryData | undefined>()
  const setNewDeliveryData = $state((data: DeliveryData) => {
    newDeliveryData = data
  })

  export { requestByIdDeliveryTabComponent }
</script>

{#snippet requestByIdDeliveryTabComponent(
  response: MyProfileDeliveryRequestResponse | undefined,
  orderId: string,
)}
  <PaperElement
    styles="flex -mx-6 md:m-0 items-center justify-center !ring-0 lg:dark:!ring-1 !p-0 !shadow-none !bg-transparent lg:!bg-light-palette-background-block lg:dark:!bg-dark-palette-background-block  lg:!shadow-md dark:!shadow-none gap-3 flex flex-col flex-grow text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
  >
    {#if response && response.status === 200}
      <RequestByIdDeliveryTabInProgresContent
        {setNewDeliveryData}
        data={compareDeliveryData(response.data, newDeliveryData)}
        {orderId}
      />
      <OrderByIdDeliveryTabCompleteContent
        renderReviewCTA={false}
        data={compareDeliveryData(response.data, newDeliveryData)}
        {orderId}
      />
      <OrderByIdDeliveryTabCanceledContent
        showRefundAndNextSteps={false}
        data={compareDeliveryData(response.data, newDeliveryData)}
      />
      <RequestByIdDeliveryTabAwaitingAcceptanceContent
        data={compareDeliveryData(response.data, newDeliveryData)}
      />
    {/if}
  </PaperElement>
{/snippet}
