<script lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import type { LayoutProps } from '@/types/LayoutProps'
  import type { OrderSubmitRequirementsFormData } from '@/types/OrderSubmitRequirementsFormData'
  import type { OrderSubmitRequirementsFormErrors } from '@/types/OrderSubmitRequirementsFormErrors'
  import type { ServiceQuestion } from '@/types/ServiceQuestion'
  import { validator } from '@felte/validator-zod'
  import { createForm } from 'felte'
  import { handleOrderSubmitRequirementsFormSubmit } from '../../handlers'
  import {
    formatOrderFormData,
    getOrderSubmitRequirementsFormInitialValues,
    orderSubmitRequirementsFormSchema,
  } from '../../helpers'
  import { orderSubmitRequirementsFormFilesStore } from '../../order-submit-requirements-store'

  type OrderSubmitRequirementsFormProps = LayoutProps & {
    freelanceQuestions: ServiceQuestion[]
    orderFormErrors: OrderSubmitRequirementsFormErrors | undefined
    orderFormData: OrderSubmitRequirementsFormData | undefined
    wasSubmitted: boolean
    orderId: string
  }

  let {
    children,
    orderFormErrors = $bindable(),
    wasSubmitted = $bindable(),
    orderFormData = $bindable(),
    freelanceQuestions,
    orderId,
  }: OrderSubmitRequirementsFormProps = $props()

  let { form, errors, isSubmitting, data, setData } = createForm<OrderSubmitRequirementsFormData>({
    onSubmit: (data) => handleOrderSubmitRequirementsFormSubmit(data),
    initialValues: getOrderSubmitRequirementsFormInitialValues(freelanceQuestions, orderId),
    extend: validator({ schema: orderSubmitRequirementsFormSchema }),
  })

  errors.subscribe((value) => (orderFormErrors = formatOrderFormData(value)))
  isSubmitting.subscribe((value) => (value ? (wasSubmitted = true) : null))
  data.subscribe((origialData) => (orderFormData = origialData))

  orderSubmitRequirementsFormFilesStore.subscribe((value) => setData('files', value))
</script>

<PaperElement
  styles="!p-0 shadow-none md:shadow-md dark:shadow-none !ring-0 md:!ring-1 !bg-transparent md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block"
>
  <form use:form>
    {@render children()}
  </form>
</PaperElement>
