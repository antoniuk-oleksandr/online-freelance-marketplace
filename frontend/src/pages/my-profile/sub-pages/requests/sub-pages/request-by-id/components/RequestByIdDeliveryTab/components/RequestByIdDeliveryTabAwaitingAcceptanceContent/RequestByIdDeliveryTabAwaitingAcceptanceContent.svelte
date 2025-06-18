<script lang="ts">
  import NoContentMessage from '@/common-components/NoContentMessage/NoContentMessage.svelte'
  import type { DeliveryData } from '@/types/DeliveryData'
  import { StatusEnum } from '@/types/StatusEnum'

  type RequestByIdDeliveryTabAwaitingAcceptanceContentProps = {
    data: DeliveryData
  }

  const { data }: RequestByIdDeliveryTabAwaitingAcceptanceContentProps = $props()
</script>

{#if data.status === StatusEnum.AwaitingAcceptance}
  <NoContentMessage text="Delivery is not available yet" />
{/if}
