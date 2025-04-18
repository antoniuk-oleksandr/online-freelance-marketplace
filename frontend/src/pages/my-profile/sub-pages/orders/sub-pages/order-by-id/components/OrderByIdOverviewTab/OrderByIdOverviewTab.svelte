<script lang="ts" module>
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import Title from '@/common-components/Title/Title.svelte'
  import DividerElement from '@/common-components/DividerElement/DividerElement.svelte'
  import OrderByIdOverviewTabReceipt from './components/OrderByIdOverviewTabReceipt/OrderByIdOverviewTabReceipt.svelte'
  import OrderByIdOverviewTabTotalPrice from './components/OrderByIdOverviewTabTotalPrice/OrderByIdOverviewTabTotalPrice.svelte'
  import OrderByIdOverviewTabDeliveryDate from './components/OrderByIdOverviewTabDeliveryDate/OrderByIdOverviewTabDeliveryDate.svelte'
  import OrderByIdOverviewTabFreelancerInfo from './components/OrderByIdOverviewTabFreelancerInfo/OrderByIdOverviewTabFreelancerInfo.svelte'
  import OrderByIdOverviewTabOrderNumber from './components/OrderByIdOverviewTabOrderNumber/OrderByIdOverviewTabOrderNumber.svelte'
  import OrderByIdOverviewTabMobileReceipt from './components/OrderByIdOverviewTabMobileReceipt/OrderByIdOverviewTabMobileReceipt.svelte'
  import type { MyProfileOverviewRequestResponse } from '@/types/MyProfileOverviewRequestResponse'

  export { orderByIdOverviewTabComponent }
</script>

{#snippet orderByIdOverviewTabComponent(response: MyProfileOverviewRequestResponse | undefined)}
  <PaperElement
    styles="!ring-0 lg:dark:!ring-1 !p-0 lg:!p-6 !shadow-none !bg-transparent lg:!bg-light-palette-background-block lg:dark:!bg-dark-palette-background-block  lg:!shadow-md dark:!shadow-none gap-6 flex flex-col flex-grow text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
  >
    {#if response && response.status === 200}
      <div class="flex flex-col md:flex-row gap-3 md:justify-between">
        <div class="flex flex-col gap-3">
          <Title text={response.data.service.title} />
          <div class="grid grid-cols-two-columns-with-divider md:flex gap-6">
            <OrderByIdOverviewTabFreelancerInfo freelancer={response.data.freelancer} />
            <DividerElement orientation="vertical" />
            <OrderByIdOverviewTabDeliveryDate deliveryDate={response.data.deliveryDate} />
          </div>
        </div>
        <OrderByIdOverviewTabTotalPrice totalPrice={response.data.totalPrice} />
      </div>
      <OrderByIdOverviewTabOrderNumber id={response.data.id} />
      <OrderByIdOverviewTabReceipt order={response.data} />
      <OrderByIdOverviewTabMobileReceipt order={response.data} />
    {/if}
  </PaperElement>
{/snippet}
