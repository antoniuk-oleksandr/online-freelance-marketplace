<script lang="ts">
  import type { DeliveryData } from '@/types/DeliveryData'
  import OrderByIdDeliveryTabCompleteContentFile from '../OrderByIdDeliveryTabCompleteContentFile/OrderByIdDeliveryTabCompleteContentFile.svelte'
  import OrderByIdDeliveryTabCompleteContentFileListLayout from './OrderByIdDeliveryTabCompleteContentFileListLayout.svelte'

  type OrderByIdDeliveryTabCompleteContentFileListProps = {
    data: DeliveryData
  }

  const { data }: OrderByIdDeliveryTabCompleteContentFileListProps = $props()
</script>

{#if data.delivery}
  <OrderByIdDeliveryTabCompleteContentFileListLayout>
    {#each data.delivery.files as file}
      <div
        class=" border border-light-palette-divider dark:border-dark-palette-divider rounded-lg p-6 transition-colors"
      >
        <OrderByIdDeliveryTabCompleteContentFile {file} />
      </div>
    {/each}
  </OrderByIdDeliveryTabCompleteContentFileListLayout>
{/if}
