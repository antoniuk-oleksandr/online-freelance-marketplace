<script lang="ts">
  import Title from '@/common-components/Title/Title.svelte'
  import { formatOrderDate } from '@/pages/order-submit-requirements/helpers'
  import type { DeliveryData } from '@/types/DeliveryData'
  import { StatusEnum } from '@/types/StatusEnum'
  import { Link } from 'svelte-routing'

  type OrderByIdDeliveryTabInProgressContentProps = {
    data: DeliveryData
  }

  const { data }: OrderByIdDeliveryTabInProgressContentProps = $props()
</script>

{#if data.status === StatusEnum.InProgress && data.freelancer && data.delivery}
  <img class="mb-3" alt="img" width="128" height="128" src="/images/box.png" />
  <Title capitalize={false} text="Your Delivery is on the Way" />
  <div class="flex gap-1">
    <Link class=" font-semibold text-cyan-500 hover:underline" to="/users/{data.freelancer.id}"
      >{data.freelancer.username}</Link
    >
    <span>should deliver this order by {formatOrderDate(data.delivery.date)}</span>
  </div>
{/if}
