<script lang="ts">
  import Input from '@/common-components/Sign/components/SignInput/Input.svelte'
  import type { FormStore } from '@/types/FormStore.ts'
  import { formStore } from '@/common-components/Sign/stores/form-store'

  let formStoreData = $state<null | FormStore>(null)
  formStore.subscribe((value) => (formStoreData = value))
</script>

{#if formStoreData && formStoreData.data && formStoreData.errors}
  <Input
    wasSubmitted={formStoreData.wasSubmitted}
    value={formStoreData.data.email}
    error={formStoreData.errors.email?.[0] || null}
    id="email"
    type="text"
    label="Email address"
  />
{/if}
