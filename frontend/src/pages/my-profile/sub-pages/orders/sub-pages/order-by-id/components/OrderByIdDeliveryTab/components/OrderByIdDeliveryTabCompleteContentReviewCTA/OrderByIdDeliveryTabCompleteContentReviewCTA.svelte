<script lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import OrderByIdDeliveryTabCompleteContentReviewCTALayout from './OrderByIdDeliveryTabCompleteContentReviewCTALayout.svelte'
  import Icon from '@iconify/svelte'
  import { getContext } from 'svelte'
  import type { Writable } from 'svelte/store'
  import { navigate } from 'svelte-routing'

  type OrderByIdDeliveryTabCompleteContentReviewCTAProps = {
    orderId: string
  }

  const tabIndex = getContext<Writable<number>>('tabIndex')

  const { orderId }: OrderByIdDeliveryTabCompleteContentReviewCTAProps = $props()
</script>

<OrderByIdDeliveryTabCompleteContentReviewCTALayout>
  <div class="flex flex-col gap-6 sm:flex-row items-center justify-between">
    <div>
      <h3 class="text-green-800 dark:text-green-300 font-medium">How was your experience?</h3>
      <p class="text-green-700 dark:text-green-200 text-sm mt-1">
        Leave a review for the freelancer
      </p>
    </div>
    <Button
      clickAction={() => {
        navigate(`/my-profile/orders/${orderId}?tabIndex=5`)
      }}
      linkStyles="!w-full sm:!w-fit"
      styles="w-full sm:w-fit !transition-colors !flex !gap-2"
      color="green"
    >
      Leave Review
      <Icon icon="heroicons:arrow-right" class="w-4 h-4" />
    </Button>
  </div>
</OrderByIdDeliveryTabCompleteContentReviewCTALayout>
