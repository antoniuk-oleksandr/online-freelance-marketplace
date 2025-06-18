<script lang="ts">
  import NewSelect from '@/common-components/NewSelect/NewSelect.svelte'
  import type { NewSelectItem } from '@/types/NewSelectItem'
  import { handleOrderStatusChange } from '../../handlers'

  type OrderByIdSidebarSelectProps = {
    orderId: string
  }

  const { orderId }: OrderByIdSidebarSelectProps = $props()

  const initialSelectValues: NewSelectItem[] = [
    { title: 'In Progress', color: 'text-red-500' },
    { title: 'Cancelled' },
  ]
</script>

<NewSelect
  noneSelectedText="Select status"
  initialSelectIndex={-1}
  onSelectChange={(value) => handleOrderStatusChange(value, orderId)}
  {initialSelectValues}
  modalHeaderTitle="Select the order status"
/>
