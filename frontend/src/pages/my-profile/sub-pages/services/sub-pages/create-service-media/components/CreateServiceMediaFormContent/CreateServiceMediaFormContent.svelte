<script lang="ts">
  import { getContext } from 'svelte'
  import ServiceCreateFilesBlock from '../CreateServiceFilesBlock/CreateServiceFilesBlock.svelte'
  import type { Writable } from 'svelte/store'
  import type { CreateServiceMediaFormData } from '@/types/CreateServiceMediaFormData'
    import Button from '@/common-components/Button/Button.svelte'

  let formDataStore: Writable<CreateServiceMediaFormData> = getContext('formDataStore')
  let formErrorsStore: Writable<Record<string, string[]>> = getContext('formErrorsStore')
  let formIsSubmittingStore: Writable<boolean> = getContext('formIsSubmittingStore')
</script>

<ServiceCreateFilesBlock {formErrorsStore} {formDataStore}/>
<Button loading={$formIsSubmittingStore} styles="w-full mt-auto lg:w-40 self-end" type="submit">
  Create Service
</Button>
