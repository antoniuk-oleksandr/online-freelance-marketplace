<script lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import Title from '@/common-components/Title/Title.svelte'
  import type { CreateServicePackageFormData } from '@/types/CreateServicePackageFormData'
  import { getContext, onDestroy } from 'svelte'
  import type { Writable } from 'svelte/store'
  import CreateServicePackagesTable from '../CreateServicePackagesTable/CreateServicePackagesTable.svelte'
  import LabeledInputError from '@/common-components/LabeledInput/components/LabeledInputError/LabeledInputError.svelte'
  import { isString } from '@tiptap/core'

  let formDataStore: Writable<CreateServicePackageFormData> = getContext('formDataStore')
  let formErrorsStore: Writable<Record<string, string[]>> = getContext('formErrorsStore')
  let formIsSubmittingStore: Writable<boolean> = getContext('formIsSubmittingStore')

  let formData = $state<CreateServicePackageFormData | undefined>()
  let formErrors = $state<Record<string, string[]>>()
  let formIsSubmitting = $state<boolean>()
  let formWasSubmitted = $state<boolean>()

  const unsubscribeFormDataStore = formDataStore.subscribe((data) => (formData = data))
  const unsubscribeFormErrorsStore = formErrorsStore.subscribe((errors) => (formErrors = errors))
  const unsubscribeFormIsSubmittingStore = formIsSubmittingStore.subscribe((isSubmitting) => {
    if (!formWasSubmitted && isSubmitting) formWasSubmitted = true
    formIsSubmitting = isSubmitting
  })

  let packagesError = $derived(formErrors && formErrors.packages && formErrors.packages[0])

  onDestroy(() => {
    unsubscribeFormDataStore()
    unsubscribeFormErrorsStore()
    unsubscribeFormIsSubmittingStore()
  })

</script>

{#if formData && formErrors && formIsSubmitting !== undefined}
  <CreateServicePackagesTable />
  {#if formWasSubmitted && isString(packagesError)}
    <LabeledInputError error={packagesError} />
  {/if}
  <Button loading={formIsSubmitting} styles="mt-auto w-full lg:w-40 self-end" type="submit">
    Save & Continue
  </Button>
{/if}
