<script lang="ts">
  import type { DiaryItem } from '@/types/DiaryItem'
  import OrderByIdDiaryTabDayListLayout from './OrderByIdDiaryTabDayListLayout.svelte'
  import OrderByIdDiaryTabDay from '../OrderByIdDiaryTabDay/OrderByIdDiaryTabDay.svelte'

  type OrderByIdDiaryTabDayListProps = {
    diaryItems: DiaryItem[]
  }

  const { diaryItems }: OrderByIdDiaryTabDayListProps = $props()
</script>

<OrderByIdDiaryTabDayListLayout>
  {#each diaryItems as item}
    <OrderByIdDiaryTabDay {...item} />
  {/each}
</OrderByIdDiaryTabDayListLayout>
