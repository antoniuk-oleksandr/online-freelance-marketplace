<script lang="ts">
  import type { DiaryItem } from '@/types/DiaryItem'
  import OrderByIdDiaryTabTableBodyLayout from './OrderByIdDiaryTabTableBodyLayout.svelte'
  import OrderByIdDiaryTabDay from '../OrderByIdDiaryTabDay/OrderByIdDiaryTabDay.svelte'

  type OrderByIdDiaryTabTableBodyProps = {
    diaryItems: DiaryItem[]
  }

  const { diaryItems }: OrderByIdDiaryTabTableBodyProps = $props()
</script>

<OrderByIdDiaryTabTableBodyLayout>
  {#each Array.from({ length: 5 }) as _, rowIndex}
  <div class="w-full grid grid-cols-7">
    {#each Array.from({ length: 7 }) as _, columnIndex}
      <OrderByIdDiaryTabDay {...diaryItems[rowIndex * 7 + columnIndex]} />
    {/each}
  </div>
  {/each}
</OrderByIdDiaryTabTableBodyLayout>
