<script lang="ts">
  import AttachFilesButton from '@/common-components/AttachFilesButton/AttachFilesButton.svelte'
  import Label from '@/common-components/Label/Label.svelte'
  import OrderSubmitRequirementsFormAttachFileList from '@/pages/order-submit-requirements/components/OrderSubmitRequirementsForm/components/OrderSubmitRequirementsFormAttachFileList/OrderSubmitRequirementsFormAttachFileList.svelte'
  import RequestByIdDeliveryTabInProgresContentAttachmentsLayout from './RequestByIdDeliveryTabInProgresContentAttachmentsLayout.svelte'
  import InputError from '@/common-components/Sign/components/SignInput/components/InputError/InputError.svelte'

  type RequestByIdDeliveryTabInProgresContentAttachmentsProps = {
    files: File[]
    setFiles: (newFiles: File[]) => void
    errors: Record<string, string[]>
    formWasSubmitted: boolean
  }

  const {
    files,
    setFiles,
    errors,
    formWasSubmitted,
  }: RequestByIdDeliveryTabInProgresContentAttachmentsProps = $props()

  const removeFile = (index: number) => {
    const updatedFiles = files.filter((_, i) => i !== index)
    setFiles(updatedFiles)
  }

  const addFiles = (newFiles: FileList) => {
    const updatedFiles = [...files, ...Array.from(newFiles)]
    setFiles(updatedFiles)
  }
</script>

<RequestByIdDeliveryTabInProgresContentAttachmentsLayout>
  <Label text="Attachments" />
  <div class="flex flex-col gap-3">
    <AttachFilesButton {addFiles} styles="!w-full" />
    <OrderSubmitRequirementsFormAttachFileList {files} {removeFile} />
  </div>
  <InputError wasSubmitted={formWasSubmitted} error={errors.files && errors.files[0]} />
</RequestByIdDeliveryTabInProgresContentAttachmentsLayout>
