<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'
  import { validator } from '@felte/validator-zod'
  import { createForm } from 'felte'
  import { deliveryFormSchema } from '../../helpers'
  import type { DeliveryFormData } from '@/types/DeliveryFormData'
  import { handleDeliveryFormSubmit } from '../../handlers'
  import { onDestroy, setContext } from 'svelte'
  import { writable } from 'svelte/store'
  import type { DeliveryData } from '@/types/DeliveryData'

  type RequestByIdDeliveryTabInProgresContentFormProps = LayoutProps & {
    orderId: string
    setNewDeliveryData: (data: DeliveryData) => void
  }

  const { orderId, children, setNewDeliveryData }: RequestByIdDeliveryTabInProgresContentFormProps =
    $props()

  const { form, data, errors, isSubmitting } = createForm<DeliveryFormData>({
    extend: validator({ schema: deliveryFormSchema }),
    onSubmit: (values) => handleDeliveryFormSubmit(values, orderId, setNewDeliveryData),
    initialValues: {
      message: '',
      files: [],
    },
  })

  const unsubscribeIsSubmitting = isSubmitting.subscribe((value) => {
    if (!value) return

    formWasSubmittedStore.set(true)
    unsubscribeIsSubmitting()
  })

  const formWasSubmittedStore = writable(false)
  setContext('formDataStore', data)
  setContext('formErrorsStore', errors)
  setContext('formIsSubmittingStore', isSubmitting)
  setContext('formWasSubmitted', formWasSubmittedStore)

  onDestroy(() => unsubscribeIsSubmitting())
</script>

<form use:form class="size-full p-6 flex flex-col gap-6">
  {@render children()}
</form>
