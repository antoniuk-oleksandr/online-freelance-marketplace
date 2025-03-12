<script lang="ts">
  import AttachFilesButton from '@/common-components/AttachFilesButton/AttachFilesButton.svelte'
  import OrderSubmitRequirementsFormAttachFilesBlockLayout from './OrderSubmitRequirementsFormAttachFilesBlockLayout.svelte'
  import OrderSubmitRequirementsFormAttachFileList from '../OrderSubmitRequirementsFormAttachFileList/OrderSubmitRequirementsFormAttachFileList.svelte'
  import { orderSubmitRequirementsFormFilesStore } from '@/pages/order-submit-requirements/order-submit-requirements-store'

  const addFiles = (newFiles: FileList) => {
    orderSubmitRequirementsFormFilesStore.update((prev) => [...prev, ...Array.from(newFiles)])
  }

  const removeFile = (index: number) => {
    orderSubmitRequirementsFormFilesStore.update((prev) => {
      const newFiles = [...prev]
      newFiles.splice(index, 1)
      return newFiles
    })
  }
</script>

<OrderSubmitRequirementsFormAttachFilesBlockLayout>
  <span>You can optionally attach files to illustrate your order more clearly.</span>
  <AttachFilesButton {addFiles} />
  <OrderSubmitRequirementsFormAttachFileList
    files={$orderSubmitRequirementsFormFilesStore}
    {removeFile}
  />
</OrderSubmitRequirementsFormAttachFilesBlockLayout>
