<script lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import LabeledInput from '@/common-components/LabeledInput/LabeledInput.svelte'
  import type { CreateServiceBasicInfoFormData } from '@/types/CreateServiceBasicInfoFormData'
  import { getContext, onDestroy } from 'svelte'
  import type { Writable } from 'svelte/store'
  import ServiceCreatePageCategorySelect from '../CreateServiceCategorySelect/CreateServiceCategorySelect.svelte'
  import ServiceCreatePageDescriptionTextArea from '../CreateServiceDescriptionTextArea/CreateServiceDescriptionTextArea.svelte'

  let formDataStore: Writable<CreateServiceBasicInfoFormData> = getContext('formDataStore')
  let formErrorsStore: Writable<Record<string, string[]>> = getContext('formErrorsStore')
  let formIsSubmittingStore: Writable<boolean> = getContext('formIsSubmittingStore')

  let formData = $state<CreateServiceBasicInfoFormData | undefined>()
  let formErrors = $state<Record<string, string[]>>()
  let formIsSubmitting = $state<boolean>()

  const unsubscribeFormDataStore = formDataStore.subscribe((data) => (formData = data))
  const unsubscribeFormErrorsStore = formErrorsStore.subscribe((errors) => (formErrors = errors))
  const unsubscribeFormIsSubmittingStore = formIsSubmittingStore.subscribe(
    (isSubmitting) => (formIsSubmitting = isSubmitting),
  )

  onDestroy(() => {
    unsubscribeFormDataStore()
    unsubscribeFormErrorsStore()
    unsubscribeFormIsSubmittingStore()
  })
</script>

{#if formData && formErrors && formIsSubmitting !== undefined}
  <LabeledInput
    id="title"
    error={formErrors && formErrors.title && formErrors.title[0]}
    placeholder="Enter service title"
    label="Title"
  />
  <ServiceCreatePageCategorySelect
    {formDataStore}
    error={formErrors && formErrors.categoryId && formErrors.categoryId[0]}
  />
  <ServiceCreatePageDescriptionTextArea
    {formDataStore}
    error={formErrors && formErrors.description && formErrors.description[0]}
  />
  <Button loading={formIsSubmitting} styles="mt-auto w-full lg:w-40 self-end" type="submit">
    Save & Continue
  </Button>
{/if}
