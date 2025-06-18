<script lang="ts">
  import NoContentMessage from '@/common-components/NoContentMessage/NoContentMessage.svelte'
  import type { DeliveryData } from '@/types/DeliveryData'
  import { StatusEnum } from '@/types/StatusEnum'
    import OrderByIdAwaitingAcceptanceReviewTab from '../../../OrderByIdReviewTab/components/OrderByIdAwaitingAcceptanceReviewTab/OrderByIdAwaitingAcceptanceReviewTab.svelte'

  type OrderByIdDeliveryTabAwaitingAcceptanceContentProps = {
    data: DeliveryData
  }

  const { data }: OrderByIdDeliveryTabAwaitingAcceptanceContentProps = $props()
</script>

{#if data.status === StatusEnum.AwaitingAcceptance}
  <OrderByIdAwaitingAcceptanceReviewTab/>
{/if}
