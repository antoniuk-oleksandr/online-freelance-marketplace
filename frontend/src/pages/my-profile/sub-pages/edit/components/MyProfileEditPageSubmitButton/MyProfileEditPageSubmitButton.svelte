<script lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import MyProfileEditPageSubmitButtonLayout from './MyProfileEditPageSubmitButtonLayout.svelte'
  import { getContext } from 'svelte'
  import type { Writable } from 'svelte/store'

  const formIsSubmittingStore = getContext<Writable<boolean>>('formIsSubmittingStore')
</script>

<MyProfileEditPageSubmitButtonLayout>
  <Button loading={$formIsSubmittingStore} type="submit" styles="w-40 ">Save Changes</Button>
</MyProfileEditPageSubmitButtonLayout>
