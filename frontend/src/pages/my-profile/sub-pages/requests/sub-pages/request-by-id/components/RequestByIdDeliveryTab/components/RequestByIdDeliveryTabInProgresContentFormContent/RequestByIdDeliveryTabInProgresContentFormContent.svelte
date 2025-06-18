<script lang="ts">
  import Textarea from '@/common-components/Textarea/Textarea.svelte'
  import Title from '@/common-components/Title/Title.svelte'
  import RequestByIdDeliveryTabInProgresContentFooter from '../RequestByIdDeliveryTabInProgresContentFooter/RequestByIdDeliveryTabInProgresContentFooter.svelte'
  import RequestByIdDeliveryTabInProgresContentAttachments from '../RequestByIdDeliveryTabInProgresContentAttachments/RequestByIdDeliveryTabInProgresContentAttachments.svelte'
  import { getContext } from 'svelte'
  import type { Writable } from 'svelte/store'
  import type { DeliveryFormData } from '@/types/DeliveryFormData'

  const formDataStore = getContext<Writable<DeliveryFormData>>('formDataStore')
  const setFiles = (newFiles: File[]) =>
    formDataStore.update((prev) => ({
      ...prev,
      files: [...prev.files, ...newFiles],
    }))
  const formErrorStore = getContext<Writable<Record<string, string[]>>>('formErrorsStore')
  const formWasSubmitted = getContext<Writable<boolean>>('formWasSubmitted')
</script>

<Title
  text="Deliver your completed work"
  description="Send files and a message to complete this order"
/>
<Textarea
  wasSubmitted={$formWasSubmitted}
  error={$formErrorStore.message && $formErrorStore.message[0]}
  id="message"
  name="message"
  useCounter={true}
  label="Delivery Message"
  placeholder="Describe what you're delivering, include instructions, or share important details..."
/>
<RequestByIdDeliveryTabInProgresContentAttachments
  errors={$formErrorStore}
  files={$formDataStore.files}
  formWasSubmitted={$formWasSubmitted}
  {setFiles}
/>
<RequestByIdDeliveryTabInProgresContentFooter
  files={$formDataStore.files}
/>
