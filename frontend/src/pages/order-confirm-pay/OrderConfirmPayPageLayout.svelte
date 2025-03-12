<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'
  import { createForm } from 'felte'
  import { setContext } from 'svelte'
  import { paymentFormSchema } from '@/pages/order-confirm-pay/helpers'
  import { validator } from '@felte/validator-zod'
  import { handlePaymentFormSubmit } from '@/pages/order-confirm-pay/handlers'
  import type { PaymentData } from '@/types/PaymentData.ts'

  type OrderConfirmPayPageLayoutProps = LayoutProps & {
    serviceId: string
    packageId: string
  }

  const { children, serviceId, packageId }: OrderConfirmPayPageLayoutProps = $props()

  const { form, setFields, data, isSubmitting, errors } = createForm<PaymentData>({
    onSubmit: (data) => handlePaymentFormSubmit(data, serviceId, packageId),
    extend: validator({ schema: paymentFormSchema }),
    // initialValues: {
    //     cardNumber: "4263 9826 4026 9299",
    //     cardHolderName: "Alex",
    //     expiryDate: "10/26",
    //     securityCode: "123",
    // },
  })

  setContext('feltData', data)
  setContext('errors', errors)
  setContext('setFields', setFields)
  setContext('orderComfirmPayFormIsSubmitting', isSubmitting)
</script>

<form use:form class="flex flex-col gap-6 text-base h-full">
  {@render children()}
</form>
