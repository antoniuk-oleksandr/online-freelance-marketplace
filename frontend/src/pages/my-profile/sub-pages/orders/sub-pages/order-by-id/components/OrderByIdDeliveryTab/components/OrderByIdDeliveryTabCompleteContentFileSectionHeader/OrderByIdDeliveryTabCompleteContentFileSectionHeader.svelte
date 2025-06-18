<script lang="ts">
  import Title from '@/common-components/Title/Title.svelte'
  import Icon from '@iconify/svelte'
  import { handleDownloadButtonClick } from '../../handlers'
  import OrderByIdDeliveryTabCompleteContentFileSectionHeaderLayout from './OrderByIdDeliveryTabCompleteContentFileSectionHeaderLayout.svelte'
  import type { DeliveryData } from '@/types/DeliveryData'

  type OrderByIdDeliveryTabCompleteContentFileSectionHeaderProps = {
    data: DeliveryData
  }

  const { data }: OrderByIdDeliveryTabCompleteContentFileSectionHeaderProps = $props()
</script>

<OrderByIdDeliveryTabCompleteContentFileSectionHeaderLayout>
  <Title text={`Delivered Files (${data.delivery.files?.length || 0})`} />
  {#if data.delivery.files && data.delivery.files.length > 0}
    <button
      onclick={() => handleDownloadButtonClick(data.delivery.files)}
      class="text-green-600 dark:text-green-500 text-sm font-medium flex items-center gap-1"
    >
      <Icon icon="hugeicons:download-04" width="24" height="24" />
      Download All
    </button>
  {/if}
</OrderByIdDeliveryTabCompleteContentFileSectionHeaderLayout>
