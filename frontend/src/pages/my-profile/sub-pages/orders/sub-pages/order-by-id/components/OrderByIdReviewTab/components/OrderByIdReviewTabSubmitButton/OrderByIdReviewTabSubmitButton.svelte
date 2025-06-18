<script lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import Icon from '@iconify/svelte'
  import OrderByIdReviewTabSubmitButtonLayout from './OrderByIdReviewTabSubmitButtonLayout.svelte'
  import { getContext } from 'svelte'
  import type { Writable } from 'svelte/store'

  const formIsSubmittingStore = getContext<Writable<boolean>>('formIsSubmittingStore')
</script>

<OrderByIdReviewTabSubmitButtonLayout>
  <Button
    loading={$formIsSubmittingStore}
    styles="flex gap-2 !w-full md:!w-44"
    type="submit"
    color="green"
  >
    <Icon icon="hugeicons:checkmark-circle-02" width="20" height="20" />
    <span>Submit Review</span>
  </Button>
</OrderByIdReviewTabSubmitButtonLayout>
