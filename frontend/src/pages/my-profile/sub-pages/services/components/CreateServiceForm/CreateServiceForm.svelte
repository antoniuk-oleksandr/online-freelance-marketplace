<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'
  import { validator } from '@felte/validator-zod'
  import { createForm } from 'felte'
  import { onDestroy, onMount, setContext } from 'svelte'
  import { createServiceStore } from '../../stores/create-service-store'
  import type { z } from 'zod'

  type CreateServiceFormProps = LayoutProps & {
    onSubmit: (data: any) => void
    schema: z.ZodObject<any>
  }

  const { children, onSubmit, schema }: CreateServiceFormProps = $props()

  let { form, isSubmitting, data, errors } = createForm({
    onSubmit,
    extend: validator({ schema }),
    initialValues: $createServiceStore,
  })

  setContext('formDataStore', data)
  setContext('formErrorsStore', errors)
  setContext('formIsSubmittingStore', isSubmitting)
</script>

<form class="flex flex-grow flex-col gap-6" use:form>
  {@render children()}
</form>
