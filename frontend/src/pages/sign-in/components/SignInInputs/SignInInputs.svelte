<script lang="ts">
  import Input from '@/common-components/Sign/components/SignInput/Input.svelte'
  import { formStore } from '@/common-components/Sign/stores/form-store'
  import type { FormStore } from '@/types/FormStore.ts'
  import BottomSignLine from '@/pages/sign-in/components/BottomSignLine/BottomSignLine.svelte'

  type SignInInputsProps = {
    keepSignedIn: boolean
  }

  let { keepSignedIn = $bindable() }: SignInInputsProps = $props()

  let formStoreData = $state<null | FormStore>(null)
  formStore.subscribe((value) => (formStoreData = value))
</script>

{#if formStoreData && formStoreData.data && formStoreData.errors}
  <Input
    wasSubmitted={formStoreData.wasSubmitted}
    value={formStoreData.data.usernameOrEmail}
    error={formStoreData.errors.usernameOrEmail?.[0] || null}
    id="usernameOrEmail"
    type="text"
    label="Username or Email"
  />
  <Input
    wasSubmitted={formStoreData.wasSubmitted}
    value={formStoreData.data.password}
    error={formStoreData.errors.password?.[0] || null}
    id="password"
    type="password"
    label="Password"
  />
  <BottomSignLine bind:keepSignedIn />
{/if}
