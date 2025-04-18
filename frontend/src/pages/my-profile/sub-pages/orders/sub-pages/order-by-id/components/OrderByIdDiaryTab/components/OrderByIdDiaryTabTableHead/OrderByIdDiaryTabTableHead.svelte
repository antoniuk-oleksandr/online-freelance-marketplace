<script lang="ts">
  import OrderByIdDiaryTabTableHeadLayout from './OrderByIdDiaryTabTableHeadLayout.svelte'

  const days = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']
</script>

<OrderByIdDiaryTabTableHeadLayout>
  {#each days as day}
    <div
      class="text-end text-lg pr-2 font-semibold text-light-palette-text-primary dark:text-dark-palette-text-primary"
    >
      {day}
    </div>
  {/each}
</OrderByIdDiaryTabTableHeadLayout>
