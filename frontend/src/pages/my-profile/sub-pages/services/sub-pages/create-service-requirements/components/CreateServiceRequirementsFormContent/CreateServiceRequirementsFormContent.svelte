<script lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import type { CreateServiceRequirementsFormData } from '@/types/CreateServiceRequirementsFormData'
  import { getContext } from 'svelte'
  import type { Writable } from 'svelte/store'
  import CreateServiceQuestionsBlock from '../CreateServiceQuestionsBlock/CreateServiceQuestionsBlock.svelte'

  let formDataStore: Writable<CreateServiceRequirementsFormData> = getContext('formDataStore')
  let formErrorsStore: Writable<Record<string, string[]>> = getContext('formErrorsStore')
  let formIsSubmittingStore: Writable<boolean> = getContext('formIsSubmittingStore')
</script>

<CreateServiceQuestionsBlock {formErrorsStore} {formDataStore} />
<Button loading={$formIsSubmittingStore} styles="w-full mt-auto lg:w-40 self-end" type="submit">
  Save & Continue
</Button>
