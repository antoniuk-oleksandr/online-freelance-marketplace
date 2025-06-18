<script lang="ts">
  import type { ReviewTabFormData } from '@/types/ReviewTabFormData'
  import { createForm } from 'felte'
  import { handleReviewTabFormSubmit } from '../../handlers'
  import { validator } from '@felte/validator-zod'
  import { reviewTabFormSchema } from '../../helpers'
  import OrderByIdReviewTabFormContent from '../OrderByIdReviewTabFormContent/OrderByIdReviewTabFormContent.svelte'
  import { onDestroy, setContext } from 'svelte'
  import { writable } from 'svelte/store'

  type OrderByIdReviewTabFormProps = {
    orderId: string
  }

  const { orderId }: OrderByIdReviewTabFormProps = $props()

  const { form, data, errors, isSubmitting } = createForm<ReviewTabFormData>({
    initialValues: {
      rating: 0,
      reviewMessage: '',
    },
    onSubmit: (data) => handleReviewTabFormSubmit(data, orderId),
    extend: validator({ schema: reviewTabFormSchema }),
  })

  const unsubIsSubmitting = isSubmitting.subscribe((value) => {
    if (!value) return

    wasSubmittedStore.set(value)
    unsubIsSubmitting()
  })

  const wasSubmittedStore = writable(false)
  setContext('formDataStore', data)
  setContext('formErrorsStore', errors)
  setContext('formIsSubmittingStore', isSubmitting)
  setContext('formWasSubmittedStore', wasSubmittedStore)

  onDestroy(() => unsubIsSubmitting())
</script>

<form class="flex flex-col gap-6 grow" use:form>
  <OrderByIdReviewTabFormContent />
</form>
