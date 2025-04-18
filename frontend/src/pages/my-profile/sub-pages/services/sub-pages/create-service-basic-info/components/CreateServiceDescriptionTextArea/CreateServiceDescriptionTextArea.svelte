<script lang="ts">
  import RichTextEditor from '@/common-components/RichTextEditor/RichTextEditor.svelte'
  import ServiceCreatePageDescriptionTextAreaLayout from './CreateServiceDescriptionTextAreaLayout.svelte'
  import Label from '@/common-components/Label/Label.svelte'
  import type { Writable } from 'svelte/store'
  import type { CreateServiceBasicInfoFormData } from '@/types/CreateServiceBasicInfoFormData'

  type ServiceCreatePageDescriptionTextAreaProps = {
    formDataStore: Writable<CreateServiceBasicInfoFormData>
    error?: string | null
  }

  const { formDataStore, error }: ServiceCreatePageDescriptionTextAreaProps = $props()

  const updateDescription = (value: string) =>
    formDataStore.update((prev) => ({ ...prev, description: value }))
</script>

<ServiceCreatePageDescriptionTextAreaLayout>
  <Label text="Description" />
  <RichTextEditor
    content={$formDataStore.description}
    {error}
    onChange={(value) => updateDescription(value)}
    limit={2000}
  />
</ServiceCreatePageDescriptionTextAreaLayout>
