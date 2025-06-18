<script module lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import type { MyProfileDeliveryRequestResponse } from '@/types/MyProfileDeliveryRequestResponse'
  import OrderByIdDeliveryTabInProgressContent from './components/OrderByIdDeliveryTabInProgressContent/OrderByIdDeliveryTabInProgressContent.svelte'
  import OrderByIdDeliveryTabInCompleteContent from './components/OrderByIdDeliveryTabCompleteContent/OrderByIdDeliveryTabCompleteContent.svelte'
  import OrderByIdDeliveryTabCanceledContent from './components/OrderByIdDeliveryTabCanceledContent/OrderByIdDeliveryTabCanceledContent.svelte'
    import OrderByIdDeliveryTabAwaitingAcceptenceContent from './components/OrderByIdDeliveryTabAwaitingAcceptanceContent/OrderByIdDeliveryTabAwaitingAcceptanceContent.svelte'

  export { orderByIdDeliveryTabComponent }
</script>

{#snippet orderByIdDeliveryTabComponent(
  response: MyProfileDeliveryRequestResponse | undefined,
  orderId: string,
)}
  <PaperElement
    styles="flex -mx-6 md:m-0 items-center justify-center !ring-0 lg:dark:!ring-1 !p-0 !shadow-none !bg-transparent lg:!bg-light-palette-background-block lg:dark:!bg-dark-palette-background-block  lg:!shadow-md dark:!shadow-none gap-3 flex flex-col flex-grow text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
  >
    {#if response && response.status === 200}
      <OrderByIdDeliveryTabInProgressContent data={response.data} />
      <OrderByIdDeliveryTabInCompleteContent data={response.data} {orderId} />
      <OrderByIdDeliveryTabCanceledContent data={response.data} />
      <OrderByIdDeliveryTabAwaitingAcceptenceContent data={response.data}/>
    {/if}
  </PaperElement>
{/snippet}
