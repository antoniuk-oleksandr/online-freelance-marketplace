<script lang="ts">
  import type { DiaryItem } from '@/types/DiaryItem'
  import OrderByIdDiaryTabTableLayout from './OrderByIdDiaryTabTableLayout.svelte'
  import OrderByIdDiaryTabTableBody from '../OrderByIdDiaryTabTableBody/OrderByIdDiaryTabTableBody.svelte'
  import OrderByIdDiaryTabTableHead from '../OrderByIdDiaryTabTableHead/OrderByIdDiaryTabTableHead.svelte'

  type OrderByIdDiaryTabTableProps = {
    diaryItems: DiaryItem[]
  }

  const props: OrderByIdDiaryTabTableProps = $props()
</script>

<OrderByIdDiaryTabTableLayout>
  <OrderByIdDiaryTabTableHead />
  <OrderByIdDiaryTabTableBody {...props} />
</OrderByIdDiaryTabTableLayout>
