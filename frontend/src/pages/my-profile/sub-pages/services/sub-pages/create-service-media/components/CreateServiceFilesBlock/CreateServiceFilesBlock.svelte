<script lang="ts">
  import AttachFilesButton from '@/common-components/AttachFilesButton/AttachFilesButton.svelte'
  import LabeledInputError from '@/common-components/LabeledInput/components/LabeledInputError/LabeledInputError.svelte'
  import OrderSubmitRequirementsFormAttachFileList from '@/pages/order-submit-requirements/components/OrderSubmitRequirementsForm/components/OrderSubmitRequirementsFormAttachFileList/OrderSubmitRequirementsFormAttachFileList.svelte'
  import type { CreateServiceMediaFormData } from '@/types/CreateServiceMediaFormData'
  import type { Writable } from 'svelte/store'

  type CreateServiceFilesBlockProps = {
    formDataStore: Writable<CreateServiceMediaFormData>
    formErrorsStore: Writable<Record<string, string[]>>
  }

  const { formDataStore, formErrorsStore }: CreateServiceFilesBlockProps = $props()

  const addFiles = (newFiles: FileList) => {
    formDataStore.update((prev) => ({
      ...prev,
      files: [...prev.files, ...newFiles],
    }))
  }

  const removeFile = (index: number) =>
    formDataStore.update((prev) => {
      const files = prev.files
      files.splice(index, 1)
      return { ...prev, files  }
    })
</script>

<div class="flex flex-col">
  <AttachFilesButton accept="image" {addFiles} />
  <LabeledInputError styles="mt-1" error={$formErrorsStore.files && $formErrorsStore.files[0]} />
  <OrderSubmitRequirementsFormAttachFileList styles="mt-6" files={$formDataStore.files} {removeFile} />
</div>
