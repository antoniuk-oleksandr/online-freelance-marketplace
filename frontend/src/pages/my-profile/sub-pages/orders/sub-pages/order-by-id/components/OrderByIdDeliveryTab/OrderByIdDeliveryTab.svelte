<script module lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import Title from '@/common-components/Title/Title.svelte'
  import { formatOrderDate } from '@/pages/order-submit-requirements/helpers'
  import type { MyProfileDeliveryRequestResponse } from '@/types/MyProfileDeliveryRequestResponse'
  import { getFileServerHost } from '@/utils/utils'
  import { Link } from 'svelte-routing'

  export { orderByIdDeliveryTabComponent }
</script>

{#snippet orderByIdDeliveryTabComponent(response: MyProfileDeliveryRequestResponse | undefined)}
  <PaperElement
    styles="flex items-center justify-center !ring-0 lg:dark:!ring-1 !p-0 lg:!p-6 !shadow-none !bg-transparent lg:!bg-light-palette-background-block lg:dark:!bg-dark-palette-background-block  lg:!shadow-md dark:!shadow-none gap-3 flex flex-col flex-grow text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
  >
    {#if response && response.status === 200}
      <img class="mb-3" alt="img" width="128" height="128" src="/images/box.png" />
      <Title capitalize={false} text="Your Delivery is on the Way" />
      <div class="flex gap-1">
        <Link
          class=" font-semibold text-cyan-500 hover:underline"
          to="/users/{response.data.freelancer.id}">{response.data.freelancer.username}</Link
        >
        <span>should deliver this order by {formatOrderDate(response.data.deliveryDate)}</span>
      </div>
    {/if}
  </PaperElement>
{/snippet}
