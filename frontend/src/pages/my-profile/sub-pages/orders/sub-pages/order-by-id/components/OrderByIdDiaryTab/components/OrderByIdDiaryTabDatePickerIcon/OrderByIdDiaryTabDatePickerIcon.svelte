<script lang="ts">
  import Icon from '@iconify/svelte'
  import OrderByIdDiaryTabDatePickerIconLayout from './OrderByIdDiaryTabDatePickerIconLayout.svelte'

  type OrderByIdDiaryTabDatePickerIconProps = {
    direction: 'left' | 'right'
  }

  const { direction }: OrderByIdDiaryTabDatePickerIconProps = $props()
</script>

<OrderByIdDiaryTabDatePickerIconLayout>
  <Icon icon="hugeicons:arrow-{direction}-01" width="24" height="24" />
</OrderByIdDiaryTabDatePickerIconLayout>
