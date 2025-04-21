<script lang="ts">
  import { getContext } from 'svelte'
  import MyProfileEditPageAvatarBlock from '../MyProfileEditPageAvatarBlock/MyProfileEditPageAvatarBlock.svelte'
  import MyProfileEditPageInputsBlock from '../MyProfileEditPageInputsBlock/MyProfileEditPageInputsBlock.svelte'
  import MyProfileEditPageSubmitButton from '../MyProfileEditPageSubmitButton/MyProfileEditPageSubmitButton.svelte'
  import type { Writable } from 'svelte/store'
  import type { MyProfileEditPageFormData } from '@/types/MyProfileEditPageFormData'
  import type { FilterParams } from '@/types/FilterParams'

  type MyProfileEditPageFormContentProps = {
    filterParams: FilterParams
  }

  const props: MyProfileEditPageFormContentProps = $props()

  const formDataStore = getContext<Writable<MyProfileEditPageFormData>>('formDataStore')
</script>

<MyProfileEditPageAvatarBlock {formDataStore} />
<MyProfileEditPageInputsBlock {formDataStore} {...props} />
<MyProfileEditPageSubmitButton />
