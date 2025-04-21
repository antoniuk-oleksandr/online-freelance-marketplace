<script lang="ts">
  import Label from '@/common-components/Label/Label.svelte'
  import RichTextEditor from '@/common-components/RichTextEditor/RichTextEditor.svelte'
  import MyProfileEditPageAboutInputLayout from './MyProfileEditPageAboutInputLayout.svelte'
  import type { MyProfileEditPageFormData } from '@/types/MyProfileEditPageFormData'
  import type { Writable } from 'svelte/store'

  type MyProfileEditPageInputsBlockProps = {
    formDataStore: Writable<MyProfileEditPageFormData>
  }

  const { formDataStore }: MyProfileEditPageInputsBlockProps = $props()

  const onChange = (value: string) => {
    formDataStore.update((prev) => ({
      ...prev,
      about: value,
    }))
  }
</script>

<MyProfileEditPageAboutInputLayout>
  <Label text="About" />
  <RichTextEditor content={$formDataStore.about} {onChange} limit={2000} />
</MyProfileEditPageAboutInputLayout>
