<script lang="ts">
  import Label from '@/common-components/Label/Label.svelte'
  import Textarea from '@/common-components/Textarea/Textarea.svelte'
  import OrderByIdReviewTabMessageSectionLayout from './OrderByIdReviewTabMessageSectionLayout.svelte'
  import { getContext } from 'svelte'
  import type { Writable } from 'svelte/store'
  import type { ReviewTabFormData } from '@/types/ReviewTabFormData'

  const formDataStore = getContext<Writable<ReviewTabFormData>>('formDataStore')
  const formWasSubmittedStore = getContext<Writable<boolean>>('formWasSubmittedStore')
  const formErrorsStore = getContext<Writable<Record<string, string[]>>>('formErrorsStore')
</script>

<OrderByIdReviewTabMessageSectionLayout>
  <Label text="Write your review (optional)" />
  <Textarea
    placeholder="Share details of your experience with this service. What did you like? What could be improved?"
    id="reviewMessage"
    name="reviewMessage"
    wasSubmitted={$formWasSubmittedStore}
    useCounter={true}
    error={$formErrorsStore.reviewMessage && $formErrorsStore.reviewMessage[0]}
    value={$formDataStore.reviewMessage}
  />
</OrderByIdReviewTabMessageSectionLayout>
