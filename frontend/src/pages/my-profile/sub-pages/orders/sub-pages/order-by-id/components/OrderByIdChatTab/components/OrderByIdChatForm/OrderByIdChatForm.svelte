<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'
  import { createForm } from 'felte'
  import { handleChatFormSubmit, handleChatKeyDown } from '../../handlers'
  import { validator } from '@felte/validator-zod'
  import { chatFormSchema } from '../../helpers'
  import type { ChatFormData } from '@/types/ChatFormData'
  import { onDestroy } from 'svelte'

  type OrderByIdChatFormProps = LayoutProps & {
    chatParentId: number
    orderId: number
  }

  const {
    children,
    chatParentId,
    orderId,
  }: OrderByIdChatFormProps = $props()

  let resetFn: () => void

  const { form, reset, handleSubmit } = createForm<ChatFormData>({
    onSubmit: (data) => handleChatFormSubmit(data, resetFn),
    extend: validator({ schema: chatFormSchema }),
    initialValues: { orderId, chatParentId },
  })

  const handleKeyDown = (event: any) => handleChatKeyDown(event, handleSubmit)
  document.addEventListener('keydown', handleKeyDown)
  onDestroy(() => document.removeEventListener('keydown', handleKeyDown))

  resetFn = reset
</script>

<form use:form class="flex gap-3 items-center py-3">
  {@render children()}
</form>
