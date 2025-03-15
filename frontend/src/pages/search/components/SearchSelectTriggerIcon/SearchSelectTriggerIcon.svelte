<script lang="ts">
  import { makeSearchSelectTitle } from '@/pages/search/helpers'
  import { OrderType } from '@/types/OrderType'
  import Icon from '@iconify/svelte'

  type SearchSelectTriggerIconProps = {
    selectedAdditionalItem: string
  }

  const { selectedAdditionalItem }: SearchSelectTriggerIconProps = $props()
</script>

{#if selectedAdditionalItem === makeSearchSelectTitle(OrderType.Descending, true)}
  <Icon icon="hugeicons:sort-by-down-01" width="24" height="24" />
{:else}
  <Icon icon="hugeicons:sort-by-up-01" width="24" height="24" />
{/if}
