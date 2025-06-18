<script lang="ts">
  import Label from '@/common-components/Label/Label.svelte'
  import type { ReviewTabFormData } from '@/types/ReviewTabFormData'
  import Icon from '@iconify/svelte'
  import { getContext } from 'svelte'
  import type { Writable } from 'svelte/store'
  import { handleReviewRatingChange } from '../../handlers'
  import OrderByIdReviewTabRatingSectionLayout from './OrderByIdReviewTabRatingSectionLayout.svelte'
  import InputError from '@/common-components/Sign/components/SignInput/components/InputError/InputError.svelte'

  let hoverRating = $state(0)
  const formDataStore = getContext<Writable<ReviewTabFormData>>('formDataStore')
  const formWasSubmittedStore = getContext<Writable<boolean>>('formWasSubmittedStore')
  const formErrorsStore = getContext<Writable<Record<string, string[]>>>('formErrorsStore')
</script>

<OrderByIdReviewTabRatingSectionLayout>
  <Label styles="w-full" text="How would you rate this service?" />
  <div class="flex items-center">
    {#each [1, 2, 3, 4, 5] as star}
      <button
        class="{star <= hoverRating || (star <= $formDataStore.rating && hoverRating === 0)
          ? 'text-yellow-500'
          : 'text-neutral-500'}
          duration-200 transition-colors"
        aria-label="Rate this service"
        type="button"
        onclick={() => handleReviewRatingChange(formDataStore, star)}
        onmouseenter={() => (hoverRating = star)}
        onmouseleave={() => (hoverRating = 0)}
      >
        <Icon icon="mingcute:star-fill" width="32" height="32" />
      </button>
    {/each}
    <span
      class="ml-3 w-16 text-lg font-semibold text-light-palette-text-primary dark:text-dark-palette-text-primary"
    >
      {$formDataStore.rating}.0/5.0
    </span>
  </div>
  <InputError
    styles="w-full"
    wasSubmitted={$formWasSubmittedStore}
    error={$formErrorsStore.rating && $formErrorsStore.rating[0]}
  />
</OrderByIdReviewTabRatingSectionLayout>
