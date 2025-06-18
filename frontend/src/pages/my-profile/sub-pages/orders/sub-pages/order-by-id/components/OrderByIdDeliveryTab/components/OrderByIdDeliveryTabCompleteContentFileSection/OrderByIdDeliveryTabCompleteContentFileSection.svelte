<script lang="ts">
  import type { DeliveryData } from '@/types/DeliveryData'
  import OrderByIdDeliveryTabCompleteContentFileSectionLayout from './OrderByIdDeliveryTabCompleteContentFileSectionLayout.svelte'
  import OrderByIdDeliveryTabCompleteContentFileList from '../OrderByIdDeliveryTabCompleteContentFileList/OrderByIdDeliveryTabCompleteContentFileList.svelte'
  import OrderByIdDeliveryTabCompleteContentFileSectionHeader from '../OrderByIdDeliveryTabCompleteContentFileSectionHeader/OrderByIdDeliveryTabCompleteContentFileSectionHeader.svelte'

  type OrderByIdDeliveryTabCompleteContentFileSectionProps = {
    data: DeliveryData
  }

  const { data }: OrderByIdDeliveryTabCompleteContentFileSectionProps = $props()
</script>

<OrderByIdDeliveryTabCompleteContentFileSectionLayout>
  <OrderByIdDeliveryTabCompleteContentFileSectionHeader {data} />
  <OrderByIdDeliveryTabCompleteContentFileList {data} />
</OrderByIdDeliveryTabCompleteContentFileSectionLayout>
