<script lang="ts">
  import type { DiaryItem } from '@/types/DiaryItem'
  import OrderByIdDiaryTabDayLayout from './OrderByIdDiaryTabDayLayout.svelte'

  type OrderByIdDiaryTabDayProps = DiaryItem

  const { text, date, bgColor = '', textColor = '000000' }: OrderByIdDiaryTabDayProps = $props()

  const day = new Date(date).getDate()
</script>

<OrderByIdDiaryTabDayLayout>
  <p
    class="text-end text-lg font-semibold text-light-palette-text-primary dark:text-dark-palette-text-primary"
  >
    {day}
  </p>
  <p
    style="
    background: #{bgColor};
    color: #{textColor}
    "
    class="line-clamp-1 rounded-md px-1"
  >
    {text}
  </p>
</OrderByIdDiaryTabDayLayout>
