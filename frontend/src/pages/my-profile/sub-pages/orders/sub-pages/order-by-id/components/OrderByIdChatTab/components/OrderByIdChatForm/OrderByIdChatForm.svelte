<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'
  import { createForm } from 'felte'
  import { handleChatFormSubmit } from '../../handlers'
  import { validator } from '@felte/validator-zod'
  import { chatFormSchema } from '../../helpers'
  import type { ChatFormData } from '@/types/ChatFormData'

  const { children }: LayoutProps = $props()

  let resetFn: () => void

  const { form, reset } = createForm<ChatFormData>({
    onSubmit: (data) => handleChatFormSubmit(data, resetFn),
    extend: validator({ schema: chatFormSchema }),
  })

  resetFn = reset
</script>

<form use:form class="flex gap-3 items-center py-3">
  {@render children()}
</form>
